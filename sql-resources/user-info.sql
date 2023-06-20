-- Active: 1686286058189@@1.116.251.2@3306@miajiodb
CREATE TABLE user_info (
    id INTEGER PRIMARY KEY AUTO_INCREMENT COMMENT '自增长id',
    uid VARCHAR(32) NOT NULL COMMENT '用户系统内uuid',
    username VARCHAR(32) NOT NULL COMMENT '用户名',
    head_pic VARCHAR(32) NULL COMMENT '用户头像id',
    email VARCHAR(128) NOT NULL COMMENT '邮箱',
    password VARCHAR(32) NOT NULL COMMENT '密码',
    status INTEGER(1) DEFAULT 1 NOT NULL COMMENT '状态 0 失效 1 正常 2 停用 3 删除',
    create_time DATETIME COMMENT '注册时间',
    update_time DATETIME COMMENT '修改时间'
) COMMENT '用户表' CHAR SET "utf8mb4";