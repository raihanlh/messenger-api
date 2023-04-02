package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitlab.com/raihanlh/messenger-api/pkg/pagination"
	"gitlab.com/raihanlh/messenger-api/pkg/validator"
	mock_user "gitlab.com/raihanlh/messenger-api/testing/mocks/user"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/delivery/handler"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

func Test_UserHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		body *payload.CreateRequest
	}

	tests := []struct {
		name            string
		args            args
		wantUsecaseResp *payload.CreateResponse
		wantUsecaseErr  error
		want            string
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "Create User Handler Success",
			args: args{
				body: &payload.CreateRequest{
					Name:     testUser.Name,
					Email:    testUser.Email,
					Password: testUser.Password,
				},
			},
			wantUsecaseResp: &payload.CreateResponse{
				User:    testUser,
				Message: "Create user success",
			},
			wantUsecaseErr: nil,
			want:           `{"status":"OK","data":{"user":{"id":"6fd33930-d76e-401a-a3ba-7a03352812c2","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","name":"Example User","email":"test@example.id"},"message":"Create user success"}}`,
			wantErr:        assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := context.TODO()

			e := echo.New()
			e.Validator = validator.New()
			data, _ := json.Marshal(tt.args.body)
			reader := bytes.NewReader(data)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/user/create", reader)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			// ctx.SetParamNames("id")
			// ctx.SetParamValues(tt.args.id)

			userUsecaseMock := mock_user.NewMockUsecase(ctrl)
			userUsecaseMock.EXPECT().Create(ct, tt.args.body).Return(tt.wantUsecaseResp, tt.wantUsecaseErr).AnyTimes()

			userHandler := handler.New(&dependency.Usecases{
				User: userUsecaseMock,
			})
			err := userHandler.Create(ctx)
			resBody := strings.Replace(rec.Body.String(), "\n", "", -1)
			if assert.NoError(t, err) {
				assert.Equalf(t, tt.want, resBody, "Create User")
			}
		})
	}
}

func Test_UserHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		body *payload.UpdateRequest
	}

	tests := []struct {
		name            string
		args            args
		wantUsecaseResp *payload.UpdateResponse
		wantUsecaseErr  error
		want            string
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "Update User Handler Success",
			args: args{
				body: &payload.UpdateRequest{
					UserID:   testUser.ID,
					Name:     testUser.Name,
					Email:    testUser.Email,
					Password: testUser.Password,
				},
			},
			wantUsecaseResp: &payload.UpdateResponse{
				User:    testUser,
				Message: "Update user success",
			},
			wantUsecaseErr: nil,
			want:           `{"status":"OK","data":{"user":{"id":"6fd33930-d76e-401a-a3ba-7a03352812c2","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","name":"Example User","email":"test@example.id"},"message":"Update user success"}}`,
			wantErr:        assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := context.TODO()

			e := echo.New()
			e.Validator = validator.New()
			data, _ := json.Marshal(tt.args.body)
			reader := bytes.NewReader(data)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/user/update", reader)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.args.body.UserID)

			userUsecaseMock := mock_user.NewMockUsecase(ctrl)
			userUsecaseMock.EXPECT().Update(ct, tt.args.body).Return(tt.wantUsecaseResp, tt.wantUsecaseErr).AnyTimes()

			userHandler := handler.New(&dependency.Usecases{
				User: userUsecaseMock,
			})
			err := userHandler.Update(ctx)
			resBody := strings.Replace(rec.Body.String(), "\n", "", -1)
			if assert.NoError(t, err) {
				assert.Equalf(t, tt.want, resBody, "Update User")
			}
		})
	}
}

func Test_UserHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		body *payload.DeleteRequest
	}

	tests := []struct {
		name            string
		args            args
		wantUsecaseResp *payload.DeleteResponse
		wantUsecaseErr  error
		want            string
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "Delete User Handler Success",
			args: args{
				body: &payload.DeleteRequest{
					UserID: testUser.ID,
				},
			},
			wantUsecaseResp: &payload.DeleteResponse{
				Message: "Delete user success",
			},
			wantUsecaseErr: nil,
			want:           `{"status":"OK","data":{"message":"Delete user success"}}`,
			wantErr:        assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := context.TODO()

			e := echo.New()
			e.Validator = validator.New()
			data, _ := json.Marshal(tt.args.body)
			reader := bytes.NewReader(data)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/user/delete", reader)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.args.body.UserID)

			userUsecaseMock := mock_user.NewMockUsecase(ctrl)
			userUsecaseMock.EXPECT().Delete(ct, tt.args.body).Return(tt.wantUsecaseResp, tt.wantUsecaseErr).AnyTimes()

			userHandler := handler.New(&dependency.Usecases{
				User: userUsecaseMock,
			})
			err := userHandler.Delete(ctx)
			resBody := strings.Replace(rec.Body.String(), "\n", "", -1)
			if assert.NoError(t, err) {
				assert.Equalf(t, tt.want, resBody, "Delete User")
			}
		})
	}
}

