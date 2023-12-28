package handler

import (
	"encoding/json"
	"net/http"

	"github.com/doniiawan/pizzahub/src/domain"
)

type MenuHTTPHandler struct {
	usecase domain.MenuUC
}

func InitMenuHTTPHandler(usecase domain.MenuUC) *MenuHTTPHandler {
	return &MenuHTTPHandler{
		usecase: usecase,
	}
}

func (handler *MenuHTTPHandler) HandlerMenu(w http.ResponseWriter, r *http.Request) {
	var response Response
	var status int

	if r.Method != http.MethodGet {
		http.NotFound(w, r)
	}

	menus, err := handler.usecase.GetMenus()
	if err != nil {
		response.Message = err.Error()
		status = http.StatusInternalServerError
	}

	response.Message = "Success getting menus"
	response.Data = menus
	status = http.StatusOK

	respJson, _ := json.Marshal(response)
	ResponseJSON(w, status, respJson)
}
