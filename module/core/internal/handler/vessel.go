package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

type VesselRepository interface {
	CreateVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ListVessels(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

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

func (h *VesselHandler) CreateVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := h.usecase.CreateVessel(r.Context(), &param.CreateVessel{
		Name:      r.URL.Query().Get("name"),
		NACCSCode: r.URL.Query().Get("naccs_code"),
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusCreated)
}

func (h *VesselHandler) ListVessels(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		fmt.Println("LIMIT")
		fmt.Println(limit, err)
		limit = util.DefaultLimit
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		fmt.Println("OFFSET")
		fmt.Println(offset, err)
		offset = util.DefaultOffset
	}

	fmt.Println(limit, offset)

	result, pagination, err := h.usecase.ListVessels(r.Context(), &param.ListVessels{
		Name:   r.URL.Query().Get("name"),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, pagination, http.StatusOK)
}

func (h *VesselHandler) GetVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := h.usecase.GetVessel(r.Context(), &param.GetVessel{
		ID:        r.URL.Query().Get("id"),
		NACCSCode: r.URL.Query().Get("naccs_code"),
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusOK)
}

func (h *VesselHandler) UpdateVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := h.usecase.UpdateVessel(r.Context(), &param.UpdateVessel{
		ID:        r.URL.Query().Get("id"),
		Name:      r.URL.Query().Get("name"),
		NACCSCode: r.URL.Query().Get("naccs_code"),
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusOK)
}
