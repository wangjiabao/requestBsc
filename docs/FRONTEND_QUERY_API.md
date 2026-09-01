# requestBsc 前端查询接口文档

## 1. 文档范围

本文档介绍以下 3 个只读 `GET` 接口：

1. 静态收益提现记录列表：`/api/get_line_claimed_list`
2. 用户全部上级列表：`/api/get_user_ancestor_list`
3. 每日全局统计列表：`/api/get_daily_statistics_list`

这 3 个接口都是前端查询接口，不是后台定时同步接口，不需要加入每分钟定时任务。

## 2. 通用约定

### 2.1 Base URL

前端应将请求域名或 IP 放在环境配置中：

```text
<BASE_URL>
```

本地直连 Go 服务的示例：

```text
http://127.0.0.1:8000
```

完整请求示例：

```text
http://127.0.0.1:8000/api/get_line_claimed_list?page=1&pageSize=20
```

### 2.2 时区和日期

- 接口时区固定为 `Asia/Shanghai`，即 UTC+8。
- 日期参数严格使用 `YYYY-MM-DD`，例如 `2026-08-31`。
- 当天的查询区间为左闭右开：`[当天 00:00:00, 次日 00:00:00)`。
- 返回中的 `timezone` 固定为 `Asia/Shanghai`。

### 2.3 `created_at` 口径

本次接口不新增链上时间字段，日期统计使用数据库已有的 `created_at`。

因此，本文档中的“今日”或“每日”表示：

> 事件写入当前数据库的日期，不是事件在 BSC 链上实际发生的日期。

如果服务停机后集中补拉历史事件，这些历史事件会记入它们实际落库的日期。

### 2.4 分页

| 参数 | 默认值 | 限制 | 说明 |
|---|---:|---:|---|
| `page` | `1` | 最大 `100000000` | 从 1 开始 |
| `pageSize` | `20` | 最大 `100` | 超过 100 时按 100 返回 |

分页返回都包含：

```json
{
  "total": "20074",
  "page": "1",
  "pageSize": "20",
  "list": []
}
```

### 2.5 JSON 数字类型

后端使用 protobuf JSON。为了避免 JavaScript 大整数精度丢失，protobuf 的 `uint64` 字段在 JSON 中会输出为字符串。

例如：

```json
{
  "id": "142066",
  "createdAt": "1788146503",
  "total": "20074"
}
```

前端注意：

- `id`、`userId`、`blockNumber`、`createdAt`、`total`、`page`、`pageSize`和各种 `Count` 字段按字符串接收。
- 需要进行大整数运算时使用 `BigInt`。
- `createdAt` 是 Unix 秒，用于页面显示时可以转成普通日期：

```ts
const date = new Date(Number(record.createdAt) * 1000);
```

### 2.6 金额类型

所有合约金额使用数据库 `DECIMAL(65,18)`，接口统一返回字符串：

```json
{
  "grossU": "2.897678843226788432",
  "currentOrderCap": "600.000000000000000000"
}
```

前端不要直接使用 JavaScript `number` 进行高精度金额运算。建议使用 `decimal.js`、`bignumber.js` 或其他高精度库。

## 3. 静态收益提现记录列表

### 3.1 请求

```http
GET /api/get_line_claimed_list
```

### 3.2 Query 参数

| 参数 | 类型 | 必填 | 默认值 | 说明 |
|---|---|---|---|---|
| `page` | `uint64` | 否 | `1` | 页码 |
| `pageSize` | `uint64` | 否 | `20` | 每页数量，最大 100 |
| `date` | `string` | 否 | 上海当天 | `YYYY-MM-DD` |
| `address` | `string` | 否 | 全部钱包 | 有效 BSC 钱包地址，精确查询 |

### 3.3 请求示例

查询今天：

```bash
curl "<BASE_URL>/api/get_line_claimed_list?page=1&pageSize=20"
```

查询指定日期：

```bash
curl "<BASE_URL>/api/get_line_claimed_list?page=1&pageSize=20&date=2026-08-31"
```

查询指定日期和钱包：

```bash
curl "<BASE_URL>/api/get_line_claimed_list?page=1&pageSize=20&date=2026-08-31&address=0xa854b4fdb5a0839ca4381064bfd13afa45abbe4e"
```

### 3.4 响应示例

