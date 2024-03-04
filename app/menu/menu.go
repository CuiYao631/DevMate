package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

type Menu struct {
	submenu *menu.Menu
	parent  *menu.MenuItem
}

// NewMenu 创建菜单
func NewMenu(Label string) *Menu {
	submenu := menu.NewMenu()
	item := menu.SubMenu(Label, submenu)
	return &Menu{
		submenu: submenu,
		parent:  item,
	}
}

// Reset 重置菜单
func (m *Menu) Reset() {
	m.parent.SubMenu = nil
}

// Add 添加菜单
func (m *Menu) Add(label string, click menu.Callback) *menu.MenuItem {
	item := menu.Text(label, nil, click)
	m.parent.Append(item)
	return item
}

// AddSubmenu 添加子菜单
func (m *Menu) AddSubmenu(label string) *menu.Menu {
	return m.submenu.AddSubmenu(label)
}

// AddText 添加文本菜单
func (m *Menu) AddText(label string) *menu.MenuItem {
	item := menu.Text(label, nil, nil)
	item.Disabled = true
	m.parent.Append(item)
	return item
}

// Remove 移除菜单
func (m *Menu) Remove(label string) {
	for i, it := range m.parent.SubMenu.Items {
		if it.Label == label {
			m.parent.SubMenu.Items = append(m.parent.SubMenu.Items[:i], m.parent.SubMenu.Items[i+1:]...)
			return
		}
	}
}

// BindKey 绑定快捷键
func (m *Menu) BindKey(label string, accelerator *keys.Accelerator) {
	for _, it := range m.parent.SubMenu.Items {
		if it.Label == label {
			it.Accelerator = accelerator
		}
	}
	m.parent.Accelerator = accelerator
}

// AddSeparator 添加分割线
func (m *Menu) AddSeparator() {
	m.submenu.AddSeparator()
}

// Build 构建菜单
func (m *Menu) Build() *menu.MenuItem {
	return m.parent
}
