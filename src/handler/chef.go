package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/doniiawan/pizzahub/src/domain"
)

type ChefHTTPHandler struct {
	usecase domain.ChefUC
}

func InitChefHTTPHandler(usecase domain.ChefUC) *ChefHTTPHandler {
	return &ChefHTTPHandler{
		usecase: usecase,
	}
}

func (handler *ChefHTTPHandler) HandlerChef(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var chef domain.Chef

	err = json.Unmarshal(payload, &chef)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := handler.usecase.AddChef(chef)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Message: "Success Adding Chef", Data: result,
	}
	responseByte, _ := json.Marshal(response)

	ResponseJSON(w, http.StatusOK, responseByte)
}