```json
{
  "total": "20074",
  "page": "1",
  "pageSize": "1",
  "list": [
    {
      "id": "142066",
      "createdAt": "1788146503",
      "txHash": "0xa74fca8bd4d6384222537cb37701c049044774da29ed28a7dffd342c7d128585",
      "userAddr": "0xa854b4fdb5a0839ca4381064bfd13afa45abbe4e",
      "orderId": "4219",
      "grossU": "2.897678843226788432",
      "feeU": "0.376698249619482496",
      "netU": "2.520980593607305936",
      "paidMs": false,
      "msAmount": "0.000000000000000000",
      "userId": "1521",
      "currentOrderCap": "600.000000000000000000",
      "currentOrderRemaining": "538.175827625570776264",
      "blockNumber": "119078577"
    }
  ],
  "date": "2026-08-31",
  "timezone": "Asia/Shanghai"
}
```

### 3.5 响应字段

| 字段 | JSON 类型 | 说明 |
|---|---|---|
| `total` | `string` | 符合日期和钱包条件的 LineClaimed 总记录数 |
| `page` | `string` | 当前页 |
| `pageSize` | `string` | 当前每页数量 |
| `date` | `string` | 当前查询日期 |
| `timezone` | `string` | 固定 `Asia/Shanghai` |
| `list[].id` | `string` | LineClaimed 表记录 ID |
| `list[].blockNumber` | `string` | 事件区块高度 |
| `list[].createdAt` | `string` | `created_at` 按上海时区解释的 Unix 秒 |
| `list[].txHash` | `string` | 事件交易哈希 |
| `list[].userAddr` | `string` | 提现用户钱包 |
| `list[].orderId` | `string` | 对应订单 ID |
| `list[].grossU` | `string` | 提现毛额 |
| `list[].feeU` | `string` | 手续费 |
| `list[].netU` | `string` | 提现净额，等于 `grossU - feeU` |
| `list[].paidMs` | `boolean` | 是否使用 MS 支付 |
| `list[].msAmount` | `string` | MS 支付数量 |
| `list[].userId` | `string` | 订单关联用户 ID |
| `list[].currentOrderCap` | `string` | 该 `orderId` 在订单表中的当前总额度 `cap` |
| `list[].currentOrderRemaining` | `string` | 该 `orderId` 在订单表中的当前剩余额度 `remaining` |

### 3.6 关联关系

```text
staking_v1_line_claimed_event.order_id
                ↓
staking_v1_order.order_id
```

接口使用数值型 `order_id` 关联，不使用钱包地址跨表关联。

返回的订单额度是请求接口时订单表中的最新状态，不是历史提现发生时的订单快照。

## 4. 用户全部上级列表

### 4.1 请求

```http
GET /api/get_user_ancestor_list
```

### 4.2 Query 参数

| 参数 | 类型 | 必填 | 说明 |
|---|---|---|---|
| `page` | `uint64` | 否 | 默认 1 |
| `pageSize` | `uint64` | 否 | 默认 20，最大 100 |
| `userId` | `uint64` | 二选一 | 目标用户 ID |
| `address` | `string` | 二选一 | 目标用户 BSC 钱包地址 |

参数规则：

- `userId` 和 `address` 至少提供一个。
- 两者同时提供时，必须定位到同一个用户。
- 当前用户自己不放入返回列表。
- 返回顺序为：直接上级 → 更高上级 → 根用户。
- 根用户返回 `total = "0"` 和空列表。

### 4.3 请求示例

根据用户 ID：

```bash
curl "<BASE_URL>/api/get_user_ancestor_list?userId=7&page=1&pageSize=20"
```

根据钱包：

```bash
curl "<BASE_URL>/api/get_user_ancestor_list?address=0xb73eb156e770356f16ef0b66a9a5a4df01f73a12&page=1&pageSize=20"
```

### 4.4 响应示例

用户 `7` 的 `recommendCode` 为 `D1D5D6`，返回的上级 ID 顺序为 `6 → 5 → 1`：

