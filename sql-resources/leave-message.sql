-- Active: 1686286058189@@1.116.251.2@3306@miajiodb
CREATE TABLE leave_message(  
    id varchar(32) primary key COMMENT 'id to uuid',
    leave_mobile VARCHAR(32) not NULL COMMENT '留言手机号',
    leave_name VARCHAR(32) NOT NULL COMMENT '留言者的姓名',
    leave_msg VARCHAR(512) NOT NULL COMMENT '留言信息',
    create_time DATETIME COMMENT 'Create Time',
    update_time DATETIME COMMENT 'Update Time',
    status int(1) COMMENT '留言状态 1 未处理 2 不予处理 3 已处理'
) COMMENT '留言信息' CHAR SET "utf8mb4";