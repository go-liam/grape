
-- home site :

DROP TABLE IF EXISTS `ws_menu`;
CREATE TABLE `ws_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `language_id` int(11) unsigned DEFAULT '0' COMMENT '语言ID',
  `site_id` int(11) unsigned DEFAULT '0' COMMENT '网站ID',
  `name` varchar(100) DEFAULT NULL COMMENT '菜单名',
  `parent_id` int(11) unsigned DEFAULT '0' COMMENT '父亲ID',
  `level` int(11) unsigned DEFAULT '0' COMMENT '级别',
  `extended` varchar(8000) DEFAULT NULL COMMENT '扩展的',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `ws_notice`;
CREATE TABLE `ws_notice` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `language_id` int(11) unsigned DEFAULT '0' COMMENT '语言ID',
  `site_id` int(11) unsigned DEFAULT '0' COMMENT '网站ID',
  `title` varchar(250) DEFAULT NULL COMMENT '标题',
  `author` varchar(100) DEFAULT NULL COMMENT '作者',
  `extended` varchar(5000) DEFAULT NULL COMMENT '扩展的',
  `content` text DEFAULT NULL COMMENT '内容',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `ws_page`;
CREATE TABLE `ws_page` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `language_id` int(11) unsigned DEFAULT '0' COMMENT '语言ID',
  `site_id` int(11) unsigned DEFAULT '0' COMMENT '网站ID',
  `menu_id` int(11) unsigned DEFAULT '0' COMMENT '菜单ID',
  `title` varchar(100) DEFAULT NULL COMMENT '标题',
  `author` varchar(100) DEFAULT NULL COMMENT '作者',
  `description` varchar(100) DEFAULT NULL COMMENT '描述',
  `extended` varchar(5000) DEFAULT NULL COMMENT '扩展的',
  `content` text DEFAULT NULL COMMENT '内容',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `ws_site`;
CREATE TABLE `ws_site` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '修改时间戳',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态：44为删除',
  `language_id` int(11) unsigned DEFAULT '0' COMMENT '语言ID',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `email` varchar(200) DEFAULT NULL COMMENT '邮箱',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `extended` text DEFAULT NULL COMMENT '扩展的',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

