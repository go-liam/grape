
--


DROP TABLE IF EXISTS `uc_user_bind`;
CREATE TABLE `uc_user_bind` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT 'uid',
  `type` tinyint(1) unsigned DEFAULT '0' COMMENT '类型',
  `name` varchar(100) DEFAULT NULL COMMENT '绑定名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `uc_user`;
CREATE TABLE `uc_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `remark` varchar(250) DEFAULT NULL COMMENT '备注',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机',
  `name` varchar(50) DEFAULT NULL COMMENT '账号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

