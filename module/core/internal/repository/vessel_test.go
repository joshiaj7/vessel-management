package repository_test

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joshiaj7/vessel-management/internal/testutil"
	"github.com/joshiaj7/vessel-management/internal/util"
	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/joshiaj7/vessel-management/module/core/fixture"
	"github.com/joshiaj7/vessel-management/module/core/param"
	"github.com/stretchr/testify/assert"
)

func TestVesselRepository_LockVessel(t *testing.T) {
	rowColumns := []string{"id", "owner_id", "name", "naccs_code", "created_at", "updated_at"}
	rowValues := []driver.Value{123, 456, "Some Name", "Some Code", testutil.CreatedAt, testutil.UpdatedAt}
	query := "SELECT `id`,`owner_id`,`name`,`naccs_code`,`created_at`,`updated_at` FROM `vessels` WHERE id = ? FOR UPDATE"

	type Request struct {
		ctx    context.Context
		params *entity.Vessel
	}

	type Response struct {
		err error
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselRepository, Request, Response)
	}{
		"success": {
			request: Request{
				ctx:    context.Background(),
				params: &entity.Vessel{},
			},
			response: Response{
				err: nil,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				rows := m.SQLMock.NewRows(rowColumns)
				rows.AddRow(rowValues...)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
			},
		},
		"db error": {
			request: Request{
				ctx:    context.Background(),
				params: &entity.Vessel{},
			},
			response: Response{
				err: testutil.ErrDB,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				rows := m.SQLMock.NewRows(rowColumns)
				rows.AddRow(rowValues...)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(testutil.ErrDB)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			repo, mocks := fixture.NewVesselRepository()
			tc.mockFn(mocks, tc.request, tc.response)
			err := repo.LockVessel(context.Background(), tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
		})
	}
}

func TestVesselRepository_CreateVessel(t *testing.T) {
	query := "INSERT INTO `vessels` (`owner_id`,`name`,`naccs_code`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)"

	type Request struct {
		ctx    context.Context
		params *param.CreateVessel
	}

	type Response struct {
		result interface{}
		err    error
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselRepository, Request, Response)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.CreateVessel{
					OwnerID:   456,
					Name:      "Some Name",
					NACCSCode: "Some Code",
				},
			},
			response: Response{
				result: map[string]interface{}{"ID": 1},
				err:    nil,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				m.SQLMock.ExpectBegin()
				m.SQLMock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(456, "Some Name", "Some Code", testutil.AnyTime{}, testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.SQLMock.ExpectCommit()
			},
		},
		"db error": {
			request: Request{
				ctx: context.Background(),
				params: &param.CreateVessel{
					OwnerID:   456,
					Name:      "Some Name",
					NACCSCode: "Some Code",
				},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrDB,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				m.SQLMock.ExpectBegin()
				m.SQLMock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(456, "Some Name", "Some Code", testutil.AnyTime{}, testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1)).
					WillReturnError(testutil.ErrDB)
				m.SQLMock.ExpectRollback()
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			repo, mocks := fixture.NewVesselRepository()
			tc.mockFn(mocks, tc.request, tc.response)
			result, err := repo.CreateVessel(context.Background(), tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			if tc.response.result != nil {
				assert.NotNil(t, result)
			}
		})
	}
}

