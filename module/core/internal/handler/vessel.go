package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
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
	par := &param.CreateVessel{}
	json.NewDecoder(r.Body).Decode(&par)

	result, err := h.usecase.CreateVessel(r.Context(), par)
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusCreated)
}

func (h *VesselHandler) ListVessels(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var ownerID int
	limit := util.DefaultLimit
	offset := util.DefaultOffset
	var err error

	if r.URL.Query().Get("owner_id") != "" {
		ownerID, err = strconv.Atoi(r.URL.Query().Get("owner_id"))
		if err != nil {
			BuildErrorResponse(w, entity.ErrorParamType)
			return
		}
	}

	if r.URL.Query().Get("limit") != "" {
		limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			BuildErrorResponse(w, entity.ErrorParamType)
			return
		}
	}

	if r.URL.Query().Get("offset") != "" {
		offset, err = strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			BuildErrorResponse(w, entity.ErrorParamType)
			return
		}
	}

	result, pagination, err := h.usecase.ListVessels(r.Context(), &param.ListVessels{
		Name:    r.URL.Query().Get("name"),
		OwnerID: ownerID,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, pagination, http.StatusOK)
}

func (h *VesselHandler) GetVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var id int
	var err error

	id, err = strconv.Atoi(params.ByName("id"))
	if err != nil {
		BuildErrorResponse(w, entity.ErrorParamType)
		return
	}

	result, err := h.usecase.GetVessel(r.Context(), &param.GetVessel{
		ID: id,
	})
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusOK)
}

func (h *VesselHandler) UpdateVessel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var id int
	var err error

	if params.ByName("id") != "" {
		id, err = strconv.Atoi(params.ByName("id"))
		if err != nil {
			BuildErrorResponse(w, entity.ErrorParamType)
			return
		}
	}

	fmt.Println("id")
	fmt.Println(id)

	par := &param.UpdateVessel{}
	json.NewDecoder(r.Body).Decode(&par)

	par.ID = id
	result, err := h.usecase.UpdateVessel(r.Context(), par)
	if err != nil {
		BuildErrorResponse(w, err)
		return
	}

	BuildSuccessResponse(w, result, nil, http.StatusOK)
}
