package usecase

import (
	"context"
	"fmt"

	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/logger"
	"go.uber.org/zap"
)

type UserUsecase struct {
	repositories *dependency.Repositories
}

func New(r *dependency.Repositories) user.Usecase {
	return &UserUsecase{
		repositories: r,
	}
}

func (u UserUsecase) Create(ctx context.Context, req *payload.CreateRequest) (*payload.CreateResponse, error) {
	log := logger.GetLogger(ctx)

	result, err := u.repositories.User.Create(ctx, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Error("Failed to create user: ", zap.Error(err))
		return nil, err
	}

	return &payload.CreateResponse{
		User:    result,
		Message: "Create user success",
	}, nil
}

func (u UserUsecase) Update(ctx context.Context, req *payload.UpdateRequest) (*payload.UpdateResponse, error) {
	log := logger.GetLogger(ctx)

	result, err := u.repositories.User.Update(ctx, &model.User{
		Model: model.Model{
			ID: req.UserID,
		},
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Error("Failed to update user: ", zap.Error(err))
		return nil, err
	}

	return &payload.UpdateResponse{
		User:    result,
		Message: "Update user success",
	}, nil
}

func (u UserUsecase) Delete(ctx context.Context, req *payload.DeleteRequest) (*payload.DeleteResponse, error) {
	log := logger.GetLogger(ctx)

	err := u.repositories.User.Delete(ctx, req.UserID)
	if err != nil {
		log.Error("Failed to delete user: ", zap.Error(err))
		return nil, err
	}

	return &payload.DeleteResponse{
		Message: "Delete user success",
	}, nil
}

func (u UserUsecase) GetById(ctx context.Context, req *payload.GetByIdRequest) (*payload.GetByIdResponse, error) {
	log := logger.GetLogger(ctx)

	user, err := u.repositories.User.GetById(ctx, req.UserID)
	if err != nil {
		log.Error("Failed to get user by id: ", zap.Error(err))
		return nil, err
	}

	return &payload.GetByIdResponse{
		User:    user,
		Message: fmt.Sprintf("Successfully get user with id %v", user.ID),
	}, nil
}

func (u UserUsecase) GetAll(ctx context.Context, req *payload.GetAllRequest) (*payload.GetAllResponse, error) {
	log := logger.GetLogger(ctx)

	pgn := &req.Pagination
	users, err := u.repositories.User.GetAll(ctx, pgn, req)
	if err != nil {
		log.Error("Failed to get all user: ", zap.Error(err))
		return nil, err
	}

	return &payload.GetAllResponse{
		Pagination:    pgn,
		PaginatedData: users,
		Message:       "Successfully get all user",
	}, nil
}
