package config

import (
	"fmt"
	"net/http"

	"github.com/joshiaj7/vessel-management/module/core/internal/handler"
	"github.com/julienschmidt/httprouter"
)

func RegisterCoreHandler(usecase *CoreUsecase) *httprouter.Router {

	voyageHandler := handler.NewVoyageHandler(
		usecase.Voyage,
	)

	vesselHandler := handler.NewVesselHandler(
		usecase.Vessel,
	)

	router := httprouter.New()
	voyageHandler.Register(router)
	vesselHandler.Register(router)

	// control
	router.HandlerFunc("GET", "/healthz", healthz)

	return router
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}
