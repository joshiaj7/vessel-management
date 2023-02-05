package fixture

import (
	"github.com/golang/mock/gomock"

	mock_repository "github.com/joshiaj7/vessel-management/module/core/internal/repository/mock"
	"github.com/joshiaj7/vessel-management/module/core/internal/usecase"
)

type MockVesselUsecase struct {
	// Repository
	VesselRepository *mock_repository.MockVesselRepository
}

func NewVesselUsecase(ctrl *gomock.Controller) (usecase.VesselUsecase, *MockVesselUsecase) {
	mocks := &MockVesselUsecase{
		VesselRepository: mock_repository.NewMockVesselRepository(ctrl),
	}
	ucs := usecase.NewVesselUsecase(mocks.VesselRepository)
	return ucs, mocks
}
