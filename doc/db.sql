DROP TABLE IF EXISTS system_user;
CREATE TABLE system_user(
    `id` BIGINT AUTO_INCREMENT COMMENT '用户编号' ,
    `nickname` VARCHAR(255)   COMMENT '昵称' ,
    `mobile` VARCHAR(255)   COMMENT '手机号码' ,
    `username` VARCHAR(64) NOT NULL  COMMENT '用户名称' ,
    `password` VARCHAR(64) NOT NULL  COMMENT '用户密码' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '用户状态（0正常 1停用）' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除(0否 1是)' ,
    `system_tenant_id` BIGINT NOT NULL DEFAULT 0 COMMENT '租户ID' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统用户';


CREATE INDEX `idx:list` ON system_user(system_tenant_id,deleted,status,nickname,mobile);
CREATE UNIQUE INDEX `item:login` ON system_user(username);

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
    `keep_alive` TINYINT NOT NULL DEFAULT 1 COMMENT '是否缓存(0不/ 1是)' ,
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
    `data_scope` TINYINT   COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）' ,
    `data_scope_dept` JSON   COMMENT '数据范围(指定部门数组)' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '角色状态（0正常 1停用）' ,
    `type` TINYINT NOT NULL DEFAULT 1 COMMENT '角色类型(1内置/2定义)' ,
    `remark` VARCHAR(255)   COMMENT '备注' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统角色';


CREATE INDEX `idx:list` ON system_role(system_tenant_id,deleted,type,status,sort,name);

DROP TABLE IF EXISTS system_role_menu;
CREATE TABLE system_role_menu(
    `id` BIGINT AUTO_INCREMENT COMMENT '自增编号' ,
    `system_role_id` BIGINT NOT NULL  COMMENT '角色编号' ,
    `system_menu_id` BIGINT NOT NULL  COMMENT '菜单编号' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统角色和系统菜单关联表';


CREATE UNIQUE INDEX `uniq:role_menu` ON system_role_menu(system_role_id,system_menu_id);
CREATE INDEX `idx:list` ON system_role_menu(system_tenant_id,deleted,system_menu_id);

DROP TABLE IF EXISTS system_user_role;
CREATE TABLE system_user_role(
    `id` BIGINT AUTO_INCREMENT COMMENT '自增编号' ,
    `system_user_id` BIGINT NOT NULL  COMMENT '用户编号' ,
    `system_role_id` BIGINT NOT NULL  COMMENT '角色编号' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建者' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新者' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统用户和系统角色关联表';


CREATE UNIQUE INDEX `uniq:user_role` ON system_user_role(system_user_id,system_role_id);
CREATE INDEX `idx:list` ON system_user_role(system_tenant_id,deleted,system_role_id);

DROP TABLE IF EXISTS user;
CREATE TABLE user(
    `id` BIGINT AUTO_INCREMENT COMMENT '用户编号' ,
    `nickname` VARCHAR(128)   COMMENT '用户昵称' ,
    `avatar` VARCHAR(255)   COMMENT '用户头像' ,
    `birthday` DATE   COMMENT '用户生日' ,
    `gender` TINYINT  DEFAULT 0 COMMENT '用户性别(0女/1男)' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '用户状态(0正常/1锁定)' ,
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

DROP TABLE IF EXISTS system_dict_type;
CREATE TABLE system_dict_type(
    `id` BIGINT AUTO_INCREMENT COMMENT '字典主键' ,
    `name` VARCHAR(64) NOT NULL  COMMENT '字典名称' ,
    `type` VARCHAR(64) NOT NULL  COMMENT '字典类型' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）' ,
    `remark` VARCHAR(255)   COMMENT '备注' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '字典类型';


CREATE UNIQUE INDEX `uniq:type` ON system_dict_type(type);
CREATE INDEX `idx:list` ON system_dict_type(status,type,name);

DROP TABLE IF EXISTS system_dict;
CREATE TABLE system_dict(
    `id` BIGINT AUTO_INCREMENT COMMENT '字典编码' ,
    `sort` INT NOT NULL DEFAULT 0 COMMENT '字典排序' ,
    `label` VARCHAR(64) NOT NULL  COMMENT '字典标签' ,
    `value` VARCHAR(64) NOT NULL  COMMENT '字典键值' ,
    `dict_type` VARCHAR(64) NOT NULL  COMMENT '字典类型' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）' ,
    `color_type` VARCHAR(255)   COMMENT '颜色类型' ,
    `css_class` VARCHAR(255)   COMMENT 'css 样式' ,
    `remark` VARCHAR(255)   COMMENT '备注' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '字典数据表';


CREATE INDEX `list:data` ON system_dict(dict_type,status,sort);

DROP TABLE IF EXISTS system_login_log;
CREATE TABLE system_login_log(
    `id` BIGINT AUTO_INCREMENT COMMENT '主键' ,
    `username` VARCHAR(64) NOT NULL  COMMENT '用户账号' ,
    `user_ip` VARCHAR(64) NOT NULL  COMMENT '用户ip' ,
    `user_agent` VARCHAR(255)   COMMENT 'UA' ,
    `login_time` DATETIME NOT NULL  COMMENT '登录时间' ,
    `channel` VARCHAR(64) NOT NULL  COMMENT '渠道' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '登录日志';


CREATE INDEX `idx:list` ON system_login_log(system_tenant_id,deleted,username,login_time,channel);

DROP TABLE IF EXISTS system_operate_log;
CREATE TABLE system_operate_log(
    `id` BIGINT AUTO_INCREMENT COMMENT '主键' ,
    `username` VARCHAR(64) NOT NULL  COMMENT '用户账号' ,
    `module` VARCHAR(64) NOT NULL  COMMENT '模块名称' ,
    `request_method` VARCHAR(64) NOT NULL  COMMENT '请求方法名' ,
    `request_url` VARCHAR(255) NOT NULL  COMMENT '请求地址' ,
    `user_ip` VARCHAR(64) NOT NULL  COMMENT '用户 ip' ,
    `user_agent` VARCHAR(255)   COMMENT 'UA' ,
    `go_method` VARCHAR(64) NOT NULL  COMMENT '方法名' ,
    `go_method_args` JSON   COMMENT '方法的参数' ,
    `start_time` DATETIME NOT NULL  COMMENT '操作开始时间' ,
    `duration` INT NOT NULL  COMMENT '执行时长' ,
    `channel` VARCHAR(64) NOT NULL  COMMENT '渠道' ,
    `result` TINYINT NOT NULL DEFAULT 0 COMMENT '结果(0 成功/1 失败)' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '操作日志';


CREATE INDEX `idx:list` ON system_operate_log(system_tenant_id,deleted,username,module,start_time,result);

DROP TABLE IF EXISTS system_tenant_package;
CREATE TABLE system_tenant_package(
    `id` BIGINT AUTO_INCREMENT COMMENT '套餐编号' ,
    `name` VARCHAR(255) NOT NULL  COMMENT '套餐名称' ,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）' ,
    `system_menu_ids` JSON NOT NULL  COMMENT '目录编号' ,
    `remark` VARCHAR(255)   COMMENT '备注' ,
    `deleted` TINYINT NOT NULL  COMMENT '是否删除(0否 1是)' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '租户套餐包';


CREATE INDEX `idx:list` ON system_tenant_package(deleted,status,name);

DROP TABLE IF EXISTS system_tenant;
CREATE TABLE system_tenant(
    `id` BIGINT AUTO_INCREMENT COMMENT '租户编号' ,
    `name` VARCHAR(255) NOT NULL  COMMENT '租户名称' ,
    `system_user_id` BIGINT   COMMENT '联系人ID' ,
    `contact_name` VARCHAR(255) NOT NULL  COMMENT '联系人' ,
    `contact_mobile` VARCHAR(255) NOT NULL  COMMENT '租户联系电话' ,
    `status` TINYINT NOT NULL  COMMENT '状态（0正常 1停用）' ,
    `domain` VARCHAR(255)   COMMENT '域名' ,
    `expire_date` DATE NOT NULL  COMMENT '过期时间' ,
    `account_count` BIGINT NOT NULL  COMMENT '账号数量' ,
    `system_tenant_package_id` BIGINT NOT NULL  COMMENT '套餐编号' ,
    `deleted` TINYINT NOT NULL  COMMENT '是否删除(0否 1是)' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '租户';


CREATE INDEX `idx:package` ON system_tenant(system_tenant_package_id);
CREATE INDEX `idx:list` ON system_tenant(deleted,status,name,expire_date);

DROP TABLE IF EXISTS system_dept;
CREATE TABLE system_dept(
    `id` BIGINT AUTO_INCREMENT COMMENT '部门ID' ,
    `name` VARCHAR(255) NOT NULL  COMMENT '部门名称' ,
    `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父部门ID' ,
    `sort` INT NOT NULL  COMMENT '显示顺序' ,
    `system_user_id` BIGINT  DEFAULT 0 COMMENT '负责人' ,
    `phone` VARCHAR(255)   COMMENT '联系电话' ,
    `email` VARCHAR(255)   COMMENT '邮件' ,
    `status` TINYINT NOT NULL  COMMENT '部门状态' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除' ,
    `system_tenant_id` BIGINT NOT NULL DEFAULT 0 COMMENT '租户ID' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '部门';


CREATE INDEX `idx:list` ON system_dept(system_tenant_id,deleted,status,parent_id,name,sort);

DROP TABLE IF EXISTS system_post;
CREATE TABLE system_post(
    `id` BIGINT AUTO_INCREMENT COMMENT '职位ID' ,
    `name` VARCHAR(255) NOT NULL  COMMENT '职位名称' ,
    `sort` INT NOT NULL  COMMENT '显示顺序' ,
    `status` TINYINT NOT NULL  COMMENT '状态' ,
    `deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户ID' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '职位';


CREATE INDEX `idx:list` ON system_post(system_tenant_id,deleted,status,name,sort);

DROP TABLE IF EXISTS system_user_post;
CREATE TABLE system_user_post(
    `id` BIGINT AUTO_INCREMENT COMMENT '编号' ,
    `system_user_id` BIGINT NOT NULL  COMMENT '系统用户 ID' ,
    `system_post_id` BIGINT NOT NULL  COMMENT '职位 id' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统用户职位';


CREATE UNIQUE INDEX `uniq:user_post` ON system_user_post(system_user_id,system_post_id);
CREATE INDEX `idx:list` ON system_user_post(system_tenant_id,deleted,system_post_id);

DROP TABLE IF EXISTS system_user_dept;
CREATE TABLE system_user_dept(
    `id` BIGINT AUTO_INCREMENT COMMENT '编号' ,
    `system_user_id` BIGINT NOT NULL  COMMENT '系统用户 id' ,
    `system_dept_id` BIGINT NOT NULL  COMMENT '部门 id' ,
    `deleted` TINYINT NOT NULL  COMMENT '删除' ,
    `system_tenant_id` BIGINT NOT NULL  COMMENT '租户' ,
    `creator` VARCHAR(64)   COMMENT '创建人' ,
    `create_time` DATETIME  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updater` VARCHAR(64)   COMMENT '更新人' ,
    `update_time` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
    PRIMARY KEY (id)
)  COMMENT = '系统用户部门';


CREATE INDEX `uniq:user_dept` ON system_user_dept(system_user_id,system_dept_id);
CREATE INDEX `idx:list` ON system_user_dept(system_tenant_id,deleted,system_dept_id);

