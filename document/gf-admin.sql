DROP TABLE  IF EXISTS `ga_administrator` ;
CREATE TABLE `ga_administrator`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`        varchar(45) NOT NULL COMMENT '用户名',
    `password`        char(32)    NOT NULL COMMENT 'MD5密码',
    `nickname`        varchar(45)  DEFAULT NULL COMMENT '昵称',
    `avatar`          varchar(200) DEFAULT NULL COMMENT '头像地址',
    `status`          varchar (10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `remark`          varchar(500) DEFAULT NULL COMMENT '备注',
    `last_login_ip`   varchar(50)  DEFAULT null COMMENT '最后登陆IP',
    `last_login_date` datetime     DEFAULT NULL COMMENT '最后登陆时间',
    `created_at`      datetime     DEFAULT NULL COMMENT '注册时间',
    `updated_at`      datetime     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员表';

DROP TABLE  IF EXISTS `ga_role` ;

CREATE TABLE `ga_role`
(
    `id`             int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(45) NOT NULL COMMENT '角色名',
    `identification` varchar(20) NOT NULL COMMENT '角色标识符',
    `sort`           tinyint(4) NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`          varchar (10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at`     datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色表';


DROP TABLE  IF EXISTS `ga_administrator_role` ;

CREATE TABLE `ga_administrator_role`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `administrator_id` int(11) NOT NULL COMMENT '管理员id',
    `role_id`          int(11) NOT NULL COMMENT '角色id',
    `created_at`       datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at`       datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员角色关联表';


DROP TABLE  IF EXISTS `ga_admin_menu` ;

CREATE TABLE `ga_admin_menu`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(20) NOT NULL COMMENT '菜单名称',
    `path` varchar (100) not null comment '请求后端路径，可以是外链',
    `parent_id`  int(11) NOT NULL default 0 comment '父id',
    `identification` varchar(20) NOT NULL COMMENT '权限标识符',
    `front_component_path` varchar(255) DEFAULT NULL COMMENT '前端组件路径',
    `icon` varchar(100) DEFAULT '#' COMMENT '菜单图标',
    `sort`           tinyint(4) NOT NULL DEFAULT 0 COMMENT '显示顺序，越小越靠前',
    `status`          varchar (10) DEFAULT 'normal' COMMENT '状态 normal 正常 disabled 禁用',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='后台菜单表';




