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

 Date: 29/02/2024 14:11:17
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典数据表';

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
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典类型';

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
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:search` (`username`,`login_time`,`channel`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='登录日志';

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
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统菜单';

-- ----------------------------
-- Records of system_menu
-- ----------------------------
BEGIN;
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (1, '首页', 'home', 2, 1, 0, '/home/index', 'HomeFilled', '/home/index', 'home', 0, 0, '', 1, 1, '', 0, '', 'admin', '2023-09-26 13:44:46', 'admin', '2024-02-27 23:48:11', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (2, '系统设置', '', 1, 2, 0, '/system', 'Tools', '', '', 0, 0, '', 1, 0, '', 0, '/system/menu', 'admin', '2023-08-08 15:17:56', 'admin', '2024-02-28 09:15:06', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (3, '菜单管理', 'menu.SystemMenuList', 2, 21, 2, '/system/menu', 'Menu', '/system/menu/index', 'systemMenu', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 20:02:04', '', '2024-02-25 13:57:44', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (4, '菜单新增', 'menu.SystemMenuCreate', 3, 1, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:22:08', '', '2024-02-25 08:22:08', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (5, '菜单编辑', 'menu.SystemMenuUpdate', 3, 2, 3, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-12 13:30:26', 'admin', '2023-05-23 04:57:14', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (6, '菜单删除', 'menu.SystemMenuDelete', 3, 3, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2023-07-12 13:29:55', 'admin', '2023-05-23 14:38:18', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (7, '菜单恢复', 'menu.SystemMenuRecover', 3, 4, 3, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2023-07-12 13:28:58', '', '1970-01-01 00:00:00', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (8, '字典管理', 'dict.SystemDictTypeList', 2, 22, 2, '/system/dict', 'Reading', '/system/dict/dict_type', 'systemDictType', 0, 0, '', 1, 0, '', 0, '', '', '2024-02-25 13:56:23', '', '2024-02-25 13:57:53', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (10, '字典添加', 'dict.SystemDictTypeCreate', 3, 1, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:23:52', '', '2024-02-25 08:23:52', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (11, '字典修改', 'dict.SystemDictTypeUpdate', 3, 2, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:24:25', '', '2024-02-25 08:24:25', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (12, '字典删除', 'dict.SystemDictTypeDelete', 3, 3, 8, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 16:24:47', '', '2024-02-25 08:24:47', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (13, '字典数据', 'dict.SystemDictList', 2, 4, 8, '/system/dict/:type/data', 'Document', '/system/dict/dict_data', 'systemDict', 0, 1, '', 1, 0, '/system/dict', 0, '', '', '2024-02-26 19:01:23', '', '2024-02-25 19:08:05', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (14, '字典数据添加', 'dict.SystemDictCreate', 3, 1, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 19:02:29', '', '2024-02-25 11:02:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (15, '字典数据修改', 'dict.SystemDictUpdate', 3, 2, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 19:02:58', '', '2024-02-25 11:02:58', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (16, '字典数据删除', 'dict.SystemDictDelete', 3, 3, 13, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-25 19:03:19', '', '2024-02-25 11:03:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (17, '角色管理', 'role.SystemRoleList', 2, 23, 2, '/system/role', 'Stamp', '/system/role/index', 'systemRole', 0, 0, '', 1, 0, '', 0, '', '', '2024-02-27 19:01:09', '', '2024-02-27 15:59:59', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (18, '角色添加', 'role.SystemRoleCreate', 3, 1, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 11:02:05', '', '2024-02-27 03:02:05', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (19, '角色修改', 'role.SystemRoleUpdate', 3, 2, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 11:02:29', '', '2024-02-27 03:02:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (20, '角色删除', 'role.SystemRoleDelete', 3, 3, 17, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 11:02:49', '', '2024-02-27 03:02:49', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (21, '用户管理', 'user.SystemUserList', 2, 20, 2, '/system/user', 'UserFilled', '/system/user/index', 'systemUser', 0, 0, '', 1, 0, '', 0, '', '', '2024-02-27 17:06:34', '', '2024-02-27 09:06:34', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (22, '用户添加', 'user.SystemUserCreate', 3, 1, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 17:07:46', '', '2024-02-27 09:07:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (23, '用户修改', 'user.SystemUserUpdate', 3, 2, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-28 01:08:08', '', '2024-02-27 17:08:34', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (24, '用户删除', 'user.SystemUserDelete', 3, 3, 21, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', '', '2024-02-27 17:08:26', '', '2024-02-27 09:08:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (68, '审计日志', '', 1, 3, 0, '/system/logger', 'Tickets', '', '', 0, 0, '', 1, 0, '', 0, NULL, 'admin', '2023-08-09 22:03:11', 'admin', '2024-02-27 23:22:48', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (69, '操作日志', 'logger.SystemOperateLogList', 2, 3, 68, '/system/logger/operate', 'DocumentCopy', '/system/logger/operate', 'systemLoggerOperate', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 18:56:50', 'admin', '2024-02-28 17:54:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (70, '登录日志', 'logger.SystemLoginLogList', 2, 2, 68, '/system/logger/login', 'FolderRemove', '/system/logger/login', 'systemLoggerLogin', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 09:53:22', 'admin', '2024-02-28 17:52:37', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (71, '操作日志删除', 'logger.SystemOperateLogDelete', 3, 1, 69, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 10:52:07', 'admin', '2024-02-28 17:54:40', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (72, '操作日志清空', 'logger.SystemOperateLogDrop', 3, 1, 69, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 10:52:32', 'admin', '2024-02-28 17:54:59', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (73, '操作日志查看', 'logger.SystemOperateLog', 3, 1, 69, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 10:53:02', 'admin', '2024-02-28 17:55:15', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (74, '登录日志删除', 'logger.SystemLoginLogDelete', 3, 1, 70, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 17:54:11', 'admin', '2024-02-28 17:53:07', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (75, '登录日志清空', 'logger.SystemLoginLogDrop', 3, 1, 70, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 17:54:39', 'admin', '2024-02-28 17:53:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (76, '登录日志查看', 'logger.SystemLoginLog', 3, 1, 70, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-17 17:55:02', 'admin', '2024-02-28 17:53:47', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (77, '系统日志', 'logger.SystemEntryLogList', 2, 1, 68, '/system/logger/entry', 'DocumentRemove', '/system/logger/entry', 'systemLoggerEntry', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 19:14:50', 'admin', '2024-02-28 17:50:30', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (78, '系统日志删除', 'logger.SystemEntryLogDelete', 3, 1, 77, '', '', '', '', 0, 0, '', 1, 0, '', 0, '', 'admin', '2023-07-18 14:28:41', 'admin', '2024-02-28 17:50:48', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (79, '系统日志查看', 'logger.SystemEntryLogDrop', 3, 0, 77, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-02-28 17:51:16', '', '2024-02-28 09:51:16', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (80, '系统日志清空', 'logger.SystemEntryLogDrop', 3, 0, 77, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-02-28 17:51:38', '', '2024-02-28 09:51:38', 0);
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
  `result` tinyint NOT NULL COMMENT '结果(0 成功/1 失败)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:log` (`username`,`module`,`start_time`,`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志';

-- ----------------------------
-- Records of system_operate_log
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
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '角色状态（0正常 1停用）',
  `type` tinyint NOT NULL COMMENT '角色类型(1内置/2定义)',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`type`,`status`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色';

-- ----------------------------
-- Records of system_role
-- ----------------------------
BEGIN;
INSERT INTO `system_role` (`id`, `name`, `code`, `sort`, `status`, `type`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '超级管理', 'super', 1, 0, 1, '', '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role` (`id`, `name`, `code`, `sort`, `status`, `type`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (2, '测试角色', 'custom', 2, 0, 2, '', '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
COMMIT;

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
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:menu:role` (`system_menu_id`,`system_role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色和系统菜单关联表';

-- ----------------------------
-- Records of system_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (62, 2, 1, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (63, 2, 8, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (64, 2, 10, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (65, 2, 11, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (66, 2, 12, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (67, 2, 13, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (68, 2, 14, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (69, 2, 15, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (70, 2, 16, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (71, 2, 2, '', '2024-02-27 17:58:13', '', '2024-02-27 17:58:13');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (95, 1, 1, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (96, 1, 2, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (97, 1, 21, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (98, 1, 22, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (99, 1, 23, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (100, 1, 24, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (101, 1, 3, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (102, 1, 4, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (103, 1, 5, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (104, 1, 6, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (105, 1, 7, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (106, 1, 8, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (107, 1, 10, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (108, 1, 11, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (109, 1, 12, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110, 1, 13, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (111, 1, 14, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (112, 1, 15, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (113, 1, 16, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (114, 1, 17, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (115, 1, 18, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (116, 1, 19, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (117, 1, 20, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (118, 1, 68, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (119, 1, 77, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120, 1, 78, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (121, 1, 70, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (122, 1, 74, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (123, 1, 75, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (124, 1, 76, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (125, 1, 69, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (126, 1, 71, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (127, 1, 72, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (128, 1, 73, '', '2024-02-29 14:04:16', 'admin', '2024-02-27 23:23:27');
COMMIT;

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
-- Records of system_user
-- ----------------------------
BEGIN;
INSERT INTO `system_user` (`id`, `nickname`, `username`, `password`, `status`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '安安', 'admin', '21232f297a57a5a743894a0e4a801fc3', 0, NULL, '2024-02-21 17:08:42', '', '2024-02-27 18:02:43');
COMMIT;

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
  PRIMARY KEY (`id`),
  KEY `list:role` (`system_role_id`),
  KEY `list:user` (`system_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户和系统角色关联表';

-- ----------------------------
-- Records of system_user_role
-- ----------------------------
BEGIN;
INSERT INTO `system_user_role` (`id`, `system_user_id`, `system_role_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (6, 1, 2, NULL, '2024-02-21 17:08:42', '', '2024-02-27 18:02:43');
INSERT INTO `system_user_role` (`id`, `system_user_id`, `system_role_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, 1, 1, NULL, '2024-02-21 17:08:42', '', '2024-02-27 18:02:43');
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
  `gender` tinyint DEFAULT NULL COMMENT '用户性别(0女/1男)',
  `status` tinyint NOT NULL COMMENT '用户状态(0正常/1锁定)',
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
