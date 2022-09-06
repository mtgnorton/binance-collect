DROP TABLE IF EXISTS `forum_users`;
CREATE TABLE `forum_users`
(
    `id`                     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`               varchar(45)      NOT NULL COMMENT '用户名',
    `email`                  varchar(50)      not null default '' comment 'email',
    `description`            varchar(255)     not null default '' comment '简介',
    `password`               char(32)         NOT NULL COMMENT 'MD5密码',
    `avatar`                 varchar(200)              DEFAULT NULL COMMENT '头像地址',
    `status`                 varchar(10)               DEFAULT 'normal' COMMENT '状态：000000  低位->高位 第一位为1 禁止登录,第二位为1 禁止发帖,第三位为1 禁止回复',
    `posts_amount`           int(11) unsigned not null default 0 comment '创建主题次数',
    `reply_amount`           int(11) unsigned not null default 0 comment '回复次数',
    `shielded_amount`        int(11) unsigned not null default 0 comment '被屏蔽次数',
    `follow_by_other_amount` int(11) unsigned not null default 0 comment '被关注次数',
    `today_activity`         int(11) unsigned not null default 0 comment '今日活跃度',
    `remark`                 varchar(500)              DEFAULT NULL COMMENT '备注',
    `last_login_ip`          varchar(50)               DEFAULT null COMMENT '最后登陆IP',
    `last_login_time`        datetime                  DEFAULT NULL COMMENT '最后登陆时间',
    `created_at`             datetime                  DEFAULT NULL COMMENT '注册时间',
    `updated_at`             datetime                  DEFAULT NULL COMMENT '更新时间',
    `deleted_at`             datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛用户表';


DROP TABLE IF EXISTS `forum_posts`;
CREATE TABLE `forum_posts`
(
    `id`                int(11) unsigned NOT NULL AUTO_INCREMENT,
    `node_id`           int(11) unsigned not null default 0 comment '节点id',
    `user_id`           int(11) unsigned not null default 0 comment '用户id',
    `username`          varchar(45)      NOT NULL COMMENT '用户名',
    `title`             varchar(255)     NOT NULL COMMENT '标题',
    `content`           longtext         NULL COMMENT '内容',
    `top_end_time`      datetime                  DEFAULT NULL COMMENT '置顶截止时间,为空说明没有置顶',
    `character_amount`  int(11) unsigned not null default 0 comment '字符长度',
    `visits_amount`     int(11) unsigned not null default 0 comment '访问次数',
    `collection_amount` int(11) unsigned not null default 0 comment '收藏次数',
    `reply_amount`      int(11) unsigned not null default 0 comment '回复次数',
    `thanks_amount`     int(11) unsigned not null default 0 comment '感谢次数',
    `shielded_amount`   int(11) unsigned not null default 0 comment '被屏蔽次数',
    `weight`            int(11)          not null default 0 comment '权重',
    `reply_last_time`   datetime                  DEFAULT NULL COMMENT '最后回复时间',
    `created_at`        datetime                  DEFAULT NULL COMMENT '主题创建时间',
    `updated_at`        datetime                  DEFAULT NULL COMMENT '主题更新时间',
    `deleted_at`        datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_node_id` (`node_id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛主题表';



DROP TABLE IF EXISTS `forum_replies`;
CREATE TABLE `forum_replies`
(
    `id`                int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`           int(11) unsigned not null default 0 comment '用户id',
    `username`          varchar(45)      NOT NULL COMMENT '用户名',
    `relation_user_ids` varchar(255)     not null default '' comment '涉及用户ids，多个以逗号分隔',
    `content`           longtext         NULL COMMENT '内容',
    `character_amount`  int(11) unsigned not null default 0 comment '字符长度',
    `thanks_amount`     int(11) unsigned not null default 0 comment '感谢次数',
    `shielded_amount`   int(11) unsigned not null default 0 comment '被屏蔽次数',
    `created_at`        datetime                  DEFAULT NULL COMMENT '创建时间',
    `updated_at`        datetime                  DEFAULT NULL COMMENT '更新时间',
    `deleted_at`        datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛回复表';

DROP TABLE IF EXISTS `forum_nodes`;
CREATE TABLE `forum_nodes`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(45)      NOT NULL COMMENT '节点名称',
    `description` text             NULL COMMENT '节点描述',
    `is_index`    tinyint(1)       not null default 0 comment '是否首页显示',
    `is_can_edit` tinyint(1)       not null default 0 comment '是否允许编辑',
    `created_at`  datetime                  DEFAULT NULL COMMENT '创建时间',
    `deleted_at`  datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛节点表';



DROP TABLE IF EXISTS `forum_messages`;
CREATE TABLE `forum_messages`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`          int(11) unsigned not null default 0 comment '用户id',
    `username`         varchar(45)      NOT NULL COMMENT '用户名',
    `replied_user_id`  int(11) unsigned not null default 0 comment '被回复用户id,用户a向用户b回复，用户b为 被回复用户id',
    `replied_username` varchar(45)      NOT NULL COMMENT '被回复用户名',
    `post_id`          int(11) unsigned not null default 0 comment '关联主题id',
    `reply_id`         int(11) unsigned not null default 0 comment '关联回复id',
    `is_read`          tinyint(1)       not null default 0 comment '是否已读，否: 0,是: 1',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `deleted_at`       datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_replied_user_id` (`replied_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛消息表';

#
# DROP TABLE IF EXISTS `forum_user_posts_histories`;
# CREATE TABLE `forum_user_posts_histories`
# (
#     `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
#     `user_id`    int(11) unsigned not null default 0 comment '用户id',
#     `username`   varchar(45)      NOT NULL COMMENT '用户名',
#     `post_id`    int(11) unsigned not null default 0 comment '关联主题id',
#     `created_at` datetime                  DEFAULT NULL COMMENT '创建时间',
#     `deleted_at` datetime                  DEFAULT NULL COMMENT '删除时间',
#     PRIMARY KEY (`id`) USING BTREE,
#     KEY `idx_user_id` (`user_id`) USING BTREE
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   ROW_FORMAT = DYNAMIC COMMENT ='论坛余额变动表';
#
#

DROP TABLE IF EXISTS `forum_balance_change_log`;
CREATE TABLE `forum_balance_change_log`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) unsigned not null default 0 comment '用户id',
    `username`   varchar(45)      NOT NULL COMMENT '用户名',
    `type`       char(10)         not null default '' comment '每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: init',
    `amount`     int(11)          not null default 0 comment '金额',
    `balance`    int(11) unsigned not null default 0 comment '余额',
    `post_id`    int(11) unsigned not null default 0 comment '关联主题id',
    `reply_id`   int(11) unsigned not null default 0 comment '关联回复id',
    `remark`     varchar(255)     not null default '' comment '备注',
    `created_at` datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛余额变动表';



DROP TABLE IF EXISTS `forum_follow_or_shield_user_relation`;
CREATE TABLE `forum_follow_or_shield_user_relation`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned not null default 0 comment '用户id',
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `target_user_id`  int(11) unsigned not null default 0 comment '被关注｜屏蔽用户id',
    `target_username` varchar(45)      NOT NULL COMMENT '被关注｜屏蔽用户名',
    `type`            char(10)         not null default '' comment '类型 关注: follow,屏蔽: shield',
    `created_at`      datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛关注|屏蔽用户关联表';



DROP TABLE IF EXISTS `forum_thanks_or_shield_or_collect_content_relation`;
CREATE TABLE `forum_thanks_or_shield_or_collect_content_relation`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned not null default 0 comment '用户id',
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `target_id`       int(11) unsigned not null default 0 comment '被感谢｜屏蔽|收藏 主题id|回复id',
    `target_user_id`  int(11) unsigned not null default 0 comment '被感谢｜屏蔽|收藏 用户id',
    `target_username` varchar(45)      NOT NULL COMMENT '被感谢用户名',
    `type`            char(10)         not null default '' comment '类型 感谢主题: thanks_posts,感谢回复: thanks_reply,屏蔽主题: shield_posts,屏蔽回复: shield_reply,收藏主题:collect_posts',
    `created_at`      datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛感谢｜屏蔽|收藏  主题｜回复 关联表';

