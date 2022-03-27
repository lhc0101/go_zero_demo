CREATE TABLE `user` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
    `open_id` varchar(255)  NOT NULL DEFAULT '' COMMENT '小程序ID',
    `nick_name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户昵称',
	`gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
	`mobile` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
	`password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `idx_mobile_unique` (`mobile`),
    UNIQUE KEY `idx_open_id_unique` (`open_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
