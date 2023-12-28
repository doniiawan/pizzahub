package usecase

import "github.com/doniiawan/pizzahub/src/domain"

type MenuUC struct {
	repo domain.MenuRepo
}

func InitMenuUC(repo domain.MenuRepo) *MenuUC {
	return &MenuUC{
		repo: repo,
	}
}

func (usecase *MenuUC) GetMenus() ([]domain.Menu, error) {
	return usecase.repo.GetMenus()
}
