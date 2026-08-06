-- 用户业绩、团队奖励和线性奖励事件表。
-- line_u 与 line_coin_u 累加的都是 LineClaimed.grossU；实际净值需用 grossU - feeU 计算。

ALTER TABLE `user_v1_bound_event`
  ADD COLUMN `amount` decimal(65,18) NOT NULL DEFAULT 0 AFTER `recommend_code`,
  ADD COLUMN `amount_history` decimal(65,18) NOT NULL DEFAULT 0 AFTER `amount`,
  ADD COLUMN `children_amount` decimal(65,18) NOT NULL DEFAULT 0 AFTER `amount_history`,
  ADD COLUMN `children_amount_history` decimal(65,18) NOT NULL DEFAULT 0 AFTER `children_amount`,
  ADD COLUMN `children_amount_extra` decimal(65,18) NOT NULL DEFAULT 0 AFTER `children_amount_history`,
  ADD COLUMN `reward_recommend_amount` decimal(65,18) NOT NULL DEFAULT 0 AFTER `children_amount_extra`,
  ADD COLUMN `reward_recommend_pay` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_amount`,
  ADD COLUMN `reward_recommend_store_amount` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_pay`,
  ADD COLUMN `reward_recommend_fee` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_store_amount`,
  ADD COLUMN `reward_recommend_team_u_amount` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_fee`,
  ADD COLUMN `reward_recommend_claimed_team_u_net` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_team_u_amount`,
  ADD COLUMN `reward_recommend_claimed_team_u_amount` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_claimed_team_u_net`,
  ADD COLUMN `reward_recommend_claimed_team_u_fee` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_claimed_team_u_amount`,
  ADD COLUMN `reward_recommend_expired` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_claimed_team_u_fee`,
  ADD COLUMN `line_u` decimal(65,18) NOT NULL DEFAULT 0 AFTER `reward_recommend_expired`,
  ADD COLUMN `line_coin_u` decimal(65,18) NOT NULL DEFAULT 0 AFTER `line_u`,
  ADD COLUMN `line_coin` decimal(65,18) NOT NULL DEFAULT 0 AFTER `line_coin_u`,
  ADD COLUMN `line_fee` decimal(65,18) NOT NULL DEFAULT 0 AFTER `line_coin`,
  ADD COLUMN `level_reward` decimal(65,18) NOT NULL DEFAULT 0 AFTER `line_fee`;

CREATE TABLE IF NOT EXISTS `user_v1_stake_changed_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_key` varchar(96) NOT NULL,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `tx_hash` varchar(66) NOT NULL DEFAULT '',
  `user_addr` varchar(42) NOT NULL DEFAULT '',
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `is_add` tinyint unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_v1_stake_changed_event_key` (`event_key`),
  KEY `idx_user_v1_stake_changed_block` (`block_number`),
  KEY `idx_user_v1_stake_changed_user` (`user_addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_v1_extra_changed_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_key` varchar(96) NOT NULL,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `tx_hash` varchar(66) NOT NULL DEFAULT '',
  `user_addr` varchar(42) NOT NULL DEFAULT '',
  `extra_amount` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_v1_extra_changed_event_key` (`event_key`),
  KEY `idx_user_v1_extra_changed_block` (`block_number`),
  KEY `idx_user_v1_extra_changed_user` (`user_addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `staking_v1_team_booked_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_key` varchar(96) NOT NULL,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `tx_hash` varchar(66) NOT NULL DEFAULT '',
  `from_addr` varchar(42) NOT NULL DEFAULT '',
  `to_addr` varchar(42) NOT NULL DEFAULT '',
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `store_amount` decimal(65,18) NOT NULL DEFAULT 0,
  `pay` decimal(65,18) NOT NULL DEFAULT 0,
  `fee` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_team_booked_event_key` (`event_key`),
  KEY `idx_staking_v1_team_booked_block` (`block_number`),
  KEY `idx_staking_v1_team_booked_from` (`from_addr`),
  KEY `idx_staking_v1_team_booked_to` (`to_addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `staking_v1_team_claimed_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_key` varchar(96) NOT NULL,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `tx_hash` varchar(66) NOT NULL DEFAULT '',
  `user_addr` varchar(42) NOT NULL DEFAULT '',
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `fee` decimal(65,18) NOT NULL DEFAULT 0,
  `net` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_team_claimed_event_key` (`event_key`),
  KEY `idx_staking_v1_team_claimed_block` (`block_number`),
  KEY `idx_staking_v1_team_claimed_user` (`user_addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `staking_v1_team_expired_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_key` varchar(96) NOT NULL,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `tx_hash` varchar(66) NOT NULL DEFAULT '',
  `from_addr` varchar(42) NOT NULL DEFAULT '',
  `to_addr` varchar(42) NOT NULL DEFAULT '',
  `amount` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_team_expired_event_key` (`event_key`),
  KEY `idx_staking_v1_team_expired_block` (`block_number`),
  KEY `idx_staking_v1_team_expired_from` (`from_addr`),
  KEY `idx_staking_v1_team_expired_to` (`to_addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `staking_v1_line_claimed_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_key` varchar(96) NOT NULL,
  `block_number` bigint unsigned NOT NULL DEFAULT 0,
  `tx_hash` varchar(66) NOT NULL DEFAULT '',
  `user_addr` varchar(42) NOT NULL DEFAULT '',
  `order_id` decimal(65,0) NOT NULL DEFAULT 0,
  `gross_u` decimal(65,18) NOT NULL DEFAULT 0,
  `fee_u` decimal(65,18) NOT NULL DEFAULT 0,
  `paid_ms` tinyint unsigned NOT NULL DEFAULT 0,
  `ms_amount` decimal(65,18) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_staking_v1_line_claimed_event_key` (`event_key`),
  KEY `idx_staking_v1_line_claimed_block` (`block_number`),
  KEY `idx_staking_v1_line_claimed_user` (`user_addr`),
  KEY `idx_staking_v1_line_claimed_order` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_v1_performance_sync_progress` (
  `stream_name` varchar(32) NOT NULL,
  `last_processed_block` bigint unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`stream_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
