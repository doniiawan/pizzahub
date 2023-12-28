package repository

import (
	"github.com/doniiawan/pizzahub/src/domain"
)

type ChefRepo struct {
}

var chefs []domain.Chef

func InitChefRepo() *ChefRepo {
	return &ChefRepo{}
}

func (repo *ChefRepo) InsertChef(chef domain.Chef) (domain.Chef, error) {
	chef.ID = len(chefs) + 1
	chefs = append(chefs, chef)
	return chef, nil
}
