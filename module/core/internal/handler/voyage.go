package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type VoyageHandler struct {
	usecase usecase.VoyageUsecase
}

func NewVoyageHandler(uc usecase.VoyageUsecase) *VoyageHandler {
	return &VoyageHandler{
		usecase: uc,
	}
}

func (h *VoyageHandler) Register(router *httprouter.Router) {
	router.POST("/v1/voyages", h.CreateVoyage)
	router.GET("/v1/voyages", h.ListVoyages)
	router.GET("/v1/voyages/:id", h.GetVoyage)
	router.PUT("/v1/voyages/:id", h.UpdateVoyage)
}

func (h *VoyageHandler) CreateVoyage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := h.usecase.CreateVoyage(r.Context(), &param.CreateVoyage{
		VesselID:        r.URL.Query().Get("vessel_id"),
		Source:          r.URL.Query().Get("source"),
		Destination:     r.URL.Query().Get("destination"),
		CurrentLocation: r.URL.Query().Get("current_location"),
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusCreated)
}

func (h *VoyageHandler) ListVoyages(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = util.DefaultLimit
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = util.DefaultOffset
	}

	result, pagination, err := h.usecase.ListVoyages(r.Context(), &param.ListVoyages{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, pagination, http.StatusOK)
}

func (h *VoyageHandler) GetVoyage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := h.usecase.GetVoyage(r.Context(), &param.GetVoyage{
		ID: r.URL.Query().Get("id"),
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusOK)
}

func (h *VoyageHandler) UpdateVoyage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := h.usecase.UpdateVoyage(r.Context(), &param.UpdateVoyage{
		CurrentLocation: r.URL.Query().Get("current_location"),
		State:           r.URL.Query().Get("state"),
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusOK)
}
