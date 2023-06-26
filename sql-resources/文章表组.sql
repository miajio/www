CREATE TABLE article_group_info (
    id int PRIMARY KEY AUTO_INCREMENT COMMENT '自增长文章组id',
    uid VARCHAR(32) NOT NULL COMMENT '随机生成uuid',
    group_name VARCHAR(32) NOT NULL DEFAULT '' COMMENT '该文章组的名称',
    group_val VARCHAR(32) NOT NULL DEFAULT '' COMMENT '该文章组的值',
    create_time DATETIME COMMENT '创建时间',
    update_time TIMESTAMP COMMENT '修改时间',
    `status` int DEFAULT 1 COMMENT '状态 0 失效 1 正常 2 删除'
) COMMENT '文章组,文章类型表' CHAR SET "utf8mb4";

CREATE TABLE
    article_info (
        id int PRIMARY KEY AUTO_INCREMENT COMMENT '自增长文章id',
        uid VARCHAR(32) NOT NULL COMMENT '随机生成uuid',
        `group` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '文章组,文章类型',
        title VARCHAR(32) NOT NULL DEFAULT '' COMMENT '标题',
        `describe` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '描述',
        `content` TEXT NOT NULL DEFAULT '' COMMENT '文章内容',
        create_user VARCHAR(32) NOT NULL DEFAULT '' COMMENT '创建者用户id',
        create_time TIMESTAMP COMMENT '创建时间',
        update_time TIMESTAMP COMMENT '修改时间',
        `status` int DEFAULT 1 COMMENT '状态 0 失效 1 正常 2 删除'
    )  COMMENT '文章表' CHAR SET "utf8mb4";

CREATE TABLE article_like (
    id int PRIMARY KEY AUTO_INCREMENT COMMENT '自增长文章点赞id',
    uid VARCHAR(32) NOT NULL COMMENT '随机生成uuid',
    article_uid VARCHAR(32) NOT NULL COMMENT '被点赞的文章uid',
    like_user_uid VARCHAR(32) NOT NULL COMMENT '点赞的用户uid',
    create_time TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP COMMENT '修改时间',
    `status` int DEFAULT 1 COMMENT '状态 1 点赞 2 取消点赞'
) COMMENT '文章点赞表' CHAR SET "utf8mb4";

CREATE TABLE article_remark (
    id int PRIMARY KEY AUTO_INCREMENT COMMENT '自增长文章评论id',
    uid VARCHAR(32) NOT NULL COMMENT '随机生成uuid',
    article_uid VARCHAR(32) NOT NULL COMMENT '被评论的文章uid',
    msg VARCHAR(256) NOT NULL DEFAULT '' COMMENT '评论内容',
    reply_uid VARCHAR(32) NOT NULL DEFAULT '' COMMENT '回复评论uid',
    remark_user_uid VARCHAR(32) NOT NULL COMMENT '评论用户uid',
    create_time TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP COMMENT '修改时间',
    `status` int DEFAULT 1 COMMENT '状态 1 可见 2 删除'
) COMMENT '文章评论表' CHAR SET 'utf8mb4';