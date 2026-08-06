-- 仅用于升级曾经包含 block_time、log_index、check_status、check_time 的旧表。
-- 新环境请直接执行 20260802_user_v1_bound_event.sql。

ALTER TABLE `user_v1_bound_event`
  ADD COLUMN `recommend_code` varchar(4096) NOT NULL DEFAULT '' AFTER `parent_addr`,
  DROP INDEX `uk_user_v1_bound_block_log`,
  ADD UNIQUE KEY `uk_user_v1_bound_user` (`user_addr`),
  ADD KEY `idx_user_v1_bound_block` (`block_number`),
  ADD KEY `idx_user_v1_bound_parent` (`parent_addr`),
  DROP COLUMN `block_time`,
  DROP COLUMN `log_index`,
  DROP COLUMN `check_status`,
  DROP COLUMN `check_time`;
