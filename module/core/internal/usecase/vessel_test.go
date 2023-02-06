package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/joshiaj7/vessel-management/internal/testutil"
	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/fixture"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

func TestVesselUsecase_CreateVessel(t *testing.T) {
	type Request struct {
		ctx    context.Context
		params *param.CreateVessel
	}

	type Response struct {
		result interface{}
		err    error
	}

	vessel := &entity.Vessel{
		ID:        1,
		OwnerID:   123,
		Name:      "Some Name",
		NACCSCode: "The code",
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselUsecase, Request)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.CreateVessel{
					Name:      "Some Name",
					OwnerID:   123,
					NACCSCode: "The code",
				},
			},
			response: Response{
				result: map[string]interface{}{"ID": 1},
				err:    nil,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().CreateVessel(req.ctx, req.params).
					Return(vessel, nil)
			},
		},
		"CreateVessel error": {
			request: Request{
				ctx: context.Background(),
				params: &param.CreateVessel{
					Name:      "Some Name",
					OwnerID:   123,
					NACCSCode: "The code",
				},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().CreateVessel(req.ctx, req.params).
					Return(nil, testutil.ErrorUnexpected)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ucs, mocks := fixture.NewVesselUsecase(ctrl)
			tc.mockFn(mocks, tc.request)

			result, err := ucs.CreateVessel(tc.request.ctx, tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			testutil.AssertStructExAc(t, tc.response.result, result)
		})
	}
}

func TestVesselUsecase_ListVessels(t *testing.T) {
	type Request struct {
		ctx    context.Context
		params *param.ListVessels
	}

	type Response struct {
		result []*entity.Vessel
		err    error
	}

	vessel := &entity.Vessel{
		ID:        1,
		OwnerID:   123,
		Name:      "Some Name",
		NACCSCode: "The code",
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselUsecase, Request)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.ListVessels{
					Name:    "Some Name",
					OwnerID: 123,
					Offset:  0,
					Limit:   10,
				},
			},
			response: Response{
				result: []*entity.Vessel{vessel},
				err:    nil,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().ListVessels(req.ctx, req.params).
					Return([]*entity.Vessel{vessel}, util.NewOffsetPagination(10, 0, 1), nil)
			},
		},
		"ListVessels error": {
			request: Request{
				ctx: context.Background(),
				params: &param.ListVessels{
					Name:    "Some Name",
					OwnerID: 123,
					Offset:  0,
					Limit:   10,
				},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().ListVessels(req.ctx, req.params).
					Return(nil, nil, testutil.ErrorUnexpected)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ucs, mocks := fixture.NewVesselUsecase(ctrl)
			tc.mockFn(mocks, tc.request)

			result, _, err := ucs.ListVessels(tc.request.ctx, tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			if len(tc.response.result) > 0 {
				testutil.AssertStructExAc(t, tc.response.result[0], result[0])
			}
		})
	}
}

func TestVesselUsecase_GetVessel(t *testing.T) {
	type Request struct {
		ctx    context.Context
		params *param.GetVessel
	}

	type Response struct {
		result interface{}
		err    error
	}

	vessel := &entity.Vessel{
		ID:        1,
		OwnerID:   123,
		Name:      "Some Name",
		NACCSCode: "The code",
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselUsecase, Request)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.GetVessel{
					ID: 1,
				},
			},
			response: Response{
				result: map[string]interface{}{"ID": 1},
				err:    nil,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().GetVessel(req.ctx, req.params).
					Return(vessel, nil)
			},
		},
		"GetVessel error": {
			request: Request{
				ctx: context.Background(),
				params: &param.GetVessel{
					ID: 1,
				},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().GetVessel(req.ctx, req.params).
					Return(nil, testutil.ErrorUnexpected)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ucs, mocks := fixture.NewVesselUsecase(ctrl)
			tc.mockFn(mocks, tc.request)

			result, err := ucs.GetVessel(tc.request.ctx, tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			testutil.AssertStructExAc(t, tc.response.result, result)
		})
	}
}

func TestVesselUsecase_UpdateVessel(t *testing.T) {
	type Request struct {
		ctx    context.Context
		params *param.UpdateVessel
	}

	type Response struct {
		result interface{}
		err    error
	}

	vessel := &entity.Vessel{
		ID:        1,
		OwnerID:   234,
		Name:      "The Name",
		NACCSCode: "The Code",
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselUsecase, Request)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.UpdateVessel{
					ID:        1,
					OwnerID:   234,
					Name:      "The Name",
					NACCSCode: "The Code",
				},
			},
			response: Response{
				result: map[string]interface{}{"ID": 1},
				err:    nil,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().LockVessel(req.ctx, &entity.Vessel{ID: 1}).
					Return(nil).
					Do(func(ctx context.Context, eObj *entity.Vessel) {
						*eObj = *vessel
					})
				m.VesselRepository.EXPECT().UpdateVessel(req.ctx, vessel, req.params).
					Return(nil)
			},
		},
		"LockVessel Error": {
			request: Request{
				ctx: context.Background(),
				params: &param.UpdateVessel{
					ID:        1,
					OwnerID:   234,
					Name:      "The Name",
					NACCSCode: "The Code",
				},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().LockVessel(req.ctx, &entity.Vessel{ID: 1}).
					Return(testutil.ErrorUnexpected)
			},
		},
		"UpdateVessel Error": {
			request: Request{
				ctx: context.Background(),
				params: &param.UpdateVessel{
					ID:        1,
					OwnerID:   234,
					Name:      "The Name",
					NACCSCode: "The Code",
				},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselUsecase, req Request) {
				m.VesselRepository.EXPECT().LockVessel(req.ctx, &entity.Vessel{ID: 1}).
					Return(nil).
					Do(func(ctx context.Context, eObj *entity.Vessel) {
						*eObj = *vessel
					})
				m.VesselRepository.EXPECT().UpdateVessel(req.ctx, vessel, req.params).
					Return(testutil.ErrorUnexpected)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ucs, mocks := fixture.NewVesselUsecase(ctrl)
			tc.mockFn(mocks, tc.request)

			result, err := ucs.UpdateVessel(tc.request.ctx, tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			testutil.AssertStructExAc(t, tc.response.result, result)
		})
	}
}
