DROP TABLE IF EXISTS `networks`;
CREATE TABLE `networks`
(
    `id`        int(11)      NOT NULL AUTO_INCREMENT,
    `url`       varchar(255) NOT NULL COMMENT '请求地址',
    `name`      varchar(255) NOT NULL COMMENT '网络名称',
    `chain_id`  int(11)      not null default 0 comment '',
    `is_use`    tinyint(1)   not null default 0 comment '是否是正在使用的网络',
    `create_at` timestamp(0) not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at` timestamp(0) not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '网络'
  ROW_FORMAT = Dynamic;
INSERT INTO `networks` (`url`, `name`, `chain_id`, `is_use`)
VALUES ('https://mainnet.infura.io/v3/231767ef48de4aa3a985c9a721699dcc', '主链', 1, 0);
INSERT INTO `networks` (`url`, `name`, `chain_id`, `is_use`)
VALUES ('https://ropsten.infura.io/v3/231767ef48de4aa3a985c9a721699dcc', 'Ropsten', 3, 1);



DROP TABLE IF EXISTS `contracts`;
CREATE TABLE `contracts`
(
    `id`              int(11)      NOT NULL AUTO_INCREMENT,
    `symbol`          varchar(255) NOT NULL COMMENT '货币类型',
    `address`         varchar(255) NOT NULL COMMENT '合约地址',
    `decimals`        int(11)      not null default 0 comment '小数位数',
    `is_collect_open` tinyint(1)   not null default 0 comment '是否开启,1是 0否',
    `create_at`       timestamp(0) not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at`       timestamp(0) not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '合约表'
  ROW_FORMAT = Dynamic;


INSERT INTO `contracts` (`symbol`, `address`, `decimals`, `is_collect_open`)
VALUES ('BNB', '0x', '18', 1);
INSERT INTO `contracts` (`symbol`, `address`, `decimals`, `is_collect_open`)
VALUES ('BSC-USD', '0x55d398326f99059fF775485246999027B3197955', '18', 1);


DROP TABLE IF EXISTS `lose_blocks`;
CREATE TABLE `lose_blocks`
(
    `id`        int(11)      NOT NULL AUTO_INCREMENT,
    `number`    int(11)      NOT NULL COMMENT '区块号',
    `status`    tinyint(1)   not null default 0 comment '检测状态，1已检测，0未检测',
    `create_at` timestamp(0) not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at` timestamp(0) not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '重新检测区块表'
  ROW_FORMAT = Dynamic;

create unique index number_idx on lose_blocks (`number`);


DROP TABLE IF EXISTS `queue_task`;
CREATE TABLE `queue_task`
(
    `id`               int(11)          NOT NULL AUTO_INCREMENT,
    `hash`             varchar(255)     not null default '' COMMENT '',
    `symbol`           varchar(255)     NOT NULL default '' COMMENT '',
    `contract_address` varchar(255)     NOT NULL default '' COMMENT '',
    `from`             varchar(255)     NOT NULL default '' COMMENT '转出地址',
    `to`               varchar(255)     NOT NULL default '' COMMENT '转入地址',
    `value`            varchar(255)     NOT NULL default '' COMMENT '金额',
    `gas_limit`        int(11)          NOT NULL default 0 COMMENT 'gas  限制',
    `gas_price`        varchar(255)     NOT NULL default '' COMMENT 'gas 预估 价格',
    `actual_gas_used`  int(11)          NOT NULL default 0 COMMENT 'gas 实际消耗',
    `actual_gas_price` varchar(255)     NOT NULL default '' COMMENT 'gas 实际价格',
    `actual_fee`       varchar(255)     NOT NULL default '' COMMENT '实际手续费',
    `nonce`            int(11)          not null default 0 comment '',
    `type`             varchar(20)      not null default '' comment '交易类型',
    `status`           varchar(20)      not null default 0 comment '转账状态: fail 转出失败,wait 等待转出,process 转出中,success转出成功',
    `fail_amount`      int(11)          not null default 0 comment '失败次数',
    `private_key`      text             null COMMENT '私钥',
    `relation_id`      int(11) unsigned NOT NULL default 0 COMMENT '关联的其他表id',
    `create_at`        timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `send_at`          timestamp(0)     null comment '发送转账时间',
    `finish_at`        timestamp(0)     null comment '转账检测成功时间',
    `update_at`        timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '队列表'
  ROW_FORMAT = Dynamic;

create index from_idx on queue_task (`from`);
create unique index hash_idx on queue_task (`hash`);


DROP TABLE IF exists `queue_task_log`;
CREATE TABLE `queue_task_log`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `queue_task_id` int(11)          not null default 0 comment '队列任务id',
    `log`           varchar(255)     NOT NULL DEFAULT '' COMMENT '错误日志',
    `fail_amount`   int(11)          not null default 0 comment '第几次失败',
    `create_at`     timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at`     timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '队列错误日志表'
  ROW_FORMAT = Dynamic;



DROP TABLE IF EXISTS `user_addresses`;
CREATE TABLE `user_addresses`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `address`          varchar(255)     NOT NULL DEFAULT '' COMMENT '以太坊地址',
    `external_user_id` varchar(255)     NOT NULL DEFAULT '' COMMENT '外部用户id',
    `private_key`      text             null COMMENT '私钥',
    `type`             tinyint(1)       NOT NULL DEFAULT 0 COMMENT '类型:0平台生成,1外部导入',
    `create_at`        timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at`        timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '用户地址表'
  ROW_FORMAT = Dynamic;

create unique index address_idx on user_addresses (`address`);



DROP TABLE IF EXISTS `collects`;
CREATE TABLE `collects`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `symbol`           varchar(255)     NOT NULL DEFAULT '' COMMENT '代币符号',
    `contract_address` varchar(255)     NOT NULL default '' COMMENT '',
    `user_id`          int(11)          not null default 0 comment '用户id',
    `user_address`     varchar(255)     NOT NULL DEFAULT '' COMMENT '用户地址',
    `recharge_hash`    varchar(255)     NOT NULL DEFAULT '' COMMENT '充值hash',
    `handfee_hash`     varchar(255)     NOT NULL DEFAULT '' COMMENT '手续费hash',
    `collect_hash`     varchar(255)     NOT NULL DEFAULT '' COMMENT '归集hash',
    `value`            varchar(255)     NOT NULL DEFAULT '' COMMENT '归集金额',
    `recharge_value`   varchar(255)     NOT NULL DEFAULT '' COMMENT '充值金额',
    `status`           varchar(20)      NOT NULL DEFAULT '' COMMENT '状态 fail 失败，wait_fee 待转手续费，process_fee 转手续费中，wait_collect 待归集,process_wait归集中，wait_notify 待通知,finish_notify通知完成',
    `create_at`        timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at`        timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '归集表'
  ROW_FORMAT = Dynamic;

create unique index recharge_hash_idx on collects (`recharge_hash`);



DROP TABLE IF EXISTS `withdraws`;
CREATE TABLE `withdraws`
(
    `id`                int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id`           int(11)          NOT NULL DEFAULT 0 COMMENT '内部用户id',
    `user_address`      varchar(255)     NOT NULL DEFAULT '''' COMMENT '用户地址',
    `external_order_id` varchar(255)     NOT NULL DEFAULT '' COMMENT '外部订单id',
    `external_user_id`  varchar(255)     NOT NULL DEFAULT '' COMMENT '外部用户id',
    `hash`              varchar(255)     NOT NULL DEFAULT '' COMMENT 'hash',
    `symbol`            varchar(255)     NOT NULL DEFAULT '' COMMENT '代币符号',
    `contract_address`  varchar(255)     NOT NULL default '' COMMENT '',
    `from`              varchar(255)     NOT NULL DEFAULT '' COMMENT '转出地址',
    `to`                varchar(255)     NOT NULL DEFAULT '' COMMENT '转入地址',
    `value`             varchar(255)     NOT NULL DEFAULT '' COMMENT '转出金额',
    `status`            varchar(20)      NOT NULL DEFAULT '' COMMENT '状态 fail 失败，wait 待转出，process 转出中，wait_notify转出完成,finish_notify通知完成',
    `create_at`         timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at`         timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '提现表'
  ROW_FORMAT = Dynamic;



DROP TABLE IF EXISTS `notify`;
CREATE TABLE `notify`
(
    `id`                   int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `type`                 varchar(20)      NOT NULL DEFAULT '' COMMENT '交易类型, recharge 充值，withdraw 提现',
    `relation_id`          int(11) unsigned NOT NULL DEFAULT 0 COMMENT '关联表id',
    `notify_data`          text             NOT NULL DEFAULT '' COMMENT '通知数据',
    `notify_address`       varchar(255)     NOT NULL DEFAULT '' COMMENT '通知地址',
    `unique_id`            varchar(255)     NOT NULL DEFAULT '' COMMENT '唯一id',
    `fail_amount`          int(11)          NOT NULL DEFAULT 0 COMMENT '失败次数',
    `status`               varchar(20)      NOT NULL DEFAULT '' COMMENT '状态 fail 失败,wait等待通知  finish通知完成',
    `is_immediately_retry` tinyint(1)       NOT NULL DEFAULT 0 comment '是否立即重试',
    `create_at`            timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `notify_at`            timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '上次通知时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '通知表'
  ROW_FORMAT = Dynamic;

DROP TABLE IF exists `notify_log`;
CREATE TABLE `notify_log`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `notify_id`   int(11)          not null default 0 comment 'notify id',
    `log`         varchar(255)     NOT NULL DEFAULT '' COMMENT '错误日志',
    `fail_amount` int(11)          not null default 0 comment '第几次失败',
    `create_at`   timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '创建时间',
    `update_at`   timestamp(0)     not null default CURRENT_TIMESTAMP(0) comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci comment '通知错误日志表'
  ROW_FORMAT = Dynamic;


