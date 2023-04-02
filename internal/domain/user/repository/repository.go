package repository

import (
	"context"
	"strings"

	"gitlab.com/raihanlh/messenger-api/internal/constant"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB *gorm.DB
}

func New(gormDB *gorm.DB) user.Repository {
	return &UserRepository{
		DB: gormDB,
	}
}

func (r UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	result := r.DB.WithContext(ctx).Model(user).Clauses(clause.OnConflict{DoNothing: true}).Create(user)
	return user, result.Error
}

func (r UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	result := r.DB.WithContext(ctx).Model(user).Where("id = ?", user.ID).Updates(user)
	return user, result.Error
}

func (r UserRepository) Delete(ctx context.Context, id string) error {
	var user *model.User
	result := r.DB.WithContext(ctx).Table(constant.UserTable).Where("id = ?", id).Delete(&user)
	return result.Error
}

func (r UserRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	result := r.DB.WithContext(ctx).Table(constant.UserTable).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (r UserRepository) GetAll(ctx context.Context, pgn *pagination.Pagination, req *payload.GetAllRequest) ([]*model.User, error) {
	var users []*model.User
	result := r.DB.WithContext(ctx).Table(constant.UserTable)

	if req.Search != "" {
		escapedSearchTerm := strings.Replace(req.Search, "%", "\\%", -1)
		result = result.Where("name ILIKE ?", "%"+escapedSearchTerm+"%")
	}

	result.Scopes(pgn.Paginate(result)).Find(&users)
	return users, result.Error
}
