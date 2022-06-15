DROP TABLE IF EXISTS `ga_administrator`;
CREATE TABLE `ga_administrator`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `password`        char(32)         NOT NULL COMMENT 'MD5密码',
    `nickname`        varchar(45)  DEFAULT NULL COMMENT '昵称',
    `avatar`          varchar(200) DEFAULT NULL COMMENT '头像地址',
    `status`          varchar(10)  DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `remark`          varchar(500) DEFAULT NULL COMMENT '备注',
    `last_login_ip`   varchar(50)  DEFAULT null COMMENT '最后登陆IP',
    `last_login_date` datetime     DEFAULT NULL COMMENT '最后登陆时间',
    `created_at`      datetime     DEFAULT NULL COMMENT '注册时间',
    `updated_at`      datetime     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='管理员表';

DROP TABLE IF EXISTS `ga_role`;

CREATE TABLE `ga_role`
(
    `id`             int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(45)      NOT NULL COMMENT '角色名',
    `identification` varchar(20)      NOT NULL COMMENT '角色标识符',
    `sort`           tinyint(4)       NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`         varchar(10)               DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`     datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='角色表';


DROP TABLE IF EXISTS `ga_administrator_role`;

CREATE TABLE `ga_administrator_role`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `role_id`          int(11) unsigned NOT NULL COMMENT '角色id',
    `created_at`       datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='管理员角色关联表';


DROP TABLE IF EXISTS `ga_admin_menu`;

CREATE TABLE `ga_admin_menu`
(
    `id`                   int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`                 varchar(20)      NOT NULL COMMENT '菜单名称',
    `type`                 varchar(12)      not null comment '菜单类型 directory:目录，link:链接,operation:操作',
    `link_type`            varchar(10)      not null default 'internal' comment '链接类型：internal：本地，external:外部',
    `front_path`           varchar(100)     not null default '' comment '前端路由地址，可以是外链',
    `parent_id`            int(11) unsigned NOT NULL default 0 comment '父id',
    `identification`       varchar(40)      NOT NULL default '' COMMENT '后端权限标识符',
    `method`               varchar(10)      NOT null default '' comment '请求方法',
    `front_component_path` varchar(255)              DEFAULT NULL COMMENT '前端组件路径',
    `icon`                 varchar(100)              DEFAULT '#' COMMENT '菜单图标',
    `sort`                 tinyint(4)       NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`               varchar(10)               DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`           datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`           datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='后台菜单表';



DROP TABLE IF EXISTS `ga_role_menu`;

CREATE TABLE `ga_role_menu`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `menu_id`    int(11) unsigned NOT NULL COMMENT '管理员id',
    `role_id`    int(11) unsigned NOT NULL COMMENT '角色id',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='角色和菜单权限关联表';



DROP Table if exists `ga_casbin_rule`;

CREATE TABLE `ga_casbin_rule`
(
    `ptype` varchar(10)  DEFAULT NULL,
    `v0`    varchar(256) DEFAULT NULL,
    `v1`    varchar(256) DEFAULT NULL,
    `v2`    varchar(256) DEFAULT NULL,
    `v3`    varchar(256) DEFAULT NULL,
    `v4`    varchar(256) DEFAULT NULL,
    `v5`    varchar(256) DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `ga_config`;

CREATE table `ga_config`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `module`     varchar(255)     not null default '' comment '所属模块',
    `key`        varchar(255)     not null default '' comment '键值',
    `value`      text comment '值',
    `created_at` datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `module_key_idx` (`module`, `key`)
)
    DEFAULT CHARSET = utf8mb4
    ROW_FORMAT = DYNAMIC COMMENT ='配置表';

drop table if exists `ga_admin_log`;
create table `ga_admin_log`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `path`             varchar(255)     not null default '' comment '请求路径',
    `method`           varchar(10)      not null default '' comment '请求方法',
    `path_name`        varchar(255)     not null default '' comment '请求路径名称',
    `params`           text                      default null comment '请求参数',
    `response`         longtext                  default null comment '响应结果',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
);


drop table if exists `ga_login_log`;
create table `ga_login_log`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) unsigned NOT NULL COMMENT '管理员id',
    `ip`               varchar(30)      not null default '' comment 'ip地址',
    `browser`          varchar(10)      not null default '' comment '浏览器',
    `os`               varchar(255)     not null default '' comment '操作系统',
    `status`           varchar(10)               DEFAULT 'normal' COMMENT '状态 success 成功 fail 失败',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime                  DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
);


drop table if exists `health_record`;
create table `health_record`
(
    `id`              int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `day`             datetime                     DEFAULT NULL COMMENT '当前日期',
    `breakfast`       varchar(255)                 DEFAULT NULL COMMENT '早餐,json格式,{type:事务名称,is_later:是否晚点,has_meat:是否含肉}',
    `launch`          varchar(255)                 DEFAULT NULL COMMENT '早餐,json格式,{type:事务名称,is_later:是否晚点,has_meat:是否含肉}',
    `dinner`          varchar(255)                 DEFAULT NULL COMMENT '中餐,json格式,json格式,{type:事务名称,is_later:是否晚点,has_meat:是否含肉}',
    `drink`           varchar(255)                 DEFAULT NULL COMMENT '饮料,json格式,{coffee_americano:美式咖啡,coffee_latte:拿铁,drink_mix_sang:西洋参等,drink_cold:冷饮,additional:其他附加的饮料}',
    `sleep`           varchar(255)                 DEFAULT NULL COMMENT '睡眠,json格式,{start:开始时间,end:结束时间}',
    `run`             varchar(255)                 DEFAULT NULL COMMENT '运动,json格式,{duration:运动时间,distance:运动距离,is_nausea:是否反胃}',
    `stomach`         varchar(255)                 DEFAULT NULL COMMENT '胃,json格式,{stomachache_level:疼痛级别(0正常，1隐痛，2轻微疼痛，3较重疼痛，4重度疼痛),is_nausea:非运动是否恶心}',
    `tongue`          varchar(255)                 DEFAULT NULL COMMENT '舌,json格式,{is_thick_ease:舌苔后是否缓解,is_dry:舌苔是否干燥}',
    `defecate_amount` int(11) unsigned    not null DEFAULT 0 COMMENT '排便次数',
    `fart_is_much`    tinyint(1) unsigned not null DEFAULT 0 COMMENT '放屁是否频繁',
    `has_new_acne`    tinyint(1) unsigned not null DEFAULT 0 COMMENT '是否有新的痤疮',
    `created_at`      datetime                     DEFAULT NULL COMMENT '创建时间',
    `updated_at`      datetime                     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
)
