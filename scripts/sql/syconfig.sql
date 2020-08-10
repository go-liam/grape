

-- config

DROP TABLE IF EXISTS `cf_configs`;
CREATE TABLE `cf_configs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  `content` varchar(8000) DEFAULT NULL COMMENT '内容',
  `extended` varchar(2000) DEFAULT NULL COMMENT '扩展的',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
