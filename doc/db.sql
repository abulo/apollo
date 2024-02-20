DROP TABLE IF EXISTS system_user;
CREATE TABLE system_user(
    `id` BIGINT AUTO_INCREMENT COMMENT '用户编号' ,
    `nickname` VARCHAR(255)   COMMENT '昵称' ,
    `username` VARCHAR(64) NOT NULL  COMMENT '用户名称' ,
    `password` VARCHAR(64) NOT NULL  COMMENT '用户密码' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '用户状态（0正常 1停用）' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统用户';


CREATE UNIQUE INDEX `item:login` ON system_user(username);
CREATE INDEX `idx:status` ON system_user(status);

DROP TABLE IF EXISTS system_menu;
CREATE TABLE system_menu(
    `id` BIGINT AUTO_INCREMENT COMMENT '菜单编号' ,
    `name` VARCHAR(64) NOT NULL  COMMENT '菜单名称' ,
    `permission` VARCHAR(128)   COMMENT '权限标识' ,
    `type` TINYINT NOT NULL DEFAULT 1 COMMENT '菜单类型(1:目录/2: 菜单/3: 按钮)' ,
    `sort` INT NOT NULL DEFAULT 1 COMMENT '显示顺序' ,
    `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父菜单ID' ,
    `path` VARCHAR(255)   COMMENT '路由地址' ,
    `icon` VARCHAR(128)   COMMENT '菜单图标' ,
    `component` VARCHAR(255)   COMMENT '组件路径' ,
    `component_name` VARCHAR(255)   COMMENT '组件名' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '菜单状态(0开启/1关闭)' ,
    `hide` TINYINT NOT NULL DEFAULT 0 COMMENT '是否隐藏(0:否/1是)' ,
    `link` VARCHAR(255)   COMMENT '路由外链时填写的访问地址' ,
    `keep_alive` TINYINT NOT NULL DEFAULT 1 COMMENT '是否缓存(0不/ 是)' ,
    `affix` TINYINT NOT NULL DEFAULT 0 COMMENT '是否总是显示(0 不显示/1 显示)' ,
    `active_path` VARCHAR(255)   COMMENT '激活链接' ,
    `full_screen` TINYINT NOT NULL DEFAULT 0 COMMENT '是否全屏' ,
    `redirect` VARCHAR(255)   COMMENT '路由重定向地址' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除' ,
    PRIMARY KEY (id)
)  COMMENT = '系统菜单';


CREATE INDEX `list:parent` ON system_menu(deleted,status,type,sort,parent_id);

DROP TABLE IF EXISTS system_role;
CREATE TABLE system_role(
    `id` BIGINT AUTO_INCREMENT COMMENT '角色编号' ,
    `name` VARCHAR(64) NOT NULL  COMMENT '角色名称' ,
    `code` VARCHAR(64) NOT NULL  COMMENT '角色权限字符串' ,
    `sort` INT NOT NULL DEFAULT 0 COMMENT '显示顺序' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '角色状态（0正常 1停用）' ,
    `type` TINYINT NOT NULL  COMMENT '角色类型(1内置/2定义)' ,
    `remark` VARCHAR(255)   COMMENT '备注' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除' ,
    PRIMARY KEY (id)
)  COMMENT = '系统角色';


CREATE INDEX `idx:list` ON system_role(deleted,type,status,sort);

DROP TABLE IF EXISTS system_role_menu;
CREATE TABLE system_role_menu(
    `id` BIGINT AUTO_INCREMENT COMMENT '自增编号' ,
    `system_role_id` BIGINT NOT NULL  COMMENT '角色编号' ,
    `system_menu_id` BIGINT NOT NULL  COMMENT '菜单编号' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除' ,
    PRIMARY KEY (id)
)  COMMENT = '系统角色和系统菜单关联表';


CREATE UNIQUE INDEX `uniq:menu:role` ON system_role_menu(system_menu_id,system_role_id);
CREATE INDEX `idx:delete` ON system_role_menu(deleted);

DROP TABLE IF EXISTS system_user_role;
CREATE TABLE system_user_role(
    `id` BIGINT AUTO_INCREMENT COMMENT '自增编号' ,
    `system_user_id` BIGINT NOT NULL  COMMENT '用户编号' ,
    `system_role_id` BIGINT NOT NULL  COMMENT '角色编号' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除' ,
    PRIMARY KEY (id)
)  COMMENT = '系统用户和系统角色关联表';


CREATE INDEX `list:role` ON system_user_role(system_role_id);
CREATE INDEX `idx:list` ON system_user_role(deleted);
CREATE INDEX `list:user` ON system_user_role(system_user_id);

DROP TABLE IF EXISTS user;
CREATE TABLE user(
    `id` BIGINT AUTO_INCREMENT COMMENT '用户编号' ,
    `nickname` VARCHAR(128)   COMMENT '用户昵称' ,
    `avatar` VARCHAR(255)   COMMENT '用户头像' ,
    `birthday` DATE   COMMENT '用户生日' ,
    `gender` TINYINT   COMMENT '用户性别(0女/1男)' ,
    `status` TINYINT NOT NULL  COMMENT '用户状态(0正常/1锁定)' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '用户';

DROP TABLE IF EXISTS user_auth;
CREATE TABLE user_auth(
    `id` BIGINT AUTO_INCREMENT COMMENT '编号' ,
    `user_id` BIGINT NOT NULL  COMMENT '用户编号' ,
    `identity_type` VARCHAR(128)   COMMENT '登录类型(手机号/邮箱) 或第三方应用名称 (微信/微博等)' ,
    `identifier` VARCHAR(128)   COMMENT '手机号/邮箱/第三方的唯一标识' ,
    `credential` VARCHAR(255)   COMMENT '密码凭证 (自建账号的保存密码, 第三方的保存 token)' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '用户三方登录授权';


CREATE INDEX `idx:user` ON user_auth(user_id);
CREATE UNIQUE INDEX `item:identity` ON user_auth(identity_type,identifier);

DROP TABLE IF EXISTS user_secret;
CREATE TABLE user_secret(
    `id` BIGINT AUTO_INCREMENT COMMENT '编号' ,
    `user_id` BIGINT NOT NULL  COMMENT '用户编号' ,
    `password` VARCHAR(255)   COMMENT '用户密码' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '用户密码';


CREATE UNIQUE INDEX `uniq:user` ON user_secret(user_id);

DROP TABLE IF EXISTS region;
CREATE TABLE region(
    `id` BIGINT AUTO_INCREMENT COMMENT '区域编号' ,
    `name` VARCHAR(255) NOT NULL  COMMENT '区域名称' ,
    `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父级编号' ,
    `weather_code` VARCHAR(255)   COMMENT '天气code' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '地区表';


CREATE INDEX `idx:parent` ON region(parent_id);
CREATE INDEX `idx:name` ON region(name);

