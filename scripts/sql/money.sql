

-- money

DROP TABLE IF EXISTS `um_change`;
CREATE TABLE `um_change` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `remark` varchar(250) DEFAULT NULL COMMENT '备注',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT 'uid',
  `amount` int(11) unsigned DEFAULT '0' COMMENT '数额',
  `last_amount` int(11) unsigned DEFAULT '0' COMMENT '上次个人账户',
  `type` tinyint(1) unsigned DEFAULT '0' COMMENT '支付类型 1**｜充值类型 **',
  `order_id` bigint(20) unsigned DEFAULT '0' COMMENT '订单ID',
  `money_type` tinyint(1) unsigned DEFAULT '0' COMMENT '1收入;2消费',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `um_pay`;
CREATE TABLE `um_pay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `remark` varchar(250) DEFAULT NULL COMMENT '备注',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT 'uid',
  `amount` int(11) unsigned DEFAULT '0' COMMENT '数额',
  `last_amount` int(11) unsigned DEFAULT '0' COMMENT '上次个人账户',
  `type` tinyint(1) unsigned DEFAULT '0' COMMENT '支付类型 1**｜充值类型 **',
  `order_id` bigint(20) unsigned DEFAULT '0' COMMENT '订单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `um_recharge`;
CREATE TABLE `um_recharge` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `remark` varchar(250) DEFAULT NULL COMMENT '备注',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT 'uid',
  `amount` int(11) unsigned DEFAULT '0' COMMENT '数额',
  `last_amount` int(11) unsigned DEFAULT '0' COMMENT '上次个人账户',
  `type` tinyint(1) unsigned DEFAULT '0' COMMENT '支付类型 1**｜充值类型 **',
  `order_id` bigint(20) unsigned DEFAULT '0' COMMENT '订单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

