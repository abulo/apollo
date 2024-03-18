package dao

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
