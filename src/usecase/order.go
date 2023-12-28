package usecase

import (
	"fmt"
	"time"

	"github.com/doniiawan/pizzahub/src/domain"
)

type OrderUC struct {
	order domain.OrderRepo
	menu  domain.MenuRepo
	queue domain.QueueRepo
}

func InitOrderUC(order domain.OrderRepo, menu domain.MenuRepo, queue domain.QueueRepo) *OrderUC {
	return &OrderUC{
		order: order,
		menu:  menu,
		queue: queue,
	}
}

func (usecase *OrderUC) AddOrder(order domain.Order) (domain.Order, error) {
	// getting ordered menu
	menu, err := usecase.menu.GetMenu(order.Menu.ID)
	if err != nil {
		return domain.Order{}, err
	}

	// get chef from queue
	chef := usecase.queue.GetChefFromQueue()

	// process order
	usecase.processOrder(menu)
	order.Chef = *chef
	usecase.queue.AddChefToQueue(chef, fmt.Sprintf("back after cook %s", order.Menu.Name))
	return usecase.order.InsertOrder(order)

}

func (usecase *OrderUC) processOrder(menu domain.Menu) {
	time.Sleep(time.Duration(menu.Duration) * time.Second)
}
