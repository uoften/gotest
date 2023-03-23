ALTER TABLE `spacex`.`reg_client_persons`
    ADD COLUMN `assign_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '分配时间' AFTER `person_id`;

ALTER TABLE `spacex`.`reg_persons`
    MODIFY COLUMN `fb_account_id` varchar(20) NOT NULL DEFAULT 0 AFTER `ol_cookies`;


CREATE TABLE `reg_auth_task` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `person_id` int(11) NOT NULL DEFAULT '0' COMMENT '账号id',
    `auth_email` varchar(128) NOT NULL DEFAULT '' COMMENT '授权账号',
    `auth_url` varchar(512) NOT NULL DEFAULT '' COMMENT '操作地址',
    `is_handle` int(1) DEFAULT '0' COMMENT '是否已经处理 1是 0 否',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_auth_email` (`auth_email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `reg_persons_tags` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `company_id` int(10) DEFAULT '0' COMMENT '公司id',
    `client_id` int(10) DEFAULT '0' COMMENT '用户id',
    `person_id` int(10) DEFAULT '0' COMMENT '账号id',
    `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户标签',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_person_client` (`company_id`,`client_id`,`person_id`,`tag`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;