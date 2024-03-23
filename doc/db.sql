/*
 Navicat Premium Data Transfer

 Source Server         : docker_mysql
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : apollo_tenant

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 23/03/2024 14:48:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
-- Records of region
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_dept
-- ----------------------------
DROP TABLE IF EXISTS `system_dept`;
CREATE TABLE `system_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` varchar(255) NOT NULL COMMENT '部门名称',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父部门ID',
  `sort` int NOT NULL COMMENT '显示顺序',
  `system_user_id` bigint DEFAULT '0' COMMENT '负责人',
  `phone` varchar(255) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(255) DEFAULT NULL COMMENT '邮件',
  `status` tinyint NOT NULL COMMENT '部门状态',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  `system_tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`status`,`parent_id`,`name`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门';

-- ----------------------------
-- Records of system_dept
-- ----------------------------
BEGIN;
COMMIT;

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
  `color_type` varchar(255) DEFAULT NULL COMMENT '颜色类型',
  `css_class` varchar(255) DEFAULT NULL COMMENT 'css 样式',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `list:data` (`dict_type`,`status`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典数据表';

-- ----------------------------
-- Records of system_dict
-- ----------------------------
BEGIN;
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, 0, '启用', '0', 'status', 0, 'success', '', '', '', '2024-02-26 03:56:06', '', '2024-02-26 03:56:49');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, 1, '停用', '1', 'status', 0, 'danger', '', '', '', '2024-02-26 03:57:47', '', '2024-02-26 03:57:47');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (3, 1, '目录', '1', 'menu.type', 0, 'primary', '', '', '', '2024-02-27 09:42:28', '', '2024-02-27 09:42:28');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (4, 2, '菜单', '2', 'menu.type', 0, 'success', '', '', '', '2024-02-26 09:43:08', '', '2024-02-26 09:43:08');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (5, 3, '按钮', '3', 'menu.type', 0, 'info', '', '', '', '2024-02-26 01:44:02', '', '2024-02-26 01:44:02');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (6, 1, '正常', '0', 'delete', 0, 'success', '', '', '', '2024-02-26 09:49:26', '', '2024-02-26 09:49:26');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, 2, '删除', '1', 'delete', 0, 'danger', '', '', '', '2024-02-26 17:49:36', '', '2024-02-26 17:49:36');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (8, 1, '总是', '1', 'menu.affix', 0, '', '', '', '', '2024-02-27 09:51:46', '', '2024-02-27 09:51:46');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (9, 2, '不是', '0', 'menu.affix', 0, '', '', '', '', '2024-02-27 01:51:58', '', '2024-02-27 01:51:58');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (10, 1, '显示', '0', 'menu.hide', 0, '', '', '', '', '2024-02-26 01:56:22', '', '2024-02-26 01:56:22');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (11, 2, '隐藏', '1', 'menu.hide', 0, '', '', '', '', '2024-02-26 01:56:33', '', '2024-02-26 01:56:33');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (12, 1, '启用', '1', 'menu.keepAlive', 0, '', '', '', '', '2024-02-26 01:58:32', '', '2024-02-26 01:58:32');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (13, 2, '禁用', '0', 'menu.keepAlive', 0, '', '', '', '', '2024-02-26 01:58:51', '', '2024-02-26 01:58:51');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (14, 1, '启用', '1', 'menu.fullScreen', 0, '', '', '', '', '2024-02-26 01:59:49', '', '2024-02-26 01:59:49');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (15, 2, '禁用', '0', 'menu.fullScreen', 0, '', '', '', '', '2024-02-26 02:00:00', '', '2024-02-26 02:00:00');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (16, 1, '内置', '1', 'role.type', 0, 'success', '', '', '', '2024-02-27 03:07:20', 'admin', '2024-02-29 09:29:09');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (17, 3, '自定义', '2', 'role.type', 0, 'info', '', '', '', '2024-02-27 03:07:47', '', '2024-02-27 03:07:47');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (18, 1, '全部数据权限', '1', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:20:03', 'admin', '2024-03-12 11:32:51');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (19, 2, '指定部门数据权限', '2', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:32:36', 'admin', '2024-03-12 11:32:59');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (20, 3, '本部门数据权限', ' 3', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:33:20', '', '2024-03-12 11:33:20');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (21, 4, '本部门及以下数据权限', '4', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:34:28', '', '2024-03-12 11:34:28');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (22, 5, '仅本人数据权限', '5', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:34:49', '', '2024-03-12 11:34:49');
COMMIT;

-- ----------------------------
-- Table structure for system_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `system_dict_type`;
CREATE TABLE `system_dict_type` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `name` varchar(64) NOT NULL COMMENT '字典名称',
  `type` varchar(64) NOT NULL COMMENT '字典类型',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:type` (`type`),
  KEY `idx:list` (`status`,`type`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典类型';

-- ----------------------------
-- Records of system_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '状态', 'status', 0, '', '', '2024-02-25 08:25:12', '', '2024-02-25 08:25:12');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, '删除', 'delete', 0, '', '', '2024-02-25 08:25:35', '', '2024-02-25 08:25:35');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (3, '菜单类型', 'menu.type', 0, '', '', '2024-02-26 01:38:37', '', '2024-02-26 01:38:37');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (4, '菜单显示', 'menu.hide', 0, '', '', '2024-02-26 09:51:25', '', '2024-02-26 09:51:25');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (5, '菜单固定', 'menu.affix', 0, '', '', '2024-02-26 01:55:42', '', '2024-02-26 01:55:42');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (6, '菜单缓存', 'menu.keepAlive', 0, '', '', '2024-02-26 01:58:14', '', '2024-02-26 01:58:14');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, '菜单全屏', 'menu.fullScreen', 0, '', '', '2024-02-26 01:59:33', 'admin', '2024-02-29 09:23:31');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (8, '角色类别', 'role.type', 0, '', '', '2024-02-27 03:06:42', 'admin', '2024-02-29 09:23:23');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (9, '角色数据范围', 'role.scope', 0, '', 'admin', '2024-03-12 11:19:26', '', '2024-03-12 11:19:26');
COMMIT;

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
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`username`,`login_time`,`channel`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='登录日志';

-- ----------------------------
-- Records of system_login_log
-- ----------------------------
BEGIN;
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统菜单';

-- ----------------------------
-- Records of system_menu
-- ----------------------------
BEGIN;
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (1, '首页', 'home', 2, 1, 0, '/home/index', 'HomeFilled', '/home/index', 'home', 0, 0, '', 1, 1, '', 0, '', 'admin', '2023-09-26 13:44:46', 'admin', '2024-02-27 23:48:11', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (2, '系统设置', '', 1, 2, 0, '/system', 'Tools', '', '', 0, 0, '', 1, 0, '', 0, '/system/menu', 'admin', '2023-08-08 15:17:56', 'admin', '2024-02-28 09:15:06', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (3, '菜单管理', '', 2, 21, 2, '/system/menu', 'Menu', '/system/menu/index', 'systemMenu', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 20:02:04', 'admin', '2024-03-01 09:42:22', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (4, '菜单新增', 'menu.SystemMenuCreate', 3, 1, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:22:08', '', '2024-02-25 08:22:08', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (5, '菜单编辑', 'menu.SystemMenuUpdate', 3, 2, 3, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-12 13:30:26', 'admin', '2023-05-23 04:57:14', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (6, '菜单删除', 'menu.SystemMenuDelete', 3, 3, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2023-07-12 13:29:55', 'admin', '2023-05-23 14:38:18', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (7, '菜单恢复', 'menu.SystemMenuRecover', 3, 4, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2023-07-12 13:28:58', '', '1970-01-01 00:00:00', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (8, '字典管理', '', 2, 22, 2, '/system/dict', 'Reading', '/system/dict/dict_type', 'systemDictType', 0, 0, '', 1, 0, '', 0, '', '', '2024-02-25 13:56:23', 'admin', '2024-03-01 09:43:02', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (10, '字典添加', 'dict.SystemDictTypeCreate', 3, 1, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:23:52', '', '2024-02-25 08:23:52', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (11, '字典修改', 'dict.SystemDictTypeUpdate', 3, 2, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:24:25', '', '2024-02-25 08:24:25', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (12, '字典删除', 'dict.SystemDictTypeDelete', 3, 3, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:24:47', '', '2024-02-25 08:24:47', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (13, '字典数据', '', 2, 4, 8, '/system/dict/:type/data', 'Document', '/system/dict/dict_data', 'systemDict', 0, 1, '', 1, 0, '/system/dict', 0, '', '', '2024-02-26 19:01:23', 'admin', '2024-03-01 09:44:00', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (14, '字典数据添加', 'dict.SystemDictCreate', 3, 1, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 19:02:29', '', '2024-02-25 11:02:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (15, '字典数据修改', 'dict.SystemDictUpdate', 3, 2, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 19:02:58', '', '2024-02-25 11:02:58', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (16, '字典数据删除', 'dict.SystemDictDelete', 3, 3, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 19:03:19', '', '2024-02-25 11:03:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (17, '角色管理', '', 2, 23, 2, '/system/role', 'Stamp', '/system/role/index', 'systemRole', 0, 0, '', 1, 0, '', 0, '', '', '2024-02-27 19:01:09', 'admin', '2024-03-01 09:44:15', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (18, '角色添加', 'role.SystemRoleCreate', 3, 1, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 11:02:05', '', '2024-02-27 03:02:05', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (19, '角色修改', 'role.SystemRoleUpdate', 3, 2, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 11:02:29', '', '2024-02-27 03:02:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (20, '角色删除', 'role.SystemRoleDelete', 3, 3, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 11:02:49', '', '2024-02-27 03:02:49', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (21, '用户管理', '', 2, 20, 2, '/system/user', 'UserFilled', '/system/user/index', 'systemUser', 0, 0, '', 1, 0, '', 0, '', '', '2024-02-27 17:06:34', 'admin', '2024-03-01 09:41:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (22, '用户添加', 'user.SystemUserCreate', 3, 1, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 17:07:46', '', '2024-02-27 09:07:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (23, '用户修改', 'user.SystemUserUpdate', 3, 2, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-28 01:08:08', '', '2024-02-27 17:08:34', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (24, '用户删除', 'user.SystemUserDelete', 3, 3, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 17:08:26', '', '2024-02-27 09:08:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (68, '审计日志', '', 1, 3, 0, '/system/logger', 'Tickets', '', '', 0, 0, '', 1, 0, '', 0, NULL, 'admin', '2023-08-09 22:03:11', 'admin', '2024-02-27 23:22:48', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (69, '操作日志', '', 2, 3, 68, '/system/logger/operate', 'DocumentCopy', '/system/logger/operate', 'systemLoggerOperate', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 18:56:50', 'admin', '2024-03-01 09:46:35', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (70, '登录日志', '', 2, 2, 68, '/system/logger/login', 'FolderRemove', '/system/logger/login', 'systemLoggerLogin', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 09:53:22', 'admin', '2024-03-01 09:45:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (71, '操作日志删除', 'logger.SystemOperateLogDelete', 3, 1, 69, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 10:52:07', 'admin', '2024-02-28 17:54:40', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (72, '操作日志清空', 'logger.SystemOperateLogDrop', 3, 1, 69, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 10:52:32', 'admin', '2024-02-28 17:54:59', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (73, '操作日志查看', 'logger.SystemOperateLog', 3, 1, 69, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 10:53:02', 'admin', '2024-02-28 17:55:15', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (74, '登录日志删除', 'logger.SystemLoginLogDelete', 3, 1, 70, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 17:54:11', 'admin', '2024-02-28 17:53:07', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (75, '登录日志清空', 'logger.SystemLoginLogDrop', 3, 1, 70, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 17:54:39', 'admin', '2024-02-28 17:53:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (76, '登录日志查看', 'logger.SystemLoginLog', 3, 1, 70, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 17:55:02', 'admin', '2024-02-28 17:53:47', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (77, '系统日志', '', 2, 1, 68, '/system/logger/entry', 'DocumentRemove', '/system/logger/entry', 'systemLoggerEntry', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 19:14:50', 'admin', '2024-03-01 09:45:09', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (78, '系统日志删除', 'logger.SystemEntryLogDelete', 3, 1, 77, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 14:28:41', 'admin', '2024-02-28 17:50:48', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (79, '系统日志查看', 'logger.SystemEntryLog', 3, 0, 77, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-02-28 17:51:16', 'admin', '2024-03-23 13:35:53', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (80, '系统日志清空', 'logger.SystemEntryLogDrop', 3, 0, 77, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-02-28 17:51:38', '', '2024-02-28 09:51:38', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (81, '用户列表', 'user.SystemUserList', 3, 0, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:41:17', 'admin', '2024-03-01 09:41:34', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (82, '菜单列表', 'menu.SystemMenuList', 3, 0, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:42:10', '', '2024-03-01 09:42:10', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (83, '字典列表', 'dict.SystemDictTypeList', 3, 0, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:43:17', '', '2024-03-01 09:43:17', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (84, '字典数据列表', 'dict.SystemDictList', 3, 0, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:43:48', '', '2024-03-01 09:43:48', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (85, '角色列表', 'role.SystemRoleList', 3, 0, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:44:32', 'admin', '2024-03-01 09:44:53', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (86, '系统日志列表', 'logger.SystemEntryLogList', 3, 0, 77, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:45:33', '', '2024-03-01 09:45:33', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (87, '登录日志列表', 'logger.SystemLoginLogList', 3, 0, 70, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:46:20', '', '2024-03-01 09:46:20', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (88, '操作日志列表', 'logger.SystemOperateLogList', 3, 0, 69, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-01 09:46:55', '', '2024-03-01 09:46:55', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (89, '部门管理', '', 2, 24, 2, '/system/dept', 'House', '/system/dept/index', 'systemDept', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-03-11 12:03:03', 'admin', '2024-03-11 12:04:50', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (90, '部门列表', 'dept.SystemDeptList', 3, 0, 89, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 12:05:52', '', '2024-03-11 12:05:52', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (91, '部门新增', 'dept.SystemDeptCreate', 3, 0, 89, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 12:06:18', '', '2024-03-11 12:06:18', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (92, '部门编辑', 'dept.SystemDeptUpdate', 3, 0, 89, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 12:06:35', '', '2024-03-11 12:06:35', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (93, '部门删除', 'dept.SystemDeptDelete', 3, 0, 89, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 12:06:54', '', '2024-03-11 12:06:54', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (94, '部门恢复', 'dept.SystemDeptRecover', 3, 0, 89, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 12:07:23', '', '2024-03-11 12:07:23', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (95, '职位管理', '', 2, 25, 2, '/system/post', 'TakeawayBox', '/system/post/index', 'systemPost', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-03-11 16:59:14', '', '2024-03-11 16:59:14', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (96, '职位列表', 'post.SystemPostList', 3, 0, 95, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 17:00:17', '', '2024-03-11 17:00:17', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (97, '职位新增', 'post.SystemPostCreate', 3, 0, 95, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 17:00:39', '', '2024-03-11 17:00:39', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (98, '职位编辑', 'post.SystemPostUpdate', 3, 0, 95, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 17:00:59', '', '2024-03-11 17:00:59', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (99, '职位删除', 'post.SystemPostDelete', 3, 0, 95, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 17:01:21', '', '2024-03-11 17:01:21', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (100, '职位恢复', 'post.SystemPostRecover', 3, 0, 95, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-11 17:01:42', '', '2024-03-11 17:01:42', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (101, '职位分配', 'user.SystemUserPostList', 3, 0, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 10:52:02', '', '2024-03-13 10:52:02', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (102, '职位处理', 'user.SystemUserPostCreate', 3, 0, 101, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 10:52:29', '', '2024-03-13 10:52:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (103, '角色分配', 'user.SystemUserRoleList', 3, 0, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 10:53:20', '', '2024-03-13 10:53:20', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (104, '角色处理', 'user.SystemUserRoleCreate', 3, 0, 103, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 10:53:44', '', '2024-03-13 10:53:44', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (105, '部门分配', 'user.SystemUserDeptList', 3, 0, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 10:54:07', '', '2024-03-13 10:54:07', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (106, '部门处理', 'user.SystemUserDeptCreate', 3, 0, 105, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 10:54:28', '', '2024-03-13 10:54:28', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (107, '用户恢复', 'user.SystemUserRecover', 3, 0, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-13 11:00:31', '', '2024-03-13 11:00:31', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (108, '租户管理', '', 2, 19, 2, '/system/tenant', 'Avatar', '', '', 0, 0, '', 1, 0, '', 0, '/system/tenant/index', '芋道源码', '2023-07-11 08:32:02', 'admin', '2024-03-16 17:13:49', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (109, '套餐管理', '', 2, 1, 108, '/system/tenant/package', 'Memo', '/system/tenant/package', 'systemTenantPackage', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-03-16 16:53:59', 'admin', '2024-03-16 17:16:38', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (110, '用户管理', '', 2, 22, 108, '/system/tenant/index', 'User', '/system/tenant/index', 'systemTenant', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-03-16 16:58:10', 'admin', '2024-03-16 17:17:01', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (111, '套餐列表', 'tenant.SystemTenantPackageList', 3, 1, 109, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 16:59:56', '', '2024-03-16 16:59:56', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (112, '套餐添加', 'tenant.SystemTenantPackageCreate', 3, 2, 109, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:01:34', '', '2024-03-16 17:01:34', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (113, '套餐修改', 'tenant.SystemTenantPackageUpdate', 3, 3, 109, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:02:03', '', '2024-03-16 17:02:03', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (114, '套餐删除', 'tenant.SystemTenantPackageDelete', 3, 4, 109, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:02:28', '', '2024-03-16 17:02:28', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (115, '套餐恢复', 'tenant.SystemTenantPackageRecover', 3, 5, 109, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:03:04', '', '2024-03-16 17:03:04', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (116, '用户列表', 'tenant.SystemTenantList', 3, 1, 110, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:03:38', '', '2024-03-16 17:03:38', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (117, '用户添加', 'tenant.SystemTenantCreate', 3, 2, 110, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:04:06', '', '2024-03-16 17:04:06', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (118, '用户修改', 'tenant.SystemTenantUpdate', 3, 3, 110, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:04:36', '', '2024-03-16 17:04:36', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (119, '用户删除', 'tenant.SystemTenantDelete', 3, 4, 110, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:05:02', '', '2024-03-16 17:05:02', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (120, '用户恢复', 'tenant.SystemTenantRecover', 3, 5, 110, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-16 17:05:26', '', '2024-03-16 17:05:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (121, '用户登录', 'tenant.SystemTenantLogin', 3, 0, 110, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'asd', '2024-03-19 00:52:59', '', '2024-03-19 00:52:59', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (122, '重置密码', 'user.SystemUserPassword', 3, 0, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'asd', '2024-03-19 00:53:47', '', '2024-03-19 00:53:47', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (123, '菜单权限', 'role.SystemRoleMenuList', 3, 0, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-23 14:00:57', '', '2024-03-23 14:00:57', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (124, '菜单权限处理', 'role.SystemRoleMenuCreate', 3, 0, 123, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-23 14:01:23', '', '2024-03-23 14:01:23', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (125, '数据权限', 'role.SystemRoleDataScope', 3, 0, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-23 14:01:50', '', '2024-03-23 14:01:50', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (126, '数据权限处理', 'role.SystemRoleDataScopeCreate', 3, 0, 125, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-03-23 14:02:11', '', '2024-03-23 14:02:11', 0);
COMMIT;

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
  `result` tinyint NOT NULL DEFAULT '0' COMMENT '结果(0 成功/1 失败)',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`username`,`module`,`start_time`,`result`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志';

-- ----------------------------
-- Records of system_operate_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_post
-- ----------------------------
DROP TABLE IF EXISTS `system_post`;
CREATE TABLE `system_post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '职位ID',
  `name` varchar(255) NOT NULL COMMENT '职位名称',
  `sort` int NOT NULL COMMENT '显示顺序',
  `status` tinyint NOT NULL COMMENT '状态',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`status`,`name`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='职位';

-- ----------------------------
-- Records of system_post
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色编号',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `code` varchar(64) NOT NULL COMMENT '角色权限字符串',
  `sort` int NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `data_scope` tinyint DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `data_scope_dept` json DEFAULT NULL COMMENT '数据范围(指定部门数组)',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '角色状态（0正常 1停用）',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '角色类型(1内置/2定义)',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`type`,`status`,`sort`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色';

-- ----------------------------
-- Records of system_role
-- ----------------------------
BEGIN;
INSERT INTO `system_role` (`id`, `name`, `code`, `sort`, `data_scope`, `data_scope_dept`, `status`, `type`, `remark`, `deleted`, `system_tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '超级管理员', 'super', 0, NULL, NULL, 0, 2, '', 0, 1, 'admin', '2024-03-23 14:02:42', NULL, '2024-03-23 14:02:42');
COMMIT;

-- ----------------------------
-- Table structure for system_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_role_menu`;
CREATE TABLE `system_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `system_role_id` bigint NOT NULL COMMENT '角色编号',
  `system_menu_id` bigint NOT NULL COMMENT '菜单编号',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:role_menu` (`system_role_id`,`system_menu_id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`system_menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色和系统菜单关联表';

-- ----------------------------
-- Records of system_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_tenant
-- ----------------------------
DROP TABLE IF EXISTS `system_tenant`;
CREATE TABLE `system_tenant` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '租户编号',
  `name` varchar(255) NOT NULL COMMENT '租户名称',
  `system_user_id` bigint DEFAULT NULL COMMENT '联系人ID',
  `contact_name` varchar(255) NOT NULL COMMENT '联系人',
  `contact_mobile` varchar(255) NOT NULL COMMENT '租户联系电话',
  `status` tinyint NOT NULL COMMENT '状态（0正常 1停用）',
  `domain` varchar(255) DEFAULT NULL COMMENT '域名',
  `expire_date` date NOT NULL COMMENT '过期时间',
  `account_count` bigint NOT NULL COMMENT '账号数量',
  `system_tenant_package_id` bigint NOT NULL COMMENT '套餐编号',
  `deleted` tinyint NOT NULL COMMENT '是否删除(0否 1是)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:package` (`system_tenant_package_id`),
  KEY `idx:list` (`deleted`,`status`,`name`,`expire_date`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户';

-- ----------------------------
-- Records of system_tenant
-- ----------------------------
BEGIN;
INSERT INTO `system_tenant` (`id`, `name`, `system_user_id`, `contact_name`, `contact_mobile`, `status`, `domain`, `expire_date`, `account_count`, `system_tenant_package_id`, `deleted`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '系统租户', 1, '周先生', '13512345678', 0, 'http://www.baidu.com', '2029-03-01', 100, 0, 0, NULL, '2024-03-20 21:14:57', 'admin', '2024-03-20 21:58:40');
COMMIT;

-- ----------------------------
-- Table structure for system_tenant_package
-- ----------------------------
DROP TABLE IF EXISTS `system_tenant_package`;
CREATE TABLE `system_tenant_package` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '套餐编号',
  `name` varchar(255) NOT NULL COMMENT '套餐名称',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `system_menu_ids` json NOT NULL COMMENT '目录编号',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `deleted` tinyint NOT NULL COMMENT '是否删除(0否 1是)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`status`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户套餐包';

-- ----------------------------
-- Records of system_tenant_package
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
  `mobile` varchar(255) DEFAULT NULL COMMENT '手机号码',
  `username` varchar(64) NOT NULL COMMENT '用户名称',
  `password` varchar(64) NOT NULL COMMENT '用户密码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '用户状态（0正常 1停用）',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除(0否 1是)',
  `system_tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:login` (`username`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`status`,`nickname`,`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户';

-- ----------------------------
-- Records of system_user
-- ----------------------------
BEGIN;
INSERT INTO `system_user` (`id`, `nickname`, `mobile`, `username`, `password`, `status`, `deleted`, `system_tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '安安', '15397620000', 'admin', '21232f297a57a5a743894a0e4a801fc3', 0, 0, 1, NULL, '2024-03-20 21:15:35', 'admin', '2024-03-20 21:52:56');
COMMIT;

-- ----------------------------
-- Table structure for system_user_dept
-- ----------------------------
DROP TABLE IF EXISTS `system_user_dept`;
CREATE TABLE `system_user_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `system_user_id` bigint NOT NULL COMMENT '系统用户 id',
  `system_dept_id` bigint NOT NULL COMMENT '部门 id',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uniq:user_dept` (`system_user_id`,`system_dept_id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`system_dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户部门';

-- ----------------------------
-- Records of system_user_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_user_post
-- ----------------------------
DROP TABLE IF EXISTS `system_user_post`;
CREATE TABLE `system_user_post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `system_user_id` bigint NOT NULL COMMENT '系统用户 ID',
  `system_post_id` bigint NOT NULL COMMENT '职位 id',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:user_post` (`system_user_id`,`system_post_id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`system_post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户职位';

-- ----------------------------
-- Records of system_user_post
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_user_role
-- ----------------------------
DROP TABLE IF EXISTS `system_user_role`;
CREATE TABLE `system_user_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `system_user_id` bigint NOT NULL COMMENT '用户编号',
  `system_role_id` bigint NOT NULL COMMENT '角色编号',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `system_tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:user_role` (`system_user_id`,`system_role_id`),
  KEY `idx:list` (`system_tenant_id`,`deleted`,`system_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户和系统角色关联表';

-- ----------------------------
-- Records of system_user_role
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `nickname` varchar(128) DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `birthday` date DEFAULT NULL COMMENT '用户生日',
  `gender` tinyint DEFAULT '0' COMMENT '用户性别(0女/1男)',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '用户状态(0正常/1锁定)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

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
-- Records of user_auth
-- ----------------------------
BEGIN;
COMMIT;

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
-- Records of user_secret
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