```json
{
  "total": "3",
  "page": "1",
  "pageSize": "20",
  "list": [
    {
      "id": "6",
      "blockNumber": "99852022",
      "userAddr": "0xcddc6cfd61305e0300d197456f9e3aaf0a062c68",
      "parentAddr": "0xf410c7c91b364a6f017b00b9a79863eb00248a48",
      "recommendCode": "D1D5",
      "amount": "0.000000000000000000",
      "amountHistory": "0.000000000000000000",
      "investmentCount": "0",
      "childrenAmount": "335996.000000000000000000",
      "childrenAmountHistory": "362425.000000000000000000",
      "childrenAmountExtra": "500000.000000000000000000",
      "rewardRecommendAmount": "0.000000000000000000",
      "rewardRecommendPay": "0.000000000000000000",
      "rewardRecommendStoreAmount": "0.000000000000000000",
      "rewardRecommendFee": "0.000000000000000000",
      "rewardRecommendTeamUAmount": "0.000000000000000000",
      "rewardRecommendClaimedTeamUNet": "0.000000000000000000",
      "rewardRecommendClaimedTeamUAmount": "0.000000000000000000",
      "rewardRecommendClaimedTeamUFee": "0.000000000000000000",
      "rewardRecommendExpired": "0.000000000000000000",
      "lineU": "0.000000000000000000",
      "lineCoinU": "0.000000000000000000",
      "lineCoin": "0.000000000000000000",
      "lineFee": "0.000000000000000000",
      "levelReward": "0.000000000000000000",
      "createdAt": "1785764665",
      "updatedAt": "1788144882",
      "name": ""
    }
  ]
}
```

> 示例只展示了列表中的第一位上级，实际响应还会继续返回用户 `5` 和根用户 `1`。

### 4.5 用户字段

每个 `list[]` 对象与现有 `/api/get_user_list` 的用户对象字段完全一致：

| 字段 | JSON 类型 | 说明 |
|---|---|---|
| `id` | `string` | 用户 ID |
| `blockNumber` | `string` | 用户注册区块 |
| `userAddr` | `string` | 用户钱包 |
| `parentAddr` | `string` | 直接上级钱包 |
| `recommendCode` | `string` | 从根用户到直接上级的路径 |
| `amount` | `string` | 当前质押金额 |
| `amountHistory` | `string` | 历史质押金额 |
| `investmentCount` | `string` | 投资次数 |
| `childrenAmount` | `string` | 伞下当前总业绩 |
| `childrenAmountHistory` | `string` | 伞下历史总业绩 |
| `childrenAmountExtra` | `string` | 伞下 Extra 业绩 |
| `rewardRecommendAmount` | `string` | TeamBooked amount 累计 |
| `rewardRecommendPay` | `string` | TeamBooked pay 累计 |
| `rewardRecommendStoreAmount` | `string` | TeamBooked storeAmount 累计 |
| `rewardRecommendFee` | `string` | TeamBooked fee 累计 |
| `rewardRecommendTeamUAmount` | `string` | 团队奖励金额 |
| `rewardRecommendClaimedTeamUNet` | `string` | TeamClaimed 净额累计 |
| `rewardRecommendClaimedTeamUAmount` | `string` | TeamClaimed amount 累计 |
| `rewardRecommendClaimedTeamUFee` | `string` | TeamClaimed fee 累计 |
| `rewardRecommendExpired` | `string` | TeamExpired 金额累计 |
| `lineU` | `string` | USDT 静态收益累计 |
| `lineCoinU` | `string` | MS 支付时的 U 金额累计 |
| `lineCoin` | `string` | MS 数量累计 |
| `lineFee` | `string` | 静态收益手续费累计 |
| `levelReward` | `string` | 当前等级奖励 |
| `createdAt` | `string` | 用户记录创建时间 Unix 秒 |
| `updatedAt` | `string` | 用户记录更新时间 Unix 秒 |
| `name` | `string` | 用户名称，可为空字符串 |

## 5. 每日全局统计列表

### 5.1 请求

```http
GET /api/get_daily_statistics_list
```

### 5.2 Query 参数

| 参数 | 类型 | 必填 | 默认值 | 说明 |
|---|---|---|---|---|
| `page` | `uint64` | 否 | `1` | 页码 |
| `pageSize` | `uint64` | 否 | `20` | 每页数量，最大 100 |
| `startDate` | `string` | 否 | 上海当天 | 开始日期，包含当天 |
| `endDate` | `string` | 否 | 与 `startDate` 相同 | 结束日期，包含当天 |

日期规则：

