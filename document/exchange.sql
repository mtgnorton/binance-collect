/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : exchange_bak

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 09/04/2021 18:43:01
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `advert_categories`;
CREATE TABLE `advert_categories`
(
    `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '广告分类名称',
    `width`       int(11) NOT NULL DEFAULT 0 COMMENT '图片宽度',
    `height`      int(11) NOT NULL DEFAULT 0 COMMENT '图片高度',
    `identifying` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '广告分类标识',
    `created_at`  timestamp(0) NULL DEFAULT NULL,
    `updated_at`  timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '广告分类' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of advert_categories
-- ----------------------------
INSERT INTO `advert_categories`
VALUES (1, '首页广告', 960, 480, 'index_carousel', '2021-04-13 11:07:21', '2021-04-13 11:07:21');
-- ----------------------------
-- Records of advert_categories
-- ----------------------------

-- ----------------------------
-- Table structure for adverts
-- ----------------------------
DROP TABLE IF EXISTS `adverts`;
CREATE TABLE `adverts`
(
    `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `category_id` int(11) NOT NULL DEFAULT 0 COMMENT '广告分类id',
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '广告名称',
    `identifying` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '广告标识或链接',
    `image_path`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图片路径',
    `sort`        int(11) NOT NULL DEFAULT 0,
    `is_disabled` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用',
    `created_at`  timestamp(0) NULL DEFAULT NULL,
    `updated_at`  timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '广告' ROW_FORMAT = Dynamic;

INSERT INTO `adverts`(`id`, `category_id`, `name`, `identifying`, `image_path`, `sort`, `is_disabled`, `created_at`,
                      `updated_at`)
VALUES (1, 1, '1', 'http://exchange_php.grayvip.com', 'adverts/e9d0daef5cd4260c7ac9e99c8964ca07.jpg', 0, 0,
        '2021-04-13 11:11:21', '2021-04-13 11:11:21');
INSERT INTO `adverts`(`id`, `category_id`, `name`, `identifying`, `image_path`, `sort`, `is_disabled`, `created_at`,
                      `updated_at`)
VALUES (2, 1, '2', 'http://exchange_php.grayvip.com', 'adverts/f684fa03cad200b0c24cf8e63a68c023.jpg', 0, 0,
        '2021-04-13 11:13:01', '2021-04-13 11:13:01');

-- ----------------------------
-- Records of adverts
-- ----------------------------

-- ----------------------------
-- Table structure for announcements
-- ----------------------------
DROP TABLE IF EXISTS `announcements`;
CREATE TABLE `announcements`
(
    `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公告标题',
    `content`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '公告内容',
    `sort`        int(11) NOT NULL DEFAULT 0,
    `is_disabled` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用',
    `created_at`  timestamp(0) NULL DEFAULT NULL,
    `updated_at`  timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '公告' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of announcements
-- ----------------------------

-- ----------------------------
-- Table structure for app_versions
-- ----------------------------
DROP TABLE IF EXISTS `app_versions`;
CREATE TABLE `app_versions`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT,
    `version`      varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '版本号',
    `title`        varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '升级标题',
    `description`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '升级描述',
    `download_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '下载链接',
    `client_type`  tinyint(1) NOT NULL DEFAULT 0 COMMENT '0 安卓,1 ios',
    `upgrade_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '2强制升级 1提醒升级  0不提醒升级',
    `created_at`   timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`   timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'app版本升级' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_versions
-- ----------------------------

-- ----------------------------
-- Table structure for bank_accounts
-- ----------------------------
DROP TABLE IF EXISTS `bank_accounts`;
CREATE TABLE `bank_accounts`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `type`       tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型 0银行卡 1微信 2支付宝',
    `name`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '真实姓名',
    `account`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号',
    `bank`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '开户银行',
    `qr_code`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '二维码',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户收款账户表' ROW_FORMAT = Dynamic;

create index user_idx on bank_accounts (user_id);
-- ----------------------------
-- Records of bank_accounts
-- ----------------------------

-- ----------------------------
-- Table structure for certifications
-- ----------------------------
DROP TABLE IF EXISTS `certifications`;
CREATE TABLE `certifications`
(
    `id`                int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`           int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`          varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `name`              varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '姓名',
    `id_card`           varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '身份证',
    `card_image_front`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '身份证正面',
    `card_image_behind` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '身份证反面',
    `type`              tinyint(1) NOT NULL DEFAULT 1 COMMENT '认证类型 1 kyc1 ,2 kyc2',
    `status`            tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态 0未审核 1已审核 -1已拒绝',
    `remark`            tinyint(1) NOT NULL DEFAULT 0 COMMENT '备注,审核失败原因',
    `created_at`        timestamp(0) NULL DEFAULT NULL,
    `updated_at`        timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '实名认证' ROW_FORMAT = Dynamic;

create index user_idx on certifications (user_id);

-- ----------------------------
-- Records of certifications
-- ----------------------------


-- ----------------------------
-- Table structure for document_categories
-- ----------------------------
DROP TABLE IF EXISTS `document_categories`;
CREATE TABLE `document_categories`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `title`      varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分类标题',
    `parent_id`  int(11) NOT NULL DEFAULT 0 COMMENT '父级id',
    `sort`       int(11) NOT NULL DEFAULT 0,
    `created_at` timestamp(0)                                           NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                           NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文档分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of document_categories
-- ----------------------------

-- ----------------------------
-- Table structure for documents
-- ----------------------------
DROP TABLE IF EXISTS `documents`;
CREATE TABLE `documents`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `title`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文档标题',
    `identify`    varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文档标识符',
    `category_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
    `content`     text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '文档内容',
    `sort`        int(11) NOT NULL DEFAULT 0,
    `is_disabled` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用',
    `created_at`  timestamp(0) NULL DEFAULT NULL,
    `updated_at`  timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文档表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of documents
-- ----------------------------

-- ----------------------------
-- Table structure for exchange_orders
-- ----------------------------
DROP TABLE IF EXISTS `exchange_orders`;
CREATE TABLE `exchange_orders`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`       int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `sn`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单号',
    `exchange_id`   int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市场id',
    `side`          tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 买入 1卖出',
    `price_type`    tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 限价 1市价',
    `price`         decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '期望成交价格',
    `amount`        decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '当为限价时,为委托数量. 当为市价买单时,为委托交易额,当为市价卖单时,为委托卖出量',
    `avg_price`     decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '实际成交平均价格',
    `remain_amount` decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '未成交数量',
    `status`        tinyint(1) NOT NULL DEFAULT 0 COMMENT '0 挂单中, 5 部分成交, 10 交易成功',
    `cancel_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0未取消 1已取消,2撤单,当为市价单时可能出现交易失败的情况',
    `created_at`    timestamp(0) NULL DEFAULT NULL,
    `updated_at`    timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '币币交易委托订单表' ROW_FORMAT = Dynamic;

create index user_idx on exchange_orders (user_id);

-- ----------------------------
-- Records of exchange_orders
-- ----------------------------

-- ----------------------------
-- Table structure for exchange_trades
-- ----------------------------
DROP TABLE IF EXISTS `exchange_trades`;
CREATE TABLE `exchange_trades`
(
    `id`             int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `taker_user_id`  int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '吃单用户id',
    `taker_username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `maker_user_id`  int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '挂单用户id',
    `maker_username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `exchange_id`    int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市场id',
    `taker_side`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '吃单方向 0 买单吃单 1卖单吃单',
    `taker_order_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '吃单id',
    `maker_order_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '挂单id',
    `taker_fee`      decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '吃单手续费',
    `maker_fee`      decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '挂单手续费',
    `price`          decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '成交价格',
    `amount`         decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '成交数量',
    `created_at`     timestamp(0) NULL DEFAULT NULL,
    `updated_at`     timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX            `taker_id_idx`(`taker_order_id`) USING BTREE,
    INDEX            `maker_id_idx`(`maker_order_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '币币交易委托订单成交表' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for exchange_trades
-- ----------------------------
DROP TABLE IF EXISTS `exchange_fail_trades`;
CREATE TABLE `exchange_fail_trades`
(
    `id`               int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content`          text         not null default '' comment '成交内容',
    `stream_key`       varchar(255) not null default '' comment '',
    `stream_id`        varchar(255) not null default '' comment '',
    `error`            text         not null default '' comment '',
    `is_market`        tinyint(1) not null default 0 comment '是否是市价交易',
    `is_handle_finish` tinyint(1) not null default 0 comment '是否手动完成0 否,1已重试,2已重新委托',

    `created_at`       timestamp(0) NULL DEFAULT NULL,
    `updated_at`       timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '成交接受失败表' ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `exchange_receive_trades`;
CREATE TABLE `exchange_receive_trades`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content`    text null default null comment '成交内容',
    `stream_key` varchar(255) not null default '' comment '',
    `stream_id`  varchar(100) not null default '' comment '',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '成交接收表' ROW_FORMAT = Dynamic;

create unique index id_idx on exchange_receive_trades (`stream_id`);


-- ----------------------------
-- Records of exchange_trades
-- ----------------------------


-- ----------------------------
-- Table structure for exchanges
-- ----------------------------
DROP TABLE IF EXISTS `exchanges`;
CREATE TABLE `exchanges`
(
    `id`                  int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `base_symbol_id`      int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '基础货币id',
    `quote_symbol_id`     int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '计价货币id',
    `is_can_trade`        tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否能够进行交易',
    `is_recommend`        tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否推荐',
    `is_disabled`         tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用',
    `sort`                int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    `decimals`            tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小数位数',
    `taker_fee`           decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '吃单手续费',
    `maker_fee`           decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '挂单手续费',
    `limit_buy_min`       decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '限价买入最小量,当为0时不限制',
    `limit_sell_min`      decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '限价卖出最小量,当为0时不限制',
    `market_buy_min`      decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '市价买入最小量,当为0时不限制',
    `market_sell_min`     decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '市价卖出最小量,当为0时不限制',
    `trade_start_time`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
    `trade_end_time`      varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
    `is_use_huobi_kline`  tinyint(1) not null default 0 comment '是否使用火币数据',
    `is_save_huobi_kline` tinyint(1) not null default 0 comment '是否保存火币数据',
    `huobi_symbol`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
    `default_price`       decimal(18, 8) not null default 0 comment '当某个币昨日没有成交时,使用该值作为今日行情的默认价格',
    `updated_at`          timestamp(0) NULL DEFAULT NULL,
    `created_at`          timestamp(0) NULL DEFAULT NULL,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '币币交易市场' ROW_FORMAT = Compact;

INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (1, 1, 2, 1, 1, 0, 1, 0, 0.20000000, 0.20000000, 0.00000000, 0.00000000, 0.00000000, 0.00000000, '00:00:00',
        '23:00:00', 1, 0, '', '2021-04-18 13:48:04', '2021-03-24 12:19:55');
INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (2, 3, 2, 1, 1, 0, 0, 0, 0.20000000, 0.20000000, 0.00000000, 0.00000000, 0.00000000, 0.00000000, '00:12:35',
        '23:12:35', 1, 0, 'zec/usdt', '2021-04-19 16:47:17', '2021-03-24 12:19:54');
INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (3, 4, 2, 1, 0, 0, 2, 0, 0.20000000, 0.20000000, 0.00000000, 0.00000000, 0.00000000, 0.00000000, '00:00:00',
        '23:59:59', 1, 1, '', '2021-04-17 16:52:26', '2021-04-14 19:19:10');
INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (9, 7, 2, 1, 0, 0, 3, 0, 0.20000000, 0.20000000, 0.00000000, 0.00000000, 0.00000000, 0.00000000, '00:00:00',
        '23:59:59', 1, 1, '', '2021-04-17 16:52:26', '2021-04-14 19:19:02');
INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (10, 8, 2, 1, 0, 0, 4, 0, 0.20000000, 0.20000000, 0.00000000, 0.00000000, 0.00000000, 0.00000000, '00:00:00',
        '23:59:59', 1, 1, '', '2021-04-18 12:17:42', '2021-03-05 16:51:13');
INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (11, 16, 2, 1, 0, 0, 5, 0, 0.20000000, 0.20000000, 10.00000000, 10.00000000, 10.00000000, 0.00000000, '00:00:00',
        '23:59:59', 1, 1, 'trx/usdt', '2021-04-20 16:38:01', '2021-03-05 16:51:20');
INSERT INTO `exchanges`(`id`, `base_symbol_id`, `quote_symbol_id`, `is_can_trade`, `is_recommend`, `is_disabled`,
                        `sort`, `decimals`, `taker_fee`, `maker_fee`, `limit_buy_min`, `limit_sell_min`,
                        `market_buy_min`, `market_sell_min`, `trade_start_time`, `trade_end_time`, `is_use_huobi_kline`,
                        `is_save_huobi_kline`, `huobi_symbol`, `updated_at`, `created_at`)
VALUES (12, 17, 2, 1, 1, 0, 0, 0, 0.30000000, 0.30000000, 10.00000000, 10.00000000, 20.00000000, 20.00000000,
        '00:00:00', '23:59:59', 0, 1, 'abd/usdt', '2021-04-28 11:15:39', '2021-04-01 09:32:20');


-- ----------------------------
-- Table structure for feedbacks
-- ----------------------------
DROP TABLE IF EXISTS `kline_control`;
CREATE TABLE `kline_control`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `exchange_id`    int(11) not null default 0 comment '交易对id',
    `lowest_price`   decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '最低价',
    `highest_price`  decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '最高价',
    `lowest_volume`  decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '最低成交量',
    `highest_volume` decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '最高成交量',
    `begin_time`     timestamp(0) NULL DEFAULT NULL comment '开始时间',
    `end_time`       timestamp(0) NULL DEFAULT NULL comment '结束时间',
    `created_at`     timestamp(0) NULL DEFAULT NULL comment '开始时间',
    `updated_at`     timestamp(0) NULL DEFAULT NULL comment '结束时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'k线控制' ROW_FORMAT = Dynamic;



-- ----------------------------
-- Table structure for symbols
-- ----------------------------
-- ----------------------------
DROP TABLE IF EXISTS `symbols`;
CREATE TABLE `symbols`
(
    `id`                   int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`                 varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '货币符号',
    `has_exchange_account` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '具有币币账户',
    `has_fiat_account`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '具有法币账户',
    `is_can_recharge`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否能够充值',
    `is_can_withdraw`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否能够提现',
    `is_can_transfer`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否支持转账',
    `is_quote`             tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否是计价货币(交易区)',
    `quote_sort`           int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '交易区显示排序',
    `is_hidden`            tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否隐藏',
    `account_sort`         int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '账号显示排序',
    `chinese_name`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '中文名',
    `release_time`         timestamp(0) NULL DEFAULT NULL COMMENT '发行时间',
    `release_total`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发行总量',
    `circulate_total`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '流通总量',
    `crowdfunding_price`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '众筹价格',
    `white_paper`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '白皮书',
    `website_url`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '官网链接',
    `block_url`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '区块查询链接',
    `introduction`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '简介',
    `symbol_address`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '货币地址',
    `logo_path`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
    `is_gateway`           tinyint(1) UNSIGNED ZEROFILL NULL DEFAULT 0,
    `min_recharge_amount`  decimal(18, 8) NULL DEFAULT 0.00000000,
    `min_withdraw_amount`  decimal(18, 8) NULL DEFAULT 0.00000000,
    `max_withdraw_amount`  decimal(18, 8) NULL DEFAULT 0.00000000,
    `withdraw_fee`         varchar(20)                                                  not null DEFAULT '0',
    `decimals`             int(11) NULL DEFAULT 0,
    `created_at`           timestamp(0) NULL DEFAULT NULL,
    `updated_at`           timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '货币符号表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of symbols
-- ----------------------------
INSERT INTO `symbols`
VALUES (1, 'BTC', 1, 1, 0, 0, 1, 0, 0, 0, 0, '比特币', '2020-12-09 07:52:33', '1000w', '1000w', '', '', '', '', '', '',
        'symbols/BTC.png', 0, 1.00000000, 10.00000000, 1000.00000000, 5.00000000, 0, '2020-12-09 07:52:48',
        '2021-04-12 14:38:35');
INSERT INTO `symbols`
VALUES (2, 'USDT', 1, 1, 1, 1, 1, 1, 0, 0, 0, '泰达币', '2020-12-09 08:18:01', '1000w', '1000w', '', '', '', '', '', '',
        'symbols/USDT.png', 0, 1.00000000, 10.00000000, 1000.00000000, 1.00000000, 0, '2020-12-09 08:18:28',
        '2021-04-12 14:38:58');
INSERT INTO `symbols`
VALUES (3, 'LINK', 1, 1, 0, 0, 1, 0, 0, 0, 0, 'link', '2020-12-09 08:18:45', '1000w', '1000w', '', '', '', '', '', '',
        NULL, 1, 0.00000000, 0.00000000, 0.00000000, 0.00000000, 0, '2020-12-09 08:18:56', '2021-03-27 16:42:18');
INSERT INTO `symbols`
VALUES (4, 'EOS', 1, 1, 0, 0, 0, 0, 0, 0, 0, 'eos', '2020-12-09 08:19:36', '1000w', '1000w', '', '', '', '', '', '',
        'symbols/EOS.png', 1, 0.00000000, 0.00000000, 0.00000000, 0.00000000, 0, '2020-12-09 08:19:47',
        '2021-04-12 14:39:18');
INSERT INTO `symbols`
VALUES (6, 'ETH', 1, 0, 0, 0, 0, 0, 0, 0, 0, '以太坊', '2021-03-04 12:26:34', '10000w', '10000w', '', '', '', '', '', '',
        'symbols/ETH.png', 0, 0.00000000, 0.00000000, 0.00000000, 0.00000000, 0, '2021-03-04 12:28:10',
        '2021-04-12 14:39:35');
INSERT INTO `symbols`
VALUES (7, 'HT', 1, 0, 0, 0, 0, 0, 0, 0, 0, 'ht', '2021-03-04 12:28:17', '1000w', '1000w', '', '', '', '', '', '', NULL,
        0, 0.00000000, 0.00000000, 0.00000000, 0.00000000, 0, '2021-03-04 12:28:59', '2021-03-04 12:29:04');
INSERT INTO `symbols`
VALUES (8, 'DOT', 1, 0, 0, 0, 1, 0, 0, 0, 0, 'dot', '2021-03-04 14:21:30', '1000w', '1000w', '', '', '', '', '', '',
        'symbols/DOT.png', 0, 0.00000000, 0.00000000, 0.00000000, 0.00000000, 0, '2021-03-04 14:22:14',
        '2021-04-12 14:50:26');
INSERT INTO `symbols`
VALUES (16, 'TRX', 1, 0, 0, 0, 1, 0, 0, 0, 0, 'trx', '2021-03-04 16:13:24', '1000w', '1000w', '', '', '', '', '', '',
        NULL, 0, 0.00000000, 0.00000000, 0.00000000, 0.00000000, 8, '2021-03-04 16:13:29', '2021-03-04 16:13:29');
INSERT INTO `symbols`
VALUES (17, 'ABD', 1, 0, 0, 0, 1, 0, 0, 0, 0, 'ABD', '2021-04-01 09:13:02', '1000w', '1000w', '', '', '', '', '', '',
        'symbols/ABD.png', 0, 10.00000000, 1.00000000, 1000000.00000000, 1.00000000, 4, '2021-04-01 09:21:03',
        '2021-04-10 15:04:39');

-- ----------------------------
-- Records of exchanges
-- ----------------------------

-- ----------------------------
-- Table structure for feedbacks
-- ----------------------------
DROP TABLE IF EXISTS `feedbacks`;
CREATE TABLE `feedbacks`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `content`    text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '反馈内容',
    `mobile`     varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `email`      varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '意见反馈表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of feedbacks
-- ----------------------------

-- ----------------------------
-- Table structure for fiat_orders
-- ----------------------------
DROP TABLE IF EXISTS `fiat_orders`;
CREATE TABLE `fiat_orders`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`       int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `sn`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单号',
    `symbol_id`     int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '货币id',
    `side`          tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 买入 1卖出',
    `type`          tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '委托类型 0 数量 1金额',
    `amount`        decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '委托数量',
    `price`         decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '委托价格',
    `money_amount`  decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '委托总金额amount*price',
    `remain_amount` decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '未成交数量',
    `fee`           decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '手续费数量',
    `status`        tinyint(1) NOT NULL DEFAULT 0 COMMENT '0 挂单中, 5 部分成交, 10 交易成功',
    `cancel_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0未取消 1已取消',
    `remark`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '交易说明',
    `created_at`    timestamp(0) NULL DEFAULT NULL,
    `updated_at`    timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '法币交易委托订单表' ROW_FORMAT = Dynamic;

create index symbol_idx on fiat_orders (symbol_id);
create index user_idx on fiat_orders (user_id);
-- ----------------------------
-- Records of fiat_orders
-- ----------------------------

-- ----------------------------
-- Table structure for fiat_trade_complaints
-- ----------------------------
DROP TABLE IF EXISTS `fiat_trade_complaints`;
CREATE TABLE `fiat_trade_complaints`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `trade_id`   int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '法币成交id',
    `user_id`    int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `content`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '申诉内容',
    `images`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '图片凭证',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '法币交易申诉表' ROW_FORMAT = Dynamic;

create index trade_user_idx on fiat_trade_complaints (trade_id, user_id);

-- ----------------------------
-- Records of fiat_trade_complaints
-- ----------------------------

-- ----------------------------
-- Table structure for fiat_trades
-- ----------------------------
DROP TABLE IF EXISTS `fiat_trades`;
CREATE TABLE `fiat_trades`
(
    `id`               int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `taker_user_id`    int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '吃单用户id',
    `taker_username`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `maker_user_id`    int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '挂单用户id',
    `maker_username`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `symbol_id`        int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '货币id',
    `taker_side`       tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '吃单方向 0 买单吃单 1卖单吃单',
    `taker_order_id`   int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '吃单id',
    `maker_order_id`   int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '挂单id',
    `price`            decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '成交价格',
    `amount`           decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '成交数量',
    `fee`              decimal(18, 8)                                                NOT NULL DEFAULT 0.00000000 COMMENT '手续费数量',
    `pay_credentials`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '支付凭证',
    `time_limit`       int(11) UNSIGNED NOT NULL DEFAULT 10 COMMENT '时间限制单位(min)',
    `pay_time`         timestamp(0) NULL DEFAULT NULL comment '支付时间',
    `send_coin_time`   timestamp(0) NULL DEFAULT NULL comment '放币时间',
    `status`           tinyint(1) NOT NULL DEFAULT 0 COMMENT '-1 取消,0 待付款, 5 待放币 ,10 交易成功',
    `complaint_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '申诉状态 -1 取消,0 未申诉 ,5 申诉中, 10申诉完成',
    `complaint_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '申诉意见',
    `created_at`       timestamp(0) NULL DEFAULT NULL,
    `updated_at`       timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '法币交易委托订单成交表' ROW_FORMAT = Dynamic;

create INDEX taker_id_idx on fiat_trades (taker_order_id);
create INDEX maker_id_idx on fiat_trades (maker_order_id);

-- ----------------------------
-- Records of fiat_trades
-- ----------------------------

-- ----------------------------
-- Table structure for kline
-- ----------------------------
DROP TABLE IF EXISTS `kline`;
CREATE TABLE `kline`
(
    `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `exchange_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '币币交易市场id',
    `period`      varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '时间区间',
    `open`        decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000,
    `close`       decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000,
    `highest`     decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000,
    `lowest`      decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000,
    `volume`      decimal(22, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '成交量',
    `time`        timestamp(0) NULL DEFAULT NULL,
    `created_at`  timestamp(0) NULL DEFAULT NULL,
    `updated_at`  timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'k线表' ROW_FORMAT = Dynamic;

create unique index exchange_id_period_time_idx on kline (`exchange_id`, `period`, `time`);

-- ----------------------------
-- Records of kline
-- ----------------------------

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `target`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '发送地址',
    `content`    varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发送内容',
    `code`       varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '验证码',
    `ip`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT 'ip',
    `type`       varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '发送类型,如注册,找回密码',
    `is_use`     tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否使用',
    `created_at` timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '短信发送记录' ROW_FORMAT = Dynamic;

create index user_idx on messages (user_id);

-- ----------------------------
-- Records of messages
-- ----------------------------

-- ----------------------------
-- Table structure for notifies
-- ----------------------------
DROP TABLE IF EXISTS `notifies`;
CREATE TABLE `notifies`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `notify_id`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '通知id',
    `data`       text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '通知内容',
    `created_at` timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '通知记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of notifies
-- ----------------------------


-- ----------------------------
-- Table structure for recharges
-- ----------------------------
DROP TABLE IF EXISTS `recharges`;
CREATE TABLE `recharges`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `symbol_id`  int(11) NOT NULL DEFAULT 0 COMMENT '货币id',
    `amount`     decimal(18, 8)                                          NOT NULL DEFAULT 0.00000000 COMMENT '充值数量',
    `remark`     varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '备注',
    `created_at` timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '充值记录表' ROW_FORMAT = Dynamic;
create index user_idx on recharges (user_id);


-- ----------------------------
-- Records of recharges
-- ----------------------------


-- ----------------------------
-- Table structure for test_frozen
-- ----------------------------
DROP TABLE IF EXISTS `test_frozen`;
CREATE TABLE `test_frozen`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `order_id`   int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `user_id`    int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `symbol_id`  int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `change`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `rs`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `flag`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'test' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of test_frozen
-- ----------------------------

-- ----------------------------
-- Table structure for user_recommend_relation
-- ----------------------------
DROP TABLE IF EXISTS `user_recommend_relation`;
CREATE TABLE `user_recommend_relation`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NULL DEFAULT NULL COMMENT '用户ID',
    `parent_id`  int(11) NULL DEFAULT NULL COMMENT '上级ID',
    `layer`      int(11) NULL DEFAULT NULL COMMENT '层数',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户推荐关系表' ROW_FORMAT = Dynamic;

create index user_parent_idx on user_recommend_relation (user_id, parent_id);
create index parent_user_idx on user_recommend_relation (parent_id, user_id);

-- ----------------------------
-- Records of user_recommend_relation
-- ----------------------------

-- ----------------------------
-- Table structure for user_symbol_gateways
-- ----------------------------
DROP TABLE IF EXISTS `user_symbol_gateways`;
CREATE TABLE `user_symbol_gateways`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `symbol_id`  int(11) NOT NULL DEFAULT 0 COMMENT '货币id',
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户货币网关' ROW_FORMAT = Dynamic;

create index symbol_user_idx on user_symbol_gateways (symbol_id, user_id);

-- ----------------------------
-- Records of user_symbol_gateways
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`                  int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username`            varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
    `pay_password`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付密码',
    `email`               varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '电子邮箱',
    `mobile`              varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '手机号',
    `certification_level` tinyint(1) NOT NULL DEFAULT 0 COMMENT '认证级别',
    `address`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '区块钱包地址,和私钥 助记词不对应',
    `is_disabled`         tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用',
    `invite_code`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '邀请码',
    `parent_id`           int(11) NOT NULL DEFAULT 0 COMMENT '上级id',
    `api_token`           varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '最后一次登录的token',
    `created_at`          timestamp(0) NULL DEFAULT NULL,
    `updated_at`          timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;


create index mobile_idx on users (mobile);


-- ----------------------------
-- Records of users
-- ----------------------------

-- ----------------------------
-- Table structure for wallet_logs
-- ----------------------------
DROP TABLE IF EXISTS `wallet_logs`;
CREATE TABLE `wallet_logs`
(
    `id`             int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`        int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `symbol_id`      varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '货币id',
    `account_type`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 币币账户 1法币账户',
    `balance_type`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '金额类型 0 余额,1冻结金额',
    `source_type`    tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 后台调整,5充值,10提现,15转入,20转出,25手续费,30 买入冻结,35卖出冻结,40买入成功,45卖出成功,50买入撤销,55卖出撤销,60买入解冻,65卖出解冻,70开通网关,75划转转出,80划转转入',
    `change_amount`  decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '变动金额',
    `before_balance` decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '变化前的余额',
    `after_balance`  decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '变化后的余额',
    `loggable_id`    int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '一对多,多态关联id,',
    `loggable_type`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '一对多,多态关联类型,转账时为user模型',
    `remark`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
    `created_at`     timestamp(0) NULL DEFAULT NULL,
    `updated_at`     timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '钱包日志表' ROW_FORMAT = Dynamic;

create index user_symbol_idx on wallet_logs (user_id, symbol_id);
-- ----------------------------
-- Records of wallet_logs
-- ----------------------------

-- ----------------------------
-- Table structure for wallets
-- ----------------------------
DROP TABLE IF EXISTS `wallets`;
CREATE TABLE `wallets`
(
    `id`           int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`      int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `account_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 币币账户 1法币账户',
    `symbol_id`    varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '货币id',
    `balance`      decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '余额',
    `frozen`       decimal(18, 8)                                               NOT NULL DEFAULT 0.00000000 COMMENT '冻结的金额',
    `created_at`   timestamp(0) NULL DEFAULT NULL,
    `updated_at`   timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_symbol_type_idx`(`user_id`, `symbol_id`, `account_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '钱包表' ROW_FORMAT = Dynamic;

create unique index user_symbol_account_idx on wallets (user_id, symbol_id, account_type);
-- ----------------------------
-- Records of wallets
-- ----------------------------

-- ----------------------------
-- Table structure for withdraws
-- ----------------------------
DROP TABLE IF EXISTS `withdraws`;
CREATE TABLE `withdraws`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `user_id`        int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `symbol_id`      int(11) NOT NULL DEFAULT 0 COMMENT '货币id',
    `amount`         decimal(18, 8)                                          NOT NULL DEFAULT 0.00000000 COMMENT '提币数量',
    `actual_amount`  decimal(18, 8)                                          NOT NULL DEFAULT 0.00000000 COMMENT '实际提币数量',
    `fee`            decimal(18, 8)                                          NOT NULL DEFAULT 0.00000000 COMMENT '提币手续费',
    `target_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '提币地址',
    `status`         tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态,-1 拒绝,0未审核,1审核通过,2提现成功',
    `remark`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
    `created_at`     timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`     timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '提现表' ROW_FORMAT = Dynamic;

create index user_symbol_idx on withdraws (user_id, symbol_id);
-- ----------------------------
-- Records of withdraws
-- ----------------------------

SET
FOREIGN_KEY_CHECKS = 1;



-- ----------------------------
-- Table structure for withdraws
-- ----------------------------
DROP TABLE IF EXISTS `zt_transfers`;
CREATE TABLE `zt_transfers`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `hash`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '',
    `from`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '',
    `to`          varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '',
    `transfer_id` int(11) NOT NULL DEFAULT 0 COMMENT 'trx转账表id',
    `amount`      decimal(18, 8)                                          NOT NULL DEFAULT 0.00000000 COMMENT '数量',
    `fee_amount`  decimal(18, 8)                                          NOT NULL DEFAULT 0.00000000 COMMENT '手续费数量',
    `status`      tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态,-1 ,0转出中,1已转出',
    `created_at`  timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                            NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'ztpay trc 转账表' ROW_FORMAT = Dynamic;

create unique index hash_idx on zt_transfers (`hash`);



-- ----------------------------
-- Table structure for withdraws
-- ----------------------------
DROP TABLE IF EXISTS `zt_collects`;
CREATE TABLE `zt_collects`
(
    `id`                   int(11) NOT NULL AUTO_INCREMENT,
    `user_id`              int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`             varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `recharge_transfer_id` int(11) NOT NULL DEFAULT 0 COMMENT '转账表充值id',
    `fee_transfer_id`      int(11) NOT NULL DEFAULT 0 COMMENT '转账表手续费id',
    `collect_transfer_id`  int(11) NOT NULL DEFAULT 0 COMMENT '转账表归集id',
    `amount`               decimal(18, 8)                                         NOT NULL DEFAULT 0.00000000 COMMENT '充值数量',
    `status`               tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态,-1失败 ,0待归集,5转出手续费中,15 归集中 20 归集完成',
    `remark`               varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
    `created_at`           timestamp(0)                                           NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`           timestamp(0)                                           NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'ztpay trc 归集表' ROW_FORMAT = Dynamic;


SET
FOREIGN_KEY_CHECKS = 1;
-- ----------------------------
-- Table structure for withdraws
-- ----------------------------
DROP TABLE IF EXISTS `crontab`;
CREATE TABLE `crontab`
(
    `id`                  int(11) NOT NULL AUTO_INCREMENT,
    `name`                varchar(255) not null default '' comment '任务名称',
    `expression`          varchar(20)  not null default '' comment '表达式',
    `type`                varchar(10)  not null default '' comment 'general 常规命令,backup 备份命令',
    `content`             text         not null default '' comment '脚本内容',
    `exec_dir`            varchar(255) not null default '' comment '执行命令目录',
    `is_immediately_exec` tinyint(1) not null default 0 comment '是否立即执行 0否,1是,执行完后将置为0',
    `is_disabled`         tinyint(1) not null default 0 comment '状态,0 启用,1禁用',
    `result`              text null default null comment '最新一次的执行结果',
    `created_at`          timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`          timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '计划任务' ROW_FORMAT = Dynamic;



-- ----------------------------
-- Records of withdraws
-- ----------------------------


DROP TABLE IF EXISTS `front_operation_log`;
CREATE TABLE `front_operation_log`
(
    `id`         int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL default 0,
    `username`   varchar(255)                                                  NOT NULL default '',
    `path`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `method`     varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL,
    `ip`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `input`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp(0) NULL DEFAULT NULL,
    `updated_at` timestamp(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX        `front_operation_log_user_id_index`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5531 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET
FOREIGN_KEY_CHECKS = 1;
