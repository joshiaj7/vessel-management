package handler_test

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/julienschmidt/httprouter"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/joshiaj7/vessel-management/internal/testutil"
// 	"github.com/joshiaj7/vessel-management/internal/util"
// 	"github.com/joshiaj7/vessel-management/module/core/entity"
// 	"github.com/joshiaj7/vessel-management/module/core/fixture"
// 	"github.com/joshiaj7/vessel-management/module/core/param"
// )

// func TestVoyageHandler_CreateVoyage(t *testing.T) {
// 	type Request struct {
// 		req    *http.Request
// 		params httprouter.Params
// 	}

// 	type Response struct {
// 		body map[string]interface{}
// 		err  error
// 	}

// 	vesselID := "123"
// 	source := "Source"
// 	destination := "Destination"
// 	currentLocation := "Sea"

// 	req, _ := http.NewRequest(http.MethodPost, "http://example.com/", nil)
// 	q := req.URL.Query()
// 	q.Add("vessel_id", vesselID)
// 	q.Add("source", source)
// 	q.Add("destination", destination)
// 	q.Add("current_location", currentLocation)
// 	req.URL.RawQuery = q.Encode()

// 	voyage := &entity.Voyage{
// 		ID:              "1",
// 		VesselID:        vesselID,
// 		Source:          source,
// 		Destination:     destination,
// 		CurrentLocation: currentLocation,
// 		State:           "docked",
// 	}

// 	testcases := map[string]struct {
// 		request  Request
// 		response Response
// 		mockFn   func(*fixture.MockVoyageHandler, Request)
// 	}{
// 		"success": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"data": "{\"ID\":\"123\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}", "status_code": 201},
// 				err:  nil,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, req Request) {
// 				m.VoyageUsecase.EXPECT().CreateVoyage(req.req.Context(), &param.CreateVoyage{
// 					VesselID:        vesselID,
// 					Source:          source,
// 					Destination:     destination,
// 					CurrentLocation: currentLocation,
// 				}).Return(voyage, nil)
// 			},
// 		},
// 		"CreateVoyage error": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"message": "status 500: err Unexpected error"},
// 				err:  testutil.ErrorUnexpected,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, req Request) {
// 				m.VoyageUsecase.EXPECT().CreateVoyage(req.req.Context(), &param.CreateVoyage{Name: name, NACCSCode: naccsCode}).
// 					Return(nil, testutil.ErrorUnexpected)
// 			},
// 		},
// 	}

// 	for name, tc := range testcases {
// 		t.Run(name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			handler, mocks := fixture.NewVoyageHandler(ctrl)
// 			tc.mockFn(mocks, tc.request)

// 			responseWriter := httptest.NewRecorder()
// 			handler.CreateVoyage(responseWriter, tc.request.req, tc.request.params)
// 			resultBody, _ := io.ReadAll(responseWriter.Body)
// 			body, _ := json.Marshal(tc.response.body)
// 			assert.Equal(t, string(body)+"\n", string(resultBody))
// 		})
// 	}
// }

// func TestVoyageHandler_ListVoyages(t *testing.T) {
// 	type Request struct {
// 		req    *http.Request
// 		params httprouter.Params
// 	}

// 	type Response struct {
// 		body map[string]interface{}
// 		err  error
// 	}

// 	name := "Some Name"
// 	limit := "5"
// 	offset := "1"

// 	req, _ := http.NewRequest(http.MethodPost, "http://example.com/", nil)
// 	q := req.URL.Query()
// 	q.Add("name", name)
// 	q.Add("limit", limit)
// 	q.Add("offset", offset)
// 	req.URL.RawQuery = q.Encode()

// 	brokenReq, _ := http.NewRequest(http.MethodPost, "http://example.com/", nil)
// 	query := brokenReq.URL.Query()
// 	query.Add("name", name)
// 	query.Add("limit", "asd")
// 	query.Add("offset", "qwe")
// 	brokenReq.URL.RawQuery = query.Encode()

// 	vessel := &entity.Voyage{
// 		ID:        "123",
// 		Name:      name,
// 		NACCSCode: "ABC123",
// 	}

// 	testcases := map[string]struct {
// 		request  Request
// 		response Response
// 		mockFn   func(*fixture.MockVoyageHandler, Request)
// 	}{
// 		"success": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"data": "[{\"ID\":\"123\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}]", "meta": "{\"limit\":5,\"offset\":1,\"total\":1}", "status_code": 200},
// 				err:  nil,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, r Request) {
// 				m.VoyageUsecase.EXPECT().ListVoyages(r.req.Context(), &param.ListVoyages{Name: name, Limit: 5, Offset: 1}).
// 					Return([]*entity.Voyage{vessel}, util.NewOffsetPagination(5, 1, 1), nil)
// 			},
// 		},
// 		"ListVoyages error": {
// 			request: Request{
// 				req:    brokenReq,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"message": "status 500: err Unexpected error"},
// 				err:  testutil.ErrorUnexpected,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, r Request) {
// 				m.VoyageUsecase.EXPECT().ListVoyages(r.req.Context(), &param.ListVoyages{Name: name, Limit: 10, Offset: 0}).
// 					Return(nil, nil, testutil.ErrorUnexpected)
// 			},
// 		},
// 	}

