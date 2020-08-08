

--

DROP TABLE IF EXISTS `rb_power`;
CREATE TABLE `rb_power` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `tag` varchar(100) DEFAULT NULL COMMENT '权限标签',
  `type` tinyint(1) unsigned DEFAULT '0' COMMENT '类型',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `rb_role`;
CREATE TABLE `rb_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `power_ids` varchar(5000) DEFAULT NULL COMMENT '权限',
  `extended` varchar(5000) DEFAULT NULL COMMENT '扩展的',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `rb_user`;
CREATE TABLE `rb_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `flag` tinyint(1) unsigned DEFAULT '0' COMMENT '1超级管理员',
  `remark` varchar(100) DEFAULT NULL COMMENT '备注',
  `role_ids` varchar(8000) DEFAULT NULL COMMENT '角色IDs',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
