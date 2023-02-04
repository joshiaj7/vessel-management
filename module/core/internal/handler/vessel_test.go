package handler_test

import (
	"encoding/json"
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
	ownerID := "234"
	naccsCode := "ABC123"

	req, _ := http.NewRequest(http.MethodPost, "http://example.com/", nil)
	q := req.URL.Query()
	q.Add("name", name)
	q.Add("owner_id", ownerID)
	q.Add("naccs_code", naccsCode)
	req.URL.RawQuery = q.Encode()

	vessel := &entity.Vessel{
		ID:        "123",
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
				req:    req,
				params: httprouter.Params{},
			},
			response: Response{
				body: map[string]interface{}{"data": "{\"ID\":\"123\",\"OwnerID\":\"\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}", "status_code": 201},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().CreateVessel(req.req.Context(), &param.CreateVessel{Name: name, NACCSCode: naccsCode, OwnerID: ownerID}).
					Return(vessel, nil)
			},
		},
		"CreateVessel error": {
			request: Request{
				req:    req,
				params: httprouter.Params{},
			},
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
			tc.mockFn(mocks, tc.request)

			responseWriter := httptest.NewRecorder()
			handler.CreateVessel(responseWriter, tc.request.req, tc.request.params)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}

func TestVesselHandler_ListVessels(t *testing.T) {
	type Request struct {
		req    *http.Request
		params httprouter.Params
	}

	type Response struct {
		body map[string]interface{}
		err  error
	}

	name := "Some Name"
	ownerID := "123"
	limit := "5"
	offset := "1"

	req, _ := http.NewRequest(http.MethodPost, "http://example.com/", nil)
	q := req.URL.Query()
	q.Add("name", name)
	q.Add("owner_id", ownerID)
	q.Add("limit", limit)
	q.Add("offset", offset)
	req.URL.RawQuery = q.Encode()

	brokenReq, _ := http.NewRequest(http.MethodPost, "http://example.com/", nil)
	query := brokenReq.URL.Query()
	query.Add("name", name)
	query.Add("owner_id", ownerID)
	query.Add("limit", "asd")
	query.Add("offset", "qwe")
	brokenReq.URL.RawQuery = query.Encode()

	vessel := &entity.Vessel{
		ID: "1",

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
				req:    req,
				params: httprouter.Params{},
			},
			response: Response{
				body: map[string]interface{}{"data": "[{\"ID\":\"1\",\"OwnerID\":\"\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}]", "meta": "{\"limit\":5,\"offset\":1,\"total\":1}", "status_code": 200},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {
				m.VesselUsecase.EXPECT().ListVessels(r.req.Context(), &param.ListVessels{Name: name, Limit: 5, Offset: 1, OwnerID: ownerID}).
					Return([]*entity.Vessel{vessel}, util.NewOffsetPagination(5, 1, 1), nil)
			},
		},
		"ListVessels error": {
			request: Request{
				req:    brokenReq,
				params: httprouter.Params{},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 500: err Unexpected error"},
				err:  testutil.ErrorUnexpected,
			},
			mockFn: func(m *fixture.MockVesselHandler, r Request) {
				m.VesselUsecase.EXPECT().ListVessels(r.req.Context(), &param.ListVessels{Name: name, Limit: 10, Offset: 0, OwnerID: ownerID}).
					Return(nil, nil, testutil.ErrorUnexpected)
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler, mocks := fixture.NewVesselHandler(ctrl)
			tc.mockFn(mocks, tc.request)

			responseWriter := httptest.NewRecorder()
			handler.ListVessels(responseWriter, tc.request.req, tc.request.params)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}

func TestVesselHandler_GetVessel(t *testing.T) {
	type Request struct {
		req    *http.Request
		params httprouter.Params
	}

	type Response struct {
		body map[string]interface{}
		err  error
	}

	id := "123"
	name := "Some Name"
	naccsCode := "ABC123"

	req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
	q := req.URL.Query()
	q.Add("id", id)
	q.Add("naccs_code", naccsCode)
	req.URL.RawQuery = q.Encode()

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
				req:    req,
				params: httprouter.Params{},
			},
			response: Response{
				body: map[string]interface{}{"data": "{\"ID\":\"123\",\"OwnerID\":\"\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}", "status_code": 200},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().GetVessel(req.req.Context(), &param.GetVessel{ID: id, NACCSCode: naccsCode}).
					Return(vessel, nil)
			},
		},
		"GetVessel error": {
			request: Request{
				req:    req,
				params: httprouter.Params{},
			},
			response: Response{
				body: map[string]interface{}{"message": "status 404: err Vessel is not found"},
				err:  entity.ErrorVesselNotFound,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().GetVessel(req.req.Context(), &param.GetVessel{ID: id, NACCSCode: naccsCode}).
					Return(nil, entity.ErrorVesselNotFound)
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler, mocks := fixture.NewVesselHandler(ctrl)
			tc.mockFn(mocks, tc.request)

			responseWriter := httptest.NewRecorder()
			handler.GetVessel(responseWriter, tc.request.req, tc.request.params)
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

	id := "123"
	ownerID := "234"
	name := "Some Name"
	naccsCode := "ABC123"

	req, _ := http.NewRequest(http.MethodPut, "http://example.com/", nil)
	q := req.URL.Query()
	q.Add("id", id)
	q.Add("name", name)
	q.Add("owner_id", ownerID)
	q.Add("naccs_code", naccsCode)
	req.URL.RawQuery = q.Encode()

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
				req:    req,
				params: httprouter.Params{},
			},
			response: Response{
				body: map[string]interface{}{"data": "{\"ID\":\"123\",\"OwnerID\":\"\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}", "status_code": 200},
				err:  nil,
			},
			mockFn: func(m *fixture.MockVesselHandler, req Request) {
				m.VesselUsecase.EXPECT().UpdateVessel(req.req.Context(), &param.UpdateVessel{ID: id, Name: name, NACCSCode: naccsCode, OwnerID: ownerID}).
					Return(vessel, nil)
			},
		},
		"UpdateVessel error": {
			request: Request{
				req:    req,
				params: httprouter.Params{},
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
			tc.mockFn(mocks, tc.request)

			responseWriter := httptest.NewRecorder()
			handler.UpdateVessel(responseWriter, tc.request.req, tc.request.params)
			resultBody, _ := io.ReadAll(responseWriter.Body)
			body, _ := json.Marshal(tc.response.body)
			assert.Equal(t, string(body)+"\n", string(resultBody))
		})
	}
}
