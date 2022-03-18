CREATE TABLE `users`
(
    `unique_id`     bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `passport_uid`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户账号',
    `passport_code` string    NOT NULL DEFAULT '' COMMENT '用户密码',
    `create_time`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
    `update_time`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_time',
    PRIMARY KEY (`unique_id`),
    UNIQUE KEY `uniq_passport_uid` (`passport_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户账号信息表格';