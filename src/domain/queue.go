package domain

type QueueRepo interface {
	AddChefToQueue(chef *Chef, msg string)
	GetChefFromQueue() *Chef
}
