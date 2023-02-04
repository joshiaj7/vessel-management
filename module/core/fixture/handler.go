package fixture

import (
	"github.com/golang/mock/gomock"

	"github.com/joshiaj7/vessel-management/module/core/internal/handler"
	mock_usecase "github.com/joshiaj7/vessel-management/module/core/internal/usecase/mock"
)

type MockVesselHandler struct {
	// Usecase
	VesselUsecase *mock_usecase.MockVesselUsecase
}

func NewVesselHandler(
	ctrl *gomock.Controller,
) (*handler.VesselHandler, *MockVesselHandler) {
	mocks := &MockVesselHandler{
		VesselUsecase: mock_usecase.NewMockVesselUsecase(ctrl),
	}

	svc := handler.NewVesselHandler(
		mocks.VesselUsecase,
	)

	return svc, mocks
}

type MockVoyageHandler struct {
	// Usecase
	VoyageUsecase *mock_usecase.MockVoyageUsecase
}

func NewVoyageHandler(
	ctrl *gomock.Controller,
) (*handler.VoyageHandler, *MockVoyageHandler) {
	mocks := &MockVoyageHandler{
		VoyageUsecase: mock_usecase.NewMockVoyageUsecase(ctrl),
	}

	svc := handler.NewVoyageHandler(
		mocks.VoyageUsecase,
	)

	return svc, mocks
}
