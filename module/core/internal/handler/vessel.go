package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
)

type VesselHandler struct {
	usecase usecase.VesselUsecase
}

func NewVesselHandler(uc usecase.VesselUsecase) *VesselHandler {
	return &VesselHandler{
		usecase: uc,
	}
}

func (h *VesselHandler) Register(router *httprouter.Router) {
	router.POST("/v1/vessels", h.CreateVessel)
	router.GET("/v1/vessels", h.ListVessels)
	router.GET("/v1/vessels/:id", h.GetVessel)
	router.PUT("/v1/vessels/:id", h.UpdateVessel)
}

func (h *VesselHandler) CreateVessel(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// h.usecase.CreateVessel(r.Context(), )

}

func (h *VesselHandler) ListVessels(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}

func (h *VesselHandler) GetVessel(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}

func (h *VesselHandler) UpdateVessel(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

}
