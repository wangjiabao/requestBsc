ALTER TABLE `user_v1_bound_event`
  ADD COLUMN `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `user_addr`;
