# ************************************************************
# Sequel Pro SQL dump
# Version 4004
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 59.110.46.243 (MySQL 5.6.35-log)
# Database: phpcasts
# Generation Time: 2021-02-06 12:57:32 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table admins
# ------------------------------------------------------------

DROP TABLE IF EXISTS `admins`;

CREATE TABLE `admins` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `name` varchar(255) NOT NULL,
                          `email` varchar(255) NOT NULL,
                          `password` varchar(60) NOT NULL,
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          `remember_token` varchar(100) DEFAULT NULL,
                          `created_at` timestamp NULL DEFAULT NULL,
                          `updated_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `admins_name_unique` (`name`),
                          UNIQUE KEY `admins_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table comments
# ------------------------------------------------------------

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
                            `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                            `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '评论所属类型 1:视频 2:wiki',
                            `related_id` int(11) NOT NULL DEFAULT '0' COMMENT '关联id',
                            `ip` varchar(15) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '评论者所在地ip',
                            `content` text CHARACTER SET utf8 NOT NULL COMMENT 'markdown评论内容',
                            `origin_content` text CHARACTER SET utf8 NOT NULL COMMENT '原始评论内容',
                            `like_count` int(11) NOT NULL,
                            `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '评论者id',
                            `device_type` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '客户端设备类型',
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table courses
# ------------------------------------------------------------

DROP TABLE IF EXISTS `courses`;

CREATE TABLE `courses` (
                           `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                           `name` varchar(255) NOT NULL COMMENT '课程名称',
                           `type` enum('backend','frontend','service','tool') NOT NULL COMMENT '课程分类',
                           `keywords` varchar(100) NOT NULL DEFAULT '' COMMENT '页面关键词',
                           `description` text NOT NULL COMMENT '页面描述',
                           `content` varchar(500) NOT NULL DEFAULT '' COMMENT '课程描述',
                           `slug` varchar(255) NOT NULL COMMENT 'slug',
                           `cover_key` varchar(255) NOT NULL DEFAULT '' COMMENT '课程封面图',
                           `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者id',
                           `is_publish` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否发布 0:否 1:是',
                           `update_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '课程的更新状态 0:初始化, 1:预告, 2:更新中, 3:已完结',
                           `created_at` timestamp NULL DEFAULT NULL,
                           `updated_at` timestamp NULL DEFAULT NULL,
                           `deleted_at` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `courses_slug_unique` (`slug`),
                           KEY `courses_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table feedbacks
# ------------------------------------------------------------

DROP TABLE IF EXISTS `feedbacks`;

CREATE TABLE `feedbacks` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `content` text NOT NULL COMMENT '反馈的内容',
                             `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
                             `created_at` timestamp NULL DEFAULT NULL,
                             `updated_at` timestamp NULL DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table followables
# ------------------------------------------------------------

DROP TABLE IF EXISTS `followables`;

