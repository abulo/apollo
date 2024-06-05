/*
 Navicat Premium Data Transfer

 Source Server         : docker_mysql
 Source Server Type    : MySQL
 Source Server Version : 80400 (8.4.0)
 Source Host           : localhost:3306
 Source Schema         : apollo_tenant

 Target Server Type    : MySQL
 Target Server Version : 80400 (8.4.0)
 File Encoding         : 65001

 Date: 05/06/2024 12:21:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
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
-- Records of member
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for member_auth
-- ----------------------------
DROP TABLE IF EXISTS `member_auth`;
CREATE TABLE `member_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `member_id` bigint NOT NULL COMMENT '用户编号',
  `identity_type` varchar(128) DEFAULT NULL COMMENT '登录类型(手机号/邮箱) 或第三方应用名称 (微信/微博等)',
  `identifier` varchar(128) DEFAULT NULL COMMENT '手机号/邮箱/第三方的唯一标识',
  `credential` varchar(255) DEFAULT NULL COMMENT '密码凭证 (自建账号的保存密码, 第三方的保存 token)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:identity` (`identity_type`,`identifier`),
  KEY `idx:user` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户三方登录授权';

-- ----------------------------
-- Records of member_auth
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for member_secret
-- ----------------------------
DROP TABLE IF EXISTS `member_secret`;
CREATE TABLE `member_secret` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `member_id` bigint NOT NULL COMMENT '用户编号',
  `password` varchar(255) DEFAULT NULL COMMENT '用户密码',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:user` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户密码';

-- ----------------------------
-- Records of member_secret
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pay_app
-- ----------------------------
DROP TABLE IF EXISTS `pay_app`;
CREATE TABLE `pay_app` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '应用编号',
  `name` varchar(64) NOT NULL COMMENT '应用名称',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '开启状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `order_notify_url` varchar(255) NOT NULL COMMENT '支付结果的回调地址',
  `refund_notify_url` varchar(255) NOT NULL COMMENT '退款结果的回调地址',
  `transfer_notify_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '转账结果的回调地址',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`status`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付应用信息';

-- ----------------------------
-- Records of pay_app
-- ----------------------------
BEGIN;
INSERT INTO `pay_app` (`id`, `name`, `status`, `remark`, `order_notify_url`, `refund_notify_url`, `transfer_notify_url`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '商城应用', 0, NULL, 'http://127.0.0.1:48080/app-api/trade/order/update-paid', 'http://127.0.0.1:48080/admin-api/trade/after-sale/update-refunded', '', 0, 1, '1', '2023-02-11 21:20:54', '1', '2023-08-18 14:54:15');
INSERT INTO `pay_app` (`id`, `name`, `status`, `remark`, `order_notify_url`, `refund_notify_url`, `transfer_notify_url`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (7, '示例应用', 0, NULL, 'http://127.0.0.1:48080/admin-api/pay/demo-order/update-paid', 'http://127.0.0.1:48080/admin-api/pay/demo-order/update-refunded', '', 0, 1, '1', '2023-02-11 21:20:54', '1', '2023-07-17 16:37:28');
INSERT INTO `pay_app` (`id`, `name`, `status`, `remark`, `order_notify_url`, `refund_notify_url`, `transfer_notify_url`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (8, '会员钱包', 0, NULL, 'http://127.0.0.1:48080/admin-api/pay/wallet-recharge/update-paid', 'http://127.0.0.1:48080/admin-api/pay/wallet-recharge/update-refunded', '', 0, 1, '1', '2023-09-30 18:43:09', '1', '2023-09-30 18:51:27');
COMMIT;

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel`;
CREATE TABLE `pay_channel` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '商户编号',
  `code` varchar(32) NOT NULL COMMENT '渠道编码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '开启状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `fee_rate` double NOT NULL DEFAULT '0' COMMENT '渠道费率，单位：百分比',
  `app_id` bigint NOT NULL COMMENT '应用编号',
  `config` json DEFAULT NULL COMMENT '支付渠道配置',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:code` (`tenant_id`,`app_id`,`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付渠道';

-- ----------------------------
-- Records of pay_channel
-- ----------------------------
BEGIN;
INSERT INTO `pay_channel` (`id`, `code`, `status`, `remark`, `fee_rate`, `app_id`, `config`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, 'mock', 0, '1', 1, 1, NULL, 0, 1, 'admin', '2024-05-15 22:24:25', 'admin', '2024-05-15 22:30:02');
COMMIT;

-- ----------------------------
-- Table structure for pay_notify_log
-- ----------------------------
DROP TABLE IF EXISTS `pay_notify_log`;
CREATE TABLE `pay_notify_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `task_id` bigint NOT NULL COMMENT '通知任务编号',
  `notify_times` tinyint NOT NULL DEFAULT '0' COMMENT '第几次被通知',
  `response` json NOT NULL COMMENT '请求参数',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '通知状态',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`status`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付通知日志';

-- ----------------------------
-- Records of pay_notify_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pay_notify_task
-- ----------------------------
DROP TABLE IF EXISTS `pay_notify_task`;
CREATE TABLE `pay_notify_task` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务编号',
  `app_id` bigint NOT NULL COMMENT '应用编号',
  `type` tinyint NOT NULL COMMENT '通知类型',
  `data_id` bigint NOT NULL COMMENT '数据编号',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '通知状态',
  `merchant_order_id` varchar(64) NOT NULL COMMENT '商户订单编号',
  `merchant_transfer_id` varchar(64) NOT NULL COMMENT '商户转账单编号',
  `next_notify_time` datetime NOT NULL COMMENT '下一次通知时间',
  `last_execute_time` datetime NOT NULL COMMENT '最后一次执行时间',
  `notify_times` tinyint NOT NULL DEFAULT '0' COMMENT '当前通知次数',
  `max_notify_times` tinyint DEFAULT '0' COMMENT '最大可通知次数',
  `notify_url` varchar(255) NOT NULL COMMENT '异步通知地址',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`app_id`,`type`,`data_id`,`status`,`merchant_order_id`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商户支付-任务通知';

-- ----------------------------
-- Records of pay_notify_task
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS `pay_order`;
CREATE TABLE `pay_order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '支付订单编号',
  `app_id` bigint NOT NULL COMMENT '应用编号',
  `channel_id` bigint NOT NULL COMMENT '渠道编号',
  `channel_code` varchar(32) NOT NULL COMMENT '渠道编码',
  `merchant_order_id` varchar(64) NOT NULL COMMENT '商户订单编号',
  `subject` varchar(64) NOT NULL COMMENT '商品标题',
  `body` varchar(128) NOT NULL COMMENT '商品描述',
  `notify_url` varchar(255) NOT NULL COMMENT '异步通知地址',
  `price` bigint NOT NULL COMMENT '支付金额，单位：分',
  `channel_fee_rate` double NOT NULL DEFAULT '0' COMMENT '渠道手续费，单位：百分比',
  `channel_fee_price` bigint NOT NULL DEFAULT '0' COMMENT '渠道手续金额，单位：分',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态',
  `user_ip` varchar(64) DEFAULT NULL COMMENT '用户 IP',
  `expire_time` datetime NOT NULL COMMENT '订单失效时间',
  `success_time` datetime DEFAULT NULL COMMENT '订单支付成功时间',
  `extension_id` bigint DEFAULT NULL COMMENT '支付成功的订单拓展单编号',
  `no` varchar(64) DEFAULT NULL COMMENT '支付订单号',
  `refund_price` bigint NOT NULL COMMENT '退款总金额，单位：分',
  `channel_user_id` varchar(255) DEFAULT NULL COMMENT '渠道用户编号',
  `channel_order_no` varchar(255) DEFAULT NULL COMMENT '渠道订单号',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`app_id`,`channel_id`,`channel_code`,`merchant_order_id`,`no`,`channel_order_no`,`status`,`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付订单';

-- ----------------------------
-- Records of pay_order
-- ----------------------------
BEGIN;
INSERT INTO `pay_order` (`id`, `app_id`, `channel_id`, `channel_code`, `merchant_order_id`, `subject`, `body`, `notify_url`, `price`, `channel_fee_rate`, `channel_fee_price`, `status`, `user_ip`, `expire_time`, `success_time`, `extension_id`, `no`, `refund_price`, `channel_user_id`, `channel_order_no`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, 1, 1, 'mock', '1', '1', '1', '1', 1, 1, 1, 0, '1', '2024-05-29 21:28:10', '2024-05-29 21:28:13', 1, '1', 1, '1', '1', 0, 1, '1', '2024-05-29 21:28:32', '1', '2024-05-29 21:28:32');
COMMIT;

-- ----------------------------
-- Table structure for pay_order_extension
-- ----------------------------
DROP TABLE IF EXISTS `pay_order_extension`;
CREATE TABLE `pay_order_extension` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `no` varchar(64) NOT NULL COMMENT '支付订单号',
  `order_id` bigint NOT NULL COMMENT '支付订单编号',
  `channel_id` bigint NOT NULL COMMENT '渠道编号',
  `channel_code` varchar(32) NOT NULL COMMENT '渠道编码',
  `user_ip` varchar(64) DEFAULT NULL COMMENT 'ip',
  `status` tinyint NOT NULL COMMENT '支付状态',
  `channel_extras` varchar(255) DEFAULT NULL COMMENT '支付渠道的额外参数',
  `channel_error_code` varchar(255) DEFAULT NULL COMMENT '渠道调用报错时，错误码',
  `channel_error_msg` varchar(255) DEFAULT NULL COMMENT '渠道调用报错时，错误信息',
  `channel_notify_data` varchar(1024) DEFAULT NULL COMMENT '支付渠道异步通知的内容',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `item:item` (`tenant_id`,`deleted`,`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付订单拓展';

-- ----------------------------
-- Records of pay_order_extension
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pay_refund
-- ----------------------------
DROP TABLE IF EXISTS `pay_refund`;
CREATE TABLE `pay_refund` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '支付退款编号',
  `no` varchar(64) NOT NULL COMMENT '退款单号',
  `app_id` bigint NOT NULL COMMENT '应用编号',
  `channel_id` bigint NOT NULL COMMENT '渠道编号',
  `channel_code` varchar(64) NOT NULL COMMENT '渠道编码',
  `order_id` bigint NOT NULL COMMENT '支付订单编号 pay_order 表id',
  `order_no` varchar(64) NOT NULL COMMENT '支付订单 no',
  `merchant_order_id` varchar(64) NOT NULL COMMENT '商户订单编号（商户系统生成）',
  `merchant_refund_id` varchar(64) NOT NULL COMMENT '商户退款订单号（商户系统生成）',
  `notify_url` varchar(255) NOT NULL COMMENT '异步通知商户地址',
  `status` tinyint NOT NULL COMMENT '退款状态',
  `pay_price` bigint NOT NULL COMMENT '支付金额,单位分',
  `refund_price` bigint NOT NULL COMMENT '退款金额,单位分',
  `reason` varchar(255) NOT NULL COMMENT '退款原因',
  `user_ip` varchar(64) DEFAULT NULL COMMENT 'ip',
  `channel_order_no` varchar(64) NOT NULL COMMENT '渠道订单号，pay_order 中的 channel_order_no 对应',
  `channel_refund_no` varchar(64) DEFAULT NULL COMMENT '渠道退款单号，渠道返',
  `success_time` datetime DEFAULT NULL COMMENT '退款成功时间',
  `channel_error_code` varchar(255) DEFAULT NULL COMMENT '渠道调用报错时，错误码',
  `channel_error_msg` varchar(255) DEFAULT NULL COMMENT '渠道调用报错时，错误信息',
  `channel_notify_data` varchar(255) DEFAULT NULL COMMENT '支付渠道异步通知的内容',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`app_id`,`channel_id`,`channel_code`,`merchant_order_id`,`order_id`,`order_no`,`channel_order_no`,`status`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='退款订单';

-- ----------------------------
-- Records of pay_refund
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for region
-- ----------------------------
DROP TABLE IF EXISTS `region`;
CREATE TABLE `region` (
  `id` bigint NOT NULL COMMENT '区域编号',
  `name` varchar(255) NOT NULL COMMENT '区域名称',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父级编号',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`status`,`name`,`parent_id`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='地区表';

-- ----------------------------
-- Records of region
-- ----------------------------
BEGIN;
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110000, '北京市', 0, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110101, '东城区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110102, '西城区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110105, '朝阳区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110106, '丰台区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110107, '石景山区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110108, '海淀区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110109, '门头沟区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110111, '房山区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110112, '通州区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110113, '顺义区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110114, '昌平区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110115, '大兴区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110116, '怀柔区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110117, '平谷区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110118, '密云区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (110119, '延庆区', 110000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120000, '天津市', 0, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120101, '和平区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120102, '河东区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120103, '河西区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120104, '南开区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120105, '河北区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120106, '红桥区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120110, '东丽区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120111, '西青区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120112, '津南区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120113, '北辰区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120114, '武清区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120115, '宝坻区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120116, '滨海新区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120117, '宁河区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120118, '静海区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (120119, '蓟州区', 120000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130000, '河北省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130100, '石家庄市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130102, '长安区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130104, '桥西区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130105, '新华区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130107, '井陉矿区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130108, '裕华区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130109, '藁城区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130110, '鹿泉区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130111, '栾城区', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130121, '井陉县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130123, '正定县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130125, '行唐县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130126, '灵寿县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130127, '高邑县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130128, '深泽县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130129, '赞皇县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130130, '无极县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130131, '平山县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130132, '元氏县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130133, '赵县', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130181, '辛集市', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130183, '晋州市', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130184, '新乐市', 130100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130200, '唐山市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130202, '路南区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130203, '路北区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130204, '古冶区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130205, '开平区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130207, '丰南区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130208, '丰润区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130209, '曹妃甸区', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130224, '滦南县', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130225, '乐亭县', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130227, '迁西县', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130229, '玉田县', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130281, '遵化市', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130283, '迁安市', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130284, '滦州市', 130200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130300, '秦皇岛市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130302, '海港区', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130303, '山海关区', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130304, '北戴河区', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130306, '抚宁区', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130321, '青龙满族自治县', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130322, '昌黎县', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130324, '卢龙县', 130300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130400, '邯郸市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130402, '邯山区', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130403, '丛台区', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130404, '复兴区', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130406, '峰峰矿区', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130407, '肥乡区', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130408, '永年区', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130423, '临漳县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130424, '成安县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130425, '大名县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130426, '涉县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130427, '磁县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130430, '邱县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130431, '鸡泽县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130432, '广平县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130433, '馆陶县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130434, '魏县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130435, '曲周县', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130481, '武安市', 130400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130500, '邢台市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130502, '襄都区', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130503, '信都区', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130505, '任泽区', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130506, '南和区', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130522, '临城县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130523, '内丘县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130524, '柏乡县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130525, '隆尧县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130528, '宁晋县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130529, '巨鹿县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130530, '新河县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130531, '广宗县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130532, '平乡县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130533, '威县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130534, '清河县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130535, '临西县', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130581, '南宫市', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130582, '沙河市', 130500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130600, '保定市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130602, '竞秀区', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130606, '莲池区', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130607, '满城区', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130608, '清苑区', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130609, '徐水区', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130623, '涞水县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130624, '阜平县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130626, '定兴县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130627, '唐县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130628, '高阳县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130630, '涞源县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130631, '望都县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130633, '易县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130634, '曲阳县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130635, '蠡县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130636, '顺平县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130637, '博野县', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130681, '涿州市', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130682, '定州市', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130683, '安国市', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130684, '高碑店市', 130600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130700, '张家口市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130702, '桥东区', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130703, '桥西区', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130705, '宣化区', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130706, '下花园区', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130708, '万全区', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130709, '崇礼区', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130722, '张北县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130723, '康保县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130724, '沽源县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130725, '尚义县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130726, '蔚县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130727, '阳原县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130728, '怀安县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130730, '怀来县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130731, '涿鹿县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130732, '赤城县', 130700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130800, '承德市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130802, '双桥区', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130803, '双滦区', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130804, '鹰手营子矿区', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130821, '承德县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130822, '兴隆县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130824, '滦平县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130825, '隆化县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130826, '丰宁满族自治县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130827, '宽城满族自治县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130828, '围场满族蒙古族自治县', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130881, '平泉市', 130800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130900, '沧州市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130902, '新华区', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130903, '运河区', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130921, '沧县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130922, '青县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130923, '东光县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130924, '海兴县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130925, '盐山县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130926, '肃宁县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130927, '南皮县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130928, '吴桥县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130929, '献县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130930, '孟村回族自治县', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130981, '泊头市', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130982, '任丘市', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130983, '黄骅市', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (130984, '河间市', 130900, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131000, '廊坊市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131002, '安次区', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131003, '广阳区', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131022, '固安县', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131023, '永清县', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131024, '香河县', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131025, '大城县', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131026, '文安县', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131028, '大厂回族自治县', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131081, '霸州市', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131082, '三河市', 131000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131100, '衡水市', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131102, '桃城区', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131103, '冀州区', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131121, '枣强县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131122, '武邑县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131123, '武强县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131124, '饶阳县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131125, '安平县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131126, '故城县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131127, '景县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131128, '阜城县', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (131182, '深州市', 131100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (133100, '雄安新区', 130000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140000, '山西省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140100, '太原市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140105, '小店区', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140106, '迎泽区', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140107, '杏花岭区', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140108, '尖草坪区', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140109, '万柏林区', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140110, '晋源区', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140121, '清徐县', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140122, '阳曲县', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140123, '娄烦县', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140181, '古交市', 140100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140200, '大同市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140212, '新荣区', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140213, '平城区', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140214, '云冈区', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140215, '云州区', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140221, '阳高县', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140222, '天镇县', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140223, '广灵县', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140224, '灵丘县', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140225, '浑源县', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140226, '左云县', 140200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140300, '阳泉市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140302, '城区', 140300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140303, '矿区', 140300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140311, '郊区', 140300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140321, '平定县', 140300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140322, '盂县', 140300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140400, '长治市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140403, '潞州区', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140404, '上党区', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140405, '屯留区', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140406, '潞城区', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140423, '襄垣县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140425, '平顺县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140426, '黎城县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140427, '壶关县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140428, '长子县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140429, '武乡县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140430, '沁县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140431, '沁源县', 140400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140500, '晋城市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140502, '城区', 140500, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140521, '沁水县', 140500, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140522, '阳城县', 140500, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140524, '陵川县', 140500, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140525, '泽州县', 140500, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140581, '高平市', 140500, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140600, '朔州市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140602, '朔城区', 140600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140603, '平鲁区', 140600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140621, '山阴县', 140600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140622, '应县', 140600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140623, '右玉县', 140600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140681, '怀仁市', 140600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140700, '晋中市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140702, '榆次区', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140703, '太谷区', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140721, '榆社县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140722, '左权县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140723, '和顺县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140724, '昔阳县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140725, '寿阳县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140727, '祁县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140728, '平遥县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140729, '灵石县', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140781, '介休市', 140700, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140800, '运城市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140802, '盐湖区', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140821, '临猗县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140822, '万荣县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140823, '闻喜县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140824, '稷山县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140825, '新绛县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140826, '绛县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140827, '垣曲县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140828, '夏县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140829, '平陆县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140830, '芮城县', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140881, '永济市', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140882, '河津市', 140800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140900, '忻州市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140902, '忻府区', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140921, '定襄县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140922, '五台县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140923, '代县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140924, '繁峙县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140925, '宁武县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140926, '静乐县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140927, '神池县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140928, '五寨县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140929, '岢岚县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140930, '河曲县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140931, '保德县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140932, '偏关县', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (140981, '原平市', 140900, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141000, '临汾市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141002, '尧都区', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141021, '曲沃县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141022, '翼城县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141023, '襄汾县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141024, '洪洞县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141025, '古县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141026, '安泽县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141027, '浮山县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141028, '吉县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141029, '乡宁县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141030, '大宁县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141031, '隰县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141032, '永和县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141033, '蒲县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141034, '汾西县', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141081, '侯马市', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141082, '霍州市', 141000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141100, '吕梁市', 140000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141102, '离石区', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141121, '文水县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141122, '交城县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141123, '兴县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141124, '临县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141125, '柳林县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141126, '石楼县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141127, '岚县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141128, '方山县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141129, '中阳县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141130, '交口县', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141181, '孝义市', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (141182, '汾阳市', 141100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150000, '内蒙古自治区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150100, '呼和浩特市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150102, '新城区', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150103, '回民区', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150104, '玉泉区', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150105, '赛罕区', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150121, '土默特左旗', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150122, '托克托县', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150123, '和林格尔县', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150124, '清水河县', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150125, '武川县', 150100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150200, '包头市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150202, '东河区', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150203, '昆都仑区', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150204, '青山区', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150205, '石拐区', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150206, '白云鄂博矿区', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150207, '九原区', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150221, '土默特右旗', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150222, '固阳县', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150223, '达尔罕茂明安联合旗', 150200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150300, '乌海市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150302, '海勃湾区', 150300, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150303, '海南区', 150300, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150304, '乌达区', 150300, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150400, '赤峰市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150402, '红山区', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150403, '元宝山区', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150404, '松山区', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150421, '阿鲁科尔沁旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150422, '巴林左旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150423, '巴林右旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150424, '林西县', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150425, '克什克腾旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150426, '翁牛特旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150428, '喀喇沁旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150429, '宁城县', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150430, '敖汉旗', 150400, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150500, '通辽市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150502, '科尔沁区', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150521, '科尔沁左翼中旗', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150522, '科尔沁左翼后旗', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150523, '开鲁县', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150524, '库伦旗', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150525, '奈曼旗', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150526, '扎鲁特旗', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150581, '霍林郭勒市', 150500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150600, '鄂尔多斯市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150602, '东胜区', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150603, '康巴什区', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150621, '达拉特旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150622, '准格尔旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150623, '鄂托克前旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150624, '鄂托克旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150625, '杭锦旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150626, '乌审旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150627, '伊金霍洛旗', 150600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150700, '呼伦贝尔市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150702, '海拉尔区', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150703, '扎赉诺尔区', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150721, '阿荣旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150722, '莫力达瓦达斡尔族自治旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150723, '鄂伦春自治旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150724, '鄂温克族自治旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150725, '陈巴尔虎旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150726, '新巴尔虎左旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150727, '新巴尔虎右旗', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150781, '满洲里市', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150782, '牙克石市', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150783, '扎兰屯市', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150784, '额尔古纳市', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150785, '根河市', 150700, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150800, '巴彦淖尔市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150802, '临河区', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150821, '五原县', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150822, '磴口县', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150823, '乌拉特前旗', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150824, '乌拉特中旗', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150825, '乌拉特后旗', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150826, '杭锦后旗', 150800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150900, '乌兰察布市', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150902, '集宁区', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150921, '卓资县', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150922, '化德县', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150923, '商都县', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150924, '兴和县', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150925, '凉城县', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150926, '察哈尔右翼前旗', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150927, '察哈尔右翼中旗', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150928, '察哈尔右翼后旗', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150929, '四子王旗', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (150981, '丰镇市', 150900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152200, '兴安盟', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152201, '乌兰浩特市', 152200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152202, '阿尔山市', 152200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152221, '科尔沁右翼前旗', 152200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152222, '科尔沁右翼中旗', 152200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152223, '扎赉特旗', 152200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152224, '突泉县', 152200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152500, '锡林郭勒盟', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152501, '二连浩特市', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152502, '锡林浩特市', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152522, '阿巴嘎旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152523, '苏尼特左旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152524, '苏尼特右旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152525, '东乌珠穆沁旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152526, '西乌珠穆沁旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152527, '太仆寺旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152528, '镶黄旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152529, '正镶白旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152530, '正蓝旗', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152531, '多伦县', 152500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152900, '阿拉善盟', 150000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152921, '阿拉善左旗', 152900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152922, '阿拉善右旗', 152900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (152923, '额济纳旗', 152900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210000, '辽宁省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210100, '沈阳市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210102, '和平区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210103, '沈河区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210104, '大东区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210105, '皇姑区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210106, '铁西区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210111, '苏家屯区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210112, '浑南区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210113, '沈北新区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210114, '于洪区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210115, '辽中区', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210123, '康平县', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210124, '法库县', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210181, '新民市', 210100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210200, '大连市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210202, '中山区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210203, '西岗区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210204, '沙河口区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210211, '甘井子区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210212, '旅顺口区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210213, '金州区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210214, '普兰店区', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210224, '长海县', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210281, '瓦房店市', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210283, '庄河市', 210200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210300, '鞍山市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210302, '铁东区', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210303, '铁西区', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210304, '立山区', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210311, '千山区', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210321, '台安县', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210323, '岫岩满族自治县', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210381, '海城市', 210300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210400, '抚顺市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210402, '新抚区', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210403, '东洲区', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210404, '望花区', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210411, '顺城区', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210421, '抚顺县', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210422, '新宾满族自治县', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210423, '清原满族自治县', 210400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210500, '本溪市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210502, '平山区', 210500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210503, '溪湖区', 210500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210504, '明山区', 210500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210505, '南芬区', 210500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210521, '本溪满族自治县', 210500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210522, '桓仁满族自治县', 210500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210600, '丹东市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210602, '元宝区', 210600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210603, '振兴区', 210600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210604, '振安区', 210600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210624, '宽甸满族自治县', 210600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210681, '东港市', 210600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210682, '凤城市', 210600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210700, '锦州市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210702, '古塔区', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210703, '凌河区', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210711, '太和区', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210726, '黑山县', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210727, '义县', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210781, '凌海市', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210782, '北镇市', 210700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210800, '营口市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210802, '站前区', 210800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210803, '西市区', 210800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210804, '鲅鱼圈区', 210800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210811, '老边区', 210800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210881, '盖州市', 210800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210882, '大石桥市', 210800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210900, '阜新市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210902, '海州区', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210903, '新邱区', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210904, '太平区', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210905, '清河门区', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210911, '细河区', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210921, '阜新蒙古族自治县', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (210922, '彰武县', 210900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211000, '辽阳市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211002, '白塔区', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211003, '文圣区', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211004, '宏伟区', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211005, '弓长岭区', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211011, '太子河区', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211021, '辽阳县', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211081, '灯塔市', 211000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211100, '盘锦市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211102, '双台子区', 211100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211103, '兴隆台区', 211100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211104, '大洼区', 211100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211122, '盘山县', 211100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211200, '铁岭市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211202, '银州区', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211204, '清河区', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211221, '铁岭县', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211223, '西丰县', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211224, '昌图县', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211281, '调兵山市', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211282, '开原市', 211200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211300, '朝阳市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211302, '双塔区', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211303, '龙城区', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211321, '朝阳县', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211322, '建平县', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211324, '喀喇沁左翼蒙古族自治县', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211381, '北票市', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211382, '凌源市', 211300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211400, '葫芦岛市', 210000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211402, '连山区', 211400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211403, '龙港区', 211400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211404, '南票区', 211400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211421, '绥中县', 211400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211422, '建昌县', 211400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (211481, '兴城市', 211400, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220000, '吉林省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220100, '长春市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220102, '南关区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220103, '宽城区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220104, '朝阳区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220105, '二道区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220106, '绿园区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220112, '双阳区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220113, '九台区', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220122, '农安县', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220182, '榆树市', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220183, '德惠市', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220184, '公主岭市', 220100, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220200, '吉林市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220202, '昌邑区', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220203, '龙潭区', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220204, '船营区', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220211, '丰满区', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220221, '永吉县', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220281, '蛟河市', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220282, '桦甸市', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220283, '舒兰市', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220284, '磐石市', 220200, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220300, '四平市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220302, '铁西区', 220300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220303, '铁东区', 220300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220322, '梨树县', 220300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220323, '伊通满族自治县', 220300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220382, '双辽市', 220300, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220400, '辽源市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220402, '龙山区', 220400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220403, '西安区', 220400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220421, '东丰县', 220400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220422, '东辽县', 220400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220500, '通化市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220502, '东昌区', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220503, '二道江区', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220521, '通化县', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220523, '辉南县', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220524, '柳河县', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220581, '梅河口市', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220582, '集安市', 220500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220600, '白山市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220602, '浑江区', 220600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220605, '江源区', 220600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220621, '抚松县', 220600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220622, '靖宇县', 220600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220623, '长白朝鲜族自治县', 220600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220681, '临江市', 220600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220700, '松原市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220702, '宁江区', 220700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220721, '前郭尔罗斯蒙古族自治县', 220700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220722, '长岭县', 220700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220723, '乾安县', 220700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220781, '扶余市', 220700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220800, '白城市', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220802, '洮北区', 220800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220821, '镇赉县', 220800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220822, '通榆县', 220800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220881, '洮南市', 220800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (220882, '大安市', 220800, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222400, '延边朝鲜族自治州', 220000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222401, '延吉市', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222402, '图们市', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222403, '敦化市', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222404, '珲春市', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222405, '龙井市', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222406, '和龙市', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222424, '汪清县', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (222426, '安图县', 222400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230000, '黑龙江省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230100, '哈尔滨市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230102, '道里区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230103, '南岗区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230104, '道外区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230108, '平房区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230109, '松北区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230110, '香坊区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230111, '呼兰区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230112, '阿城区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230113, '双城区', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230123, '依兰县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230124, '方正县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230125, '宾县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230126, '巴彦县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230127, '木兰县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230128, '通河县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230129, '延寿县', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230183, '尚志市', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230184, '五常市', 230100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230200, '齐齐哈尔市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230202, '龙沙区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230203, '建华区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230204, '铁锋区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230205, '昂昂溪区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230206, '富拉尔基区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230207, '碾子山区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230208, '梅里斯达斡尔族区', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230221, '龙江县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230223, '依安县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230224, '泰来县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230225, '甘南县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230227, '富裕县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230229, '克山县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230230, '克东县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230231, '拜泉县', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230281, '讷河市', 230200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230300, '鸡西市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230302, '鸡冠区', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230303, '恒山区', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230304, '滴道区', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230305, '梨树区', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230306, '城子河区', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230307, '麻山区', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230321, '鸡东县', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230381, '虎林市', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230382, '密山市', 230300, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230400, '鹤岗市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230402, '向阳区', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230403, '工农区', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230404, '南山区', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230405, '兴安区', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230406, '东山区', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230407, '兴山区', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230421, '萝北县', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230422, '绥滨县', 230400, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230500, '双鸭山市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230502, '尖山区', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230503, '岭东区', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230505, '四方台区', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230506, '宝山区', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230521, '集贤县', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230522, '友谊县', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230523, '宝清县', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230524, '饶河县', 230500, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230600, '大庆市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230602, '萨尔图区', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230603, '龙凤区', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230604, '让胡路区', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230605, '红岗区', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230606, '大同区', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230621, '肇州县', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230622, '肇源县', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230623, '林甸县', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230624, '杜尔伯特蒙古族自治县', 230600, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230700, '伊春市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230717, '伊美区', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230718, '乌翠区', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230719, '友好区', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230722, '嘉荫县', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230723, '汤旺县', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230724, '丰林县', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230725, '大箐山县', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230726, '南岔县', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230751, '金林区', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230781, '铁力市', 230700, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230800, '佳木斯市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230803, '向阳区', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230804, '前进区', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230805, '东风区', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230811, '郊区', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230822, '桦南县', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230826, '桦川县', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230828, '汤原县', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230881, '同江市', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230882, '富锦市', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230883, '抚远市', 230800, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230900, '七台河市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230902, '新兴区', 230900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230903, '桃山区', 230900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230904, '茄子河区', 230900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (230921, '勃利县', 230900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231000, '牡丹江市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231002, '东安区', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231003, '阳明区', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231004, '爱民区', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231005, '西安区', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231025, '林口县', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231081, '绥芬河市', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231083, '海林市', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231084, '宁安市', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231085, '穆棱市', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231086, '东宁市', 231000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231100, '黑河市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231102, '爱辉区', 231100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231123, '逊克县', 231100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231124, '孙吴县', 231100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231181, '北安市', 231100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231182, '五大连池市', 231100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231183, '嫩江市', 231100, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231200, '绥化市', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231202, '北林区', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231221, '望奎县', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231222, '兰西县', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231223, '青冈县', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231224, '庆安县', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231225, '明水县', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231226, '绥棱县', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231281, '安达市', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231282, '肇东市', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (231283, '海伦市', 231200, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (232700, '大兴安岭地区', 230000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (232701, '漠河市', 232700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (232721, '呼玛县', 232700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (232722, '塔河县', 232700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310000, '上海市', 0, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310101, '黄浦区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310104, '徐汇区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310105, '长宁区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310106, '静安区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310107, '普陀区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310109, '虹口区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310110, '杨浦区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310112, '闵行区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310113, '宝山区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310114, '嘉定区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310115, '浦东新区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310116, '金山区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310117, '松江区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310118, '青浦区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310120, '奉贤区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (310151, '崇明区', 310000, 0, 0, 0, NULL, '2024-03-27 23:31:30', NULL, '2024-03-27 23:31:30');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320000, '江苏省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320100, '南京市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320102, '玄武区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320104, '秦淮区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320105, '建邺区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320106, '鼓楼区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320111, '浦口区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320113, '栖霞区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320114, '雨花台区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320115, '江宁区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320116, '六合区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320117, '溧水区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320118, '高淳区', 320100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320200, '无锡市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320205, '锡山区', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320206, '惠山区', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320211, '滨湖区', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320213, '梁溪区', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320214, '新吴区', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320281, '江阴市', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320282, '宜兴市', 320200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320300, '徐州市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320302, '鼓楼区', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320303, '云龙区', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320305, '贾汪区', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320311, '泉山区', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320312, '铜山区', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320321, '丰县', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320322, '沛县', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320324, '睢宁县', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320381, '新沂市', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320382, '邳州市', 320300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320400, '常州市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320402, '天宁区', 320400, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320404, '钟楼区', 320400, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320411, '新北区', 320400, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320412, '武进区', 320400, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320413, '金坛区', 320400, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320481, '溧阳市', 320400, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320500, '苏州市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320505, '虎丘区', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320506, '吴中区', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320507, '相城区', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320508, '姑苏区', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320509, '吴江区', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320581, '常熟市', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320582, '张家港市', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320583, '昆山市', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320585, '太仓市', 320500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320600, '南通市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320612, '通州区', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320613, '崇川区', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320614, '海门区', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320623, '如东县', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320681, '启东市', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320682, '如皋市', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320685, '海安市', 320600, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320700, '连云港市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320703, '连云区', 320700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320706, '海州区', 320700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320707, '赣榆区', 320700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320722, '东海县', 320700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320723, '灌云县', 320700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320724, '灌南县', 320700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320800, '淮安市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320803, '淮安区', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320804, '淮阴区', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320812, '清江浦区', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320813, '洪泽区', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320826, '涟水县', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320830, '盱眙县', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320831, '金湖县', 320800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320900, '盐城市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320902, '亭湖区', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320903, '盐都区', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320904, '大丰区', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320921, '响水县', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320922, '滨海县', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320923, '阜宁县', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320924, '射阳县', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320925, '建湖县', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (320981, '东台市', 320900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321000, '扬州市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321002, '广陵区', 321000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321003, '邗江区', 321000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321012, '江都区', 321000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321023, '宝应县', 321000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321081, '仪征市', 321000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321084, '高邮市', 321000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321100, '镇江市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321102, '京口区', 321100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321111, '润州区', 321100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321112, '丹徒区', 321100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321181, '丹阳市', 321100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321182, '扬中市', 321100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321183, '句容市', 321100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321200, '泰州市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321202, '海陵区', 321200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321203, '高港区', 321200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321204, '姜堰区', 321200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321281, '兴化市', 321200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321282, '靖江市', 321200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321283, '泰兴市', 321200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321300, '宿迁市', 320000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321302, '宿城区', 321300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321311, '宿豫区', 321300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321322, '沭阳县', 321300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321323, '泗阳县', 321300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (321324, '泗洪县', 321300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330000, '浙江省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330100, '杭州市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330102, '上城区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330105, '拱墅区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330106, '西湖区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330108, '滨江区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330109, '萧山区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330110, '余杭区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330111, '富阳区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330112, '临安区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330113, '临平区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330114, '钱塘区', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330122, '桐庐县', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330127, '淳安县', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330182, '建德市', 330100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330200, '宁波市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330203, '海曙区', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330205, '江北区', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330206, '北仑区', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330211, '镇海区', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330212, '鄞州区', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330213, '奉化区', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330225, '象山县', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330226, '宁海县', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330281, '余姚市', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330282, '慈溪市', 330200, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330300, '温州市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330302, '鹿城区', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330303, '龙湾区', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330304, '瓯海区', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330305, '洞头区', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330324, '永嘉县', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330326, '平阳县', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330327, '苍南县', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330328, '文成县', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330329, '泰顺县', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330381, '瑞安市', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330382, '乐清市', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330383, '龙港市', 330300, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330400, '嘉兴市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330402, '南湖区', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330411, '秀洲区', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330421, '嘉善县', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330424, '海盐县', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330481, '海宁市', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330482, '平湖市', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330483, '桐乡市', 330400, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330500, '湖州市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330502, '吴兴区', 330500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330503, '南浔区', 330500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330521, '德清县', 330500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330522, '长兴县', 330500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330523, '安吉县', 330500, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330600, '绍兴市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330602, '越城区', 330600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330603, '柯桥区', 330600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330604, '上虞区', 330600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330624, '新昌县', 330600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330681, '诸暨市', 330600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330683, '嵊州市', 330600, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330700, '金华市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330702, '婺城区', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330703, '金东区', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330723, '武义县', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330726, '浦江县', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330727, '磐安县', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330781, '兰溪市', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330782, '义乌市', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330783, '东阳市', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330784, '永康市', 330700, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330800, '衢州市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330802, '柯城区', 330800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330803, '衢江区', 330800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330822, '常山县', 330800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330824, '开化县', 330800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330825, '龙游县', 330800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330881, '江山市', 330800, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330900, '舟山市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330902, '定海区', 330900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330903, '普陀区', 330900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330921, '岱山县', 330900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (330922, '嵊泗县', 330900, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331000, '台州市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331002, '椒江区', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331003, '黄岩区', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331004, '路桥区', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331022, '三门县', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331023, '天台县', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331024, '仙居县', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331081, '温岭市', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331082, '临海市', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331083, '玉环市', 331000, 0, 0, 0, NULL, '2024-03-27 23:31:33', NULL, '2024-03-27 23:31:33');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331100, '丽水市', 330000, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331102, '莲都区', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331121, '青田县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331122, '缙云县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331123, '遂昌县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331124, '松阳县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331125, '云和县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331126, '庆元县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331127, '景宁畲族自治县', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (331181, '龙泉市', 331100, 0, 0, 0, NULL, '2024-03-27 23:31:34', NULL, '2024-03-27 23:31:34');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340000, '安徽省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340100, '合肥市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340102, '瑶海区', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340103, '庐阳区', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340104, '蜀山区', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340111, '包河区', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340121, '长丰县', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340122, '肥东县', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340123, '肥西县', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340124, '庐江县', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340181, '巢湖市', 340100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340200, '芜湖市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340202, '镜湖区', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340207, '鸠江区', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340209, '弋江区', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340210, '湾沚区', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340212, '繁昌区', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340223, '南陵县', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340281, '无为市', 340200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340300, '蚌埠市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340302, '龙子湖区', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340303, '蚌山区', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340304, '禹会区', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340311, '淮上区', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340321, '怀远县', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340322, '五河县', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340323, '固镇县', 340300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340400, '淮南市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340402, '大通区', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340403, '田家庵区', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340404, '谢家集区', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340405, '八公山区', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340406, '潘集区', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340421, '凤台县', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340422, '寿县', 340400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340500, '马鞍山市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340503, '花山区', 340500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340504, '雨山区', 340500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340506, '博望区', 340500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340521, '当涂县', 340500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340522, '含山县', 340500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340523, '和县', 340500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340600, '淮北市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340602, '杜集区', 340600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340603, '相山区', 340600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340604, '烈山区', 340600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340621, '濉溪县', 340600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340700, '铜陵市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340705, '铜官区', 340700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340706, '义安区', 340700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340711, '郊区', 340700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340722, '枞阳县', 340700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340800, '安庆市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340802, '迎江区', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340803, '大观区', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340811, '宜秀区', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340822, '怀宁县', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340825, '太湖县', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340826, '宿松县', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340827, '望江县', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340828, '岳西县', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340881, '桐城市', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (340882, '潜山市', 340800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341000, '黄山市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341002, '屯溪区', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341003, '黄山区', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341004, '徽州区', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341021, '歙县', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341022, '休宁县', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341023, '黟县', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341024, '祁门县', 341000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341100, '滁州市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341102, '琅琊区', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341103, '南谯区', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341122, '来安县', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341124, '全椒县', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341125, '定远县', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341126, '凤阳县', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341181, '天长市', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341182, '明光市', 341100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341200, '阜阳市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341202, '颍州区', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341203, '颍东区', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341204, '颍泉区', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341221, '临泉县', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341222, '太和县', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341225, '阜南县', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341226, '颍上县', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341282, '界首市', 341200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341300, '宿州市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341302, '埇桥区', 341300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341321, '砀山县', 341300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341322, '萧县', 341300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341323, '灵璧县', 341300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341324, '泗县', 341300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341500, '六安市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341502, '金安区', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341503, '裕安区', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341504, '叶集区', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341522, '霍邱县', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341523, '舒城县', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341524, '金寨县', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341525, '霍山县', 341500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341600, '亳州市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341602, '谯城区', 341600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341621, '涡阳县', 341600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341622, '蒙城县', 341600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341623, '利辛县', 341600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341700, '池州市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341702, '贵池区', 341700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341721, '东至县', 341700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341722, '石台县', 341700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341723, '青阳县', 341700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341800, '宣城市', 340000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341802, '宣州区', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341821, '郎溪县', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341823, '泾县', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341824, '绩溪县', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341825, '旌德县', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341881, '宁国市', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (341882, '广德市', 341800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350000, '福建省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350100, '福州市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350102, '鼓楼区', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350103, '台江区', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350104, '仓山区', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350105, '马尾区', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350111, '晋安区', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350112, '长乐区', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350121, '闽侯县', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350122, '连江县', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350123, '罗源县', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350124, '闽清县', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350125, '永泰县', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350128, '平潭县', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350181, '福清市', 350100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350200, '厦门市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350203, '思明区', 350200, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350205, '海沧区', 350200, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350206, '湖里区', 350200, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350211, '集美区', 350200, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350212, '同安区', 350200, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350213, '翔安区', 350200, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350300, '莆田市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350302, '城厢区', 350300, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350303, '涵江区', 350300, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350304, '荔城区', 350300, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350305, '秀屿区', 350300, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350322, '仙游县', 350300, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350400, '三明市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350404, '三元区', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350405, '沙县区', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350421, '明溪县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350423, '清流县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350424, '宁化县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350425, '大田县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350426, '尤溪县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350428, '将乐县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350429, '泰宁县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350430, '建宁县', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350481, '永安市', 350400, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350500, '泉州市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350502, '鲤城区', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350503, '丰泽区', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350504, '洛江区', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350505, '泉港区', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350521, '惠安县', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350524, '安溪县', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350525, '永春县', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350526, '德化县', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350527, '金门县', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350581, '石狮市', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350582, '晋江市', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350583, '南安市', 350500, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350600, '漳州市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350602, '芗城区', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350603, '龙文区', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350604, '龙海区', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350605, '长泰区', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350622, '云霄县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350623, '漳浦县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350624, '诏安县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350626, '东山县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350627, '南靖县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350628, '平和县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350629, '华安县', 350600, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350700, '南平市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350702, '延平区', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350703, '建阳区', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350721, '顺昌县', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350722, '浦城县', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350723, '光泽县', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350724, '松溪县', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350725, '政和县', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350781, '邵武市', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350782, '武夷山市', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350783, '建瓯市', 350700, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350800, '龙岩市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350802, '新罗区', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350803, '永定区', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350821, '长汀县', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350823, '上杭县', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350824, '武平县', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350825, '连城县', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350881, '漳平市', 350800, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350900, '宁德市', 350000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350902, '蕉城区', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350921, '霞浦县', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350922, '古田县', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350923, '屏南县', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350924, '寿宁县', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350925, '周宁县', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350926, '柘荣县', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350981, '福安市', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (350982, '福鼎市', 350900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360000, '江西省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360100, '南昌市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360102, '东湖区', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360103, '西湖区', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360104, '青云谱区', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360111, '青山湖区', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360112, '新建区', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360113, '红谷滩区', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360121, '南昌县', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360123, '安义县', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360124, '进贤县', 360100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360200, '景德镇市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360202, '昌江区', 360200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360203, '珠山区', 360200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360222, '浮梁县', 360200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360281, '乐平市', 360200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360300, '萍乡市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360302, '安源区', 360300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360313, '湘东区', 360300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360321, '莲花县', 360300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360322, '上栗县', 360300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360323, '芦溪县', 360300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360400, '九江市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360402, '濂溪区', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360403, '浔阳区', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360404, '柴桑区', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360423, '武宁县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360424, '修水县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360425, '永修县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360426, '德安县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360428, '都昌县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360429, '湖口县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360430, '彭泽县', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360481, '瑞昌市', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360482, '共青城市', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360483, '庐山市', 360400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360500, '新余市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360502, '渝水区', 360500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360521, '分宜县', 360500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360600, '鹰潭市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360602, '月湖区', 360600, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360603, '余江区', 360600, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360681, '贵溪市', 360600, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360700, '赣州市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360702, '章贡区', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360703, '南康区', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360704, '赣县区', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360722, '信丰县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360723, '大余县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360724, '上犹县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360725, '崇义县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360726, '安远县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360728, '定南县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360729, '全南县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360730, '宁都县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360731, '于都县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360732, '兴国县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360733, '会昌县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360734, '寻乌县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360735, '石城县', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360781, '瑞金市', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360783, '龙南市', 360700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360800, '吉安市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360802, '吉州区', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360803, '青原区', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360821, '吉安县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360822, '吉水县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360823, '峡江县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360824, '新干县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360825, '永丰县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360826, '泰和县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360827, '遂川县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360828, '万安县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360829, '安福县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360830, '永新县', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360881, '井冈山市', 360800, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360900, '宜春市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360902, '袁州区', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360921, '奉新县', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360922, '万载县', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360923, '上高县', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360924, '宜丰县', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360925, '靖安县', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360926, '铜鼓县', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360981, '丰城市', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360982, '樟树市', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (360983, '高安市', 360900, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361000, '抚州市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361002, '临川区', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361003, '东乡区', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361021, '南城县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361022, '黎川县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361023, '南丰县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361024, '崇仁县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361025, '乐安县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361026, '宜黄县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361027, '金溪县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361028, '资溪县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361030, '广昌县', 361000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361100, '上饶市', 360000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361102, '信州区', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361103, '广丰区', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361104, '广信区', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361123, '玉山县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361124, '铅山县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361125, '横峰县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361126, '弋阳县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361127, '余干县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361128, '鄱阳县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361129, '万年县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361130, '婺源县', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (361181, '德兴市', 361100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370000, '山东省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370100, '济南市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370102, '历下区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370103, '市中区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370104, '槐荫区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370105, '天桥区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370112, '历城区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370113, '长清区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370114, '章丘区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370115, '济阳区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370116, '莱芜区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370117, '钢城区', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370124, '平阴县', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370126, '商河县', 370100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370200, '青岛市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370202, '市南区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370203, '市北区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370211, '黄岛区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370212, '崂山区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370213, '李沧区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370214, '城阳区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370215, '即墨区', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370281, '胶州市', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370283, '平度市', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370285, '莱西市', 370200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370300, '淄博市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370302, '淄川区', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370303, '张店区', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370304, '博山区', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370305, '临淄区', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370306, '周村区', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370321, '桓台县', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370322, '高青县', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370323, '沂源县', 370300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370400, '枣庄市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370402, '市中区', 370400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370403, '薛城区', 370400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370404, '峄城区', 370400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370405, '台儿庄区', 370400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370406, '山亭区', 370400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370481, '滕州市', 370400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370500, '东营市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370502, '东营区', 370500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370503, '河口区', 370500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370505, '垦利区', 370500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370522, '利津县', 370500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370523, '广饶县', 370500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370600, '烟台市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370602, '芝罘区', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370611, '福山区', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370612, '牟平区', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370613, '莱山区', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370614, '蓬莱区', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370681, '龙口市', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370682, '莱阳市', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370683, '莱州市', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370685, '招远市', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370686, '栖霞市', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370687, '海阳市', 370600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370700, '潍坊市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370702, '潍城区', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370703, '寒亭区', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370704, '坊子区', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370705, '奎文区', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370724, '临朐县', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370725, '昌乐县', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370781, '青州市', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370782, '诸城市', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370783, '寿光市', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370784, '安丘市', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370785, '高密市', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370786, '昌邑市', 370700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370800, '济宁市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370811, '任城区', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370812, '兖州区', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370826, '微山县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370827, '鱼台县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370828, '金乡县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370829, '嘉祥县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370830, '汶上县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370831, '泗水县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370832, '梁山县', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370881, '曲阜市', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370883, '邹城市', 370800, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370900, '泰安市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370902, '泰山区', 370900, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370911, '岱岳区', 370900, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370921, '宁阳县', 370900, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370923, '东平县', 370900, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370982, '新泰市', 370900, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (370983, '肥城市', 370900, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371000, '威海市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371002, '环翠区', 371000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371003, '文登区', 371000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371082, '荣成市', 371000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371083, '乳山市', 371000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371100, '日照市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371102, '东港区', 371100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371103, '岚山区', 371100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371121, '五莲县', 371100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371122, '莒县', 371100, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371300, '临沂市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371302, '兰山区', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371311, '罗庄区', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371312, '河东区', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371321, '沂南县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371322, '郯城县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371323, '沂水县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371324, '兰陵县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371325, '费县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371326, '平邑县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371327, '莒南县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371328, '蒙阴县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371329, '临沭县', 371300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371400, '德州市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371402, '德城区', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371403, '陵城区', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371422, '宁津县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371423, '庆云县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371424, '临邑县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371425, '齐河县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371426, '平原县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371427, '夏津县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371428, '武城县', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371481, '乐陵市', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371482, '禹城市', 371400, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371500, '聊城市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371502, '东昌府区', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371503, '茌平区', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371521, '阳谷县', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371522, '莘县', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371524, '东阿县', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371525, '冠县', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371526, '高唐县', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371581, '临清市', 371500, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371600, '滨州市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371602, '滨城区', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371603, '沾化区', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371621, '惠民县', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371622, '阳信县', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371623, '无棣县', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371625, '博兴县', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371681, '邹平市', 371600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371700, '菏泽市', 370000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371702, '牡丹区', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371703, '定陶区', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371721, '曹县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371722, '单县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371723, '成武县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371724, '巨野县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371725, '郓城县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371726, '鄄城县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (371728, '东明县', 371700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410000, '河南省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410100, '郑州市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410102, '中原区', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410103, '二七区', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410104, '管城回族区', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410105, '金水区', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410106, '上街区', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410108, '惠济区', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410122, '中牟县', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410181, '巩义市', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410182, '荥阳市', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410183, '新密市', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410184, '新郑市', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410185, '登封市', 410100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410200, '开封市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410202, '龙亭区', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410203, '顺河回族区', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410204, '鼓楼区', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410205, '禹王台区', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410212, '祥符区', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410221, '杞县', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410222, '通许县', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410223, '尉氏县', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410225, '兰考县', 410200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410300, '洛阳市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410302, '老城区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410303, '西工区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410304, '瀍河回族区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410305, '涧西区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410307, '偃师区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410308, '孟津区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410311, '洛龙区', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410323, '新安县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410324, '栾川县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410325, '嵩县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410326, '汝阳县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410327, '宜阳县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410328, '洛宁县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410329, '伊川县', 410300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410400, '平顶山市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410402, '新华区', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410403, '卫东区', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410404, '石龙区', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410411, '湛河区', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410421, '宝丰县', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410422, '叶县', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410423, '鲁山县', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410425, '郏县', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410481, '舞钢市', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410482, '汝州市', 410400, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410500, '安阳市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410502, '文峰区', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410503, '北关区', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410505, '殷都区', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410506, '龙安区', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410522, '安阳县', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410523, '汤阴县', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410526, '滑县', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410527, '内黄县', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410581, '林州市', 410500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410600, '鹤壁市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410602, '鹤山区', 410600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410603, '山城区', 410600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410611, '淇滨区', 410600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410621, '浚县', 410600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410622, '淇县', 410600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410700, '新乡市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410702, '红旗区', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410703, '卫滨区', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410704, '凤泉区', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410711, '牧野区', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410721, '新乡县', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410724, '获嘉县', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410725, '原阳县', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410726, '延津县', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410727, '封丘县', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410781, '卫辉市', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410782, '辉县市', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410783, '长垣市', 410700, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410800, '焦作市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410802, '解放区', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410803, '中站区', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410804, '马村区', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410811, '山阳区', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410821, '修武县', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410822, '博爱县', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410823, '武陟县', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410825, '温县', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410882, '沁阳市', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410883, '孟州市', 410800, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410900, '濮阳市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410902, '华龙区', 410900, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410922, '清丰县', 410900, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410923, '南乐县', 410900, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410926, '范县', 410900, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410927, '台前县', 410900, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (410928, '濮阳县', 410900, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411000, '许昌市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411002, '魏都区', 411000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411003, '建安区', 411000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411024, '鄢陵县', 411000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411025, '襄城县', 411000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411081, '禹州市', 411000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411082, '长葛市', 411000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411100, '漯河市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411102, '源汇区', 411100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411103, '郾城区', 411100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411104, '召陵区', 411100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411121, '舞阳县', 411100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411122, '临颍县', 411100, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411200, '三门峡市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411202, '湖滨区', 411200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411203, '陕州区', 411200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411221, '渑池县', 411200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411224, '卢氏县', 411200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411281, '义马市', 411200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411282, '灵宝市', 411200, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411300, '南阳市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411302, '宛城区', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411303, '卧龙区', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411321, '南召县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411322, '方城县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411323, '西峡县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411324, '镇平县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411325, '内乡县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411326, '淅川县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411327, '社旗县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411328, '唐河县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411329, '新野县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411330, '桐柏县', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411381, '邓州市', 411300, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411400, '商丘市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411402, '梁园区', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411403, '睢阳区', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411421, '民权县', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411422, '睢县', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411423, '宁陵县', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411424, '柘城县', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411425, '虞城县', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411426, '夏邑县', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411481, '永城市', 411400, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411500, '信阳市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411502, '浉河区', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411503, '平桥区', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411521, '罗山县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411522, '光山县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411523, '新县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411524, '商城县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411525, '固始县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411526, '潢川县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411527, '淮滨县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411528, '息县', 411500, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411600, '周口市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411602, '川汇区', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411603, '淮阳区', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411621, '扶沟县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411622, '西华县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411623, '商水县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411624, '沈丘县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411625, '郸城县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411627, '太康县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411628, '鹿邑县', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411681, '项城市', 411600, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411700, '驻马店市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411702, '驿城区', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411721, '西平县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411722, '上蔡县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411723, '平舆县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411724, '正阳县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411725, '确山县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411726, '泌阳县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411727, '汝南县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411728, '遂平县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (411729, '新蔡县', 411700, 0, 0, 0, NULL, '2024-03-27 23:31:46', NULL, '2024-03-27 23:31:46');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (419001, '济源市', 410000, 0, 0, 0, NULL, '2024-03-27 23:31:45', NULL, '2024-03-27 23:31:45');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420000, '湖北省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420100, '武汉市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420102, '江岸区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420103, '江汉区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420104, '硚口区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420105, '汉阳区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420106, '武昌区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420107, '青山区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420111, '洪山区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420112, '东西湖区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420113, '汉南区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420114, '蔡甸区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420115, '江夏区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420116, '黄陂区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420117, '新洲区', 420100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420200, '黄石市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420202, '黄石港区', 420200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420203, '西塞山区', 420200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420204, '下陆区', 420200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420205, '铁山区', 420200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420222, '阳新县', 420200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420281, '大冶市', 420200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420300, '十堰市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420302, '茅箭区', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420303, '张湾区', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420304, '郧阳区', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420322, '郧西县', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420323, '竹山县', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420324, '竹溪县', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420325, '房县', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420381, '丹江口市', 420300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420500, '宜昌市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420502, '西陵区', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420503, '伍家岗区', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420504, '点军区', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420505, '猇亭区', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420506, '夷陵区', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420525, '远安县', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420526, '兴山县', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420527, '秭归县', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420528, '长阳土家族自治县', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420529, '五峰土家族自治县', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420581, '宜都市', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420582, '当阳市', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420583, '枝江市', 420500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420600, '襄阳市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420602, '襄城区', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420606, '樊城区', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420607, '襄州区', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420624, '南漳县', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420625, '谷城县', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420626, '保康县', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420682, '老河口市', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420683, '枣阳市', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420684, '宜城市', 420600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420700, '鄂州市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420702, '梁子湖区', 420700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420703, '华容区', 420700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420704, '鄂城区', 420700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420800, '荆门市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420802, '东宝区', 420800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420804, '掇刀区', 420800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420822, '沙洋县', 420800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420881, '钟祥市', 420800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420882, '京山市', 420800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420900, '孝感市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420902, '孝南区', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420921, '孝昌县', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420922, '大悟县', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420923, '云梦县', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420981, '应城市', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420982, '安陆市', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (420984, '汉川市', 420900, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421000, '荆州市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421002, '沙市区', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421003, '荆州区', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421022, '公安县', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421024, '江陵县', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421081, '石首市', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421083, '洪湖市', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421087, '松滋市', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421088, '监利市', 421000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421100, '黄冈市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421102, '黄州区', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421121, '团风县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421122, '红安县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421123, '罗田县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421124, '英山县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421125, '浠水县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421126, '蕲春县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421127, '黄梅县', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421181, '麻城市', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421182, '武穴市', 421100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421200, '咸宁市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421202, '咸安区', 421200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421221, '嘉鱼县', 421200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421222, '通城县', 421200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421223, '崇阳县', 421200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421224, '通山县', 421200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421281, '赤壁市', 421200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421300, '随州市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421303, '曾都区', 421300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421321, '随县', 421300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (421381, '广水市', 421300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422800, '恩施土家族苗族自治州', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422801, '恩施市', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422802, '利川市', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422822, '建始县', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422823, '巴东县', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422825, '宣恩县', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422826, '咸丰县', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422827, '来凤县', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (422828, '鹤峰县', 422800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (429004, '仙桃市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (429005, '潜江市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (429006, '天门市', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (429021, '神农架林区', 420000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430000, '湖南省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430100, '长沙市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430102, '芙蓉区', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430103, '天心区', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430104, '岳麓区', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430105, '开福区', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430111, '雨花区', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430112, '望城区', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430121, '长沙县', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430181, '浏阳市', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430182, '宁乡市', 430100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430200, '株洲市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430202, '荷塘区', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430203, '芦淞区', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430204, '石峰区', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430211, '天元区', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430212, '渌口区', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430223, '攸县', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430224, '茶陵县', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430225, '炎陵县', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430281, '醴陵市', 430200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430300, '湘潭市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430302, '雨湖区', 430300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430304, '岳塘区', 430300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430321, '湘潭县', 430300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430381, '湘乡市', 430300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430382, '韶山市', 430300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430400, '衡阳市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430405, '珠晖区', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430406, '雁峰区', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430407, '石鼓区', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430408, '蒸湘区', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430412, '南岳区', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430421, '衡阳县', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430422, '衡南县', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430423, '衡山县', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430424, '衡东县', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430426, '祁东县', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430481, '耒阳市', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430482, '常宁市', 430400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430500, '邵阳市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430502, '双清区', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430503, '大祥区', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430511, '北塔区', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430522, '新邵县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430523, '邵阳县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430524, '隆回县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430525, '洞口县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430527, '绥宁县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430528, '新宁县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430529, '城步苗族自治县', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430581, '武冈市', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430582, '邵东市', 430500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430600, '岳阳市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430602, '岳阳楼区', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430603, '云溪区', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430611, '君山区', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430621, '岳阳县', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430623, '华容县', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430624, '湘阴县', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430626, '平江县', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430681, '汨罗市', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430682, '临湘市', 430600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430700, '常德市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430702, '武陵区', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430703, '鼎城区', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430721, '安乡县', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430722, '汉寿县', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430723, '澧县', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430724, '临澧县', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430725, '桃源县', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430726, '石门县', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430781, '津市市', 430700, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430800, '张家界市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430802, '永定区', 430800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430811, '武陵源区', 430800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430821, '慈利县', 430800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430822, '桑植县', 430800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430900, '益阳市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430902, '资阳区', 430900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430903, '赫山区', 430900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430921, '南县', 430900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430922, '桃江县', 430900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430923, '安化县', 430900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (430981, '沅江市', 430900, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431000, '郴州市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431002, '北湖区', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431003, '苏仙区', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431021, '桂阳县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431022, '宜章县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431023, '永兴县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431024, '嘉禾县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431025, '临武县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431026, '汝城县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431027, '桂东县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431028, '安仁县', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431081, '资兴市', 431000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431100, '永州市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431102, '零陵区', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431103, '冷水滩区', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431122, '东安县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431123, '双牌县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431124, '道县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431125, '江永县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431126, '宁远县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431127, '蓝山县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431128, '新田县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431129, '江华瑶族自治县', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431181, '祁阳市', 431100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431200, '怀化市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431202, '鹤城区', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431221, '中方县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431222, '沅陵县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431223, '辰溪县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431224, '溆浦县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431225, '会同县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431226, '麻阳苗族自治县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431227, '新晃侗族自治县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431228, '芷江侗族自治县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431229, '靖州苗族侗族自治县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431230, '通道侗族自治县', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431281, '洪江市', 431200, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431300, '娄底市', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431302, '娄星区', 431300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431321, '双峰县', 431300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431322, '新化县', 431300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431381, '冷水江市', 431300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (431382, '涟源市', 431300, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433100, '湘西土家族苗族自治州', 430000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433101, '吉首市', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433122, '泸溪县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433123, '凤凰县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433124, '花垣县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433125, '保靖县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433126, '古丈县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433127, '永顺县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:44', NULL, '2024-03-27 23:31:44');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (433130, '龙山县', 433100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440000, '广东省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440100, '广州市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440103, '荔湾区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440104, '越秀区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440105, '海珠区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440106, '天河区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440111, '白云区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440112, '黄埔区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440113, '番禺区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440114, '花都区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440115, '南沙区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440117, '从化区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440118, '增城区', 440100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440200, '韶关市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440203, '武江区', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440204, '浈江区', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440205, '曲江区', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440222, '始兴县', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440224, '仁化县', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440229, '翁源县', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440232, '乳源瑶族自治县', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440233, '新丰县', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440281, '乐昌市', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440282, '南雄市', 440200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440300, '深圳市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440303, '罗湖区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440304, '福田区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440305, '南山区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440306, '宝安区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440307, '龙岗区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440308, '盐田区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440309, '龙华区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440310, '坪山区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440311, '光明区', 440300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440400, '珠海市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440402, '香洲区', 440400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440403, '斗门区', 440400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440404, '金湾区', 440400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440500, '汕头市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440507, '龙湖区', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440511, '金平区', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440512, '濠江区', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440513, '潮阳区', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440514, '潮南区', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440515, '澄海区', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440523, '南澳县', 440500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440600, '佛山市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440604, '禅城区', 440600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440605, '南海区', 440600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440606, '顺德区', 440600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440607, '三水区', 440600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440608, '高明区', 440600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440700, '江门市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440703, '蓬江区', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440704, '江海区', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440705, '新会区', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440781, '台山市', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440783, '开平市', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440784, '鹤山市', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440785, '恩平市', 440700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440800, '湛江市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440802, '赤坎区', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440803, '霞山区', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440804, '坡头区', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440811, '麻章区', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440823, '遂溪县', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440825, '徐闻县', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440881, '廉江市', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440882, '雷州市', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440883, '吴川市', 440800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440900, '茂名市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440902, '茂南区', 440900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440904, '电白区', 440900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440981, '高州市', 440900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440982, '化州市', 440900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (440983, '信宜市', 440900, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441200, '肇庆市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441202, '端州区', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441203, '鼎湖区', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441204, '高要区', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441223, '广宁县', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441224, '怀集县', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441225, '封开县', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441226, '德庆县', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441284, '四会市', 441200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441300, '惠州市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441302, '惠城区', 441300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441303, '惠阳区', 441300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441322, '博罗县', 441300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441323, '惠东县', 441300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441324, '龙门县', 441300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441400, '梅州市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441402, '梅江区', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441403, '梅县区', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441422, '大埔县', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441423, '丰顺县', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441424, '五华县', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441426, '平远县', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441427, '蕉岭县', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441481, '兴宁市', 441400, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441500, '汕尾市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441502, '城区', 441500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441521, '海丰县', 441500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441523, '陆河县', 441500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441581, '陆丰市', 441500, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441600, '河源市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441602, '源城区', 441600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441621, '紫金县', 441600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441622, '龙川县', 441600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441623, '连平县', 441600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441624, '和平县', 441600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441625, '东源县', 441600, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441700, '阳江市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441702, '江城区', 441700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441704, '阳东区', 441700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441721, '阳西县', 441700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441781, '阳春市', 441700, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441800, '清远市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441802, '清城区', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441803, '清新区', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441821, '佛冈县', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441823, '阳山县', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441825, '连山壮族瑶族自治县', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441826, '连南瑶族自治县', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441881, '英德市', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441882, '连州市', 441800, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (441900, '东莞市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (442000, '中山市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445100, '潮州市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445102, '湘桥区', 445100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445103, '潮安区', 445100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445122, '饶平县', 445100, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445200, '揭阳市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445202, '榕城区', 445200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445203, '揭东区', 445200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445222, '揭西县', 445200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445224, '惠来县', 445200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445281, '普宁市', 445200, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445300, '云浮市', 440000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445302, '云城区', 445300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445303, '云安区', 445300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445321, '新兴县', 445300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445322, '郁南县', 445300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (445381, '罗定市', 445300, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450000, '广西壮族自治区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450100, '南宁市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450102, '兴宁区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450103, '青秀区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450105, '江南区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450107, '西乡塘区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450108, '良庆区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450109, '邕宁区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450110, '武鸣区', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450123, '隆安县', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450124, '马山县', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450125, '上林县', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450126, '宾阳县', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450181, '横州市', 450100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450200, '柳州市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450202, '城中区', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450203, '鱼峰区', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450204, '柳南区', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450205, '柳北区', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450206, '柳江区', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450222, '柳城县', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450223, '鹿寨县', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450224, '融安县', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450225, '融水苗族自治县', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450226, '三江侗族自治县', 450200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450300, '桂林市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450302, '秀峰区', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450303, '叠彩区', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450304, '象山区', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450305, '七星区', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450311, '雁山区', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450312, '临桂区', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450321, '阳朔县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450323, '灵川县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450324, '全州县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450325, '兴安县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450326, '永福县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450327, '灌阳县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450328, '龙胜各族自治县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450329, '资源县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450330, '平乐县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450332, '恭城瑶族自治县', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450381, '荔浦市', 450300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450400, '梧州市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450403, '万秀区', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450405, '长洲区', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450406, '龙圩区', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450421, '苍梧县', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450422, '藤县', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450423, '蒙山县', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450481, '岑溪市', 450400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450500, '北海市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450502, '海城区', 450500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450503, '银海区', 450500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450512, '铁山港区', 450500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450521, '合浦县', 450500, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450600, '防城港市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450602, '港口区', 450600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450603, '防城区', 450600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450621, '上思县', 450600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450681, '东兴市', 450600, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450700, '钦州市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450702, '钦南区', 450700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450703, '钦北区', 450700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450721, '灵山县', 450700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450722, '浦北县', 450700, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450800, '贵港市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450802, '港北区', 450800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450803, '港南区', 450800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450804, '覃塘区', 450800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450821, '平南县', 450800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450881, '桂平市', 450800, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450900, '玉林市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450902, '玉州区', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450903, '福绵区', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450921, '容县', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450922, '陆川县', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450923, '博白县', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450924, '兴业县', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (450981, '北流市', 450900, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451000, '百色市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451002, '右江区', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451003, '田阳区', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451022, '田东县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451024, '德保县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451026, '那坡县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451027, '凌云县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451028, '乐业县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451029, '田林县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451030, '西林县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451031, '隆林各族自治县', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451081, '靖西市', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451082, '平果市', 451000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451100, '贺州市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451102, '八步区', 451100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451103, '平桂区', 451100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451121, '昭平县', 451100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451122, '钟山县', 451100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451123, '富川瑶族自治县', 451100, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451200, '河池市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451202, '金城江区', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451203, '宜州区', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451221, '南丹县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451222, '天峨县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451223, '凤山县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451224, '东兰县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451225, '罗城仫佬族自治县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451226, '环江毛南族自治县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451227, '巴马瑶族自治县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451228, '都安瑶族自治县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451229, '大化瑶族自治县', 451200, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451300, '来宾市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451302, '兴宾区', 451300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451321, '忻城县', 451300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451322, '象州县', 451300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451323, '武宣县', 451300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451324, '金秀瑶族自治县', 451300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451381, '合山市', 451300, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451400, '崇左市', 450000, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451402, '江州区', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451421, '扶绥县', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451422, '宁明县', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451423, '龙州县', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451424, '大新县', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451425, '天等县', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (451481, '凭祥市', 451400, 0, 0, 0, NULL, '2024-03-27 23:31:36', NULL, '2024-03-27 23:31:36');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460000, '海南省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460100, '海口市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460105, '秀英区', 460100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460106, '龙华区', 460100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460107, '琼山区', 460100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460108, '美兰区', 460100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460200, '三亚市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460202, '海棠区', 460200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460203, '吉阳区', 460200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460204, '天涯区', 460200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460205, '崖州区', 460200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460300, '三沙市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460321, '西沙区', 460300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460322, '南沙区', 460300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460323, '中沙群岛的岛礁及其海域', 460300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460324, '永乐群岛的岛礁及其海域', 460300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (460400, '儋州市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469001, '五指山市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469002, '琼海市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469005, '文昌市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469006, '万宁市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469007, '东方市', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469021, '定安县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469022, '屯昌县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469023, '澄迈县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469024, '临高县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469025, '白沙黎族自治县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469026, '昌江黎族自治县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469027, '乐东黎族自治县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469028, '陵水黎族自治县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469029, '保亭黎族苗族自治县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (469030, '琼中黎族苗族自治县', 460000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500000, '重庆市', 0, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500101, '万州区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500102, '涪陵区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500103, '渝中区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500104, '大渡口区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500105, '江北区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500106, '沙坪坝区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500107, '九龙坡区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500108, '南岸区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500109, '北碚区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500110, '綦江区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500111, '大足区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500112, '渝北区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500113, '巴南区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500114, '黔江区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500115, '长寿区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500116, '江津区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500117, '合川区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500118, '永川区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500119, '南川区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500120, '璧山区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500151, '铜梁区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500152, '潼南区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500153, '荣昌区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500154, '开州区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500155, '梁平区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500156, '武隆区', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500229, '城口县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500230, '丰都县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500231, '垫江县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500233, '忠县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500235, '云阳县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500236, '奉节县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500237, '巫山县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500238, '巫溪县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500240, '石柱土家族自治县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500241, '秀山土家族苗族自治县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500242, '酉阳土家族苗族自治县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (500243, '彭水苗族土家族自治县', 500000, 0, 0, 0, NULL, '2024-03-27 23:31:29', NULL, '2024-03-27 23:31:29');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510000, '四川省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510100, '成都市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510104, '锦江区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510105, '青羊区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510106, '金牛区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510107, '武侯区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510108, '成华区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510112, '龙泉驿区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510113, '青白江区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510114, '新都区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510115, '温江区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510116, '双流区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510117, '郫都区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510118, '新津区', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510121, '金堂县', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510129, '大邑县', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510131, '蒲江县', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510181, '都江堰市', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510182, '彭州市', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510183, '邛崃市', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510184, '崇州市', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510185, '简阳市', 510100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510300, '自贡市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510302, '自流井区', 510300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510303, '贡井区', 510300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510304, '大安区', 510300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510311, '沿滩区', 510300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510321, '荣县', 510300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510322, '富顺县', 510300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510400, '攀枝花市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510402, '东区', 510400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510403, '西区', 510400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510411, '仁和区', 510400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510421, '米易县', 510400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510422, '盐边县', 510400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510500, '泸州市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510502, '江阳区', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510503, '纳溪区', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510504, '龙马潭区', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510521, '泸县', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510522, '合江县', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510524, '叙永县', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510525, '古蔺县', 510500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510600, '德阳市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510603, '旌阳区', 510600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510604, '罗江区', 510600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510623, '中江县', 510600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510681, '广汉市', 510600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510682, '什邡市', 510600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510683, '绵竹市', 510600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510700, '绵阳市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510703, '涪城区', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510704, '游仙区', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510705, '安州区', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510722, '三台县', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510723, '盐亭县', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510725, '梓潼县', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510726, '北川羌族自治县', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510727, '平武县', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510781, '江油市', 510700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510800, '广元市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510802, '利州区', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510811, '昭化区', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510812, '朝天区', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510821, '旺苍县', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510822, '青川县', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510823, '剑阁县', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510824, '苍溪县', 510800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510900, '遂宁市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510903, '船山区', 510900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510904, '安居区', 510900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510921, '蓬溪县', 510900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510923, '大英县', 510900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (510981, '射洪市', 510900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511000, '内江市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511002, '市中区', 511000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511011, '东兴区', 511000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511024, '威远县', 511000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511025, '资中县', 511000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511083, '隆昌市', 511000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511100, '乐山市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511102, '市中区', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511111, '沙湾区', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511112, '五通桥区', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511113, '金口河区', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511123, '犍为县', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511124, '井研县', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511126, '夹江县', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511129, '沐川县', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511132, '峨边彝族自治县', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511133, '马边彝族自治县', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511181, '峨眉山市', 511100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511300, '南充市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511302, '顺庆区', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511303, '高坪区', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511304, '嘉陵区', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511321, '南部县', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511322, '营山县', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511323, '蓬安县', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511324, '仪陇县', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511325, '西充县', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511381, '阆中市', 511300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511400, '眉山市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511402, '东坡区', 511400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511403, '彭山区', 511400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511421, '仁寿县', 511400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511423, '洪雅县', 511400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511424, '丹棱县', 511400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511425, '青神县', 511400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511500, '宜宾市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511502, '翠屏区', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511503, '南溪区', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511504, '叙州区', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511523, '江安县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511524, '长宁县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511525, '高县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511526, '珙县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511527, '筠连县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511528, '兴文县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511529, '屏山县', 511500, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511600, '广安市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511602, '广安区', 511600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511603, '前锋区', 511600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511621, '岳池县', 511600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511622, '武胜县', 511600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511623, '邻水县', 511600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511681, '华蓥市', 511600, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511700, '达州市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511702, '通川区', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511703, '达川区', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511722, '宣汉县', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511723, '开江县', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511724, '大竹县', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511725, '渠县', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511781, '万源市', 511700, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511800, '雅安市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511802, '雨城区', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511803, '名山区', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511822, '荥经县', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511823, '汉源县', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511824, '石棉县', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511825, '天全县', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511826, '芦山县', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511827, '宝兴县', 511800, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511900, '巴中市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511902, '巴州区', 511900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511903, '恩阳区', 511900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511921, '通江县', 511900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511922, '南江县', 511900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (511923, '平昌县', 511900, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (512000, '资阳市', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (512002, '雁江区', 512000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (512021, '安岳县', 512000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (512022, '乐至县', 512000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513200, '阿坝藏族羌族自治州', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513201, '马尔康市', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513221, '汶川县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513222, '理县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513223, '茂县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513224, '松潘县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513225, '九寨沟县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513226, '金川县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513227, '小金县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513228, '黑水县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513230, '壤塘县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513231, '阿坝县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513232, '若尔盖县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513233, '红原县', 513200, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513300, '甘孜藏族自治州', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513301, '康定市', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513322, '泸定县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513323, '丹巴县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513324, '九龙县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513325, '雅江县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513326, '道孚县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513327, '炉霍县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513328, '甘孜县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513329, '新龙县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513330, '德格县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513331, '白玉县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513332, '石渠县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513333, '色达县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513334, '理塘县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513335, '巴塘县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513336, '乡城县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513337, '稻城县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513338, '得荣县', 513300, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513400, '凉山彝族自治州', 510000, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513401, '西昌市', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513402, '会理市', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513422, '木里藏族自治县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513423, '盐源县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513424, '德昌县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513426, '会东县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513427, '宁南县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513428, '普格县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513429, '布拖县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513430, '金阳县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513431, '昭觉县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513432, '喜德县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513433, '冕宁县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513434, '越西县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513435, '甘洛县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513436, '美姑县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (513437, '雷波县', 513400, 0, 0, 0, NULL, '2024-03-27 23:31:41', NULL, '2024-03-27 23:31:41');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520000, '贵州省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520100, '贵阳市', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520102, '南明区', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520103, '云岩区', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520111, '花溪区', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520112, '乌当区', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520113, '白云区', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520115, '观山湖区', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520121, '开阳县', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520122, '息烽县', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520123, '修文县', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520181, '清镇市', 520100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520200, '六盘水市', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520201, '钟山区', 520200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520203, '六枝特区', 520200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520204, '水城区', 520200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520281, '盘州市', 520200, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520300, '遵义市', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520302, '红花岗区', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520303, '汇川区', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520304, '播州区', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520322, '桐梓县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520323, '绥阳县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520324, '正安县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520325, '道真仡佬族苗族自治县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520326, '务川仡佬族苗族自治县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520327, '凤冈县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520328, '湄潭县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520329, '余庆县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520330, '习水县', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520381, '赤水市', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520382, '仁怀市', 520300, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520400, '安顺市', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520402, '西秀区', 520400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520403, '平坝区', 520400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520422, '普定县', 520400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520423, '镇宁布依族苗族自治县', 520400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520424, '关岭布依族苗族自治县', 520400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520425, '紫云苗族布依族自治县', 520400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520500, '毕节市', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520502, '七星关区', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520521, '大方县', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520523, '金沙县', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520524, '织金县', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520525, '纳雍县', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520526, '威宁彝族回族苗族自治县', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520527, '赫章县', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520581, '黔西市', 520500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520600, '铜仁市', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520602, '碧江区', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520603, '万山区', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520621, '江口县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520622, '玉屏侗族自治县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520623, '石阡县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520624, '思南县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520625, '印江土家族苗族自治县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520626, '德江县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520627, '沿河土家族自治县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (520628, '松桃苗族自治县', 520600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522300, '黔西南布依族苗族自治州', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522301, '兴义市', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522302, '兴仁市', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522323, '普安县', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522324, '晴隆县', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522325, '贞丰县', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522326, '望谟县', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522327, '册亨县', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522328, '安龙县', 522300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522600, '黔东南苗族侗族自治州', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522601, '凯里市', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522622, '黄平县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522623, '施秉县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522624, '三穗县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522625, '镇远县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522626, '岑巩县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522627, '天柱县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522628, '锦屏县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522629, '剑河县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522630, '台江县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522631, '黎平县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522632, '榕江县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522633, '从江县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522634, '雷山县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522635, '麻江县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522636, '丹寨县', 522600, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522700, '黔南布依族苗族自治州', 520000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522701, '都匀市', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522702, '福泉市', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522722, '荔波县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522723, '贵定县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522725, '瓮安县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522726, '独山县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522727, '平塘县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522728, '罗甸县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522729, '长顺县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522730, '龙里县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522731, '惠水县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (522732, '三都水族自治县', 522700, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530000, '云南省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530100, '昆明市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530102, '五华区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530103, '盘龙区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530111, '官渡区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530112, '西山区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530113, '东川区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530114, '呈贡区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530115, '晋宁区', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530124, '富民县', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530125, '宜良县', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530126, '石林彝族自治县', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530127, '嵩明县', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530128, '禄劝彝族苗族自治县', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530129, '寻甸回族彝族自治县', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530181, '安宁市', 530100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530300, '曲靖市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530302, '麒麟区', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530303, '沾益区', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530304, '马龙区', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530322, '陆良县', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530323, '师宗县', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530324, '罗平县', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530325, '富源县', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530326, '会泽县', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530381, '宣威市', 530300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530400, '玉溪市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530402, '红塔区', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530403, '江川区', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530423, '通海县', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530424, '华宁县', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530425, '易门县', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530426, '峨山彝族自治县', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530427, '新平彝族傣族自治县', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530428, '元江哈尼族彝族傣族自治县', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530481, '澄江市', 530400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530500, '保山市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530502, '隆阳区', 530500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530521, '施甸县', 530500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530523, '龙陵县', 530500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530524, '昌宁县', 530500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530581, '腾冲市', 530500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530600, '昭通市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530602, '昭阳区', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530621, '鲁甸县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530622, '巧家县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530623, '盐津县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530624, '大关县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530625, '永善县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530626, '绥江县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530627, '镇雄县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530628, '彝良县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530629, '威信县', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530681, '水富市', 530600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530700, '丽江市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530702, '古城区', 530700, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530721, '玉龙纳西族自治县', 530700, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530722, '永胜县', 530700, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530723, '华坪县', 530700, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530724, '宁蒗彝族自治县', 530700, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530800, '普洱市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530802, '思茅区', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530821, '宁洱哈尼族彝族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530822, '墨江哈尼族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530823, '景东彝族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530824, '景谷傣族彝族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530825, '镇沅彝族哈尼族拉祜族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530826, '江城哈尼族彝族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530827, '孟连傣族拉祜族佤族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530828, '澜沧拉祜族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530829, '西盟佤族自治县', 530800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530900, '临沧市', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530902, '临翔区', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530921, '凤庆县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530922, '云县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530923, '永德县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530924, '镇康县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530925, '双江拉祜族佤族布朗族傣族自治县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530926, '耿马傣族佤族自治县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (530927, '沧源佤族自治县', 530900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532300, '楚雄彝族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532301, '楚雄市', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532302, '禄丰市', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532322, '双柏县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532323, '牟定县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532324, '南华县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532325, '姚安县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532326, '大姚县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532327, '永仁县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532328, '元谋县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532329, '武定县', 532300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532500, '红河哈尼族彝族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532501, '个旧市', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532502, '开远市', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532503, '蒙自市', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532504, '弥勒市', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532523, '屏边苗族自治县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532524, '建水县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532525, '石屏县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532527, '泸西县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532528, '元阳县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532529, '红河县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532530, '金平苗族瑶族傣族自治县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532531, '绿春县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532532, '河口瑶族自治县', 532500, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532600, '文山壮族苗族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532601, '文山市', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532622, '砚山县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532623, '西畴县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532624, '麻栗坡县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532625, '马关县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532626, '丘北县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532627, '广南县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532628, '富宁县', 532600, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532800, '西双版纳傣族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532801, '景洪市', 532800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532822, '勐海县', 532800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532823, '勐腊县', 532800, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532900, '大理白族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532901, '大理市', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532922, '漾濞彝族自治县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532923, '祥云县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532924, '宾川县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532925, '弥渡县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532926, '南涧彝族自治县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532927, '巍山彝族回族自治县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532928, '永平县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532929, '云龙县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532930, '洱源县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532931, '剑川县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (532932, '鹤庆县', 532900, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533100, '德宏傣族景颇族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533102, '瑞丽市', 533100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533103, '芒市', 533100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533122, '梁河县', 533100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533123, '盈江县', 533100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533124, '陇川县', 533100, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533300, '怒江傈僳族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533301, '泸水市', 533300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533323, '福贡县', 533300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533324, '贡山独龙族怒族自治县', 533300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533325, '兰坪白族普米族自治县', 533300, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533400, '迪庆藏族自治州', 530000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533401, '香格里拉市', 533400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533422, '德钦县', 533400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (533423, '维西傈僳族自治县', 533400, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540000, '西藏自治区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540100, '拉萨市', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540102, '城关区', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540103, '堆龙德庆区', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540104, '达孜区', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540121, '林周县', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540122, '当雄县', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540123, '尼木县', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540124, '曲水县', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540127, '墨竹工卡县', 540100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540200, '日喀则市', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540202, '桑珠孜区', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540221, '南木林县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540222, '江孜县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540223, '定日县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540224, '萨迦县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540225, '拉孜县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540226, '昂仁县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540227, '谢通门县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540228, '白朗县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540229, '仁布县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540230, '康马县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540231, '定结县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540232, '仲巴县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540233, '亚东县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540234, '吉隆县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540235, '聂拉木县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540236, '萨嘎县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540237, '岗巴县', 540200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540300, '昌都市', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540302, '卡若区', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540321, '江达县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540322, '贡觉县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540323, '类乌齐县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540324, '丁青县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540325, '察雅县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540326, '八宿县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540327, '左贡县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540328, '芒康县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540329, '洛隆县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540330, '边坝县', 540300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540400, '林芝市', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540402, '巴宜区', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540421, '工布江达县', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540423, '墨脱县', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540424, '波密县', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540425, '察隅县', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540426, '朗县', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540481, '米林市', 540400, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540500, '山南市', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540502, '乃东区', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540521, '扎囊县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540522, '贡嘎县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540523, '桑日县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540524, '琼结县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540525, '曲松县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540526, '措美县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540527, '洛扎县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540528, '加查县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540529, '隆子县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540531, '浪卡子县', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540581, '错那市', 540500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540600, '那曲市', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540602, '色尼区', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540621, '嘉黎县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540622, '比如县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540623, '聂荣县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540624, '安多县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540625, '申扎县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540626, '索县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540627, '班戈县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540628, '巴青县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540629, '尼玛县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (540630, '双湖县', 540600, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542500, '阿里地区', 540000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542521, '普兰县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542522, '札达县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542523, '噶尔县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542524, '日土县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542525, '革吉县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542526, '改则县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (542527, '措勤县', 542500, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610000, '陕西省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610100, '西安市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610102, '新城区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610103, '碑林区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610104, '莲湖区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610111, '灞桥区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610112, '未央区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610113, '雁塔区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610114, '阎良区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610115, '临潼区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610116, '长安区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610117, '高陵区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610118, '鄠邑区', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610122, '蓝田县', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610124, '周至县', 610100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610200, '铜川市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610202, '王益区', 610200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610203, '印台区', 610200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610204, '耀州区', 610200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610222, '宜君县', 610200, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610300, '宝鸡市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610302, '渭滨区', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610303, '金台区', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610304, '陈仓区', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610305, '凤翔区', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610323, '岐山县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610324, '扶风县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610326, '眉县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610327, '陇县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610328, '千阳县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610329, '麟游县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610330, '凤县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610331, '太白县', 610300, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610400, '咸阳市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610402, '秦都区', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610403, '杨陵区', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610404, '渭城区', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610422, '三原县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610423, '泾阳县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610424, '乾县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610425, '礼泉县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610426, '永寿县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610428, '长武县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610429, '旬邑县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610430, '淳化县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610431, '武功县', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610481, '兴平市', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610482, '彬州市', 610400, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610500, '渭南市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610502, '临渭区', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610503, '华州区', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610522, '潼关县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610523, '大荔县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610524, '合阳县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610525, '澄城县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610526, '蒲城县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610527, '白水县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610528, '富平县', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610581, '韩城市', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610582, '华阴市', 610500, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610600, '延安市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610602, '宝塔区', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610603, '安塞区', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610621, '延长县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610622, '延川县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610625, '志丹县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610626, '吴起县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610627, '甘泉县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610628, '富县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610629, '洛川县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610630, '宜川县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610631, '黄龙县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610632, '黄陵县', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610681, '子长市', 610600, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610700, '汉中市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610702, '汉台区', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610703, '南郑区', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610722, '城固县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610723, '洋县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610724, '西乡县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610725, '勉县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610726, '宁强县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610727, '略阳县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610728, '镇巴县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610729, '留坝县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610730, '佛坪县', 610700, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610800, '榆林市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610802, '榆阳区', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610803, '横山区', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610822, '府谷县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610824, '靖边县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610825, '定边县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610826, '绥德县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610827, '米脂县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610828, '佳县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610829, '吴堡县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610830, '清涧县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610831, '子洲县', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610881, '神木市', 610800, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610900, '安康市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610902, '汉滨区', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610921, '汉阴县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610922, '石泉县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610923, '宁陕县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610924, '紫阳县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610925, '岚皋县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610926, '平利县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610927, '镇坪县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610929, '白河县', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (610981, '旬阳市', 610900, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611000, '商洛市', 610000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611002, '商州区', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611021, '洛南县', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611022, '丹凤县', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611023, '商南县', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611024, '山阳县', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611025, '镇安县', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (611026, '柞水县', 611000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620000, '甘肃省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620100, '兰州市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620102, '城关区', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620103, '七里河区', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620104, '西固区', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620105, '安宁区', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620111, '红古区', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620121, '永登县', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620122, '皋兰县', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620123, '榆中县', 620100, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620200, '嘉峪关市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620300, '金昌市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620302, '金川区', 620300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620321, '永昌县', 620300, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620400, '白银市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620402, '白银区', 620400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620403, '平川区', 620400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620421, '靖远县', 620400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620422, '会宁县', 620400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620423, '景泰县', 620400, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620500, '天水市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620502, '秦州区', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620503, '麦积区', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620521, '清水县', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620522, '秦安县', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620523, '甘谷县', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620524, '武山县', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620525, '张家川回族自治县', 620500, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620600, '武威市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620602, '凉州区', 620600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620621, '民勤县', 620600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620622, '古浪县', 620600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620623, '天祝藏族自治县', 620600, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620700, '张掖市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620702, '甘州区', 620700, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620721, '肃南裕固族自治县', 620700, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620722, '民乐县', 620700, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620723, '临泽县', 620700, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620724, '高台县', 620700, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620725, '山丹县', 620700, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620800, '平凉市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620802, '崆峒区', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620821, '泾川县', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620822, '灵台县', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620823, '崇信县', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620825, '庄浪县', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620826, '静宁县', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620881, '华亭市', 620800, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620900, '酒泉市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620902, '肃州区', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620921, '金塔县', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620922, '瓜州县', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620923, '肃北蒙古族自治县', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620924, '阿克塞哈萨克族自治县', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620981, '玉门市', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (620982, '敦煌市', 620900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621000, '庆阳市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621002, '西峰区', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621021, '庆城县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621022, '环县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621023, '华池县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621024, '合水县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621025, '正宁县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621026, '宁县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621027, '镇原县', 621000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621100, '定西市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621102, '安定区', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621121, '通渭县', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621122, '陇西县', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621123, '渭源县', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621124, '临洮县', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621125, '漳县', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621126, '岷县', 621100, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621200, '陇南市', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621202, '武都区', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621221, '成县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621222, '文县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621223, '宕昌县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621224, '康县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621225, '西和县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621226, '礼县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621227, '徽县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (621228, '两当县', 621200, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622900, '临夏回族自治州', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622901, '临夏市', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622921, '临夏县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622922, '康乐县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622923, '永靖县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622924, '广河县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622925, '和政县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622926, '东乡族自治县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (622927, '积石山保安族东乡族撒拉族自治县', 622900, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623000, '甘南藏族自治州', 620000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623001, '合作市', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623021, '临潭县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623022, '卓尼县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623023, '舟曲县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623024, '迭部县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:43', NULL, '2024-03-27 23:31:43');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623025, '玛曲县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623026, '碌曲县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (623027, '夏河县', 623000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630000, '青海省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630100, '西宁市', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630102, '城东区', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630103, '城中区', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630104, '城西区', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630105, '城北区', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630106, '湟中区', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630121, '大通回族土族自治县', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630123, '湟源县', 630100, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630200, '海东市', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630202, '乐都区', 630200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630203, '平安区', 630200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630222, '民和回族土族自治县', 630200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630223, '互助土族自治县', 630200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630224, '化隆回族自治县', 630200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (630225, '循化撒拉族自治县', 630200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632200, '海北藏族自治州', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632221, '门源回族自治县', 632200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632222, '祁连县', 632200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632223, '海晏县', 632200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632224, '刚察县', 632200, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632300, '黄南藏族自治州', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632301, '同仁市', 632300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632322, '尖扎县', 632300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632323, '泽库县', 632300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632324, '河南蒙古族自治县', 632300, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632500, '海南藏族自治州', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632521, '共和县', 632500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632522, '同德县', 632500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632523, '贵德县', 632500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632524, '兴海县', 632500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632525, '贵南县', 632500, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632600, '果洛藏族自治州', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632621, '玛沁县', 632600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632622, '班玛县', 632600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632623, '甘德县', 632600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632624, '达日县', 632600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632625, '久治县', 632600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632626, '玛多县', 632600, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632700, '玉树藏族自治州', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632701, '玉树市', 632700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632722, '杂多县', 632700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632723, '称多县', 632700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632724, '治多县', 632700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632725, '囊谦县', 632700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632726, '曲麻莱县', 632700, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632800, '海西蒙古族藏族自治州', 630000, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632801, '格尔木市', 632800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632802, '德令哈市', 632800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632803, '茫崖市', 632800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632821, '乌兰县', 632800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632822, '都兰县', 632800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (632823, '天峻县', 632800, 0, 0, 0, NULL, '2024-03-27 23:31:35', NULL, '2024-03-27 23:31:35');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640000, '宁夏回族自治区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640100, '银川市', 640000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640104, '兴庆区', 640100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640105, '西夏区', 640100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640106, '金凤区', 640100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640121, '永宁县', 640100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640122, '贺兰县', 640100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640181, '灵武市', 640100, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640200, '石嘴山市', 640000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640202, '大武口区', 640200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640205, '惠农区', 640200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640221, '平罗县', 640200, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640300, '吴忠市', 640000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640302, '利通区', 640300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640303, '红寺堡区', 640300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640323, '盐池县', 640300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640324, '同心县', 640300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640381, '青铜峡市', 640300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640400, '固原市', 640000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640402, '原州区', 640400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640422, '西吉县', 640400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640423, '隆德县', 640400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640424, '泾源县', 640400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640425, '彭阳县', 640400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640500, '中卫市', 640000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640502, '沙坡头区', 640500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640521, '中宁县', 640500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (640522, '海原县', 640500, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650000, '新疆维吾尔自治区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650100, '乌鲁木齐市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650102, '天山区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650103, '沙依巴克区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650104, '新市区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650105, '水磨沟区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650106, '头屯河区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650107, '达坂城区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650109, '米东区', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650121, '乌鲁木齐县', 650100, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650200, '克拉玛依市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650202, '独山子区', 650200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650203, '克拉玛依区', 650200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650204, '白碱滩区', 650200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650205, '乌尔禾区', 650200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650400, '吐鲁番市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650402, '高昌区', 650400, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650421, '鄯善县', 650400, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650422, '托克逊县', 650400, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650500, '哈密市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650502, '伊州区', 650500, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650521, '巴里坤哈萨克自治县', 650500, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (650522, '伊吾县', 650500, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652300, '昌吉回族自治州', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652301, '昌吉市', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652302, '阜康市', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652323, '呼图壁县', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652324, '玛纳斯县', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652325, '奇台县', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652327, '吉木萨尔县', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652328, '木垒哈萨克自治县', 652300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652700, '博尔塔拉蒙古自治州', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652701, '博乐市', 652700, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652702, '阿拉山口市', 652700, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652722, '精河县', 652700, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652723, '温泉县', 652700, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652800, '巴音郭楞蒙古自治州', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652801, '库尔勒市', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652822, '轮台县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652823, '尉犁县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652824, '若羌县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652825, '且末县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652826, '焉耆回族自治县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652827, '和静县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652828, '和硕县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652829, '博湖县', 652800, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652900, '阿克苏地区', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652901, '阿克苏市', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652902, '库车市', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652922, '温宿县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652924, '沙雅县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652925, '新和县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652926, '拜城县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652927, '乌什县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652928, '阿瓦提县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (652929, '柯坪县', 652900, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653000, '克孜勒苏柯尔克孜自治州', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653001, '阿图什市', 653000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653022, '阿克陶县', 653000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653023, '阿合奇县', 653000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653024, '乌恰县', 653000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653100, '喀什地区', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653101, '喀什市', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653121, '疏附县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653122, '疏勒县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653123, '英吉沙县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653124, '泽普县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653125, '莎车县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653126, '叶城县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653127, '麦盖提县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653128, '岳普湖县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653129, '伽师县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653130, '巴楚县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653131, '塔什库尔干塔吉克自治县', 653100, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653200, '和田地区', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653201, '和田市', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653221, '和田县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653222, '墨玉县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653223, '皮山县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653224, '洛浦县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653225, '策勒县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653226, '于田县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (653227, '民丰县', 653200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654000, '伊犁哈萨克自治州', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654002, '伊宁市', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654003, '奎屯市', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654004, '霍尔果斯市', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654021, '伊宁县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654022, '察布查尔锡伯自治县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654023, '霍城县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654024, '巩留县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654025, '新源县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654026, '昭苏县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654027, '特克斯县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654028, '尼勒克县', 654000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654200, '塔城地区', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654201, '塔城市', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654202, '乌苏市', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654203, '沙湾市', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654221, '额敏县', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654224, '托里县', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654225, '裕民县', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654226, '和布克赛尔蒙古自治县', 654200, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654300, '阿勒泰地区', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654301, '阿勒泰市', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654321, '布尔津县', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654322, '富蕴县', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654323, '福海县', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654324, '哈巴河县', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654325, '青河县', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (654326, '吉木乃县', 654300, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659001, '石河子市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659002, '阿拉尔市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659003, '图木舒克市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659004, '五家渠市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659005, '北屯市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659006, '铁门关市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659007, '双河市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659008, '可克达拉市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659009, '昆玉市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659010, '胡杨河市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659011, '新星市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:37', NULL, '2024-03-27 23:31:37');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (659012, '白杨市', 650000, 0, 0, 0, NULL, '2024-03-27 23:31:38', NULL, '2024-03-27 23:31:38');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810000, '香港特别行政区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810101, '中西区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810102, '湾仔区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810103, '东区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810104, '南区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810105, '油尖旺区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810106, '深水埗区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810107, '九龙城区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810108, '黄大仙区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810109, '观塘区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810110, '北区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810111, '大埔区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810112, '沙田区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810113, '西贡区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810114, '荃湾区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810115, '屯门区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810116, '元朗区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810117, '葵青区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (810118, '离岛区', 810000, 0, 0, 0, NULL, '2024-03-27 23:31:42', NULL, '2024-03-27 23:31:42');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820000, '澳门特别行政区', 0, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820101, '花地玛堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820102, '圣安多尼堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820103, '大堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820104, '望德堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820105, '风顺堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820106, '嘉模堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820107, '圣方济各堂区', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:39', NULL, '2024-03-27 23:31:39');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820108, '路氹城', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (820109, '澳门新城', 820000, 0, 0, 0, NULL, '2024-03-27 23:31:40', NULL, '2024-03-27 23:31:40');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830000, '台湾省', 0, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830100, '台北市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830101, '中正区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830102, '大同区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830103, '中山区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830104, '万华区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830105, '信义区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830106, '松山区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830107, '大安区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830108, '南港区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830109, '北投区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830110, '内湖区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830111, '士林区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830112, '文山区', 830100, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830200, '新北市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830201, '板桥区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830202, '土城区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830203, '新庄区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830204, '新店区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830205, '深坑区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830206, '石碇区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830207, '坪林区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830208, '乌来区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830209, '五股区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830210, '八里区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830211, '林口区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830212, '淡水区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830213, '中和区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830214, '永和区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830215, '三重区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830216, '芦洲区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830217, '泰山区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830218, '树林区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830219, '莺歌区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830220, '三峡区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830221, '汐止区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830222, '金山区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830223, '万里区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830224, '三芝区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830225, '石门区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830226, '瑞芳区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830227, '贡寮区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830228, '双溪区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830229, '平溪区', 830200, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830300, '桃园市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830301, '桃园区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830302, '中坜区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830303, '平镇区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830304, '八德区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830305, '杨梅区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830306, '芦竹区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830307, '大溪区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830308, '龙潭区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830309, '龟山区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830310, '大园区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830311, '观音区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830312, '新屋区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830313, '复兴区', 830300, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830400, '台中市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830401, '中区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830402, '东区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830403, '西区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830404, '南区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830405, '北区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830406, '西屯区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830407, '南屯区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830408, '北屯区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830409, '丰原区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830410, '大里区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830411, '太平区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830412, '东势区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830413, '大甲区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830414, '清水区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830415, '沙鹿区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830416, '梧栖区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830417, '后里区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830418, '神冈区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830419, '潭子区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830420, '大雅区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830421, '新小区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830422, '石冈区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830423, '外埔区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830424, '大安区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830425, '乌日区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830426, '大肚区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830427, '龙井区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830428, '雾峰区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830429, '和平区', 830400, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830500, '台南市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830501, '中西区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830502, '东区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830503, '南区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830504, '北区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830505, '安平区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830506, '安南区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830507, '永康区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830508, '归仁区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830509, '新化区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830510, '左镇区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830511, '玉井区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830512, '楠西区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830513, '南化区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830514, '仁德区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830515, '关庙区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830516, '龙崎区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830517, '官田区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830518, '麻豆区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830519, '佳里区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830520, '西港区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830521, '七股区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830522, '将军区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830523, '学甲区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830524, '北门区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830525, '新营区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830526, '后壁区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830527, '白河区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830528, '东山区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830529, '六甲区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830530, '下营区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830531, '柳营区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830532, '盐水区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830533, '善化区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830534, '大内区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830535, '山上区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830536, '新市区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830537, '安定区', 830500, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830600, '高雄市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830601, '楠梓区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830602, '左营区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830603, '鼓山区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830604, '三民区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830605, '盐埕区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830606, '前金区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830607, '新兴区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830608, '苓雅区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830609, '前镇区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830610, '旗津区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830611, '小港区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830612, '凤山区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830613, '大寮区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830614, '鸟松区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830615, '林园区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830616, '仁武区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830617, '大树区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830618, '大社区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830619, '冈山区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830620, '路竹区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830621, '桥头区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830622, '梓官区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830623, '弥陀区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830624, '永安区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830625, '燕巢区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830626, '阿莲区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830627, '茄萣区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830628, '湖内区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830629, '田寮区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830630, '旗山区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830631, '美浓区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830632, '内门区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830633, '杉林区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830634, '甲仙区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830635, '六龟区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830636, '茂林区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830637, '桃源区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830638, '那玛夏区', 830600, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830700, '基隆市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830701, '中正区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830702, '七堵区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830703, '暖暖区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830704, '仁爱区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830705, '中山区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830706, '安乐区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830707, '信义区', 830700, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830800, '新竹市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830801, '东区', 830800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830802, '北区', 830800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830803, '香山区', 830800, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830900, '嘉义市', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830901, '东区', 830900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (830902, '西区', 830900, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839001, '宜兰县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839002, '新竹县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839003, '苗栗县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839004, '彰化县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839005, '南投县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839006, '嘉义县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839007, '云林县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839008, '屏东县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839009, '台东县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839010, '花莲县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839011, '澎湖县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839012, '金门县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:31', NULL, '2024-03-27 23:31:31');
INSERT INTO `region` (`id`, `name`, `parent_id`, `status`, `deleted`, `sort`, `creator`, `create_time`, `updater`, `update_time`) VALUES (839013, '连江县', 830000, 0, 0, 0, NULL, '2024-03-27 23:31:32', NULL, '2024-03-27 23:31:32');
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
  `user_id` bigint DEFAULT '0' COMMENT '负责人',
  `phone` varchar(255) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(255) DEFAULT NULL COMMENT '邮件',
  `status` tinyint NOT NULL COMMENT '部门状态',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`status`,`parent_id`,`name`,`sort`)
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
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典数据表';

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
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (20, 3, '本部门数据权限', '3', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:33:20', '', '2024-06-05 11:43:31');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (21, 4, '本部门及以下数据权限', '4', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:34:28', '', '2024-03-12 11:34:28');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (22, 5, '仅本人数据权限', '5', 'role.scope', 0, 'default', '', '', 'admin', '2024-03-12 11:34:49', '', '2024-03-12 11:34:49');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (23, 1, '通知', '1', 'notice.type', 0, '', '', '', 'admin', '2024-04-24 14:54:14', '', '2024-04-24 14:54:14');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (24, 2, '公告', '2', 'notice.type', 0, '', '', '', 'admin', '2024-04-24 14:54:24', '', '2024-04-24 14:54:24');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (25, 1, '通知公告', '0', 'notifyTemplate.type', 0, 'primary', '', '', 'admin', '2024-04-28 15:47:31', 'admin', '2024-04-28 19:16:55');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (26, 2, '系统消息', '1', 'notifyTemplate.type', 0, 'success', '', '', 'admin', '2024-04-28 15:47:49', 'admin', '2024-04-28 19:16:42');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (27, 0, '管理员', '0', 'notifyMessage.userType', 0, 'default', '', '', 'admin', '2024-05-03 23:32:52', '', '2024-05-03 23:32:52');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (28, 1, '会员', '1', 'notifyMessage.userType', 0, 'success', '', '', 'admin', '2024-05-03 23:33:05', '', '2024-06-05 11:43:26');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (29, 0, '未读', '0', 'notifyMessage.readStatus', 0, 'success', '', '', 'admin', '2024-05-03 23:38:26', '', '2024-05-03 23:38:26');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (30, 1, '已读', '1', 'notifyMessage.readStatus', 0, 'default', '', '', 'admin', '2024-05-03 23:38:37', '', '2024-05-03 23:38:37');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (31, 0, '否', '0', 'mailAccount.ssl', 0, '', '', '', 'admin', '2024-05-08 15:42:55', '', '2024-05-08 15:42:55');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (32, 1, '是', '1', 'mailAccount.ssl', 0, '', '', '', 'admin', '2024-05-08 15:43:08', '', '2024-05-08 15:43:08');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (33, 0, '初始化', '0', 'mailLog.sendStatus', 0, 'primary', '', '', 'admin', '2024-05-08 18:20:56', 'admin', '2024-05-08 18:21:27');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (34, 0, '发送成功', '1', 'mailLog.sendStatus', 0, 'success', '', '', 'admin', '2024-05-08 18:21:19', '', '2024-05-08 18:21:19');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (35, 0, '发送失败', '2', 'mailLog.sendStatus', 0, 'danger', '', '', 'admin', '2024-05-08 18:21:45', '', '2024-05-08 18:21:45');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (36, 0, '不发送', '3', 'mailLog.sendStatus', 0, 'info', '', '', 'admin', '2024-05-08 18:22:00', '', '2024-05-08 18:22:00');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (37, 10, '模拟支付', 'mock', 'payChannel.code', 0, 'default', '', '', 'admin', '2024-05-13 23:33:27', 'admin', '2024-05-13 23:38:07');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (38, 1, '支付宝扫码支付', 'alipay_qr', 'payChannel.code', 0, 'primary', '', '', 'admin', '2024-05-13 23:35:05', '', '2024-05-13 23:35:06');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (39, 2, '支付宝条码支付', 'alipay_bar', 'payChannel.code', 0, 'primary', '', '', 'admin', '2024-05-13 23:35:26', '', '2024-05-13 23:35:26');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (40, 3, '支付宝App支付', 'alipay_app', 'payChannel.code', 0, 'primary', '', '', 'admin', '2024-05-13 23:35:46', 'admin', '2024-05-18 13:21:57');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (41, 4, '支付宝Wap网站支付', 'alipay_wap', 'payChannel.code', 0, 'primary', '', '', 'admin', '2024-05-13 23:36:08', 'admin', '2024-05-18 13:22:05');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (42, 5, '支付宝PC网站支付', 'alipay_pc', 'payChannel.code', 0, 'primary', '', '', 'admin', '2024-05-13 23:36:28', 'admin', '2024-05-18 13:22:15');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (43, 6, '微信App支付', 'wx_app', 'payChannel.code', 0, 'success', '', '', 'admin', '2024-05-13 23:37:09', 'admin', '2024-05-18 13:22:24');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (44, 7, '微信小程序支付', 'wx_lite', 'payChannel.code', 0, 'success', '', '', 'admin', '2024-05-13 23:37:29', '', '2024-05-13 23:37:29');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (45, 8, '微信公众号支付', 'wx_pub', 'payChannel.code', 0, 'success', '', '', 'admin', '2024-05-13 23:37:47', '', '2024-05-13 23:37:47');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (46, 0, '图片', '1', 'file.type', 0, '', '', '', 'admin', '2024-05-17 18:48:19', '', '2024-05-17 18:48:19');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (47, 0, '文档', '2', 'file.type', 0, '', '', '', 'admin', '2024-05-17 18:48:30', '', '2024-05-17 18:48:30');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (48, 0, '表格', '3', 'file.type', 0, '', '', '', 'admin', '2024-05-17 18:48:43', '', '2024-05-17 18:48:43');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (49, 0, 'PPT', '4', 'file.type', 0, '', '', '', 'admin', '2024-05-17 18:48:54', '', '2024-05-17 18:48:54');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (50, 0, 'PDF', '5', 'file.type', 0, '', '', '', 'admin', '2024-05-17 18:49:05', '', '2024-05-17 18:49:05');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (51, 0, '等待支付', '0', 'pay.orderStatus', 0, 'info', '', '', 'admin', '2024-05-29 21:15:29', '', '2024-05-29 21:15:29');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (52, 1, '支付成功', '1', 'pay.orderStatus', 0, 'success', '', '', 'admin', '2024-05-29 21:15:44', '', '2024-05-29 21:15:44');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (53, 2, '已退款', '2', 'pay.orderStatus', 0, 'danger', '', '', 'admin', '2024-05-29 21:16:04', '', '2024-05-29 21:16:04');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (54, 3, '已退款', '3', 'pay.orderStatus', 0, 'info', '', '', 'admin', '2024-05-29 21:16:19', '', '2024-05-29 21:16:19');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (55, 1, '等待退款', '0', 'pay.refundStatus', 0, 'info', '', '', 'admin', '2024-05-30 09:08:05', '', '2024-05-30 09:08:05');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (56, 2, '退款成功', '1', 'pay.refundStatus', 0, 'success', '', '', 'admin', '2024-05-30 09:08:26', '', '2024-05-30 09:08:26');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (57, 3, '退款失败', '2', 'pay.refundStatus', 0, 'danger', '', '', 'admin', '2024-05-30 09:08:44', '', '2024-05-30 09:08:44');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (58, 0, '等待通知', '0', 'pay.notifyStatus', 0, 'info', '', '', 'admin', '2024-06-05 11:37:03', '', '2024-06-05 11:37:03');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (59, 1, '通知成功', '1', 'pay.notifyStatus', 0, 'success', '', '', 'admin', '2024-06-05 11:37:18', '', '2024-06-05 11:37:18');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (60, 2, '通知失败', '2', 'pay.notifyStatus', 0, 'danger', '', '', 'admin', '2024-06-05 11:37:35', '', '2024-06-05 11:37:35');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (61, 3, '请求成功结果失败', '3', 'pay.notifyStatus', 0, 'warning', '', '', 'admin', '2024-06-05 11:38:01', '', '2024-06-05 11:38:01');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (62, 4, '请求失败', '4', 'pay.notifyStatus', 0, 'warning', '', '', 'admin', '2024-06-05 11:38:21', '', '2024-06-05 11:38:21');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (63, 0, '支付订单', '1', 'pay.notifyType', 0, 'primary', '', '', 'admin', '2024-06-05 11:40:13', '', '2024-06-05 11:43:03');
INSERT INTO `system_dict` (`id`, `sort`, `label`, `value`, `dict_type`, `status`, `color_type`, `css_class`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (64, 0, '退款订单', '2', 'pay.notifyType', 0, 'default', '', '', 'admin', '2024-06-05 11:40:25', '', '2024-06-05 11:42:56');
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
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典类型';

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
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (11, '公告类型', 'notice.type', 0, '', 'admin', '2024-04-24 14:53:36', '', '2024-04-24 14:53:36');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (12, '模板类型', 'notifyTemplate.type', 0, '', 'admin', '2024-04-28 15:46:53', '', '2024-04-28 15:46:53');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (13, '用户类型', 'notifyMessage.userType', 0, '', 'admin', '2024-05-03 23:32:08', 'admin', '2024-05-03 23:32:24');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (14, '阅读状态', 'notifyMessage.readStatus', 0, '', 'admin', '2024-05-03 23:38:03', '', '2024-05-03 23:38:03');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (15, '是否开启 SSL', 'mailAccount.ssl', 0, '', 'admin', '2024-05-08 15:42:14', '', '2024-05-08 15:42:14');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (16, '邮件发送状态', 'mailLog.sendStatus', 0, '', 'admin', '2024-05-08 18:20:29', '', '2024-05-08 18:20:29');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (17, '支付渠道', 'payChannel.code', 0, '', 'admin', '2024-05-13 23:32:45', '', '2024-05-13 23:32:45');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (18, '文件类型', 'file.type', 0, '', 'admin', '2024-05-17 18:47:43', '', '2024-05-17 18:47:43');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (19, '支付状态', 'pay.orderStatus', 0, '', 'admin', '2024-05-29 21:15:01', '', '2024-05-29 21:15:01');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (20, '退款订单状态', 'pay.refundStatus', 0, '', 'admin', '2024-05-30 09:07:36', '', '2024-05-30 09:07:36');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (21, '支付通知状态', 'pay.notifyStatus', 0, '', 'admin', '2024-06-05 11:36:36', '', '2024-06-05 11:36:36');
INSERT INTO `system_dict_type` (`id`, `name`, `type`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`) VALUES (22, '支付通知类型', 'pay.notifyType', 0, '', 'admin', '2024-06-05 11:39:44', '', '2024-06-05 11:42:35');
COMMIT;

-- ----------------------------
-- Table structure for system_file
-- ----------------------------
DROP TABLE IF EXISTS `system_file`;
CREATE TABLE `system_file` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名称',
  `file_type` tinyint NOT NULL DEFAULT '0' COMMENT '文件类型',
  `file_mime_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Mime类型',
  `file_size` bigint NOT NULL DEFAULT '0' COMMENT '文件大小',
  `file_path` varchar(255) NOT NULL COMMENT '文件路径',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`file_type`,`file_name`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件管理';

-- ----------------------------
-- Records of system_file
-- ----------------------------
BEGIN;
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
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`username`,`login_time`,`channel`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='登录日志';

-- ----------------------------
-- Records of system_login_log
-- ----------------------------
BEGIN;

COMMIT;

-- ----------------------------
-- Table structure for system_mail_account
-- ----------------------------
DROP TABLE IF EXISTS `system_mail_account`;
CREATE TABLE `system_mail_account` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `mail` varchar(255) NOT NULL COMMENT '邮箱',
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `host` varchar(255) NOT NULL COMMENT 'SMTP 服务器域名',
  `port` int NOT NULL COMMENT 'SMTP 服务器端口',
  `ssl_enable` tinyint NOT NULL DEFAULT '0' COMMENT '是否开启 SSL',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`mail`,`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='邮箱账号表';

-- ----------------------------
-- Records of system_mail_account
-- ----------------------------
BEGIN;
INSERT INTO `system_mail_account` (`id`, `mail`, `username`, `password`, `host`, `port`, `ssl_enable`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (1, 'asda', 'asdasd', 'asdasd', 'asdas', 0, 0, 'admin', '2024-05-08 16:11:54', NULL, '2024-05-08 16:11:54', 0);
COMMIT;

-- ----------------------------
-- Table structure for system_mail_log
-- ----------------------------
DROP TABLE IF EXISTS `system_mail_log`;
CREATE TABLE `system_mail_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint DEFAULT NULL COMMENT '用户编号',
  `user_type` tinyint DEFAULT NULL COMMENT '用户类型',
  `to_mail` varchar(255) NOT NULL COMMENT '接收邮箱地址',
  `account_id` bigint NOT NULL COMMENT '邮箱账号编号',
  `from_mail` varchar(255) NOT NULL COMMENT '发送邮箱地址',
  `template_id` bigint NOT NULL COMMENT '模板编号',
  `template_code` varchar(64) NOT NULL COMMENT '模板编码',
  `template_nickname` varchar(255) DEFAULT NULL COMMENT '模版发送人名称',
  `template_title` varchar(255) NOT NULL COMMENT '邮件标题',
  `template_content` varchar(10240) NOT NULL COMMENT '邮件内容',
  `template_params` json DEFAULT NULL COMMENT '邮件参数',
  `send_status` tinyint NOT NULL DEFAULT '0' COMMENT '发送状态',
  `send_time` datetime DEFAULT NULL COMMENT '发送时间',
  `send_message_id` varchar(255) DEFAULT NULL COMMENT '发送返回的消息 ID',
  `send_exception` varchar(4096) DEFAULT NULL COMMENT '发送异常',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`send_status`,`send_time`,`template_title`,`user_id`,`user_type`,`account_id`,`template_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='邮件日志表';

-- ----------------------------
-- Records of system_mail_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_mail_template
-- ----------------------------
DROP TABLE IF EXISTS `system_mail_template`;
CREATE TABLE `system_mail_template` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '模板名称',
  `code` varchar(64) NOT NULL COMMENT '模板编码',
  `account_id` bigint NOT NULL COMMENT '发送的邮箱账号编号',
  `nickname` varchar(255) DEFAULT NULL COMMENT '发送人名称',
  `title` varchar(255) NOT NULL COMMENT '模板标题',
  `content` varchar(10240) NOT NULL COMMENT '模板内容',
  `params` json DEFAULT NULL COMMENT '参数数组',
  `status` tinyint NOT NULL COMMENT '开启状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`status`,`title`,`name`,`code`,`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='邮件模版表';

-- ----------------------------
-- Records of system_mail_template
-- ----------------------------
BEGIN;
INSERT INTO `system_mail_template` (`id`, `name`, `code`, `account_id`, `nickname`, `title`, `content`, `params`, `status`, `remark`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (1, '阿萨德', ' sss', 1, 'sss', '', 'ssss', '{}', 0, 'ssssss', 'admin', '2024-05-08 18:15:35', NULL, '2024-05-08 18:15:36', 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统菜单';

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
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (68, '审计日志', '', 1, 4, 0, '/system/logger', 'Tickets', '', '', 0, 0, '', 1, 0, '', 0, NULL, 'admin', '2023-08-09 22:03:11', 'admin', '2024-05-13 23:38:59', 0);
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
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (108, '租户管理', '', 2, 19, 2, '/system/tenant', 'Folder', '', '', 0, 0, '', 1, 0, '', 0, '/system/tenant/index', '芋道源码', '2023-07-11 08:32:02', 'admin', '2024-04-24 14:39:06', 0);
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
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (127, '文件管理', '', 2, 26, 2, '/system/file', 'Files', '/system/file/index', 'systemFile', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-03-23 15:44:21', 'admin', '2024-05-17 18:42:05', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (128, '地区管理', '', 2, 26, 2, '/region', 'MapLocation', '/region/index', 'region', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-04-05 16:46:37', 'admin', '2024-04-05 17:02:50', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (129, '地区列表', 'region.RegionList', 3, 0, 128, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-05 16:48:10', '', '2024-04-05 16:48:10', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (130, '地区新增', 'region.RegionCreate', 3, 0, 128, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-05 16:48:41', '', '2024-04-05 16:48:41', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (131, '地区编辑', 'region.RegionUpdate', 3, 0, 128, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-05 16:49:05', '', '2024-04-05 16:49:05', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (132, '地区删除', 'region.RegionDelete', 3, 0, 128, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-05 16:49:26', '', '2024-04-05 16:49:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (133, '地区恢复', 'region.RegionRecover', 3, 0, 128, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-05 16:49:46', '', '2024-04-05 16:49:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (134, '公告管理', '', 2, 26, 2, '/system/notice', 'Notification', '/system/notice/index', 'systemNotice', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-04-24 14:28:46', '', '2024-04-24 14:28:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (135, '公告列表', 'notice.SystemNoticeList', 3, 0, 134, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-24 14:29:56', '', '2024-04-24 14:29:56', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (136, '公告新增', 'notice.SystemNoticeCreate', 3, 0, 134, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-24 14:30:36', '', '2024-04-24 14:30:36', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (137, '公告编辑', 'notice.SystemNoticeUpdate', 3, 0, 134, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-24 14:31:34', '', '2024-04-24 14:31:34', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (138, '公告删除', 'notice.SystemNoticeDelete', 3, 0, 134, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-24 14:31:57', '', '2024-04-24 14:31:57', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (139, '公告恢复', 'notice.SystemNoticeRecover', 3, 0, 134, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-24 14:32:23', '', '2024-04-24 14:32:23', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (140, '站内信管理', '', 2, 27, 2, '/system/notify', 'Folder', '', '', 0, 0, '', 1, 0, '', 0, '/system/notify/message', 'admin', '2024-04-28 15:49:21', 'admin', '2024-04-28 15:50:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (141, '模板管理', '', 2, 0, 140, '/system/notify/template', 'Cellphone', '/system/notify/template', 'systemNotifyTemplate', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-04-28 15:52:30', '', '2024-04-28 15:52:30', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (142, '模板列表', 'notify.SystemNotifyTemplateList', 3, 0, 141, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 15:56:19', '', '2024-04-28 15:56:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (143, '模板新增', 'notify.SystemNotifyTemplateCreate', 3, 0, 141, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 15:58:07', '', '2024-04-28 15:58:07', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (144, '模板编辑', 'notify.SystemNotifyTemplateUpdate', 3, 0, 141, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 15:59:16', '', '2024-04-28 15:59:16', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (145, '模板删除', 'notify.SystemNotifyTemplateDelete', 3, 0, 141, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 15:59:43', '', '2024-04-28 15:59:43', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (146, '模板恢复', 'notify.SystemNotifyTemplateRecover', 3, 0, 141, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 16:00:29', '', '2024-04-28 16:00:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (147, '消息管理', '', 2, 0, 140, '/system/notify/message', 'Message', '/system/notify/message', 'systemNotifyMessage', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-04-28 16:02:54', '', '2024-04-28 16:02:54', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (148, '消息列表', 'notify.SystemNotifyMessageList', 3, 0, 147, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 15:56:19', '', '2024-04-28 15:56:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (151, '消息删除', 'notify.SystemNotifyMessageDelete', 3, 0, 147, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 15:59:43', '', '2024-04-28 15:59:43', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (152, '消息恢复', 'notify.SystemNotifyMessageRecover', 3, 0, 147, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-04-28 16:00:29', '', '2024-04-28 16:00:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (153, '消息查看', 'notify.SystemNotifyTemplate', 3, 0, 147, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 09:21:46', '', '2024-05-08 09:21:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (154, '邮箱管理', '', 2, 28, 2, '/system/mail', 'Folder', '', '', 0, 0, '', 1, 0, '', 0, '/system/mail/log', 'admin', '2024-05-08 09:26:06', 'admin', '2024-05-08 10:07:49', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (155, '账号管理', '', 2, 1, 154, '/system/mail/account', 'Avatar', '/system/mail/account', 'systemMailAccount', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-05-08 09:28:46', '', '2024-05-08 09:28:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (156, '模板管理', '', 2, 1, 154, '/system/mail/template', 'Cellphone', '/system/mail/template', 'systemMailTemplate', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-05-08 10:10:19', '', '2024-05-08 10:10:19', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (157, '记录管理', '', 2, 2, 154, '/system/mail/log', 'Message', '/system/mail/log', 'systemMailLog', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-05-08 10:11:42', '', '2024-05-08 10:11:42', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (158, '账号列表', 'mail.SystemMailAccountList', 3, 0, 155, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:12:45', '', '2024-05-08 10:12:45', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (159, '账号新增', 'mail.SystemMailAccountCreate', 3, 0, 155, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:13:11', 'admin', '2024-05-08 10:13:56', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (160, '账号编辑', 'mail.SystemMailAccountUpdate', 3, 0, 155, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:13:37', '', '2024-05-08 10:13:37', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (161, '账号删除', 'mail.SystemMailAccountDelete', 3, 0, 155, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:14:25', '', '2024-05-08 10:14:25', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (162, '账号恢复', 'mail.SystemMailAccountRecover', 3, 0, 155, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:14:46', '', '2024-05-08 10:14:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (163, '模板列表', 'mail.SystemMailTemplateList', 3, 0, 156, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:15:08', '', '2024-05-08 10:15:08', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (164, '模板新增', 'mail.SystemMailTemplateCreate', 3, 0, 156, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:15:33', '', '2024-05-08 10:15:33', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (165, '模板编辑', 'mail.SystemMailTemplateUpdate', 3, 0, 156, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:15:53', '', '2024-05-08 10:15:53', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (166, '模板删除', 'mail.SystemMailTemplateDelete', 3, 0, 156, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:16:11', '', '2024-05-08 10:16:11', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (167, '模板恢复', 'mail.SystemMailTemplateRecover', 3, 0, 156, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:16:31', '', '2024-05-08 10:16:31', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (168, '记录列表', 'mail.SystemMailLogList', 3, 0, 157, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:16:57', '', '2024-05-08 10:16:57', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (169, '记录删除', 'mail.SystemMailLogDelete', 3, 0, 157, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:17:30', '', '2024-05-08 10:17:30', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (170, '记录恢复', 'mail.SystemMailLogRecover', 3, 0, 157, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:17:56', '', '2024-05-08 10:17:56', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (171, '记录查看', 'mail.SystemMailLog', 3, 0, 157, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-08 10:18:30', '', '2024-05-08 10:18:30', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (172, '支付管理', '', 1, 3, 0, '/pay', 'Money', '', '', 0, 0, '', 0, 0, '', 0, '/pay/app', 'admin', '2024-05-13 23:40:09', 'admin', '2024-05-13 23:45:05', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (173, '应用管理', '', 2, 1, 172, '/pay/app', 'Cellphone', '/pay/app/index', 'payApp', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-05-13 23:42:50', 'admin', '2024-05-13 23:44:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (174, '应用列表', 'app.PayAppList', 3, 0, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:47:17', '', '2024-05-13 23:47:17', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (175, '应用新增', 'app.PayAppCreate', 3, 1, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:47:49', 'admin', '2024-05-13 23:48:25', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (176, '应用编辑', 'app.PayAppUpdate', 3, 0, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:48:50', '', '2024-05-13 23:48:50', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (177, '应用删除', 'app.PayAppDelete', 3, 0, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:49:15', '', '2024-05-13 23:49:15', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (178, '应用恢复', 'app.PayAppRecover', 3, 0, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:49:50', '', '2024-05-13 23:49:50', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (179, '渠道查看', 'channel.PayChannel', 3, 0, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:50:47', '', '2024-05-13 23:50:47', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (180, '渠道设置', 'channel.PayChannelCreate', 3, 0, 173, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-13 23:51:10', '', '2024-05-13 23:51:10', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (181, '文件列表', 'file.SystemFileList', 3, 0, 127, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-17 18:42:46', '', '2024-05-17 18:42:46', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (182, '文件新增', 'file.SystemFileCreate', 3, 0, 127, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-17 18:43:07', '', '2024-05-17 18:43:07', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (183, '文件编辑', 'file.SystemFileUpdate', 3, 0, 127, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-17 18:43:25', '', '2024-05-17 18:43:25', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (184, '文件删除', 'file.SystemFileDelete', 3, 0, 127, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-17 18:43:41', '', '2024-05-17 18:43:41', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (185, '支付订单', '', 2, 1, 172, '/pay/order', 'TakeawayBox', '/pay/order/index', 'payOrder', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-05-29 21:21:54', '', '2024-05-29 21:21:54', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (186, '支付订单列表', 'order.PayOrderList', 3, 0, 185, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-29 21:22:14', 'admin', '2024-05-29 21:22:44', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (187, '支付订单删除', 'order.PayOrderDelete', 3, 0, 185, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-29 21:23:08', '', '2024-05-29 21:23:08', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (188, '支付订单恢复', 'order.PayOrderRecover', 3, 0, 185, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-29 21:23:30', '', '2024-05-29 21:23:30', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (189, '退款订单', '', 2, 1, 172, '/pay/refund', 'Wallet', '/pay/refund/index', 'payRefund', 0, 0, '', 1, 0, '', 0, '', 'admin', '2024-05-30 09:12:20', '', '2024-05-30 09:12:20', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (190, '退款订单列表', 'refund.PayRefundList', 3, 0, 189, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-30 09:13:29', '', '2024-05-30 09:13:29', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (191, '退款订单删除', 'refund.PayRefundDelete', 3, 0, 189, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-30 09:13:53', '', '2024-05-30 09:13:53', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (192, '退款订单恢复', 'order.PayOrderRecover', 3, 0, 189, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-05-30 09:14:11', '', '2024-05-30 09:14:11', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (193, '通知回调', '', 2, 4, 172, '/pay/notify', 'DocumentCopy', '/pay/notify/index', '', 0, 0, '', 1, 0, 'payNotifyTask', 0, '', 'admin', '2024-06-05 12:18:25', '', '2024-06-05 12:18:26', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (194, '通知回调列表', 'notify.PayNotifyTaskList', 3, 0, 193, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-06-05 12:19:08', '', '2024-06-05 12:19:08', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (195, '通知回调删除', 'notify.PayNotifyTaskDelete', 3, 0, 193, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-06-05 12:19:33', '', '2024-06-05 12:19:33', 0);
INSERT INTO `system_menu` (`id`, `name`, `permission`, `type`, `sort`, `parent_id`, `path`, `icon`, `component`, `component_name`, `status`, `hide`, `link`, `keep_alive`, `affix`, `active_path`, `full_screen`, `redirect`, `creator`, `create_time`, `updater`, `update_time`, `deleted`) VALUES (196, '通知回调恢复', 'notify.PayNotifyTaskRecover', 3, 0, 193, '', '', '', '', 0, 0, '', 0, 0, '', 0, '', 'admin', '2024-06-05 12:19:55', '', '2024-06-05 12:19:55', 0);
COMMIT;

-- ----------------------------
-- Table structure for system_notice
-- ----------------------------
DROP TABLE IF EXISTS `system_notice`;
CREATE TABLE `system_notice` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `title` varchar(64) NOT NULL COMMENT '公告标题',
  `content` text NOT NULL COMMENT '公告内容',
  `type` tinyint NOT NULL COMMENT '公告类型（1通知 2公告）',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`status`,`type`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='通知公告表';

-- ----------------------------
-- Records of system_notice
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_notify_message
-- ----------------------------
DROP TABLE IF EXISTS `system_notify_message`;
CREATE TABLE `system_notify_message` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '消息',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `user_type` tinyint NOT NULL COMMENT '用户类型',
  `template_id` bigint NOT NULL COMMENT '模版编号',
  `template_code` varchar(64) NOT NULL COMMENT '模板编码',
  `template_nickname` varchar(64) NOT NULL COMMENT '模版发送人名称',
  `template_content` varchar(1024) NOT NULL COMMENT '模版内容',
  `template_type` int NOT NULL COMMENT '模版类型',
  `template_params` json NOT NULL COMMENT '模版参数',
  `read_status` tinyint NOT NULL COMMENT '是否已读',
  `read_time` datetime DEFAULT NULL COMMENT '阅读时间',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`template_type`,`read_status`,`read_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='站内信消息表';

-- ----------------------------
-- Records of system_notify_message
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_notify_template
-- ----------------------------
DROP TABLE IF EXISTS `system_notify_template`;
CREATE TABLE `system_notify_template` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL COMMENT '模板名称',
  `code` varchar(64) NOT NULL COMMENT '模版编码',
  `nickname` varchar(255) NOT NULL COMMENT '发送人名称',
  `content` varchar(1024) NOT NULL COMMENT '模版内容',
  `type` tinyint NOT NULL COMMENT '类型',
  `params` json DEFAULT NULL COMMENT '参数数组',
  `status` tinyint NOT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`deleted`,`status`,`type`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='站内信模板表';

-- ----------------------------
-- Records of system_notify_template
-- ----------------------------
BEGIN;
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
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`username`,`module`,`start_time`,`result`)
) ENGINE=InnoDB AUTO_INCREMENT=1029 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志';

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
  `tenant_id` bigint NOT NULL COMMENT '租户ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`status`,`name`,`sort`)
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
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`type`,`status`,`sort`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统角色';

-- ----------------------------
-- Records of system_role
-- ----------------------------
BEGIN;
INSERT INTO `system_role` (`id`, `name`, `code`, `sort`, `data_scope`, `data_scope_dept`, `status`, `type`, `remark`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '超级管理员', 'super', 0, NULL, NULL, 0, 2, '', 1, 1, 'admin', '2024-03-23 14:02:42', NULL, '2024-03-23 14:02:42');
INSERT INTO `system_role` (`id`, `name`, `code`, `sort`, `data_scope`, `data_scope_dept`, `status`, `type`, `remark`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (5, '测试', 'test', 1, NULL, NULL, 0, 2, '', 0, 1, 'admin', '2024-03-30 20:41:48', NULL, '2024-03-30 20:41:48');
COMMIT;

-- ----------------------------
-- Table structure for system_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_role_menu`;
CREATE TABLE `system_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `role_id` bigint NOT NULL COMMENT '角色编号',
  `menu_id` bigint NOT NULL COMMENT '菜单编号',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:role_menu` (`role_id`,`menu_id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`menu_id`)
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
  `user_id` bigint DEFAULT NULL COMMENT '联系人ID',
  `contact_name` varchar(255) NOT NULL COMMENT '联系人',
  `contact_mobile` varchar(255) NOT NULL COMMENT '租户联系电话',
  `status` tinyint NOT NULL COMMENT '状态（0正常 1停用）',
  `domain` varchar(255) DEFAULT NULL COMMENT '域名',
  `expire_date` date NOT NULL COMMENT '过期时间',
  `account_count` bigint NOT NULL COMMENT '账号数量',
  `tenant_package_id` bigint NOT NULL COMMENT '套餐编号',
  `deleted` tinyint NOT NULL COMMENT '是否删除(0否 1是)',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx:package` (`tenant_package_id`),
  KEY `idx:list` (`deleted`,`status`,`name`,`expire_date`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户';

-- ----------------------------
-- Records of system_tenant
-- ----------------------------
BEGIN;
INSERT INTO `system_tenant` (`id`, `name`, `user_id`, `contact_name`, `contact_mobile`, `status`, `domain`, `expire_date`, `account_count`, `tenant_package_id`, `deleted`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '系统租户', 1, '周先生', '13512345678', 0, 'http://www.baidu.com', '2029-03-01', 100, 0, 0, NULL, '2024-03-20 21:14:57', 'admin', '2024-04-21 19:04:41');
COMMIT;

-- ----------------------------
-- Table structure for system_tenant_package
-- ----------------------------
DROP TABLE IF EXISTS `system_tenant_package`;
CREATE TABLE `system_tenant_package` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '套餐编号',
  `name` varchar(255) NOT NULL COMMENT '套餐名称',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `menu_ids` json NOT NULL COMMENT '目录编号',
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
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `item:login` (`username`),
  KEY `idx:list` (`tenant_id`,`deleted`,`status`,`nickname`,`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户';

-- ----------------------------
-- Records of system_user
-- ----------------------------
BEGIN;
INSERT INTO `system_user` (`id`, `nickname`, `mobile`, `username`, `password`, `status`, `deleted`, `tenant_id`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, '安安', '15397620000', 'admin', '21232f297a57a5a743894a0e4a801fc3', 0, 0, 1, NULL, '2024-03-20 21:15:35', 'admin', '2024-04-21 19:03:45');
COMMIT;

-- ----------------------------
-- Table structure for system_user_dept
-- ----------------------------
DROP TABLE IF EXISTS `system_user_dept`;
CREATE TABLE `system_user_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '系统用户 id',
  `dept_id` bigint NOT NULL COMMENT '部门 id',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uniq:user_dept` (`user_id`,`dept_id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`dept_id`)
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
  `user_id` bigint NOT NULL COMMENT '系统用户 ID',
  `post_id` bigint NOT NULL COMMENT '职位 id',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:user_post` (`user_id`,`post_id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`post_id`)
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
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `role_id` bigint NOT NULL COMMENT '角色编号',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq:user_role` (`user_id`,`role_id`),
  KEY `idx:list` (`tenant_id`,`deleted`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户和系统角色关联表';

-- ----------------------------
-- Records of system_user_role
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for system_user_tenant
-- ----------------------------
DROP TABLE IF EXISTS `system_user_tenant`;
CREATE TABLE `system_user_tenant` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '系统用户 ID',
  `tenant_id` bigint NOT NULL COMMENT '租户 id',
  `deleted` tinyint NOT NULL COMMENT '删除',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uniq:user_tenant` (`user_id`,`tenant_id`),
  KEY `idx:tenant` (`tenant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统租户用户';

-- ----------------------------
-- Records of system_user_tenant
-- ----------------------------
BEGIN;
INSERT INTO `system_user_tenant` (`id`, `user_id`, `tenant_id`, `deleted`, `creator`, `create_time`, `updater`, `update_time`) VALUES (1, 1, 1, 0, NULL, '2024-03-31 20:33:48', NULL, '2024-04-21 19:04:14');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
