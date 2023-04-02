package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"gitlab.com/raihanlh/messenger-api/pkg/pagination"
	mock_user "gitlab.com/raihanlh/messenger-api/testing/mocks/user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/usecase"
	"gitlab.com/raihanlh/messenger-api/internal/model"
)

func Test_UserUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		req *payload.CreateRequest
	}
	tests := []struct {
		name            string
		args            args
		wantRepoResp    *model.User
		wantErrRepoResp error
		want            *payload.CreateResponse
		wantErr         assert.ErrorAssertionFunc
	}{{
		name: "Create User Usecase Success",
		args: args{
			req: &payload.CreateRequest{
				Name:     testUser.Name,
				Email:    testUser.Email,
				Password: testUser.Password,
			},
		},
		wantRepoResp:    testUser,
		wantErrRepoResp: nil,
		want: &payload.CreateResponse{
			User:    testUser,
			Message: "Create user success",
		},
		wantErr: assert.NoError,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := mock_user.NewMockRepository(ctrl)
			ctx := context.TODO()
			userRepoMock.EXPECT().Create(ctx, &model.User{
				Name:     tt.args.req.Name,
				Email:    tt.args.req.Email,
				Password: tt.args.req.Password,
			}).Return(tt.wantRepoResp, tt.wantErrRepoResp).AnyTimes()

			userUsecase := usecase.New(&dependency.Repositories{
				User: userRepoMock,
			})

			res, err := userUsecase.Create(ctx, tt.args.req)
			tt.wantErr(t, err)
			if err == nil {
				assert.Equalf(t, res, tt.want, "Create User")
			}
		})
	}
}

func Test_UserUsecase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		req *payload.UpdateRequest
	}
	tests := []struct {
		name            string
		args            args
		wantRepoResp    *model.User
		wantErrRepoResp error
		want            *payload.UpdateResponse
		wantErr         assert.ErrorAssertionFunc
	}{{
		name: "Update User Usecase Success",
		args: args{
			req: &payload.UpdateRequest{
				UserID:   testUser.ID,
				Name:     testUser.Name,
				Email:    testUser.Email,
				Password: testUser.Password,
			},
		},
		wantRepoResp:    testUser,
		wantErrRepoResp: nil,
		want: &payload.UpdateResponse{
			User:    testUser,
			Message: "Update user success",
		},
		wantErr: assert.NoError,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := mock_user.NewMockRepository(ctrl)
			ctx := context.TODO()
			userRepoMock.EXPECT().Update(ctx, &model.User{
				Model: model.Model{
					ID: tt.args.req.UserID,
				},
				Name:     tt.args.req.Name,
				Email:    tt.args.req.Email,
				Password: tt.args.req.Password,
			}).Return(tt.wantRepoResp, tt.wantErrRepoResp).AnyTimes()

			userUsecase := usecase.New(&dependency.Repositories{
				User: userRepoMock,
			})

			res, err := userUsecase.Update(ctx, tt.args.req)
			tt.wantErr(t, err)
			if err == nil {
				assert.Equalf(t, res, tt.want, "Update User")
			}
		})
	}
}

func Test_UserUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *payload.DeleteRequest
	}
	tests := []struct {
		name            string
		args            args
		wantRepoResp    *model.User
		wantErrRepoResp error
		want            *payload.DeleteResponse
		wantErr         assert.ErrorAssertionFunc
	}{{
		name: "Delete User Usecase Success",
		args: args{
			req: &payload.DeleteRequest{
				UserID: "6fd33930-d76e-401a-a3ba-7a03352812c2",
			},
		},
		wantErrRepoResp: nil,
		want: &payload.DeleteResponse{
			Message: "Delete user success",
		},
		wantErr: assert.NoError,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := mock_user.NewMockRepository(ctrl)
			ctx := context.TODO()
			userRepoMock.EXPECT().Delete(ctx, tt.args.req.UserID).Return(tt.wantErrRepoResp).AnyTimes()

			userUsecase := usecase.New(&dependency.Repositories{
				User: userRepoMock,
			})

			res, err := userUsecase.Delete(ctx, tt.args.req)
			tt.wantErr(t, err)
			if err == nil {
				assert.Equalf(t, res, tt.want, "Delete User")
			}
		})
	}
}

func Test_UserUsecase_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		req *payload.GetByIdRequest
	}
	tests := []struct {
		name            string
		args            args
		wantRepoResp    *model.User
		wantErrRepoResp error
		want            *payload.GetByIdResponse
		wantErr         assert.ErrorAssertionFunc
	}{{
		name: "Get User By Id Usecase Success",
		args: args{
			req: &payload.GetByIdRequest{
				UserID: "6fd33930-d76e-401a-a3ba-7a03352812c2",
			},
		},
		wantRepoResp:    testUser,
		wantErrRepoResp: nil,
		want: &payload.GetByIdResponse{
			User:    testUser,
			Message: fmt.Sprintf("Successfully get user with id %v", testUser.ID),
		},
		wantErr: assert.NoError,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := mock_user.NewMockRepository(ctrl)
			ctx := context.TODO()
			userRepoMock.EXPECT().GetById(ctx, tt.args.req.UserID).Return(tt.wantRepoResp, tt.wantErrRepoResp).AnyTimes()

			userUsecase := usecase.New(&dependency.Repositories{
				User: userRepoMock,
			})

			res, err := userUsecase.GetById(ctx, tt.args.req)
			tt.wantErr(t, err)
			if err == nil {
				assert.Equalf(t, res, tt.want, "Get User By Id")
			}
		})
	}
}

func Test_UserUsecase_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := &model.User{
		Model:    model.Model{ID: "6fd33930-d76e-401a-a3ba-7a03352812c2"},
		Name:     "Example User",
		Email:    "test@example.id",
		Password: "password",
	}

	type args struct {
		req *payload.GetAllRequest
	}
	tests := []struct {
		name            string
		args            args
		wantRepoResp    []*model.User
		wantErrRepoResp error
		want            *payload.GetAllResponse
		wantErr         assert.ErrorAssertionFunc
	}{{
		name: "Get User By Id Usecase Success",
		args: args{
			req: &payload.GetAllRequest{
				Pagination: pagination.Pagination{
					Page:         1,
					ItemsPerPage: 10,
					TotalPages:   1,
					TotalItems:   1,
					Sort:         "Id desc",
				},
				Search: "",
			},
		},
		wantRepoResp: []*model.User{
			testUser,
		},
		wantErrRepoResp: nil,
		want: &payload.GetAllResponse{
			Pagination: &pagination.Pagination{
				Page:         1,
				ItemsPerPage: 10,
				TotalPages:   1,
				TotalItems:   1,
				Sort:         "Id desc",
			},
			PaginatedData: []*model.User{
				testUser,
			},
			Message: "Successfully get all user",
		},
		wantErr: assert.NoError,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := mock_user.NewMockRepository(ctrl)
			ctx := context.TODO()
			userRepoMock.EXPECT().GetAll(ctx, &tt.args.req.Pagination, tt.args.req).Return(tt.wantRepoResp, tt.wantErrRepoResp).AnyTimes()

			userUsecase := usecase.New(&dependency.Repositories{
				User: userRepoMock,
			})

			res, err := userUsecase.GetAll(ctx, tt.args.req)
			tt.wantErr(t, err)
			if err == nil {
				assert.Equalf(t, res, tt.want, "Get All User")
			}
		})
	}
}
