# requestBsc 宝塔线上更新指南（MySQL 5.7）

本次是在现有线上数据库上执行的增量更新。`requestbsc_big_update_mysql57.sql` 只新增或补充表、字段和索引，不应清空现有数据。SQL 采用 MySQL 5.7 支持的 `utf8mb4_unicode_ci`，不使用 MySQL 8 的 `utf8mb4_0900_ai_ci`。

## 一、更新前

1. 在宝塔「计划任务」中暂停 requestBsc 的所有定时请求，避免在备份和改表时写入数据。
2. 确认服务目录是 `/www/wwwroot/requestBsc`，数据库名是 `bscdata`。如果线上实际名称不同，下面命令中仅替换数据库名。
3. 保留现有的配置和环境变量文件，不要把节点地址、密码或密钥写入 Git、SQL 包或本指南。

### 1. 备份数据库

可以直接使用宝塔「数据库 -> 备份」功能，并确认备份文件大小不为 0。也可在终端执行：

```bash
install -d -m 700 /www/backup/database/requestbsc
BACKUP_FILE="/www/backup/database/requestbsc/bscdata_before_big_update_$(date +%F_%H%M%S).sql"
mysqldump --default-character-set=utf8mb4 \
  --single-transaction --quick --routines --triggers --events \
  -u root -p bscdata > "$BACKUP_FILE"
test -s "$BACKUP_FILE" && ls -lh "$BACKUP_FILE"
```

`-p` 会让 MySQL 交互式询问密码，不要把密码直接写在命令中。备份未成功时不要继续更新。

## 二、拉取 main、导入 SQL、构建并重启

### 1. 拉取 main

```bash
cd /www/wwwroot/requestBsc
git status --short
git pull --ff-only origin main
```

如果 `git status --short` 显示线上有未确认的手工改动，先备份并确认它们，不要直接覆盖。

### 2. 导入 MySQL 5.7 增量 SQL

```bash
cd /www/wwwroot/requestBsc
mysql --default-character-set=utf8mb4 -u root -p bscdata \
  < ./deploy/requestbsc_big_update_mysql57.sql
```

导入完成后命令应返回 0，不应出现 `ERROR 1146`、`ERROR 1273` 或其他 SQL 错误。该 SQL 按可重复执行设计，但每次重新执行前仍应先备份。

### 3. 构建并重启 requestBsc

```bash
cd /www/wwwroot/requestBsc
mkdir -p ./bin ./log
go build -o ./bin/requestEth ./cmd/requestEth

old_pid=$(lsof -t -iTCP:8000 -sTCP:LISTEN)
if [ -n "$old_pid" ]; then
  kill $old_pid
  sleep 2
fi

set -a
source /www/server/requestBsc/request-bsc.env
set +a

nohup ./bin/requestEth -conf ./configs \
  > ./log/requestEth.out 2>&1 &

sleep 3
lsof -nP -iTCP:8000 -sTCP:LISTEN
tail -n 30 ./log/requestEth.out
```

只要 `127.0.0.1:8000` 正在监听，且启动日志没有配置、数据库或合约初始化错误，就可以继续。节点配置仍使用线上现有的私密配置，本文不记录其内容。

## 三、一次性恢复和补数据

下列操作在增量 SQL 导入、新服务启动后执行。这一阶段保持宝塔定时任务暂停，所有请求只走本机 `127.0.0.1`。

### 1. 先同步 Bound

```bash
curl -fsS --connect-timeout 10 --max-time 75 \
  http://127.0.0.1:8000/api/get_user_bound_event
echo
```

订单和投资数据依赖用户已经存在，因此 Bound 必须先执行。如果线上仍有较大的历史积压，或后续恢复接口提示「用户不存在」，继续重复请求本接口；它会按数据库断点续跑，无须从头开始。

### 2. 反复执行投资数据修复，直到 `finished=true`

```bash
curl -fsS --connect-timeout 10 --max-time 75 \
  http://127.0.0.1:8000/api/recover_user_investment_data
echo
```

每次检查返回 JSON 中的 `finished`：

- `finished=false`：继续执行同一个 GET。
- `finished=true`：全历史的投资序号和用户投资次数已修复；最近 500,000 个区块内的事件时间也已补齐，足以覆盖今日、昨日金额与订单数统计。
- HTTP 失败或超时：不需要从头开始，处理节点或服务错误后再执行，程序会从数据库断点继续。

