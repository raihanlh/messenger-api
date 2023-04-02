package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/raihanlh/messenger-api/config"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/internal/utils"
	"gitlab.com/raihanlh/messenger-api/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Error("Failed to hash password: ", zap.Error(err))
		return nil, err
	}
	result, err := u.repositories.User.Create(ctx, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
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

func (u UserUsecase) GetByToken(ctx context.Context, req *payload.GetByTokenRequest) (*payload.GetByTokenResponse, error) {
	log := logger.GetLogger(ctx)
	conf := config.New()

	tokenStr := req.Token
	claims := &utils.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Secret), nil
	})
	if err != nil {
		log.Error("Failed to parse jwt token: ", zap.Error(err))
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Unauthorized")
	}

	user, err := u.repositories.User.GetById(ctx, claims.UserID)
	if err != nil {
		log.Error("Failed to get user by id: ", zap.Error(err))
		return nil, err
	}
	log.Info(fmt.Sprintf("%+v", user))

	return &payload.GetByTokenResponse{
		User:    user,
		Message: "Successfully get user by token",
	}, nil
}

func (u UserUsecase) Login(ctx context.Context, req *payload.LoginRequest) (*payload.LoginResponse, error) {
	log := logger.GetLogger(ctx)
	log.Info(fmt.Sprintf("%+v", req))

	user, err := u.repositories.User.GetByEmail(ctx, req.Email)
	if err != nil {
		log.Error("Failed to log in: ", zap.Error(err))
		return nil, err
	}
	log.Info(fmt.Sprintf("%+v", user))

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error("Failed to log in: ", zap.Error(err))
		return nil, err
	}

	exp := time.Now().Add(time.Hour * 72)
	token, err := utils.GenerateToken(req.Email, user.ID, exp)
	if err != nil {
		log.Error("Failed to generate token: ", zap.Error(err))
		return nil, err
	}
	return &payload.LoginResponse{
		Token:   token,
		Message: "Login success",
		Exp:     exp,
	}, nil
}
