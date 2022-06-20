/*
 Navicat Premium Data Transfer

 Source Server         : docker 本地
 Source Server Type    : MySQL
 Source Server Version : 100422
 Source Host           : localhost:3306
 Source Schema         : binance

 Target Server Type    : MySQL
 Target Server Version : 100422
 File Encoding         : 65001

 Date: 18/06/2022 10:11:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for collects
-- ----------------------------
DROP TABLE IF EXISTS `collects`;
CREATE TABLE `collects` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `symbol` varchar(255) NOT NULL DEFAULT '' COMMENT '代币符号',
  `contract_address` varchar(255) NOT NULL DEFAULT '',
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `user_address` varchar(255) NOT NULL DEFAULT '' COMMENT '用户地址',
  `recharge_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '充值hash',
  `handfee_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '手续费hash',
  `collect_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '归集hash',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '归集金额',
  `recharge_value` varchar(255) NOT NULL DEFAULT '' COMMENT '充值金额',
  `status` varchar(20) NOT NULL DEFAULT '' COMMENT '状态 fail 失败，wait_fee 待转手续费，process_fee 转手续费中，wait_collect 待归集,process_wait归集中，wait_notify 待通知,finish_notify通知完成',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `recharge_hash_idx` (`recharge_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='归集表';

-- ----------------------------
-- Records of collects
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for contracts
-- ----------------------------
DROP TABLE IF EXISTS `contracts`;
CREATE TABLE `contracts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(255) NOT NULL COMMENT '货币类型',
  `address` varchar(255) NOT NULL COMMENT '合约地址',
  `decimals` int(11) NOT NULL DEFAULT 0 COMMENT '小数位数',
  `is_collect_open` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启,1是 0否',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='合约表';

-- ----------------------------
-- Records of contracts
-- ----------------------------
BEGIN;
INSERT INTO `contracts` VALUES (1, 'BNB', '0x', 18, 0, '2022-05-30 09:13:36', '2022-06-17 15:43:14');
INSERT INTO `contracts` VALUES (2, 'BSC-USD', '0x55d398326f99059fF775485246999027B3197955', 18, 1, '2022-05-30 09:13:36', '2022-05-30 09:13:36');
INSERT INTO `contracts` VALUES (3, 'MTG-USD', '0x2151F2B84134C6df6690E8E3E11AEf1AC3594145', 18, 1, '2022-06-05 08:55:58', '2022-06-05 08:55:58');
COMMIT;

-- ----------------------------
-- Table structure for ga_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `ga_admin_log`;
CREATE TABLE `ga_admin_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '请求路径',
  `method` varchar(10) NOT NULL DEFAULT '' COMMENT '请求方法',
  `path_name` varchar(255) NOT NULL DEFAULT '' COMMENT '请求路径名称',
  `params` text DEFAULT NULL COMMENT '请求参数',
  `response` longtext DEFAULT NULL COMMENT '响应结果',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of ga_admin_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ga_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `ga_admin_menu`;
CREATE TABLE `ga_admin_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '菜单名称',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '前端路由地址，可以是外链',
  `parent_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父id',
  `identification` varchar(40) NOT NULL DEFAULT '' COMMENT '后端权限标识符',
  `method` varchar(10) NOT NULL DEFAULT '' COMMENT '请求方法',
  `front_component_path` varchar(255) DEFAULT NULL COMMENT '前端组件路径',
  `icon` varchar(100) DEFAULT '#' COMMENT '菜单图标',
  `sort` tinyint(4) NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
  `status` varchar(10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `type` varchar(12) NOT NULL COMMENT '菜单类型',
  `link_type` varchar(12) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='后台菜单表';

-- ----------------------------
-- Records of ga_admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `ga_admin_menu` VALUES (1, '系统管理', 'system', 0, '', '', '', 'system', 0, 'normal', '2022-01-09 19:44:50', '2022-02-21 16:06:17', 'directory', 'internal');
INSERT INTO `ga_admin_menu` VALUES (2, '管理员列表', 'administrator', 1, '/administrator-list', 'get', 'system/administrator/index', 'user', 0, 'normal', '2022-01-09 19:46:32', '2022-02-23 16:51:44', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (3, '角色列表', 'role', 1, '/role-list', 'get', 'system/role/index', 'peoples', 0, 'normal', '2022-01-09 19:47:59', '2022-02-23 17:02:24', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (4, '角色新增', '', 3, '/role-store', 'post', '', '', 0, 'normal', '2022-01-09 19:49:15', '2022-01-09 19:49:15', 'operation', '');
INSERT INTO `ga_admin_menu` VALUES (5, '角色更新', '', 3, '/role-update', 'put', '', '', 0, 'normal', '2022-01-09 19:49:36', '2022-01-09 19:49:36', 'operation', '');
INSERT INTO `ga_admin_menu` VALUES (6, '角色删除', '', 3, '/role-destroy', 'delete', '', '', 0, 'normal', '2022-01-09 19:49:46', '2022-01-09 19:49:46', 'operation', '');
INSERT INTO `ga_admin_menu` VALUES (7, '管理员新增', '', 2, '/administrator-store', 'post', '', '', 0, 'normal', '2022-01-09 19:50:30', '2022-02-21 15:47:38', 'operation', '');
INSERT INTO `ga_admin_menu` VALUES (8, '管理员删除', '', 2, '/administrator-destroy', 'delete', '', '', 0, 'normal', '2022-01-09 19:54:44', '2022-01-09 19:54:44', 'operation', '');
INSERT INTO `ga_admin_menu` VALUES (9, '管理员更新', '', 2, '/administrator-update', 'put', '', '', 0, 'normal', '2022-01-09 19:55:13', '2022-01-09 19:55:13', 'operation', '');
INSERT INTO `ga_admin_menu` VALUES (10, '菜单列表', 'menu', 1, '/menu-list', 'get', 'system/menu/index', 'tree-table', 0, 'normal', NULL, NULL, 'link', 'tab');
INSERT INTO `ga_admin_menu` VALUES (11, '菜单新增', '', 10, '/menu-store', 'POST', '', '#', 0, 'normal', NULL, '2022-02-21 16:10:46', 'operation', 'tab');
INSERT INTO `ga_admin_menu` VALUES (12, '菜单更新', '', 10, '/menu-update', 'PUT', '', '#', 0, 'normal', NULL, '2022-02-21 16:10:54', 'operation', 'tab');
INSERT INTO `ga_admin_menu` VALUES (13, '菜单删除', '', 10, '/menu-destroy', 'DELETE', '', '#', 0, 'normal', NULL, '2022-02-23 16:55:40', 'operation', 'tab');
INSERT INTO `ga_admin_menu` VALUES (15, 'aaa', 'test', 14, 'fffff', 'GET', '', '', 1, 'normal', '2022-02-18 13:31:26', '2022-02-21 17:34:05', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (18, '管理员信息', '', 2, '/administrator-info', 'GET', '', '', 1, 'normal', '2022-02-21 16:12:38', '2022-02-23 17:32:55', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (19, '角色信息', '', 3, '/role-info', 'GET', '', '', 1, 'normal', '2022-02-21 16:18:39', '2022-02-23 15:17:30', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (20, '菜单信息', '', 10, '/menu-info', 'GET', '', '', 1, 'normal', '2022-02-21 16:19:16', '2022-02-23 15:05:50', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (22, '配置管理', 'config', 1, '/config-list', 'GET', 'system/config/index', 'edit', 0, 'normal', '2022-03-19 17:40:06', '2022-04-27 18:39:20', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (23, '配置更新', '', 22, '/config-update', 'PUT', '', '', 1, 'normal', '2022-04-28 14:29:18', '2022-04-28 14:32:53', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (24, '操作日志', '/operationlog', 1, '/operation-log-list', 'GET', 'system/operationlog/index', 'documentation', 6, 'normal', '2022-05-05 12:02:45', '2022-05-05 15:14:38', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (26, '出入金管理', 'binance', 0, '', 'GET', '', 'money', 5, 'normal', '2022-06-15 16:03:24', '2022-06-15 16:03:24', 'directory', 'internal');
INSERT INTO `ga_admin_menu` VALUES (27, '用户钱包地址', 'binance_user_address', 26, '/binance-user-address-list', 'GET', 'binance/useraddress/index', 'list', 1, 'normal', '2022-06-15 16:09:12', '2022-06-16 15:43:17', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (28, '归集列表', 'collect-list', 26, '/binance-collect-list', 'GET', 'binance/collect/index', 'list', 3, 'normal', '2022-06-16 14:53:57', '2022-06-16 15:42:56', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (29, '提现列表', 'bianance-withdraw', 26, '/binance-withdraw-list', 'GET', 'binance/withdraw/index', 'list', 4, 'normal', '2022-06-16 16:01:32', '2022-06-16 16:01:32', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (30, '转账队列', 'binance-queue-task-list', 26, '/binance-queue-task-list', 'GET', 'binance/queue_task/index', 'list', 5, 'normal', '2022-06-16 16:50:54', '2022-06-16 17:45:59', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (31, '通知列表', 'binance-notify-list', 26, '/binance-notify-list', 'GET', 'binance/notify/index', 'list', 6, 'normal', '2022-06-16 17:45:52', '2022-06-16 17:45:52', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (32, '合约管理', 'binance-contract-list', 26, '/binance-contract-list', 'GET', 'binance/contract/index', 'education', 0, 'normal', '2022-06-17 14:27:26', '2022-06-17 14:27:26', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (33, '新增合约', '', 32, '/binance-contract-store', 'POST', '', '', 0, 'normal', '2022-06-17 15:11:22', '2022-06-17 15:17:10', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (34, '更新合约', '', 32, '/binance-contract-update', 'PUT', '', '', 1, 'normal', '2022-06-17 15:11:35', '2022-06-17 15:50:17', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (35, '合约信息', '', 32, '/binance-contract-info', 'GET', '', '', 22, 'normal', '2022-06-17 15:24:22', '2022-06-17 15:24:22', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (36, '删除合约', '', 32, '/binance-contract-destroy', 'DELETE', '', '', 6, 'normal', '2022-06-17 15:50:11', '2022-06-17 15:50:11', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (37, '丢失区块列表', 'binance-lose-block-list', 26, '/binance-lose-block-list', 'GET', 'binance/lose_block/index', 'list', 2, 'normal', '2022-06-17 16:28:59', '2022-06-17 16:28:59', 'link', 'internal');
INSERT INTO `ga_admin_menu` VALUES (38, '新增区块', '', 37, '/binance-lose-block-store', 'POST', '', '', 0, 'normal', '2022-06-17 16:29:19', '2022-06-17 16:29:19', 'operation', 'internal');
INSERT INTO `ga_admin_menu` VALUES (39, '删除区块', '', 37, '/binance-lose-block-destroy', 'DELETE', '', '', 1, 'normal', '2022-06-17 16:30:10', '2022-06-17 16:31:23', 'operation', 'internal');
COMMIT;

-- ----------------------------
-- Table structure for ga_administrator
-- ----------------------------
DROP TABLE IF EXISTS `ga_administrator`;
CREATE TABLE `ga_administrator` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL COMMENT '用户名',
  `password` char(32) NOT NULL COMMENT 'MD5密码',
  `nickname` varchar(45) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(200) DEFAULT NULL COMMENT '头像地址',
  `status` varchar(10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `last_login_ip` varchar(50) DEFAULT NULL COMMENT '最后登陆IP',
  `last_login_date` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员表';

-- ----------------------------
-- Records of ga_administrator
-- ----------------------------
BEGIN;
INSERT INTO `ga_administrator` VALUES (1, 'admin', 'a66abb5684c45962d887564f08346e8d', 'admin', '', 'normal', 'sssss', '', NULL, '2022-01-09 19:38:04', '2022-04-29 18:33:00');
COMMIT;

-- ----------------------------
-- Table structure for ga_administrator_role
-- ----------------------------
DROP TABLE IF EXISTS `ga_administrator_role`;
CREATE TABLE `ga_administrator_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员角色关联表';

-- ----------------------------
-- Records of ga_administrator_role
-- ----------------------------
BEGIN;
INSERT INTO `ga_administrator_role` VALUES (4, 1, 1, '2022-02-19 19:33:49', '2022-02-19 19:33:49');
COMMIT;

-- ----------------------------
-- Table structure for ga_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `ga_casbin_rule`;
CREATE TABLE `ga_casbin_rule` (
  `ptype` varchar(10) DEFAULT NULL,
  `v0` varchar(256) DEFAULT NULL,
  `v1` varchar(256) DEFAULT NULL,
  `v2` varchar(256) DEFAULT NULL,
  `v3` varchar(256) DEFAULT NULL,
  `v4` varchar(256) DEFAULT NULL,
  `v5` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ga_casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `ga_casbin_rule` VALUES ('g', 'admin', 'super', '', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/administrator-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/role-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/role-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/role-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/role-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/administrator-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/administrator-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/administrator-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/menu-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/menu-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/menu-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/menu-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/administrator-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/role-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/menu-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/config-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/config-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/operation-log-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-user-address-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-collect-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-withdraw-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-queue-task-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-notify-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-contract-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-contract-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-contract-update', 'put', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-contract-info', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-contract-destroy', 'delete', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-lose-block-list', 'get', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-lose-block-store', 'post', '', '', '');
INSERT INTO `ga_casbin_rule` VALUES ('p', 'super', '/binance-lose-block-destroy', 'delete', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for ga_config
-- ----------------------------
DROP TABLE IF EXISTS `ga_config`;
CREATE TABLE `ga_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `module` varchar(80) NOT NULL DEFAULT '' COMMENT '所属模块',
  `key` varchar(80) NOT NULL DEFAULT '' COMMENT '键值',
  `value` text DEFAULT NULL COMMENT '值',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `module_key_idx` (`module`,`key`)
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='配置表';

INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (12, 'backend', 'is_open_verify_captcha', '0', NULL, '2022-05-05 16:05:50');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (33, 'binance', 'collect_address', '0x8520e2ea780e400ab87322d04c158267f36f733a', NULL, '2022-06-18 15:26:35');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (34, 'binance', 'fee_withdraw_address', '0x92b1e4c92c506a95fc6b1af465eab8dab8f39ab7', NULL, '2022-06-18 15:26:35');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (191, 'binance', 'fee_withdraw_private_key', '0xb9d36a8552bdb5e2cf78b9908a9569911539417d743d34d73e7d52be3ae49b61', NULL, '2022-06-18 15:26:35');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (192, 'binance', 'notify_address', 'http://127.0.0.1:8787', NULL, '2022-06-18 15:26:35');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (197, 'binance', 'min_collect_amount', '0.001', NULL, '2022-06-18 15:26:35');
INSERT INTO `ga_config` (`id`, `module`, `key`, `value`, `created_at`, `updated_at`) VALUES (203, 'binance', 'net_type', 'test_net', NULL, '2022-06-18 15:26:35');

-- ----------------------------
-- Table structure for ga_login_log
-- ----------------------------
DROP TABLE IF EXISTS `ga_login_log`;
CREATE TABLE `ga_login_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
  `ip` varchar(30) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `browser` varchar(10) NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(255) NOT NULL DEFAULT '' COMMENT '操作系统',
  `status` varchar(10) DEFAULT 'normal' COMMENT '状态 success 成功 fail 失败',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of ga_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ga_role
-- ----------------------------
DROP TABLE IF EXISTS `ga_role`;
CREATE TABLE `ga_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL COMMENT '角色名',
  `identification` varchar(20) NOT NULL COMMENT '角色标识符',
  `sort` tinyint(4) NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
  `status` varchar(10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色表';

-- ----------------------------
-- Records of ga_role
-- ----------------------------
BEGIN;
INSERT INTO `ga_role` VALUES (1, '超级管理员', 'super', 0, 'normal', '2022-01-09 22:32:23', '2022-06-17 16:31:50');
COMMIT;

-- ----------------------------
-- Table structure for ga_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `ga_role_menu`;
CREATE TABLE `ga_role_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` int(11) unsigned NOT NULL COMMENT '管理员id',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=489 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色和菜单权限关联表';

-- ----------------------------
-- Records of ga_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `ga_role_menu` VALUES (456, 1, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (457, 2, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (458, 7, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (459, 8, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (460, 9, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (461, 18, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (462, 3, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (463, 4, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (464, 5, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (465, 6, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (466, 19, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (467, 10, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (468, 11, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (469, 12, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (470, 13, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (471, 20, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (472, 22, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (473, 23, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (474, 24, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (475, 26, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (476, 32, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (477, 33, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (478, 34, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (479, 36, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (480, 35, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (481, 27, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (482, 37, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (483, 38, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (484, 39, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (485, 28, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (486, 29, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (487, 30, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
INSERT INTO `ga_role_menu` VALUES (488, 31, 1, '2022-06-17 16:31:50', '2022-06-17 16:31:50');
COMMIT;

-- ----------------------------
-- Table structure for lose_blocks
-- ----------------------------
DROP TABLE IF EXISTS `lose_blocks`;
CREATE TABLE `lose_blocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `number` int(11) NOT NULL COMMENT '区块号',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '检测状态，1已检测，0未检测',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `number_idx` (`number`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='重新检测区块表';

-- ----------------------------
-- Records of lose_blocks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for networks
-- ----------------------------
DROP TABLE IF EXISTS `networks`;
CREATE TABLE `networks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL COMMENT '请求地址',
  `name` varchar(255) NOT NULL COMMENT '网络名称',
  `chain_id` int(11) NOT NULL DEFAULT 0,
  `is_use` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是正在使用的网络',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='网络';

-- ----------------------------
-- Records of networks
-- ----------------------------
BEGIN;
INSERT INTO `networks` VALUES (1, 'https://mainnet.infura.io/v3/231767ef48de4aa3a985c9a721699dcc', '主链', 1, 0, '2022-05-30 03:53:36', '2022-05-30 03:53:36');
INSERT INTO `networks` VALUES (2, 'https://ropsten.infura.io/v3/231767ef48de4aa3a985c9a721699dcc', 'Ropsten', 3, 1, '2022-05-30 03:53:36', '2022-05-30 03:53:36');
COMMIT;

-- ----------------------------
-- Table structure for notify
-- ----------------------------
DROP TABLE IF EXISTS `notify`;
CREATE TABLE `notify` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` varchar(20) NOT NULL DEFAULT '' COMMENT '交易类型, recharge 充值，withdraw 提现',
  `relation_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联表id',
  `notify_data` text NOT NULL DEFAULT '' COMMENT '通知数据',
  `notify_address` varchar(255) NOT NULL DEFAULT '' COMMENT '通知地址',
  `unique_id` varchar(255) NOT NULL DEFAULT '' COMMENT '唯一id',
  `fail_amount` int(11) NOT NULL DEFAULT 0 COMMENT '失败次数',
  `status` varchar(20) NOT NULL DEFAULT '' COMMENT '状态 fail 失败,wait等待通知  finish通知完成',
  `is_immediately_retry` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否立即重试',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `notify_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '上次通知时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='通知表';

-- ----------------------------
-- Records of notify
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for notify_log
-- ----------------------------
DROP TABLE IF EXISTS `notify_log`;
CREATE TABLE `notify_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `notify_id` int(11) NOT NULL DEFAULT 0 COMMENT 'notify id',
  `log` varchar(255) NOT NULL DEFAULT '' COMMENT '错误日志',
  `fail_amount` int(11) NOT NULL DEFAULT 0 COMMENT '第几次失败',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='通知错误日志表';

-- ----------------------------
-- Records of notify_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for queue_task
-- ----------------------------
DROP TABLE IF EXISTS `queue_task`;
CREATE TABLE `queue_task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hash` varchar(255) NOT NULL DEFAULT '',
  `symbol` varchar(255) NOT NULL DEFAULT '',
  `contract_address` varchar(255) NOT NULL DEFAULT '',
  `from` varchar(255) NOT NULL DEFAULT '' COMMENT '转出地址',
  `to` varchar(255) NOT NULL DEFAULT '' COMMENT '转入地址',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '金额',
  `gas_limit` int(11) NOT NULL DEFAULT 0 COMMENT 'gas  限制',
  `gas_price` varchar(255) NOT NULL DEFAULT '' COMMENT 'gas 预估 价格',
  `actual_gas_used` int(11) NOT NULL DEFAULT 0 COMMENT 'gas 实际消耗',
  `actual_gas_price` varchar(255) NOT NULL DEFAULT '' COMMENT 'gas 实际价格',
  `actual_fee` varchar(255) NOT NULL DEFAULT '' COMMENT '实际手续费',
  `nonce` int(11) NOT NULL DEFAULT 0,
  `type` varchar(20) NOT NULL DEFAULT '' COMMENT '交易类型',
  `status` varchar(20) NOT NULL DEFAULT '0' COMMENT '转账状态: fail 转出失败,wait 等待转出,process 转出中,success转出成功',
  `fail_amount` int(11) NOT NULL DEFAULT 0 COMMENT '失败次数',
  `private_key` text DEFAULT NULL COMMENT '私钥',
  `relation_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联的其他表id',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `send_at` timestamp NULL DEFAULT NULL COMMENT '发送转账时间',
  `finish_at` timestamp NULL DEFAULT NULL COMMENT '转账检测成功时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `hash_idx` (`hash`),
  KEY `from_idx` (`from`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='队列表';

-- ----------------------------
-- Records of queue_task
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for queue_task_log
-- ----------------------------
DROP TABLE IF EXISTS `queue_task_log`;
CREATE TABLE `queue_task_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `queue_task_id` int(11) NOT NULL DEFAULT 0 COMMENT '队列任务id',
  `log` varchar(255) NOT NULL DEFAULT '' COMMENT '错误日志',
  `fail_amount` int(11) NOT NULL DEFAULT 0 COMMENT '第几次失败',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='队列错误日志表';

-- ----------------------------
-- Records of queue_task_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_addresses
-- ----------------------------
DROP TABLE IF EXISTS `user_addresses`;
CREATE TABLE `user_addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '以太坊地址',
  `external_user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '外部用户id',
  `private_key` text DEFAULT NULL COMMENT '私钥',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型:0平台生成,1外部导入',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `address_idx` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户地址表';

-- ----------------------------
-- Records of user_addresses
-- ----------------------------
BEGIN;
INSERT INTO `user_addresses` VALUES (115, '0x81023633832221b512018a21f8a3c6a6fe774913', '2', '0x841176e4948f90ea6f237d405ebd25eb0eec4b860f1e5cf76f6541c886353241', 0, '2022-06-07 10:16:45', '2022-06-07 10:16:45');
INSERT INTO `user_addresses` VALUES (118, '0x84b2d9c9b870ca47719e17e8cf790d66686743c8', '6', '0x3e525c9e6a3687f342021efee8258a21f66eee5bdfe6975e981a7ac438b95425', 0, '2022-06-11 17:43:09', '2022-06-11 17:43:09');
COMMIT;

-- ----------------------------
-- Table structure for withdraws
-- ----------------------------
DROP TABLE IF EXISTS `withdraws`;
CREATE TABLE `withdraws` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '内部用户id',
  `user_address` varchar(255) NOT NULL DEFAULT '''' COMMENT '用户地址',
  `external_order_id` varchar(255) NOT NULL DEFAULT '' COMMENT '外部订单id',
  `external_user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '外部用户id',
  `hash` varchar(255) NOT NULL DEFAULT '' COMMENT 'hash',
  `symbol` varchar(255) NOT NULL DEFAULT '' COMMENT '代币符号',
  `contract_address` varchar(255) NOT NULL DEFAULT '',
  `from` varchar(255) NOT NULL DEFAULT '' COMMENT '转出地址',
  `to` varchar(255) NOT NULL DEFAULT '' COMMENT '转入地址',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '转出金额',
  `status` varchar(20) NOT NULL DEFAULT '' COMMENT '状态 fail 失败，wait 待转出，process 转出中，wait_notify转出完成,finish_notify通知完成',
  `create_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='提现表';

-- ----------------------------
-- Records of withdraws
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
