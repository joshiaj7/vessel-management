package fixture

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/joshiaj7/vessel-management/internal/testutil"
	"github.com/joshiaj7/vessel-management/module/core/internal/repository"
)

type MockVesselRepository struct {
	SQLMock sqlmock.Sqlmock
}

func NewVesselRepository() (repository.VesselRepository, *MockVesselRepository) {
	db, sqlMock := testutil.NewDatabase()
	mocks := &MockVesselRepository{SQLMock: sqlMock}
	repo := repository.NewVesselRepository("", db)
	return repo, mocks
}