### 3. 反复执行订单历史恢复，直到 `finished=true`

```bash
curl -fsS --connect-timeout 10 --max-time 75 \
  http://127.0.0.1:8000/api/recover_staking_order_event
echo
```

每次检查返回中的 `eventFinished`、`remainingSnapshotUserCount` 和 `finished`。只有 `finished=true` 才代表历史事件、当前订单快照和完整性校验都已完成。中断后直接重新请求即可，会从数据库断点继续。

如果恢复订单时明确提示「用户不存在」，先再执行一次 `/api/get_user_bound_event`，然后继续恢复订单。不要手工跳过该错误。

> `/api/recover_user_investment_data` 和 `/api/recover_staking_order_event` 都是一次性补数据接口，**不得加入永久定时任务**。

## 四、永久增量同步定时任务

一次性恢复全部返回 `finished=true` 后，再启用宝塔定时任务。推荐只建立一条「Shell 脚本」任务，每分钟执行，在同一把 `flock` 锁内按顺序请求下列 5 个同步接口：

1. `/api/get_user_bound_event`
2. `/api/get_user_stake_changed_event`
3. `/api/get_user_extra_changed_event`
4. `/api/get_staking_reward_event`
5. `/api/get_staking_order_event`

宝塔 Shell 任务内容：

```bash
#!/bin/bash

exec 9>/tmp/requestbsc_incremental_sync.lock
flock -n 9 || exit 0

BASE_URL="http://127.0.0.1:8000"

curl -fsS --connect-timeout 10 --max-time 75 "$BASE_URL/api/get_user_bound_event" >/dev/null || exit 1

sync_status=0
for endpoint in \
  get_user_stake_changed_event \
  get_user_extra_changed_event \
  get_staking_reward_event \
  get_staking_order_event
do
  curl -fsS --connect-timeout 10 --max-time 75 "$BASE_URL/api/$endpoint" >/dev/null || sync_status=1
done

exit "$sync_status"
```

说明：

- `flock -n` 保证上一轮未结束时本轮直接跳过，不会重入。
- Bound 始终先执行；Bound 失败时停止当轮，因为后续数据依赖用户表。
- 其余某个接口失败时会记录当轮失败，但仍让其他接口获得执行机会；下一分钟各接口都从自己的数据库断点自然继续。
- 每个 curl 最大等待 75 秒；一轮总耗时超过一分钟时，新一轮会被 `flock` 跳过，不影响数据正确性。
- `-f` 会把 HTTP 非 2xx 视为失败；不配置自动重试，避免在同一轮中重入。

### 保留现有 `/api/p_queue`

`/api/p_queue` 是现有认购队列任务，不属于上面 5 个数据同步接口，但原有每分钟定时任务应继续保留。其自身也使用独立锁防止重入：

```bash
flock -n /tmp/requestbsc_p_queue.lock \
  curl -fsS --connect-timeout 10 --max-time 75 \
  http://127.0.0.1:8000/api/p_queue >/dev/null
```

## 五、订单保留规则与上线后检查

### 订单状态

- `status=1`：排队中。
- `status=2`：正在进行。
- `status=3`：已结束/已离场。

订单结束时只把主表记录标记为 `status=3`，**永不物理删除订单记录**。即使该订单已经不能从合约当前订单数组读到，程序也会依据历史结束事件补齐可恢复数据并保留该行。不要对 `staking_v1_order` 或订单事件表执行 `DELETE`/`TRUNCATE`。

上线后可做以下只读检查：

```bash
curl -fsS --connect-timeout 10 --max-time 75 \
  "http://127.0.0.1:8000/api/get_user_overview"
echo

curl -fsS --connect-timeout 10 --max-time 75 \
  "http://127.0.0.1:8000/api/get_staking_order_list?page=1&pageSize=20"
echo
```

在 MySQL 中检查三种状态数量以及已结束订单的收尾值：

```sql
SELECT status, COUNT(*) AS total
FROM staking_v1_order
GROUP BY status
ORDER BY status;

SELECT COUNT(*) AS invalid_exited_orders
FROM staking_v1_order
WHERE status = 3
  AND (remaining <> 0 OR line_claimable <> 0);
```

`invalid_exited_orders` 应为 0。如果接口返回错误、订单完整性校验未通过或该数值不为 0，不要删数据强行跳过；保留断点和错误日志，处理原因后再从原接口继续。
