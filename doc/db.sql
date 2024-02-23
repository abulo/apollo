/*
 Navicat Premium Data Transfer

 Source Server         : docker_mysql
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : apollo

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 21/02/2024 16:22:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
  `username` varchar(64) NOT NULL COMMENT '用户名称',
  `password` varchar(64) NOT NULL COMMENT '用户密码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '用户状态（0正常 1停用）',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:login` (`username`),
  KEY `idx:status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户';

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色编号',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `code` varchar(64) NOT NULL COMMENT '角色权限字符串',
  `sort` int NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '角色状态（0正常 1停用）',
  `type` tinyint NOT NULL COMMENT '角色类型(1内置/2定义)',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`type`,`status`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色';

-- ----------------------------
-- Table structure for system_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_role_menu`;
CREATE TABLE `system_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `system_role_id` bigint NOT NULL COMMENT '角色编号',
  `system_menu_id` bigint NOT NULL COMMENT '菜单编号',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:menu:role` (`system_menu_id`,`system_role_id`),
  KEY `idx:delete` (`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色和系统菜单关联表';

-- ----------------------------
-- Table structure for system_user_role
-- ----------------------------
DROP TABLE IF EXISTS `system_user_role`;
CREATE TABLE `system_user_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `system_user_id` bigint NOT NULL COMMENT '用户编号',
  `system_role_id` bigint NOT NULL COMMENT '角色编号',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `list:role` (`system_role_id`),
  KEY `idx:list` (`deleted`),
  KEY `list:user` (`system_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户和系统角色关联表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `nickname` varchar(128) DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `birthday` date DEFAULT NULL COMMENT '用户生日',
  `gender` tinyint DEFAULT NULL COMMENT '用户性别(0女/1男)',
  `status` tinyint NOT NULL COMMENT '用户状态(0正常/1锁定)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `identity_type` varchar(128) DEFAULT NULL COMMENT '登录类型(手机号/邮箱) 或第三方应用名称 (微信/微博等)',
  `identifier` varchar(128) DEFAULT NULL COMMENT '手机号/邮箱/第三方的唯一标识',
  `credential` varchar(255) DEFAULT NULL COMMENT '密码凭证 (自建账号的保存密码, 第三方的保存 token)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:identity` (`identity_type`,`identifier`),
  KEY `idx:user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户三方登录授权';

-- ----------------------------
-- Table structure for region
-- ----------------------------
DROP TABLE IF EXISTS `region`;
CREATE TABLE `region` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '区域编号',
  `name` varchar(255) NOT NULL COMMENT '区域名称',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父级编号',
  `weather_code` varchar(255) DEFAULT NULL COMMENT '天气code',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:parent` (`parent_id`),
  KEY `idx:name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='地区表';

-- ----------------------------
-- Table structure for user_secret
-- ----------------------------
DROP TABLE IF EXISTS `user_secret`;
CREATE TABLE `user_secret` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `password` varchar(255) DEFAULT NULL COMMENT '用户密码',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户密码';

-- ----------------------------
-- Table structure for system_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `system_dict_type`;
CREATE TABLE `system_dict_type` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `name` varchar(64) NOT NULL COMMENT '字典名称',
  `type` varchar(64) NOT NULL COMMENT '字典类型',
  `status` tinyint NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:type` (`type`),
  KEY `idx:status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典类型';

-- ----------------------------
-- Table structure for system_login_log
-- ----------------------------
DROP TABLE IF EXISTS `system_login_log`;
CREATE TABLE `system_login_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(64) NOT NULL COMMENT '用户账号',
  `user_ip` varchar(64) NOT NULL COMMENT '用户ip',
  `user_agent` varchar(255) DEFAULT NULL COMMENT 'UA',
  `login_time` datetime NOT NULL COMMENT '登录时间',
  `channel` varchar(64) NOT NULL COMMENT '渠道',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:search` (`username`,`login_time`,`channel`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='登录日志';

-- ----------------------------
-- Table structure for system_operate_log
-- ----------------------------
DROP TABLE IF EXISTS `system_operate_log`;
CREATE TABLE `system_operate_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(64) NOT NULL COMMENT '用户账号',
  `module` varchar(64) NOT NULL COMMENT '模块名称',
  `request_method` varchar(64) NOT NULL COMMENT '请求方法名',
  `request_url` varchar(255) NOT NULL COMMENT '请求地址',
  `user_ip` varchar(64) NOT NULL COMMENT '用户 ip',
  `user_agent` varchar(255) DEFAULT NULL COMMENT 'UA',
  `go_method` varchar(64) NOT NULL COMMENT '方法名',
  `go_method_args` json DEFAULT NULL COMMENT '方法的参数',
  `start_time` datetime NOT NULL COMMENT '操作开始时间',
  `duration` int NOT NULL COMMENT '执行时长',
  `channel` varchar(64) NOT NULL COMMENT '渠道',
  `result` tinyint NOT NULL COMMENT '结果(0 成功/1 失败)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:log` (`username`,`module`,`start_time`,`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志';

-- ----------------------------
-- Table structure for system_dict
-- ----------------------------
DROP TABLE IF EXISTS `system_dict`;
CREATE TABLE `system_dict` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `sort` int NOT NULL DEFAULT '0' COMMENT '字典排序',
  `label` varchar(64) NOT NULL COMMENT '字典标签',
  `value` varchar(64) NOT NULL COMMENT '字典键值',
  `dict_type` varchar(64) NOT NULL COMMENT '字典类型',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `list:data` (`dict_type`,`status`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典数据表';

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单编号',
  `name` varchar(64) NOT NULL COMMENT '菜单名称',
  `permission` varchar(128) DEFAULT NULL COMMENT '权限标识',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '菜单类型(1:目录/2: 菜单/3: 按钮)',
  `sort` int NOT NULL DEFAULT '1' COMMENT '显示顺序',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父菜单ID',
  `path` varchar(255) DEFAULT NULL COMMENT '路由地址',
  `icon` varchar(128) DEFAULT NULL COMMENT '菜单图标',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `component_name` varchar(255) DEFAULT NULL COMMENT '组件名',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '菜单状态(0开启/1关闭)',
  `hide` tinyint NOT NULL DEFAULT '0' COMMENT '是否隐藏(0:否/1是)',
  `link` varchar(255) DEFAULT NULL COMMENT '路由外链时填写的访问地址',
  `keep_alive` tinyint NOT NULL DEFAULT '1' COMMENT '是否缓存(0不/ 1是)',
  `affix` tinyint NOT NULL DEFAULT '0' COMMENT '是否总是显示(0 不显示/1 显示)',
  `active_path` varchar(255) DEFAULT NULL COMMENT '激活链接',
  `full_screen` tinyint NOT NULL DEFAULT '0' COMMENT '是否全屏',
  `redirect` varchar(255) DEFAULT NULL COMMENT '路由重定向地址',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `list:parent` (`deleted`,`status`,`type`,`sort`,`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统菜单';

SET FOREIGN_KEY_CHECKS = 1;
