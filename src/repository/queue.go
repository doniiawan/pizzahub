package repository

import (
	"github.com/doniiawan/pizzahub/src/domain"
)

type QueueRepo struct {
	queue chan *domain.Chef
}

func InitQueueRepo(queue chan *domain.Chef) *QueueRepo {
	return &QueueRepo{
		queue: queue,
	}
}

func (repo *QueueRepo) AddChefToQueue(chef *domain.Chef, msg string) {
	repo.queue <- chef
}

func (repo *QueueRepo) GetChefFromQueue() *domain.Chef {
	chef := <-repo.queue
	return chef
}
