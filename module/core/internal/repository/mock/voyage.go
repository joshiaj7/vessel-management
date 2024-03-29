// Code generated by MockGen. DO NOT EDIT.
// Source: voyage.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	util "github.com/joshiaj7/vessel-management/internal/util"
	entity "github.com/joshiaj7/vessel-management/module/core/entity"
	param "github.com/joshiaj7/vessel-management/module/core/param"
)

// MockVoyageRepository is a mock of VoyageRepository interface.
type MockVoyageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVoyageRepositoryMockRecorder
}

// MockVoyageRepositoryMockRecorder is the mock recorder for MockVoyageRepository.
type MockVoyageRepositoryMockRecorder struct {
	mock *MockVoyageRepository
}

// NewMockVoyageRepository creates a new mock instance.
func NewMockVoyageRepository(ctrl *gomock.Controller) *MockVoyageRepository {
	mock := &MockVoyageRepository{ctrl: ctrl}
	mock.recorder = &MockVoyageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVoyageRepository) EXPECT() *MockVoyageRepositoryMockRecorder {
	return m.recorder
}

// CreateVoyage mocks base method.
func (m *MockVoyageRepository) CreateVoyage(ctx context.Context, params *param.CreateVoyage) (*entity.Voyage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVoyage", ctx, params)
	ret0, _ := ret[0].(*entity.Voyage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVoyage indicates an expected call of CreateVoyage.
func (mr *MockVoyageRepositoryMockRecorder) CreateVoyage(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVoyage", reflect.TypeOf((*MockVoyageRepository)(nil).CreateVoyage), ctx, params)
}

// GetVoyage mocks base method.
func (m *MockVoyageRepository) GetVoyage(ctx context.Context, params *param.GetVoyage) (*entity.Voyage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVoyage", ctx, params)
	ret0, _ := ret[0].(*entity.Voyage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoyage indicates an expected call of GetVoyage.
func (mr *MockVoyageRepositoryMockRecorder) GetVoyage(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoyage", reflect.TypeOf((*MockVoyageRepository)(nil).GetVoyage), ctx, params)
}

// ListVoyages mocks base method.
func (m *MockVoyageRepository) ListVoyages(ctx context.Context, params *param.ListVoyages) ([]*entity.Voyage, *util.OffsetPagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVoyages", ctx, params)
	ret0, _ := ret[0].([]*entity.Voyage)
	ret1, _ := ret[1].(*util.OffsetPagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVoyages indicates an expected call of ListVoyages.
func (mr *MockVoyageRepositoryMockRecorder) ListVoyages(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVoyages", reflect.TypeOf((*MockVoyageRepository)(nil).ListVoyages), ctx, params)
}

// UpdateVoyage mocks base method.
func (m *MockVoyageRepository) UpdateVoyage(ctx context.Context, eObj *entity.Voyage) (*entity.Voyage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVoyage", ctx, eObj)
	ret0, _ := ret[0].(*entity.Voyage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVoyage indicates an expected call of UpdateVoyage.
func (mr *MockVoyageRepositoryMockRecorder) UpdateVoyage(ctx, eObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVoyage", reflect.TypeOf((*MockVoyageRepository)(nil).UpdateVoyage), ctx, eObj)
}