- `startDate`和 `endDate` 都不传：查询上海当天。
- 只传其中一个：按该日查询。
- 两者都传：日期范围包含开始日和结束日。
- `endDate` 不能早于 `startDate`。
- 一次最多查询 366 个自然日。
- 返回日期按倒序排列。
- 日期范围内没有数据的日期也会返回，所有数值为 `"0"`。

### 5.3 请求示例

查询今天：

```bash
curl "<BASE_URL>/api/get_daily_statistics_list"
```

查询单日：

```bash
curl "<BASE_URL>/api/get_daily_statistics_list?startDate=2026-08-31&endDate=2026-08-31"
```

查询日期范围：

```bash
curl "<BASE_URL>/api/get_daily_statistics_list?startDate=2026-08-01&endDate=2026-08-31&page=1&pageSize=31"
```

### 5.4 响应示例

```json
{
  "total": "1",
  "page": "1",
  "pageSize": "20",
  "list": [
    {
      "date": "2026-08-31",
      "lineClaimedCount": "20074",
      "lineClaimedGrossU": "440492.236784542107422935",
      "lineClaimedFeeU": "57263.990781990473955187",
      "lineClaimedNetU": "383228.246002551633467748",
      "reinvestmentCount": "517",
      "reinvestmentAmount": "862736.480000000000000000",
      "newOrderCount": "622",
      "newOrderAmount": "1113386.480000000000000000"
    }
  ],
  "timezone": "Asia/Shanghai"
}
```

### 5.5 响应字段

| 字段 | JSON 类型 | 说明 |
|---|---|---|
| `total` | `string` | 请求日期范围的自然日总数，不是有数据的日期数 |
| `page` | `string` | 当前页 |
| `pageSize` | `string` | 每页数量 |
| `timezone` | `string` | 固定 `Asia/Shanghai` |
| `list[].date` | `string` | 统计日期 |
| `list[].lineClaimedCount` | `string` | 当日落库的 LineClaimed 记录数 |
| `list[].lineClaimedGrossU` | `string` | LineClaimed `gross_u` 汇总 |
| `list[].lineClaimedFeeU` | `string` | LineClaimed `fee_u` 汇总 |
| `list[].lineClaimedNetU` | `string` | `gross_u - fee_u` 汇总 |
| `list[].reinvestmentCount` | `string` | `is_add=1 AND investment_number>1` 的质押事件数 |
| `list[].reinvestmentAmount` | `string` | 上述复投事件的 `amount` 汇总 |
| `list[].newOrderCount` | `string` | `staking_v1_order_created_event` 实际订单创建数 |
| `list[].newOrderAmount` | `string` | 实际订单创建事件的 `amount` 汇总 |

### 5.6 无数据日期示例

```json
{
  "date": "2026-09-01",
  "lineClaimedCount": "0",
  "lineClaimedGrossU": "0",
  "lineClaimedFeeU": "0",
  "lineClaimedNetU": "0",
  "reinvestmentCount": "0",
  "reinvestmentAmount": "0",
  "newOrderCount": "0",
  "newOrderAmount": "0"
}
```

## 6. TypeScript 类型建议

```ts
export type UInt64String = string;
export type DecimalString = string;

export interface PageReply<T> {
  total: UInt64String;
  page: UInt64String;
  pageSize: UInt64String;
  list: T[];
}

export interface LineClaimedRecord {
  id: UInt64String;
  blockNumber: UInt64String;
  createdAt: UInt64String;
  txHash: string;
  userAddr: string;
  orderId: string;
  grossU: DecimalString;
  feeU: DecimalString;
  netU: DecimalString;
  paidMs: boolean;
  msAmount: DecimalString;
  userId: UInt64String;
  currentOrderCap: DecimalString;
  currentOrderRemaining: DecimalString;
}

export interface LineClaimedListReply extends PageReply<LineClaimedRecord> {
  date: string;
  timezone: "Asia/Shanghai";
}

export interface UserRecord {
  id: UInt64String;
  blockNumber: UInt64String;
  userAddr: string;
  parentAddr: string;
  recommendCode: string;
  amount: DecimalString;
  amountHistory: DecimalString;
  investmentCount: UInt64String;
  childrenAmount: DecimalString;
  childrenAmountHistory: DecimalString;
  childrenAmountExtra: DecimalString;
  rewardRecommendAmount: DecimalString;
  rewardRecommendPay: DecimalString;
  rewardRecommendStoreAmount: DecimalString;
  rewardRecommendFee: DecimalString;
  rewardRecommendTeamUAmount: DecimalString;
  rewardRecommendClaimedTeamUNet: DecimalString;
  rewardRecommendClaimedTeamUAmount: DecimalString;
  rewardRecommendClaimedTeamUFee: DecimalString;
  rewardRecommendExpired: DecimalString;
  lineU: DecimalString;
  lineCoinU: DecimalString;
  lineCoin: DecimalString;
  lineFee: DecimalString;
  levelReward: DecimalString;
  createdAt: UInt64String;
  updatedAt: UInt64String;
  name: string;
}

export type UserAncestorListReply = PageReply<UserRecord>;

export interface DailyStatisticsRecord {
  date: string;
  lineClaimedCount: UInt64String;
  lineClaimedGrossU: DecimalString;
  lineClaimedFeeU: DecimalString;
  lineClaimedNetU: DecimalString;
  reinvestmentCount: UInt64String;
  reinvestmentAmount: DecimalString;
  newOrderCount: UInt64String;
  newOrderAmount: DecimalString;
}

export interface DailyStatisticsListReply extends PageReply<DailyStatisticsRecord> {
  timezone: "Asia/Shanghai";
}
```

