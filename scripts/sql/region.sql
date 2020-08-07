

-- 用户区域管理

DROP TABLE IF EXISTS `ar_user_region`;
CREATE TABLE `ar_user_region` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `region_id` int(11) unsigned DEFAULT '0' COMMENT '区域ID',

  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

