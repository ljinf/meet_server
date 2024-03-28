DROP TABLE IF EXISTS `im_register`;
CREATE TABLE `im_register`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL COMMENT '用户id',
    `phone`      varchar(11) DEFAULT NULL COMMENT '手机号',
    `email`      varchar(64) DEFAULT NULL COMMENT '邮箱',
    `password`   varchar(64) DEFAULT NULL COMMENT '密码',
    `created_at` datetime    DEFAULT NULL,
    `updated_at` datetime    DEFAULT NULL,
    `deleted_at` datetime    DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT '注册信息表';


DROP TABLE IF EXISTS `im_user_info`;
CREATE TABLE `im_user_info`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`           int(11) NOT NULL COMMENT '用户id',
    `nick_name`         varchar(100) DEFAULT NULL COMMENT '昵称',
    `avatar`            varchar(255) DEFAULT NULL COMMENT '头像',
    `gender`            int(10) DEFAULT NULL COMMENT '性别',
    `birth_day`         varchar(50)  DEFAULT NULL COMMENT '生日',
    `self_signature`    varchar(255) DEFAULT NULL COMMENT '个性签名',
    `friend_allow_type` int(10) NOT NULL DEFAULT '1' COMMENT '加好友验证类型（Friend_AllowType） 1无需验证 2需要验证',
    `silent_flag`       int(10) NOT NULL DEFAULT '0' COMMENT '禁言标识 1禁言',
    `user_type`         int(10) NOT NULL DEFAULT '1' COMMENT '用户类型 1普通用户 2客服 3机器人',
    `del_flag`          int(20) NOT NULL DEFAULT '0' COMMENT '删除标识',
    `created_at`        datetime     DEFAULT NULL,
    `updated_at`        datetime     DEFAULT NULL,
    `deleted_at`        datetime     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT '用户信息表';