func Test_UserHandler_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		body *payload.GetByIdRequest
	}

	tests := []struct {
		name            string
		args            args
		wantUsecaseResp *payload.GetByIdResponse
		wantUsecaseErr  error
		want            string
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "Delete User Handler Success",
			args: args{
				body: &payload.GetByIdRequest{
					UserID: testUser.ID,
				},
			},
			wantUsecaseResp: &payload.GetByIdResponse{
				User:    testUser,
				Message: fmt.Sprintf("Successfully get user with id %v", testUser.ID),
			},
			wantUsecaseErr: nil,
			want:           fmt.Sprintf(`{"status":"OK","data":{"user":{"id":"6fd33930-d76e-401a-a3ba-7a03352812c2","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","name":"Example User","email":"test@example.id"},"message":"Successfully get user with id %v"}}`, testUser.ID),
			wantErr:        assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := context.TODO()

			e := echo.New()
			e.Validator = validator.New()
			data, _ := json.Marshal(tt.args.body)
			reader := bytes.NewReader(data)
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/user/%s", tt.args.body.UserID), reader)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.args.body.UserID)

			userUsecaseMock := mock_user.NewMockUsecase(ctrl)
			userUsecaseMock.EXPECT().GetById(ct, tt.args.body).Return(tt.wantUsecaseResp, tt.wantUsecaseErr).AnyTimes()

			userHandler := handler.New(&dependency.Usecases{
				User: userUsecaseMock,
			})
			err := userHandler.GetById(ctx)
			resBody := strings.Replace(rec.Body.String(), "\n", "", -1)
			if assert.NoError(t, err) {
				assert.Equalf(t, tt.want, resBody, "Get User By Id")
			}
		})
	}
}

func Test_UserHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	pgn := pagination.Pagination{
		Page:         1,
		ItemsPerPage: 10,
		TotalPages:   1,
		TotalItems:   1,
		Sort:         "Id desc",
	}

	type args struct {
		body *payload.GetAllRequest
	}

	tests := []struct {
		name            string
		args            args
		wantUsecaseResp *payload.GetAllResponse
		wantUsecaseErr  error
		want            string
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "Delete User Handler Success",
			args: args{
				body: &payload.GetAllRequest{
					Pagination: pgn,
					Search:     "",
				},
			},
			wantUsecaseResp: &payload.GetAllResponse{
				Pagination:    &pgn,
				PaginatedData: []*model.User{testUser},
				Message:       "Successfully get all user",
			},
			wantUsecaseErr: nil,
			want:           `{"status":"OK","data":{"page":1,"perPage":10,"totalPages":1,"totalItems":1,"sort":"Id desc","paginatedData":[{"id":"6fd33930-d76e-401a-a3ba-7a03352812c2","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","name":"Example User","email":"test@example.id"}],"message":"Successfully get all user"}}`,
			wantErr:        assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := context.TODO()

			e := echo.New()
			e.Validator = validator.New()
			data, _ := json.Marshal(tt.args.body)
			reader := bytes.NewReader(data)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users", reader)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("search")
			ctx.SetParamValues(tt.args.body.Search)

			userUsecaseMock := mock_user.NewMockUsecase(ctrl)
			userUsecaseMock.EXPECT().GetAll(ct, tt.args.body).Return(tt.wantUsecaseResp, tt.wantUsecaseErr).AnyTimes()

			userHandler := handler.New(&dependency.Usecases{
				User: userUsecaseMock,
			})
			err := userHandler.GetAll(ctx)
			resBody := strings.Replace(rec.Body.String(), "\n", "", -1)
			if assert.NoError(t, err) {
				assert.Equalf(t, tt.want, resBody, "Get All User")
			}
		})
	}
}