## 7. Fetch 请求示例

```ts
const baseURL = import.meta.env.VITE_API_BASE_URL;

const params = new URLSearchParams({
  page: "1",
  pageSize: "20",
  date: "2026-08-31",
});

const response = await fetch(
  `${baseURL}/api/get_line_claimed_list?${params.toString()}`,
  { method: "GET" },
);

if (!response.ok) {
  const error = await response.json();
  throw new Error(error.message || `HTTP ${response.status}`);
}

const data = (await response.json()) as LineClaimedListReply;
```

## 8. 错误响应

当参数错误、用户不存在或推荐链数据不完整时，当前项目的错误响应格式为：

```json
{
  "code": 500,
  "reason": "",
  "message": "具体错误信息",
  "metadata": {}
}
```

前端建议：

1. 先检查 HTTP 状态码。
2. HTTP 非 2xx 时读取 JSON 中的 `message`。
3. 不要只根据 `reason` 判断，因为当前参数错误的 `reason` 可能为空字符串。

常见错误场景：

| 场景 | `message` 示例 |
|---|---|
| 日期格式错误 | `日期 "2026-08-31x" 格式错误，应为 YYYY-MM-DD` |
| 结束日早于开始日 | `endDate 不能早于 startDate` |
| 日期范围过大 | `一次最多查询 366 天` |
| 钱包地址非法 | `address 不是有效的 BSC 地址` |
| 上级查询缺少用户 | `userId 和 address 至少传一个` |
| `userId/address` 不一致 | `userId 和 address 不属于同一个用户` |
| 用户不存在 | `用户 999999 不存在` |
| 推荐路径不完整 | `用户 0x... 的推荐路径不完整，缺少上级 ...` |

## 9. 统计指标数据来源

| 接口字段 | 数据表 | 条件/算法 |
|---|---|---|
| `lineClaimedCount` | `staking_v1_line_claimed_event` | 按 `created_at` 日期计数 |
| `lineClaimedGrossU` | `staking_v1_line_claimed_event` | `SUM(gross_u)` |
| `lineClaimedFeeU` | `staking_v1_line_claimed_event` | `SUM(fee_u)` |
| `lineClaimedNetU` | `staking_v1_line_claimed_event` | `SUM(gross_u - fee_u)` |
| `reinvestmentCount` | `user_v1_stake_changed_event` | `is_add=1 AND investment_number>1` |
| `reinvestmentAmount` | `user_v1_stake_changed_event` | 上述复投记录的 `SUM(amount)` |
| `newOrderCount` | `staking_v1_order_created_event` | 按 `created_at` 日期计数 |
| `newOrderAmount` | `staking_v1_order_created_event` | `SUM(amount)` |

由于新订单和复投来自两个不同的事件表，而本次又明确使用各自的 `created_at` 落库日，在服务停机补数据或两条同步流落库时间跨日时，同一个日历日的新订单和复投数量不一定存在严格的包含关系。

## 10. 上线说明

- 不需要新增数据库字段。
- 不需要新增数据库表。
- 不需要执行数据修复接口。
- 不需要为这 3 个查询接口新增定时任务。
- 金额字段不允许在前端用普通浮点数直接累加。