// 	for name, tc := range testcases {
// 		t.Run(name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			handler, mocks := fixture.NewVoyageHandler(ctrl)
// 			tc.mockFn(mocks, tc.request)

// 			responseWriter := httptest.NewRecorder()
// 			handler.ListVoyages(responseWriter, tc.request.req, tc.request.params)
// 			resultBody, _ := io.ReadAll(responseWriter.Body)
// 			body, _ := json.Marshal(tc.response.body)
// 			assert.Equal(t, string(body)+"\n", string(resultBody))
// 		})
// 	}
// }

// func TestVoyageHandler_GetVoyage(t *testing.T) {
// 	type Request struct {
// 		req    *http.Request
// 		params httprouter.Params
// 	}

// 	type Response struct {
// 		body map[string]interface{}
// 		err  error
// 	}

// 	id := "123"
// 	name := "Some Name"
// 	naccsCode := "ABC123"

// 	req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
// 	q := req.URL.Query()
// 	q.Add("id", id)
// 	q.Add("naccs_code", naccsCode)
// 	req.URL.RawQuery = q.Encode()

// 	vessel := &entity.Voyage{
// 		ID:        id,
// 		Name:      name,
// 		NACCSCode: naccsCode,
// 	}

// 	testcases := map[string]struct {
// 		request  Request
// 		response Response
// 		mockFn   func(*fixture.MockVoyageHandler, Request)
// 	}{
// 		"success": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"data": "{\"ID\":\"123\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}", "status_code": 200},
// 				err:  nil,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, req Request) {
// 				m.VoyageUsecase.EXPECT().GetVoyage(req.req.Context(), &param.GetVoyage{ID: id, NACCSCode: naccsCode}).
// 					Return(vessel, nil)
// 			},
// 		},
// 		"GetVoyage error": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"message": "status 404: err Voyage is not found"},
// 				err:  entity.ErrorVoyageNotFound,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, req Request) {
// 				m.VoyageUsecase.EXPECT().GetVoyage(req.req.Context(), &param.GetVoyage{ID: id, NACCSCode: naccsCode}).
// 					Return(nil, entity.ErrorVoyageNotFound)
// 			},
// 		},
// 	}

// 	for name, tc := range testcases {
// 		t.Run(name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			handler, mocks := fixture.NewVoyageHandler(ctrl)
// 			tc.mockFn(mocks, tc.request)

// 			responseWriter := httptest.NewRecorder()
// 			handler.GetVoyage(responseWriter, tc.request.req, tc.request.params)
// 			resultBody, _ := io.ReadAll(responseWriter.Body)
// 			body, _ := json.Marshal(tc.response.body)
// 			assert.Equal(t, string(body)+"\n", string(resultBody))
// 		})
// 	}
// }

// func TestVoyageHandler_UpdateVoyage(t *testing.T) {
// 	type Request struct {
// 		req    *http.Request
// 		params httprouter.Params
// 	}

// 	type Response struct {
// 		body map[string]interface{}
// 		err  error
// 	}

// 	id := "123"
// 	name := "Some Name"
// 	naccsCode := "ABC123"

// 	req, _ := http.NewRequest(http.MethodPut, "http://example.com/", nil)
// 	q := req.URL.Query()
// 	q.Add("id", id)
// 	q.Add("name", name)
// 	q.Add("naccs_code", naccsCode)
// 	req.URL.RawQuery = q.Encode()

// 	vessel := &entity.Voyage{
// 		ID:        id,
// 		Name:      name,
// 		NACCSCode: naccsCode,
// 	}

// 	testcases := map[string]struct {
// 		request  Request
// 		response Response
// 		mockFn   func(*fixture.MockVoyageHandler, Request)
// 	}{
// 		"success": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"data": "{\"ID\":\"123\",\"Name\":\"Some Name\",\"NACCSCode\":\"ABC123\",\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"}", "status_code": 200},
// 				err:  nil,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, req Request) {
// 				m.VoyageUsecase.EXPECT().UpdateVoyage(req.req.Context(), &param.UpdateVoyage{ID: id, Name: name, NACCSCode: naccsCode}).
// 					Return(vessel, nil)
// 			},
// 		},
// 		"UpdateVoyage error": {
// 			request: Request{
// 				req:    req,
// 				params: httprouter.Params{},
// 			},
// 			response: Response{
// 				body: map[string]interface{}{"message": "status 404: err Voyage is not found"},
// 				err:  entity.ErrorVoyageNotFound,
// 			},
// 			mockFn: func(m *fixture.MockVoyageHandler, req Request) {
// 				m.VoyageUsecase.EXPECT().UpdateVoyage(req.req.Context(), &param.UpdateVoyage{ID: id, Name: name, NACCSCode: naccsCode}).
// 					Return(nil, entity.ErrorVoyageNotFound)
// 			},
// 		},
// 	}

// 	for name, tc := range testcases {
// 		t.Run(name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			handler, mocks := fixture.NewVoyageHandler(ctrl)
// 			tc.mockFn(mocks, tc.request)

// 			responseWriter := httptest.NewRecorder()
// 			handler.UpdateVoyage(responseWriter, tc.request.req, tc.request.params)
// 			resultBody, _ := io.ReadAll(responseWriter.Body)
// 			body, _ := json.Marshal(tc.response.body)
// 			assert.Equal(t, string(body)+"\n", string(resultBody))
// 		})
// 	}
// }
