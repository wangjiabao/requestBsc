-- ExtraChanged can be emitted for an address that is not registered yet.
-- 1=applied to user_v1_bound_event, 2=unregistered address and not applied.

ALTER TABLE `user_v1_extra_changed_event`
  ADD COLUMN `apply_status` tinyint unsigned NOT NULL DEFAULT 1
    COMMENT '1=applied,2=unregistered_pending'
    AFTER `extra_amount`,
  ADD KEY `idx_user_v1_extra_apply_block` (`apply_status`, `block_number`, `id`);
