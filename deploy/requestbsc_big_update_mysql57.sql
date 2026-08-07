-- requestBsc 用户统计与订单同步增量升级
-- 适用：MySQL 5.7
-- 特性：不删表、不清空数据、不导入本地数据，可重复执行。
-- 要求：导入前必须在宝塔中选中 requestBsc 实际使用的数据库。
-- MySQL 5.7 不支持 ADD COLUMN/INDEX IF NOT EXISTS，下面使用
-- information_schema + PREPARE 避免重复创建已有字段和索引。

SET NAMES utf8mb4 COLLATE utf8mb4_unicode_ci;
SET @requestbsc_schema := DATABASE();

-- --------------------------------------------------------------------------
-- 1. 用户投资次数及用户列表查询索引
-- --------------------------------------------------------------------------

SET @requestbsc_sql := IF(
  EXISTS (
    SELECT 1 FROM information_schema.TABLES
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event'
  ) AND NOT EXISTS (
    SELECT 1 FROM information_schema.COLUMNS
    WHERE TABLE_SCHEMA = @requestbsc_schema
      AND TABLE_NAME = 'user_v1_bound_event'
      AND COLUMN_NAME = 'investment_count'
  ),
  'ALTER TABLE `user_v1_bound_event` ADD COLUMN `investment_count` bigint unsigned NOT NULL DEFAULT 0',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event'
      AND INDEX_NAME = 'idx_user_v1_bound_amount_id'
  ),
  'ALTER TABLE `user_v1_bound_event` ADD INDEX `idx_user_v1_bound_amount_id` (`amount`, `id`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event'
      AND INDEX_NAME = 'idx_user_v1_bound_amount_history_id'
  ),
  'ALTER TABLE `user_v1_bound_event` ADD INDEX `idx_user_v1_bound_amount_history_id` (`amount_history`, `id`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event'
      AND INDEX_NAME = 'idx_user_v1_bound_children_amount_id'
  ),
  'ALTER TABLE `user_v1_bound_event` ADD INDEX `idx_user_v1_bound_children_amount_id` (`children_amount`, `id`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event'
      AND INDEX_NAME = 'idx_user_v1_bound_investment_count'
  ),
  'ALTER TABLE `user_v1_bound_event` ADD INDEX `idx_user_v1_bound_investment_count` (`investment_count`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

-- recommend_code 最长 4096，MySQL 5.7 utf8mb4 使用 191 字符前缀兼容旧版 InnoDB 索引长度。
SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_bound_event'
      AND INDEX_NAME = 'idx_user_v1_bound_recommend_code'
  ),
  'ALTER TABLE `user_v1_bound_event` ADD INDEX `idx_user_v1_bound_recommend_code` (`recommend_code`(191))',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

-- --------------------------------------------------------------------------
-- 2. 质押（投资）事件的区块时间、第几次投资及查询索引
-- --------------------------------------------------------------------------

SET @requestbsc_sql := IF(
  EXISTS (
    SELECT 1 FROM information_schema.TABLES
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event'
  ) AND NOT EXISTS (
    SELECT 1 FROM information_schema.COLUMNS
    WHERE TABLE_SCHEMA = @requestbsc_schema
      AND TABLE_NAME = 'user_v1_stake_changed_event'
      AND COLUMN_NAME = 'block_time'
  ),
  'ALTER TABLE `user_v1_stake_changed_event` ADD COLUMN `block_time` bigint unsigned NOT NULL DEFAULT 0',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (
    SELECT 1 FROM information_schema.TABLES
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event'
  ) AND NOT EXISTS (
    SELECT 1 FROM information_schema.COLUMNS
    WHERE TABLE_SCHEMA = @requestbsc_schema
      AND TABLE_NAME = 'user_v1_stake_changed_event'
      AND COLUMN_NAME = 'investment_number'
  ),
  'ALTER TABLE `user_v1_stake_changed_event` ADD COLUMN `investment_number` bigint unsigned NOT NULL DEFAULT 0',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event'
      AND INDEX_NAME = 'idx_user_v1_stake_add_time_block'
  ),
  'ALTER TABLE `user_v1_stake_changed_event` ADD INDEX `idx_user_v1_stake_add_time_block` (`is_add`, `block_time`, `block_number`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event'
      AND INDEX_NAME = 'idx_user_v1_stake_add_number'
  ),
  'ALTER TABLE `user_v1_stake_changed_event` ADD INDEX `idx_user_v1_stake_add_number` (`is_add`, `investment_number`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

