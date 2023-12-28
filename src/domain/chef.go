package domain

type Chef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ChefUC interface {
	AddChef(chef Chef) (Chef, error)
}

type ChefRepo interface {
	InsertChef(chef Chef) (Chef, error)
}
