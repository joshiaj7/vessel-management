package handler

import (
	"encoding/json"
	"net/http"

	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
)

type SuccessResponse struct {
	Data       interface{}            `json:"data"`
	Meta       *util.OffsetPagination `json:"meta,omitempty"`
	StatusCode int                    `json:"status_code"`
}

func BuildErrorResponse(w http.ResponseWriter, err error) {
	e, ok := err.(entity.RequestError)
	if !ok {
		WriteHTTPResponse(w, map[string]string{"message": e.Error()}, http.StatusInternalServerError)
		return
	}

	WriteHTTPResponse(w, map[string]string{"message": e.Error()}, e.StatusCode)
}

func BuildSuccessResponse(w http.ResponseWriter, r interface{}, m *util.OffsetPagination, c int) {
	resp := &SuccessResponse{
		Data:       r,
		Meta:       m,
		StatusCode: c,
	}

	WriteHTTPResponse(w, resp, c)
}

func WriteHTTPResponse(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(body)
}
