CREATE TABLE IF NOT EXISTS `user_v1_bound_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `block_number` bigint unsigned NOT NULL DEFAULT '0',
  `user_addr` varchar(42) NOT NULL DEFAULT '',
  `parent_addr` varchar(42) NOT NULL DEFAULT '',
  `recommend_code` varchar(4096) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_v1_bound_user` (`user_addr`),
  KEY `idx_user_v1_bound_block` (`block_number`),
  KEY `idx_user_v1_bound_parent` (`parent_addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_v1_bound_sync_progress` (
  `id` tinyint unsigned NOT NULL,
  `last_processed_block` bigint unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
