package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
)

type ShipmentHandler struct {
	usecase usecase.ShipmentUsecase
}

func NewShipmentHandler(uc usecase.ShipmentUsecase) *ShipmentHandler {
	return &ShipmentHandler{
		usecase: uc,
	}
}

func (h *ShipmentHandler) Register(router *httprouter.Router) {
	router.POST("/v1/shipments", h.CreateShipment)
	router.GET("/v1/shipments", h.ListShipments)
	router.GET("/v1/shipments/:id", h.GetShipment)
	router.PUT("/v1/shipments/:id", h.UpdateShipment)
}

func (h *ShipmentHandler) CreateShipment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}

func (h *ShipmentHandler) ListShipments(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}

func (h *ShipmentHandler) GetShipment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}

func (h *ShipmentHandler) UpdateShipment(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}
