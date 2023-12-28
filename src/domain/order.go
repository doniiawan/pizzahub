package domain

type Order struct {
	ID       int    `json:"id"`
	Chef     Chef   `json:"chef,omitempty"`
	Menu     Menu   `json:"menu,omitempty"`
	CustName string `json:"cust_name"`
}

type OrderUC interface {
	AddOrder(order Order) (Order, error)
}

type OrderRepo interface {
	InsertOrder(order Order) (Order, error)
}
