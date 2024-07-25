package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemMenu 系统菜单 system_menu
type SystemMenu struct {
	Id            *int64        `db:"id,-" json:"id"`                      //bigint 菜单编号,PRI
	Name          *string       `db:"name" json:"name"`                    //varchar 菜单名称
	Permission    null.String   `db:"permission" json:"permission"`        //varchar 权限标识
	Type          *int32        `db:"type" json:"type"`                    //tinyint 菜单类型(1:目录/2: 菜单/3: 按钮)
	Sort          *int32        `db:"sort" json:"sort"`                    //int 显示顺序
	ParentId      *int64        `db:"parent_id" json:"parentId"`           //bigint 父菜单ID
	Path          null.String   `db:"path" json:"path"`                    //varchar 路由地址
	Icon          null.String   `db:"icon" json:"icon"`                    //varchar 菜单图标
	Component     null.String   `db:"component" json:"component"`          //varchar 组件路径
	ComponentName null.String   `db:"component_name" json:"componentName"` //varchar 组件名
	Status        *int32        `db:"status" json:"status"`                //tinyint 菜单状态(0开启/1关闭)
	Hide          *int32        `db:"hide" json:"hide"`                    //tinyint 是否隐藏(0:否/1是)
	Link          null.String   `db:"link" json:"link"`                    //varchar 路由外链时填写的访问地址
	KeepAlive     *int32        `db:"keep_alive" json:"keepAlive"`         //tinyint 是否缓存(0不/ 1是)
	Affix         *int32        `db:"affix" json:"affix"`                  //tinyint 是否总是显示(0 不显示/1 显示)
	ActivePath    null.String   `db:"active_path" json:"activePath"`       //varchar 激活链接
	FullScreen    *int32        `db:"full_screen" json:"fullScreen"`       //tinyint 是否全屏
	Redirect      null.String   `db:"redirect" json:"redirect"`            //varchar 路由重定向地址
	Creator       null.String   `db:"creator" json:"creator"`              //varchar 创建者
	CreateTime    null.DateTime `db:"create_time" json:"createTime"`       //datetime 创建时间
	Updater       null.String   `db:"updater" json:"updater"`              //varchar 更新者
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime"`       //datetime 更新时间
	Deleted       *int32        `db:"deleted" json:"deleted"`              //tinyint 是否删除
	Children      []*SystemMenu `db:"-" json:"children"`
}

// SystemMenuTreeMetaRes 菜单元信息
type SystemMenuTreeMeta struct {
	Icon        string `json:"icon"`                 // 菜单图标
	Title       string `json:"title"`                // 菜单标题
	IsLink      string `json:"isLink,omitempty"`     // 是否外链
	IsHide      bool   `json:"isHide"`               // 是否隐藏
	IsFull      bool   `json:"isFull"`               // 是否全屏
	IsAffix     bool   `json:"isAffix"`              // 是否固定
	IsKeepAlive bool   `json:"isKeepAlive"`          // 是否缓存
	ActiveMenu  string `json:"activeMenu,omitempty"` // 激活菜单
}

// SystemMenuTreeRes 菜单数据
type SystemMenuTree struct {
	Id        int64               `json:"-"`                   // 菜单ID
	ParentId  int64               `json:"-"`                   // 父菜单ID
	Type      int32               `json:"-"`                   // 菜单类型
	Path      string              `json:"path"`                // 路由地址
	Name      string              `json:"name"`                // 路由名称
	Component string              `json:"component,omitempty"` // 组件路径
	Redirect  string              `json:"redirect,omitempty"`  // 重定向地址
	Meta      *SystemMenuTreeMeta `json:"meta"`                // 菜单元信息
	Children  []*SystemMenuTree   `json:"children,omitempty"`  // 子菜单
}
