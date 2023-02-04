package config

import (
	"fmt"
	"net/http"

	"github.com/joshiaj7/vessel-management/module/core/internal/handler"
	"github.com/julienschmidt/httprouter"
)

func RegisterCoreHandler(usecase *CoreUsecase) {

	shipmentHandler := handler.NewShipmentHandler(
		usecase.Shipment,
	)

	vesselHandler := handler.NewVesselHandler(
		usecase.Vessel,
	)

	router := httprouter.New()
	shipmentHandler.Register(router)
	vesselHandler.Register(router)

	router.HandlerFunc("GET", "/healthz", healthz)
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}

// func notFound(w http.ResponseWriter, _ *http.Request) {
// 	meta := response.MetaInfo{HTTPStatus: 404}
// 	res := response.BuildSuccess(nil, "not found", meta)

// 	_ = response.Write(w, res, meta.HTTPStatus)
// }