func TestVesselRepository_ListVessels(t *testing.T) {
	rowColumns := []string{"id", "owner_id", "name", "naccs_code", "created_at", "updated_at"}
	rowValues := []driver.Value{123, 456, "Some Name", "Some Code", testutil.CreatedAt, testutil.UpdatedAt}
	query1 := "SELECT `id`,`owner_id`,`name`,`naccs_code`,`created_at`,`updated_at` FROM `vessels`"
	query2 := "SELECT count(*) FROM `vessels`"

	type Request struct {
		ctx    context.Context
		params *param.ListVessels
	}

	type Response struct {
		result     []*entity.Vessel
		pagination *util.OffsetPagination
		err        error
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselRepository, Request, Response)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.ListVessels{
					Name:    "name",
					OwnerID: 123,
					Limit:   10,
					Offset:  0,
				},
			},
			response: Response{
				result: []*entity.Vessel{
					{
						ID:        123,
						OwnerID:   456,
						Name:      "Some Name",
						NACCSCode: "Some Code",
						CreatedAt: testutil.CreatedAt,
						UpdatedAt: testutil.UpdatedAt,
					},
				},
				pagination: util.NewOffsetPagination(10, 0, 1),
				err:        nil,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				additionalQuery1 := "WHERE name LIKE ? AND owner_id = ? LIMIT 10"
				rows1 := m.SQLMock.NewRows(rowColumns)
				rows1.AddRow(rowValues...)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query1 + " " + additionalQuery1)).WillReturnRows(rows1)

				additionalQuery2 := "WHERE name LIKE ? AND owner_id = ?"
				rows2 := m.SQLMock.NewRows([]string{"count"})
				rows2.AddRow(1)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query2 + " " + additionalQuery2)).WillReturnRows(rows2)
			},
		},
		"db error first": {
			request: Request{
				ctx: context.Background(),
				params: &param.ListVessels{
					Name:    "name",
					OwnerID: 123,
					Limit:   10,
					Offset:  1,
				},
			},
			response: Response{
				result:     nil,
				pagination: nil,
				err:        testutil.ErrDB,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				additionalQuery1 := "WHERE name LIKE ? AND owner_id = ? LIMIT 10"
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query1 + " " + additionalQuery1)).
					WillReturnError(testutil.ErrDB)
			},
		},
		"db error second": {
			request: Request{
				ctx: context.Background(),
				params: &param.ListVessels{
					Name:    "name",
					OwnerID: 123,
					Limit:   10,
					Offset:  1,
				},
			},
			response: Response{
				result:     nil,
				pagination: nil,
				err:        testutil.ErrDB,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				additionalQuery1 := "WHERE name LIKE ? AND owner_id = ? LIMIT 10"
				rows1 := m.SQLMock.NewRows(rowColumns)
				rows1.AddRow(rowValues...)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query1 + " " + additionalQuery1)).WillReturnRows(rows1)

				additionalQuery2 := "WHERE name LIKE ? AND owner_id = ?"
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query2 + " " + additionalQuery2)).
					WillReturnError(testutil.ErrDB)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			repo, mocks := fixture.NewVesselRepository()
			tc.mockFn(mocks, tc.request, tc.response)
			result, _, err := repo.ListVessels(context.Background(), tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			if len(tc.response.result) > 0 {
				testutil.AssertStructExAc(t, tc.response.result[0], result[0])
			}
		})
	}
}

func TestVesselRepository_GetVessel(t *testing.T) {
	rowColumns := []string{"id", "owner_id", "name", "naccs_code", "created_at", "updated_at"}
	rowValues := []driver.Value{123, 456, "Some Name", "Some Code", testutil.CreatedAt, testutil.UpdatedAt}
	query := "SELECT `id`,`owner_id`,`name`,`naccs_code`,`created_at`,`updated_at` FROM `vessels`"

	type Request struct {
		ctx    context.Context
		params *param.GetVessel
	}

	type Response struct {
		result interface{}
		err    error
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselRepository, Request, Response)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				params: &param.GetVessel{
					ID: 123,
				},
			},
			response: Response{
				result: map[string]interface{}{"ID": 123},
				err:    nil,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				rows := m.SQLMock.NewRows(rowColumns)
				rows.AddRow(rowValues...)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
			},
		},
		"db error": {
			request: Request{
				ctx:    context.Background(),
				params: &param.GetVessel{},
			},
			response: Response{
				result: nil,
				err:    testutil.ErrDB,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				rows := m.SQLMock.NewRows(rowColumns)
				rows.AddRow(rowValues...)
				m.SQLMock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(testutil.ErrDB)
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			repo, mocks := fixture.NewVesselRepository()
			tc.mockFn(mocks, tc.request, tc.response)
			result, err := repo.GetVessel(context.Background(), tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
			testutil.AssertStructExAc(t, tc.response.result, result)
		})
	}
}

func TestVesselRepository_UpdateVessel(t *testing.T) {
	query := "UPDATE `vessels` SET `naccs_code`=?,`name`=?,`owner_id`=?,`updated_at`=? WHERE `id` = ?"

	type Request struct {
		ctx    context.Context
		obj    *entity.Vessel
		params *param.UpdateVessel
	}

	type Response struct {
		err error
	}

	vessel := &entity.Vessel{
		ID:        123,
		OwnerID:   456,
		Name:      "Some Name",
		NACCSCode: "Some Code",
	}

	testcases := map[string]struct {
		request  Request
		response Response
		mockFn   func(*fixture.MockVesselRepository, Request, Response)
	}{
		"success": {
			request: Request{
				ctx: context.Background(),
				obj: vessel,
				params: &param.UpdateVessel{
					ID:        123,
					OwnerID:   789,
					Name:      "New Name",
					NACCSCode: "New Code",
				},
			},
			response: Response{

				err: nil,
			},
			mockFn: func(m *fixture.MockVesselRepository, req Request, res Response) {
				m.SQLMock.ExpectBegin()
				m.SQLMock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs("New Code", "New Name", 789, testutil.AnyTime{}, 123).
					WillReturnResult(sqlmock.NewResult(123, 1))
				m.SQLMock.ExpectCommit()
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			repo, mocks := fixture.NewVesselRepository()
			tc.mockFn(mocks, tc.request, tc.response)
			err := repo.UpdateVessel(context.Background(), tc.request.obj, tc.request.params)
			testutil.AssertErrorExAc(t, tc.response.err, err)
		})
	}
}
