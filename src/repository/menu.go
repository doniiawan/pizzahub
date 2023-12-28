package repository

import (
	"errors"
	"fmt"

	"github.com/doniiawan/pizzahub/src/domain"
)

type MenuRepo struct{}

var menus = []domain.Menu{
	{
		ID:       1,
		Name:     "Cheese Pizza",
		Duration: 3,
	},
	{
		ID:       2,
		Name:     "BBQ Pizza",
		Duration: 5,
	},
}

func InitMenuRepo() *MenuRepo {
	return &MenuRepo{}
}

func (repo *MenuRepo) GetMenus() ([]domain.Menu, error) {
	return menus, nil
}

func (repo *MenuRepo) GetMenu(id int) (domain.Menu, error) {
	for _, menu := range menus {
		if menu.ID == id {
			return menu, nil
		}
	}
	return domain.Menu{}, errors.New("no menu registered with id " + fmt.Sprintf("%d", id))
}
