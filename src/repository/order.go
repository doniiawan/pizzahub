package repository

import "github.com/doniiawan/pizzahub/src/domain"

type OrderRepo struct{}

var orders []domain.Order

func InitOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

func (repo *OrderRepo) GetOrders() ([]domain.Order, error) {
	return orders, nil
}

func (repo *OrderRepo) InsertOrder(order domain.Order) (domain.Order, error) {
	order.ID = len(orders) + 1
	orders = append(orders, order)
	return order, nil
}
