

-- log

DROP TABLE IF EXISTS `lg_log`;
CREATE TABLE `lg_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `level` int(11) unsigned DEFAULT '0' COMMENT '级别',
  `extended` varchar(1000) DEFAULT NULL COMMENT '扩展的',
  `msg` varchar(8000) DEFAULT NULL COMMENT '信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

