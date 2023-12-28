package usecase

import (
	"fmt"

	"github.com/doniiawan/pizzahub/src/domain"
)

type ChefUC struct {
	chef  domain.ChefRepo
	queue domain.QueueRepo
}

func InitChefUC(chef domain.ChefRepo, queue domain.QueueRepo) *ChefUC {
	return &ChefUC{
		chef:  chef,
		queue: queue,
	}
}

func (usecase *ChefUC) AddChef(chef domain.Chef) (domain.Chef, error) {
	chef, err := usecase.chef.InsertChef(chef)
	if err != nil {
		return chef, err
	}

	usecase.queue.AddChefToQueue(&chef, fmt.Sprintf("new chef %s", chef.Name))

	return chef, nil
}
