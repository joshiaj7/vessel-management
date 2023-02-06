package config

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/joshiaj7/vessel-management/module/core/internal/handler"
)

func RegisterCoreHandler(router *httprouter.Router, usecase *CoreUsecase) {

	voyageHandler := handler.NewVoyageHandler(
		usecase.Voyage,
	)

	vesselHandler := handler.NewVesselHandler(
		usecase.Vessel,
	)

	voyageHandler.Register(router)
	vesselHandler.Register(router)

	// control
	router.HandlerFunc("GET", "/healthz", healthz)
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}
