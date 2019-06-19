
ALTER TABLE `phpcasts`.`sections`
CHANGE COLUMN `order` `weight` tinyint(2) NOT NULL DEFAULT 0 COMMENT '排序值' AFTER `name`

CREATE TABLE `comments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) NOT NULL COMMENT '评论所属类型 1:视频 2:wiki',
  `related_id` int(11) NOT NULL DEFAULT '0' COMMENT '关联id',
  `origin_content` text NOT NULL COMMENT '原始markdown评论内容',
  `content` text NOT NULL COMMENT '评论内容',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT '评论者所在地ip',
  `like_count` int(11) NOT NULL default 0 '喜欢数',
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '评论者id',
  `device_type` varchar(255) NOT NULL DEFAULT '' COMMENT '客户端设备类型',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4;