SET @requestbsc_sql := IF(
  EXISTS (SELECT 1 FROM information_schema.TABLES WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event')
  AND NOT EXISTS (
    SELECT 1 FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = @requestbsc_schema AND TABLE_NAME = 'user_v1_stake_changed_event'
      AND INDEX_NAME = 'idx_user_v1_stake_add_user_order'
  ),
  'ALTER TABLE `user_v1_stake_changed_event` ADD INDEX `idx_user_v1_stake_add_user_order` (`is_add`, `user_addr`, `block_number`, `id`)',
  'DO 0'
);
PREPARE requestbsc_stmt FROM @requestbsc_sql;
EXECUTE requestbsc_stmt;
DEALLOCATE PREPARE requestbsc_stmt;

-- --------------------------------------------------------------------------
-- 3. 订单主表：数据库永久保留订单，status=3 表示已结束，不删除记录
-- --------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `staking_v1_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_order_index` decimal(65,0) NOT NULL DEFAULT 0,
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `base_cap` decimal(65,18) NOT NULL DEFAULT 0,
  `cap` decimal(65,18) NOT NULL DEFAULT 0,
  `used` decimal(65,18) NOT NULL DEFAULT 0,
  `remaining` decimal(65,18) NOT NULL DEFAULT 0,
  `compensation` decimal(65,18) NOT NULL DEFAULT 0,
  `line_paid` decimal(65,18) NOT NULL DEFAULT 0,
  `line_claimable` decimal(65,18) NOT NULL DEFAULT 0,
  `plan_id` decimal(65,0) NOT NULL DEFAULT 0,
  `created_time` bigint unsigned NOT NULL DEFAULT 0,
  `start_time` bigint unsigned NOT NULL DEFAULT 0,
  `claim_effective` bigint unsigned NOT NULL DEFAULT 0,
  `days_count` int unsigned NOT NULL DEFAULT 0,
  `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '1=queued,2=running,3=exited',
  `queue_index` decimal(65,0) NOT NULL DEFAULT 0,
  `queue_liq_u` decimal(65,18) NOT NULL DEFAULT 0,
  `queued_at` bigint unsigned NOT NULL DEFAULT 0,
  `queue_done` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `created_block` bigint unsigned NOT NULL DEFAULT 0,
  `entered_block` bigint unsigned NOT NULL DEFAULT 0,
  `exited_block` bigint unsigned NOT NULL DEFAULT 0,
  `last_synced_block` bigint unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_order_id` (`order_id`),
  KEY `idx_staking_v1_order_user_id_order` (`user_id`, `order_id`),
  KEY `idx_staking_v1_order_user_addr_order` (`user_addr`, `order_id`),
  KEY `idx_staking_v1_order_status_order` (`status`, `order_id`),
  KEY `idx_staking_v1_order_snapshot` (`status`, `last_synced_block`, `user_id`),
  KEY `idx_staking_v1_order_last_synced` (`last_synced_block`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------------------------
-- 4. 订单生命周期事件表
-- --------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `staking_v1_order_created_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `cap` decimal(65,18) NOT NULL DEFAULT 0,
  `plan_id` decimal(65,0) NOT NULL DEFAULT 0,
  `days_count` int unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_created_event_key` (`event_key`),
  KEY `idx_staking_v1_order_created_block` (`block_number`),
  KEY `idx_staking_v1_order_created_order` (`order_id`),
  KEY `idx_staking_v1_order_created_user` (`user_id`, `order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `staking_v1_order_entered_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `start_time` bigint unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_entered_event_key` (`event_key`),
  KEY `idx_staking_v1_order_entered_block` (`block_number`),
  KEY `idx_staking_v1_order_entered_order` (`order_id`),
  KEY `idx_staking_v1_order_entered_user` (`user_id`, `order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `staking_v1_order_exited_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `used` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_exited_event_key` (`event_key`),
  KEY `idx_staking_v1_order_exited_block` (`block_number`),
  KEY `idx_staking_v1_order_exited_order` (`order_id`),
  KEY `idx_staking_v1_order_exited_user` (`user_id`, `order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `staking_v1_order_cap_set_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `user_order_index` decimal(65,0) NOT NULL DEFAULT 0,
  `old_cap` decimal(65,18) NOT NULL DEFAULT 0,
  `new_cap` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_cap_set_event_key` (`event_key`),
  KEY `idx_staking_v1_order_cap_set_block` (`block_number`),
  KEY `idx_staking_v1_order_cap_set_order` (`order_id`),
  KEY `idx_staking_v1_order_cap_set_user` (`user_id`, `order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `staking_v1_order_queued_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `queue_index` decimal(65,0) NOT NULL DEFAULT 0,
  `queue_liq_u` decimal(65,18) NOT NULL DEFAULT 0,
  `queued_at` bigint unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_queued_event_key` (`event_key`),
  KEY `idx_staking_v1_order_queued_block` (`block_number`),
  KEY `idx_staking_v1_order_queued_order` (`order_id`),
  KEY `idx_staking_v1_order_queued_user` (`user_id`, `order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `staking_v1_order_queue_done_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `user_addr` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `queue_index` decimal(65,0) NOT NULL DEFAULT 0,
  `queue_liq_u` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_order_queue_done_event_key` (`event_key`),
  KEY `idx_staking_v1_order_queue_done_block` (`block_number`),
  KEY `idx_staking_v1_order_queue_done_order` (`order_id`),
  KEY `idx_staking_v1_order_queue_done_user` (`user_id`, `order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `staking_v1_plan_set_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `event_key` varchar(96) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `plan_id` decimal(65,0) NOT NULL DEFAULT 0,
  `min_amount` decimal(65,18) NOT NULL DEFAULT 0,
  `max_amount` decimal(65,18) NOT NULL DEFAULT 0,
  `out_amount` decimal(65,18) NOT NULL DEFAULT 0,
  `days_count` int unsigned NOT NULL DEFAULT 0,
  `enabled` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_plan_set_event_key` (`event_key`),
  KEY `idx_staking_v1_plan_set_block` (`block_number`),
  KEY `idx_staking_v1_plan_set_plan_id` (`plan_id`, `id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------------------------
-- 5. 同步断点表；stream_name 可直接保存 staking_order
-- 迁移不预写断点，避免跳过历史事件，由恢复接口初始化。
-- --------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `user_v1_performance_sync_progress` (
  `stream_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_processed_block` bigint unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`stream_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 只清理本文使用的会话变量，不会改动业务数据。
SET @requestbsc_sql := NULL;
SET @requestbsc_schema := NULL;
