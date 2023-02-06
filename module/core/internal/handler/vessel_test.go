package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	"github.com/joshiaj7/vessel-management/internal/testutil"
	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/fixture"
	"github.com/joshiaj7/vessel-management/module/core/param"
)

func TestVesselHandler_CreateVessel(t *testing.T) {
	type Request struct {
		req    *http.Request
		params httprouter.Params
	}

	type Response struct {
		body map[string]interface{}
		err  error
	}

	name := "Some Name"
	ownerID := 234
	naccsCode := "ABC123"

	par := param.CreateVessel{
		Name:      name,
		OwnerID:   ownerID,
		NACCSCode: naccsCode,
	}
	body, _ := json.Marshal(par)

	vessel := &entity.Vessel{
		ID:        123,
		Name:      name,
		NACCSCode: naccsCode,
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselHandler, Request)
	}{
		"success": {
			response: Response{
				body: map[string]interface{}{"data": vessel, "status_code": 201},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().CreateVessel(req.req.Context(), &param.CreateVessel{Name: name, NACCSCode: naccsCode, OwnerID: ownerID}).
					Return(vessel, nil)
			},
		},
		"CreateVessel error": {
			response: Response{
				body: map[string]interface{}{"message": "status 500: err Unexpected error"},
				err:  testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().CreateVessel(req.req.Context(), &param.CreateVessel{Name: name, NACCSCode: naccsCode, OwnerID: ownerID}).
					Return(nil, testutil.ErrorUnexpected)
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler, mocks := fixture.NewVesselHandler(ctrl)
			req, _ := http.NewRequest(http.MethodPost, "http://example.com/", bytes.NewReader(body))
			tc.mockFn(mocks, Request{
				req:    req,
				params: httprouter.Params{},
			})

			responseWriter := httptest.NewRecorder()
			handler.CreateVessel(responseWriter, req, nil)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}

func TestVesselHandler_ListVessels(t *testing.T) {
	type Request struct {
		req         *http.Request
		params      httprouter.Params
		queryParams map[string]interface{}
	}

	type Response struct {
		body map[string]interface{}
		err  error
	}

	name := "Some Name"
	ownerID := 123
	limit := "5"
	offset := "1"

	vessel := &entity.Vessel{
		ID:        1,
		OwnerID:   ownerID,
		Name:      name,
		NACCSCode: "ABC123",
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselHandler, Request)
	}{
		"success": {
			request: Request{
				queryParams: map[string]interface{}{
					"name":     name,
					"owner_id": ownerID,
					"limit":    limit,
					"offset":   offset,
				},
			},
			response: Response{
				body: map[string]interface{}{"data": []*entity.Vessel{vessel}, "meta": util.NewOffsetPagination(5, 1, 1), "status_code": 200},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {
				m.VesselUsecase.EXPECT().ListVessels(r.req.Context(), &param.ListVessels{Name: name, Limit: 5, Offset: 1, OwnerID: ownerID}).
					Return([]*entity.Vessel{vessel}, util.NewOffsetPagination(5, 1, 1), nil)
			},
		},
		"error param owner_id": {
			request: Request{
				queryParams: map[string]interface{}{
					"owner_id": "asd",
				},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 422: err Wrong param type"},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {},
		},
		"error param limit": {
			request: Request{
				queryParams: map[string]interface{}{
					"limit": "asd",
				},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 422: err Wrong param type"},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {},
		},
		"error param offset": {
			request: Request{
				queryParams: map[string]interface{}{
					"offset": "asd",
				},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 422: err Wrong param type"},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {},
		},
		"ListVessels error": {
			request: Request{
				queryParams: map[string]interface{}{
					"name":     name,
					"owner_id": ownerID,
					"limit":    limit,
					"offset":   offset,
				},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 500: err Unexpected error"},
				err:  testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {
				m.VesselUsecase.EXPECT().ListVessels(r.req.Context(), &param.ListVessels{Name: name, Limit: 5, Offset: 1, OwnerID: ownerID}).
					Return(nil, nil, testutil.ErrorUnexpected)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
			q := req.URL.Query()
			for k, v := range tc.request.queryParams {
				q.Add(k, fmt.Sprintf("%v", v))
			}
			req.URL.RawQuery = q.Encode()

			handler, mocks := fixture.NewVesselHandler(ctrl)
			tc.mockFn(mocks, Request{
				req:         req,
				params:      nil,
				queryParams: tc.request.queryParams,
			})

			responseWriter := httptest.NewRecorder()
			handler.ListVessels(responseWriter, req, nil)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}

func TestVesselHandler_GetVessel(t *testing.T) {
	type Request struct {
		req         *http.Request
		params      httprouter.Params
		queryParams map[string]interface{}
	}

	type Response struct {
		body map[string]interface{}
		err  error
	}

	id := 123
	name := "Some Name"
	naccsCode := "ABC123"

	vessel := &entity.Vessel{
		ID:        id,
		Name:      name,
		NACCSCode: naccsCode,
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselHandler, Request)
	}{
		"success": {
			request: Request{
				params: httprouter.Params{httprouter.Param{
					Key:   "id",
					Value: "123",
				}},
			},
			response: Response{
				body: map[string]interface{}{"data": vessel, "status_code": 200},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().GetVessel(req.req.Context(), &param.GetVessel{ID: id}).
					Return(vessel, nil)
			},
		},
		"error param id": {
			request: Request{
				params: httprouter.Params{httprouter.Param{
					Key:   "id",
					Value: "asd",
				}},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 422: err Wrong param type"},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {},
		},
		"error GetVessel": {
			request: Request{
				params: httprouter.Params{httprouter.Param{
					Key:   "id",
					Value: "123",
				}},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 404: err Vessel is not found"},
				err:  entity.ErrorVesselNotFound,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().GetVessel(req.req.Context(), &param.GetVessel{ID: id}).
					Return(nil, entity.ErrorVesselNotFound)
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
			handler, mocks := fixture.NewVesselHandler(ctrl)
			tc.mockFn(mocks, Request{
				req:         req,
				params:      nil,
				queryParams: tc.request.queryParams,
			})
			responseWriter := httptest.NewRecorder()
			handler.GetVessel(responseWriter, req, tc.request.params)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}

func TestVesselHandler_UpdateVessel(t *testing.T) {
	type Request struct {
		req    *http.Request
		params httprouter.Params
	}

	type Response struct {
		body map[string]interface{}
		err  error
	}

	id := 123
	ownerID := 234
	name := "Some Name"
	naccsCode := "ABC123"

	par := param.UpdateVessel{
		Name:      name,
		OwnerID:   ownerID,
		NACCSCode: naccsCode,
	}
	body, _ := json.Marshal(par)

	vessel := &entity.Vessel{
		ID:        id,
		Name:      name,
		OwnerID:   ownerID,
		NACCSCode: naccsCode,
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselHandler, Request)
	}{
		"success": {
			request: Request{
				params: httprouter.Params{httprouter.Param{
					Key:   "id",
					Value: "123",
				}},
			},
			response: Response{
				body: map[string]interface{}{"data": vessel, "status_code": 200},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().UpdateVessel(req.req.Context(), &param.UpdateVessel{ID: id, Name: name, NACCSCode: naccsCode, OwnerID: ownerID}).
					Return(vessel, nil)
			},
		},
		"error param id": {
			request: Request{
				params: httprouter.Params{httprouter.Param{
					Key:   "id",
					Value: "asd",
				}},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 422: err Wrong param type"},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {},
		},
		"UpdateVessel error": {
			request: Request{
				params: httprouter.Params{httprouter.Param{
					Key:   "id",
					Value: "123",
				}},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 404: err Vessel is not found"},
				err:  entity.ErrorVesselNotFound,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().UpdateVessel(req.req.Context(), &param.UpdateVessel{ID: id, Name: name, NACCSCode: naccsCode, OwnerID: ownerID}).
					Return(nil, entity.ErrorVesselNotFound)
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler, mocks := fixture.NewVesselHandler(ctrl)
			req, _ := http.NewRequest(http.MethodPost, "http://example.com/", bytes.NewReader(body))
			tc.mockFn(mocks, Request{
				req:    req,
				params: httprouter.Params{},
			})

			responseWriter := httptest.NewRecorder()
			handler.UpdateVessel(responseWriter, req, tc.request.params)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}
