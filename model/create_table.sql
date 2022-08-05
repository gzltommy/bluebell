create table if not exists user
(
    `id`          bigint(20)  not null auto_increment,
    `user_id`     bigint(20)  not null,
    `username`    varchar(64) not null,
    `password`    varchar(64) not null,
    `email`       varchar(64) not null default '',
    `gender`      tinyint(4)  not null default 0,
    `create_time` timestamp   not null default current_timestamp,
    `update_time` timestamp   not null default current_timestamp,
    primary key (`id`),
    unique key `idx_username` (`username`),
    unique key `idx_user_id` (`user_id`)
);

create table if not exists `post`
(
    `id`           bigint(20)    not null auto_increment,
    `post_id`      bigint(20)    not null comment '帖子id',
    `title`        varchar(128)  not null comment '标题',
    `content`      varchar(8192) not null comment '内容',
    `author_id`    bigint(20)    not null comment '作者的用户id',
    `community_id` bigint(20)    not null comment '所属社区',
    `status`       tinyint(4)    not null default '1' comment '帖子状态',
    `create_time`  timestamp     null     default current_timestamp comment '创建时间',
    `update_time`  timestamp     null     default current_timestamp on update current_timestamp comment '更新时间',
    primary key (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_author_id` (`author_id`),
    key `idx_community_id` (`community_id`)
);

create table if not exists `community`
(
    `id`             int(11)          not null auto_increment,
    `community_id`   int(10) unsigned not null,
    `community_name` varchar(128)     not null,
    `introduction`   varchar(256)     not null,
    `create_time`    timestamp        not null default current_timestamp,
    `update_time`    timestamp        not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
);
insert into `community`
values ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10'),
       ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00'),
       ('3', '3', 'PUBG', '大吉大利，今晚吃鸡。', '2018-08-07 08:30:00', '2018-08-07 08:30:00'),
       ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');

create table if not exists `post`
(
    `id`           bigint(20)    not null auto_increment,
    `post_id`      bigint(20)    not null comment '帖子id',
    `title`        varchar(128)  not null comment '标题',
    `content`      varchar(8192) not null comment '内容',
    `author_id`    bigint(20)    not null comment '作者的用户id',
    `community_id` bigint(20)    not null comment '所属社区',
    `status`       tinyint(4)    not null default '1' comment '帖子状态',
    `create_time`  timestamp     null     default current_timestamp comment '创建时间',
    `update_time`  timestamp     null     default current_timestamp on update current_timestamp comment '更新时间',
    primary key (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_author_id` (`author_id`),
    key `idx_community_id` (`community_id`)
);

create table if not exists `comment`
(
    `id`          bigint(20)          not null auto_increment,
    `comment_id`  bigint(20) unsigned not null,
    `content`     text                not null,
    `post_id`     bigint(20)          not null,
    `author_id`   bigint(20)          not null,
    `parent_id`   bigint(20)          not null default '0',
    `status`      tinyint(3) unsigned not null default '1',
    `create_time` timestamp           null     default current_timestamp,
    `update_time` timestamp           null     default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_comment_id` (`comment_id`),
    key `idx_author_id` (`author_id`)
);