CREATE TABLE `followables` (
                               `user_id` int(10) unsigned NOT NULL,
                               `followable_id` int(10) unsigned NOT NULL,
                               `followable_type` varchar(255) NOT NULL,
                               `relation` varchar(255) NOT NULL DEFAULT 'follow' COMMENT 'folllow/like/subscribe/favorite/',
                               `created_at` timestamp NOT NULL,
                               KEY `followables_user_id_foreign` (`user_id`),
                               KEY `followables_followable_type_index` (`followable_type`),
                               CONSTRAINT `followables_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table followers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `followers`;

CREATE TABLE `followers` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `user_id` int(10) unsigned NOT NULL,
                             `follow_id` int(10) unsigned NOT NULL,
                             `created_at` timestamp NULL DEFAULT NULL,
                             `updated_at` timestamp NULL DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table forum_categories
# ------------------------------------------------------------

DROP TABLE IF EXISTS `forum_categories`;

CREATE TABLE `forum_categories` (
                                    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                    `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
                                    `name` varchar(255) NOT NULL COMMENT '分类名',
                                    `slug` varchar(60) NOT NULL COMMENT '缩略名',
                                    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
                                    `weight` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '权重',
                                    `topic_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '帖子数',
                                    `user_id` int(10) unsigned NOT NULL DEFAULT '0',
                                    `created_at` timestamp NULL DEFAULT NULL,
                                    `updated_at` timestamp NULL DEFAULT NULL,
                                    `deleted_at` timestamp NULL DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `forum_categories_slug_unique` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table forum_replies
# ------------------------------------------------------------

DROP TABLE IF EXISTS `forum_replies`;

CREATE TABLE `forum_replies` (
                                 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                 `topic_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '帖子id',
                                 `origin_body` text CHARACTER SET utf8 NOT NULL,
                                 `body` text CHARACTER SET utf8 NOT NULL,
                                 `like_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '投票数',
                                 `is_blocked` enum('yes','no') CHARACTER SET utf8 NOT NULL DEFAULT 'no' COMMENT '是否block帖子',
                                 `source` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '访问来源 iOS，Android, PC',
                                 `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者id',
                                 `created_at` timestamp NULL DEFAULT NULL,
                                 `updated_at` timestamp NULL DEFAULT NULL,
                                 `deleted_at` timestamp NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `forum_replies_topic_id_index` (`topic_id`),
                                 KEY `forum_replies_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table forum_topics
# ------------------------------------------------------------

DROP TABLE IF EXISTS `forum_topics`;

CREATE TABLE `forum_topics` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分类id',
                                `title` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '标题',
                                `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '帖子类型 1:markdown 2:html',
                                `origin_body` mediumtext CHARACTER SET utf8 NOT NULL COMMENT '帖子内容',
                                `body` mediumtext CHARACTER SET utf8 NOT NULL COMMENT '帖子内容',
                                `view_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
                                `reply_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '回复数',
                                `vote_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '投票数',
                                `is_excellent` enum('yes','no') CHARACTER SET utf8 NOT NULL DEFAULT 'no' COMMENT '是否加精帖子',
                                `is_blocked` enum('yes','no') CHARACTER SET utf8 NOT NULL DEFAULT 'no' COMMENT '是否block帖子',
                                `last_reply_user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后回复用户',
                                `last_reply_time_at` datetime NOT NULL COMMENT '最后回复时间',
                                `source` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '访问来源 iOS，Android, PC',
                                `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者id',
                                `created_at` timestamp NULL DEFAULT NULL,
                                `updated_at` timestamp NULL DEFAULT NULL,
                                `deleted_at` timestamp NULL DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                KEY `forum_topics_last_reply_user_id_index` (`last_reply_user_id`),
                                KEY `forum_topics_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table images
# ------------------------------------------------------------

DROP TABLE IF EXISTS `images`;

CREATE TABLE `images` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `image_name` varchar(255) NOT NULL COMMENT '图片名，经过str_random函数处理',
                          `image_path` varchar(255) NOT NULL DEFAULT '' COMMENT '图片存储的实际路径',
                          `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
                          `created_at` timestamp NULL DEFAULT NULL,
                          `updated_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table migrations
# ------------------------------------------------------------

DROP TABLE IF EXISTS `migrations`;

CREATE TABLE `migrations` (
                              `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                              `migration` varchar(255) NOT NULL,
                              `batch` int(11) NOT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table notifications
# ------------------------------------------------------------

DROP TABLE IF EXISTS `notifications`;

CREATE TABLE `notifications` (
                                 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                 `from_user_id` int(11) NOT NULL,
                                 `user_id` int(11) NOT NULL,
                                 `topic_id` int(11) NOT NULL,
                                 `video_id` int(11) NOT NULL,
                                 `post_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章id',
                                 `reply_id` int(11) DEFAULT NULL,
                                 `body` text,
                                 `type` varchar(255) NOT NULL COMMENT '类型: new_reply,at,follow,new_video_reply,video_at',
                                 `created_at` timestamp NULL DEFAULT NULL,
                                 `updated_at` timestamp NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `notifications_from_user_id_index` (`from_user_id`),
                                 KEY `notifications_user_id_index` (`user_id`),
                                 KEY `notifications_topic_id_index` (`topic_id`),
                                 KEY `notifications_video_id_index` (`video_id`),
                                 KEY `notifications_reply_id_index` (`reply_id`),
                                 KEY `notifications_type_index` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table order_items
# ------------------------------------------------------------

DROP TABLE IF EXISTS `order_items`;

CREATE TABLE `order_items` (
                               `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                               `order_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '订单id',
                               `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
                               `item_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品id',
                               `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
                               `price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品单价',
                               `quantity` int(11) NOT NULL DEFAULT '0' COMMENT '购买数量',
                               `amount` decimal(8,2) NOT NULL COMMENT '商品总金额',
                               `deleted_at` timestamp NULL DEFAULT NULL,
                               `created_at` timestamp NULL DEFAULT NULL,
                               `updated_at` timestamp NULL DEFAULT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table orders
# ------------------------------------------------------------

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
                          `id` bigint(20) unsigned NOT NULL,
                          `order_amount` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '订单总金额',
                          `pay_amount` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '应付总金额',
                          `pay_method` enum('alipay','wechat','youzan') NOT NULL DEFAULT 'youzan' COMMENT '支付方式',
                          `paid_at` timestamp NULL DEFAULT NULL COMMENT '支付时间',
                          `completed_at` timestamp NULL DEFAULT NULL COMMENT '完成时间',
                          `canceled_at` timestamp NULL DEFAULT NULL,
                          `qrcode_id` int(11) NOT NULL DEFAULT '0' COMMENT '二维码id',
                          `trade_id` varchar(32) NOT NULL DEFAULT '' COMMENT '交易id',
                          `status` enum('pending','paid','canceled','completed') NOT NULL DEFAULT 'pending' COMMENT '订单状态',
                          `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '买家uid',
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          `created_at` timestamp NULL DEFAULT NULL,
                          `updated_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table password_resets
# ------------------------------------------------------------

DROP TABLE IF EXISTS `password_resets`;

CREATE TABLE `password_resets` (
                                   `email` varchar(255) NOT NULL,
                                   `token` varchar(255) NOT NULL,
                                   `created_at` timestamp NOT NULL,
                                   KEY `password_resets_email_index` (`email`),
                                   KEY `password_resets_token_index` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table permission_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `permission_role`;

CREATE TABLE `permission_role` (
                                   `permission_id` int(10) unsigned NOT NULL,
                                   `role_id` int(10) unsigned NOT NULL,
                                   PRIMARY KEY (`permission_id`,`role_id`),
                                   KEY `permission_role_role_id_foreign` (`role_id`),
                                   CONSTRAINT `permission_role_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                                   CONSTRAINT `permission_role_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `permissions`;

CREATE TABLE `permissions` (
                               `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                               `name` varchar(255) NOT NULL,
                               `display_name` varchar(255) DEFAULT NULL,
                               `description` varchar(255) DEFAULT NULL,
                               `created_at` timestamp NULL DEFAULT NULL,
                               `updated_at` timestamp NULL DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `permissions_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table plans
# ------------------------------------------------------------

DROP TABLE IF EXISTS `plans`;

CREATE TABLE `plans` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `alias` varchar(255) NOT NULL COMMENT '别名',
                         `name` varchar(255) NOT NULL COMMENT '名称',
                         `description` varchar(255) NOT NULL COMMENT '描述',
                         `price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '价格',
                         `promo_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '促销价格',
                         `promo_start` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '促销开始时间',
                         `promo_end` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '促销结束时间',
                         `valid_days` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '有效天数',
                         `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:正常 0:默认',
                         `user_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '买家uid',
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table posts
# ------------------------------------------------------------

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `title` varchar(255) NOT NULL,
                         `slug` varchar(255) NOT NULL DEFAULT '',
                         `summary` varchar(255) NOT NULL DEFAULT '',
                         `origin_content` text NOT NULL,
                         `view_count` int(11) NOT NULL DEFAULT '0' COMMENT '浏览总数',
                         `content` text NOT NULL,
                         `user_id` int(11) NOT NULL DEFAULT '0',
                         `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:草稿 1:已发布',
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table role_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `role_user`;

CREATE TABLE `role_user` (
                             `user_id` int(10) unsigned NOT NULL,
                             `role_id` int(10) unsigned NOT NULL,
                             PRIMARY KEY (`user_id`,`role_id`),
                             KEY `role_user_role_id_foreign` (`role_id`),
                             CONSTRAINT `role_user_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                             CONSTRAINT `role_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `roles`;

CREATE TABLE `roles` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) NOT NULL,
                         `display_name` varchar(255) DEFAULT NULL,
                         `description` varchar(255) DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `roles_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sections
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sections`;

CREATE TABLE `sections` (
                            `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                            `course_id` int(11) NOT NULL DEFAULT '0' COMMENT '课程id',
                            `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'section名称',
                            `weight` tinyint(2) NOT NULL DEFAULT '0' COMMENT '排序值',
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程章节';



# Dump of table tags
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tags`;

CREATE TABLE `tags` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                        `tag` varchar(255) NOT NULL COMMENT 'tag名称',
                        `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '新闻id',
                        `tag_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '标签名',
                        `created_at` timestamp NULL DEFAULT NULL,
                        `updated_at` timestamp NULL DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table user_activations
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_activations`;

CREATE TABLE `user_activations` (
                                    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                    `user_id` int(10) unsigned NOT NULL,
                                    `token` varchar(255) NOT NULL,
                                    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`),
                                    KEY `user_activations_user_id_foreign` (`user_id`),
                                    CONSTRAINT `user_activations_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table user_members
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_members`;

CREATE TABLE `user_members` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'uid',
                                `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '会员类型 1:月度 2:季度 3:半年 4:年卡 5:2年 6:3年',
                                `start_time` datetime NOT NULL COMMENT '开始时间',
                                `end_time` datetime NOT NULL COMMENT '结束时间',
                                `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 2:无效 1:有效',
                                `created_at` timestamp NULL DEFAULT NULL,
                                `updated_at` timestamp NULL DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `uniq_uid` (`user_id`),
                                KEY `idx_uid` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户会员表，一个用户只有一条记录';



# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `username` varchar(255) NOT NULL DEFAULT '',
                         `email` varchar(255) NOT NULL,
                         `password` varchar(60) NOT NULL,
                         `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
                         `real_name` varchar(255) NOT NULL DEFAULT '' COMMENT '真实姓名',
                         `city` varchar(255) NOT NULL DEFAULT '' COMMENT '所在城市',
                         `company` varchar(255) NOT NULL DEFAULT '' COMMENT '所在公司',
                         `weibo_url` varchar(255) NOT NULL DEFAULT '' COMMENT '微博',
                         `wechat_id` varchar(255) NOT NULL DEFAULT '' COMMENT '微信',
                         `personal_website` varchar(255) NOT NULL DEFAULT '' COMMENT '个人网站',
                         `introduction` varchar(255) NOT NULL DEFAULT '' COMMENT '自我介绍',
                         `topic_count` int(11) NOT NULL DEFAULT '0',
                         `reply_count` int(11) NOT NULL DEFAULT '0',
                         `follower_count` int(11) NOT NULL DEFAULT '0',
                         `notification_count` int(11) NOT NULL DEFAULT '0',
                         `status` tinyint(4) NOT NULL DEFAULT '1',
                         `last_login_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                         `last_login_ip` varchar(45) NOT NULL DEFAULT '',
                         `github_id` varchar(100) NOT NULL DEFAULT '' COMMENT 'github ID',
                         `github_name` varchar(50) NOT NULL DEFAULT '',
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         `remember_token` varchar(100) DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         `is_activated` tinyint(1) NOT NULL COMMENT '是否激活 1:是 0:否',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `users_name_unique` (`username`),
                         UNIQUE KEY `users_email_unique` (`email`),
                         KEY `users_topic_count_index` (`topic_count`),
                         KEY `users_reply_count_index` (`reply_count`),
                         KEY `users_follower_count_index` (`follower_count`),
                         KEY `users_notification_count_index` (`notification_count`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table videos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `videos`;

CREATE TABLE `videos` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `course_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属课程id',
                          `section_id` int(11) NOT NULL DEFAULT '0',
                          `episode_id` int(10) NOT NULL DEFAULT '0' COMMENT 'episode id',
                          `name` varchar(255) DEFAULT NULL COMMENT '名称',
                          `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '页面关键词',
                          `description` varchar(255) DEFAULT NULL COMMENT '描述',
                          `cover_key` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图',
                          `mp4_key` varchar(255) NOT NULL DEFAULT '' COMMENT '视频u地址',
                          `is_free` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否免费 0:收费, 1:免费',
                          `length` varchar(10) NOT NULL DEFAULT '0' COMMENT '视频长度',
                          `duration` int(11) NOT NULL COMMENT '视频总长度，单位:秒',
                          `is_publish` tinyint(1) NOT NULL DEFAULT '0',
                          `published_at` timestamp NULL DEFAULT NULL COMMENT '发布时间',
                          `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
                          `created_at` timestamp NULL DEFAULT NULL,
                          `updated_at` timestamp NULL DEFAULT NULL,
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table votes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `votes`;

CREATE TABLE `votes` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '投票者id',
                         `votable_id` int(10) unsigned NOT NULL DEFAULT '9' COMMENT '被投票的关联id',
                         `votable_type` varchar(255) NOT NULL,
                         `is` varchar(255) NOT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `votes_votable_type_index` (`votable_type`),
                         KEY `votes_is_index` (`is`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table wiki_categories
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wiki_categories`;

CREATE TABLE `wiki_categories` (
                                   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                   `name` varchar(255) NOT NULL COMMENT '分类名',
                                   `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
                                   `weight` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '权重',
                                   `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:下线 1:上线',
                                   `created_at` timestamp NULL DEFAULT NULL,
                                   `updated_at` timestamp NULL DEFAULT NULL,
                                   `deleted_at` timestamp NULL DEFAULT NULL,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table wiki_pages
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wiki_pages`;

CREATE TABLE `wiki_pages` (
                              `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                              `is_parent` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:不显示在目录 1:显示在目录',
                              `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属分类id',
                              `title` varchar(255) NOT NULL DEFAULT '',
                              `slug` varchar(128) NOT NULL DEFAULT '',
                              `summary` varchar(255) NOT NULL DEFAULT '',
                              `origin_content` text NOT NULL,
                              `content` text NOT NULL,
                              `weight` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '排序权重',
                              `view_count` int(11) NOT NULL DEFAULT '0' COMMENT '浏览总数',
                              `fix_count` int(11) NOT NULL DEFAULT '0' COMMENT '修正数',
                              `comment_count` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
                              `user_id` int(11) NOT NULL DEFAULT '0',
                              `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:草稿 1:已发布',
                              `created_at` timestamp NULL DEFAULT NULL,
                              `updated_at` timestamp NULL DEFAULT NULL,
                              `deleted_at` timestamp NULL DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `uniq_slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
