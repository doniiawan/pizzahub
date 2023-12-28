package domain

type Menu struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Duration int64  `json:"duration"`
}

type MenuUC interface {
	GetMenus() ([]Menu, error)
}

type MenuRepo interface {
	GetMenus() ([]Menu, error)
	GetMenu(int) (Menu, error)
}
