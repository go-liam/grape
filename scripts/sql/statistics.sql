
-- 统计系统

DROP TABLE IF EXISTS `sa_server`;
CREATE TABLE `sa_server` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `extended` varchar(2000) DEFAULT NULL COMMENT '扩展的',
  `url` text DEFAULT NULL COMMENT '地址',
  `run_status` text DEFAULT NULL COMMENT '健康状态',
  `health` varchar(5000) DEFAULT NULL COMMENT '健康信息',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

