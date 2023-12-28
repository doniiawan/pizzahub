package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/doniiawan/pizzahub/src/domain"
)

type OrderHTTPHandler struct {
	usecase domain.OrderUC
}

func InitOrderHTTPHandler(usecase domain.OrderUC) *OrderHTTPHandler {
	return &OrderHTTPHandler{
		usecase: usecase,
	}
}

func (handler *OrderHTTPHandler) HandlerOrder(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var order domain.Order

	err = json.Unmarshal(payload, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := handler.usecase.AddOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Message: "Success Processing order", Data: result,
	}
	responseByte, _ := json.Marshal(response)
	ResponseJSON(w, http.StatusOK, responseByte)
	fmt.Printf("Order from %v at : %s Time used: from %s\n", order.CustName, startTime, time.Since(startTime))
}
