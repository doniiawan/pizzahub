package main

import (
	"fmt"
	"net/http"

	"github.com/doniiawan/pizzahub/src/domain"
	"github.com/doniiawan/pizzahub/src/handler"
	"github.com/doniiawan/pizzahub/src/repository"
	"github.com/doniiawan/pizzahub/src/usecase"
)

func main() {

	queueChannel := make(chan *domain.Chef, 1000)
	queueRepository := repository.InitQueueRepo(queueChannel)

	menuRepo := repository.InitMenuRepo()
	menuUC := usecase.InitMenuUC(menuRepo)
	menuHandler := handler.InitMenuHTTPHandler(menuUC)
	http.HandleFunc("/menu", menuHandler.HandlerMenu)

	chefRepo := repository.InitChefRepo()
	chefUC := usecase.InitChefUC(chefRepo, queueRepository)
	chefHandler := handler.InitChefHTTPHandler(chefUC)
	http.HandleFunc("/chefs", chefHandler.HandlerChef)

	orderRepo := repository.InitOrderRepo()
	orderUC := usecase.InitOrderUC(orderRepo, menuRepo, queueRepository)
	orderHandler := handler.InitOrderHTTPHandler(orderUC)
	http.HandleFunc("/orders", orderHandler.HandlerOrder)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
