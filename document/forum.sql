DROP TABLE IF EXISTS `forum_users`;
CREATE TABLE `forum_users`
(
    `id`                     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`               varchar(45)      NOT NULL COMMENT '用户名',
    `email`                  varchar(50)      not null default '' comment 'email',
    `description`            varchar(255)     not null default '' comment '简介',
    `password`               char(32)         NOT NULL COMMENT 'MD5密码',
    `avatar`                 varchar(200)              DEFAULT NULL COMMENT '头像地址',
    `status`                 varchar(255)              DEFAULT '' COMMENT '状态：disable_login | disable_posts | disable_reply',
    `posts_amount`           int(11) unsigned not null default 0 comment '创建主题次数',
    `reply_amount`           int(11) unsigned not null default 0 comment '回复次数',
    `shielded_amount`        int(11) unsigned not null default 0 comment '被屏蔽次数',
    `follow_by_other_amount` int(11) unsigned not null default 0 comment '被关注次数',
    `today_activity`         int(11) unsigned not null default 0 comment '今日活跃度',
    `balance`                bigint unsigned  not null default 0 comment '余额',
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
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `node_id`             int(11) unsigned not null default 0 comment '节点id',
    `user_id`             int(11) unsigned not null default 0 comment '用户id',
    `username`            varchar(45)      NOT NULL COMMENT '用户名',
    `title`               varchar(255)     NOT NULL COMMENT '标题',
    `content`             longtext         NULL COMMENT '内容',
    `top_end_time`        datetime                  DEFAULT NULL COMMENT '置顶截止时间,为空说明没有置顶',
    `character_amount`    int(11) unsigned not null default 0 comment '字符长度',
    `visits_amount`       int(11) unsigned not null default 0 comment '访问次数',
    `collection_amount`   int(11) unsigned not null default 0 comment '收藏次数',
    `reply_amount`        int(11) unsigned not null default 0 comment '回复次数',
    `thanks_amount`       int(11) unsigned not null default 0 comment '感谢次数',
    `shielded_amount`     int(11) unsigned not null default 0 comment '被屏蔽次数',
    `status`              varchar(255)              DEFAULT '' COMMENT '状态：no_audit, normal, shielded',
    `weight`              int(11)          not null default 0 comment '权重',
    `reply_last_user_id`  int(11) unsigned not null default 0 comment '最后回复用户id',
    `reply_last_username` varchar(45)      NOT NULL COMMENT '最后回复用户名',
    `reply_last_time`     datetime                  DEFAULT NULL COMMENT '最后回复时间',
    `created_at`          datetime                  DEFAULT NULL COMMENT '主题创建时间',
    `updated_at`          datetime                  DEFAULT NULL COMMENT '主题更新时间',
    `deleted_at`          datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_node_id` (`node_id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛主题表';


insert into forum_posts (id, node_id, user_id, username, title, content, top_end_time, character_amount, visits_amount,
                         collection_amount, reply_amount, thanks_amount, shielded_amount, status, weight,
                         reply_last_user_id,
                         reply_last_username, reply_last_time, created_at, updated_at, deleted_at)
values (2, 1, 1, 'mtgnorton', '现在润美帝还是个好选择吗？
', 'LZ 目前人在日本，随着最近 1 个十年汽油车的逐渐淘汰，日本的衰败感真是写在脸上了， 按刀计 lz 工资已接近腰斩。所以目前考虑离开日本在北美找找机会 因为有家人朋友在美帝所以加拿大不考虑 但现在唯一也是最担心的就是美帝的治安问题， 感觉这几年美帝枪击案真是越来越随机，越来越有抽奖的味道了。。家人朋友也是最担心哪天被抽中了大奖。。。 有没有目前就在美帝的前辈现身说法下，你们是怎么看待治安问题的？是否因为治安问题产生过离开美帝的想法？

', null, 0, 0, 0, 0, 0, 0, 'normal', 0, 0, '', null,
        '2022-10-22 00:00:00', '2022-10-22 00:00:00', null);


DROP TABLE IF EXISTS `forum_replies`;
CREATE TABLE `forum_replies`
(
    `id`                int(11) unsigned NOT NULL AUTO_INCREMENT,
    `posts_id`          int(11) unsigned not null default 0 comment '主题id',
    `user_id`           int(11) unsigned not null default 0 comment '用户id',
    `username`          varchar(45)      NOT NULL COMMENT '用户名',
    `relation_user_ids` varchar(255)     not null default '' comment '涉及用户ids，多个以逗号分隔',
    `content`           longtext         NULL COMMENT '内容',
    `character_amount`  int(11) unsigned not null default 0 comment '字符长度',
    `thanks_amount`     int(11) unsigned not null default 0 comment '感谢次数',
    `shielded_amount`   int(11) unsigned not null default 0 comment '被屏蔽次数',
    `status`            varchar(255)              DEFAULT '' COMMENT '状态：no_audit, normal, shielded',
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
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`             varchar(45)      NOT NULL COMMENT '节点名称',
    `description`      text             NULL COMMENT '节点描述',
    `is_index`         tinyint(1)       not null default 0 comment '是否首页显示',
    `is_disabled_edit` tinyint(1)       not null default 0 comment '是否禁用编辑和删除,1是 0否',
    `sort`             int(11)          not null default 0 comment '显示顺序越小越靠前',
    `created_at`       datetime                  DEFAULT NULL COMMENT '创建时间',
    `deleted_at`       datetime                  DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛节点表';

insert into forum_nodes (name, description, is_index, is_disabled_edit, sort, created_at)
values ('技术', '分享创造，分享发现', 1, 0, 1, now());
insert into forum_nodes (name, description, is_index, is_disabled_edit, sort, created_at) value ('创意', '请保持友善，互帮互助', 1, 0, 2, now());
insert into forum_nodes (name, description, is_index, is_disabled_edit, sort, created_at) value ('好玩', '站点公告', 1, 1, 3, now());
insert into forum_nodes (name, description, is_index, is_disabled_edit, sort, created_at) value ('Apple', 'APPLE', 1, 0, 4, now());
# 酷工作
insert into forum_nodes (name, description, is_index, is_disabled_edit, sort, created_at) value ('酷工作', '外包/兼职/全职', 1, 0, 5, now());


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



DROP TABLE IF EXISTS `forum_balance_change_log`;
CREATE TABLE `forum_balance_change_log`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     int(11) unsigned not null default 0 comment '用户id',
    `username`    varchar(45)      NOT NULL COMMENT '用户名',
    `type`        char(20)         not null default '' comment '每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: register',
    `amount`      int(11)          not null default 0 comment '金额',
    `before`      int(11) unsigned not null default 0 comment '变动前余额',
    `after`       int(11) unsigned not null default 0 comment '变动后余额',
    `relation_id` int(11) unsigned not null default 0 comment '关联主题id或关联回复id',
    `remark`      varchar(255)     not null default '' comment '备注',
    `created_at`  datetime                  DEFAULT NULL COMMENT '创建时间',
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
  ROW_FORMAT = DYNAMIC COMMENT ='论坛 关注|屏蔽  用户关联表';



DROP TABLE IF EXISTS `forum_thanks_or_shield_or_collect_content_relation`;
CREATE TABLE `forum_thanks_or_shield_or_collect_content_relation`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`         int(11) unsigned not null default 0 comment '用户id',
    `username`        varchar(45)      NOT NULL COMMENT '用户名',
    `target_id`       int(11) unsigned not null default 0 comment '被感谢｜屏蔽|收藏 主题id|回复id',
    `target_user_id`  int(11) unsigned not null default 0 comment '被感谢｜屏蔽|收藏 用户id',
    `target_username` varchar(45)      NOT NULL COMMENT '被感谢用户名',
    `type`            char(15)         not null default '' comment '类型 感谢主题: thanks_posts,感谢回复: thanks_reply,屏蔽主题: shield_posts,屏蔽回复: shield_reply,收藏主题:collect_posts',
    `created_at`      datetime                  DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_target_user_id` (`target_user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='论坛 感谢｜屏蔽|收藏  主题｜回复 关联表';

