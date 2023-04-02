package repository_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
	repo "gitlab.com/raihanlh/messenger-api/internal/domain/user/repository"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/pagination"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Setup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		PreferSimpleProtocol: true,
		Conn:                 db,
		DriverName:           "postgres",
	})

	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	return gormDB, mock
}

func Test_UserRepository_Create(t *testing.T) {
	db, mock := Setup()

	query := `INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email","password","id") VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT DO NOTHING RETURNING "id"`

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	column := []string{"id"}
	rows := []driver.Value{id}
	reqUser := model.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		rowsMock *sqlmock.Rows
		want     *model.User
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "Create User Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...),
			want:     &reqUser,
			wantErr:  assert.NoError,
		},
		{
			name: "Create User Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			want:     &reqUser,
			wantErr:  assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(query)).
				WithArgs(AnyTime{}, AnyTime{}, nil, reqUser.Name, reqUser.Email, reqUser.Password, sqlmock.AnyArg()).
				WillReturnRows(tt.rowsMock)
			mock.ExpectCommit()

			r := &repo.UserRepository{
				DB: tt.fields.DB,
			}

			res, err := r.Create(tt.args.ctx, &reqUser)
			tt.want.ID = res.ID
			tt.wantErr(t, err)
			assert.Equalf(t, tt.want, res, "Error on: Create(%v, %v)", tt.args.ctx, &reqUser)
		})
	}
}

func Test_UserRepository_Update(t *testing.T) {
	db, mock := Setup()

	query := `UPDATE "users" SET "updated_at"=$1,"name"=$2,"email"=$3,"password"=$4 WHERE id = $5 AND "users"."deleted_at" IS NULL AND "id" = $6`

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	column := []string{"id"}
	rows := []driver.Value{id}
	reqUser := model.User{
		Model: model.Model{
			ID: id,
		},
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		rowsMock    *sqlmock.Rows
		want        *model.User
		wantErr     assert.ErrorAssertionFunc
		wantErrExec error
	}{
		{
			name: "Update User Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
			},
			rowsMock:    sqlmock.NewRows(column).AddRow(rows...),
			want:        &reqUser,
			wantErr:     assert.NoError,
			wantErrExec: nil,
		},
		{
			name: "Update User Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
			},
			rowsMock:    sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			want:        &reqUser,
			wantErr:     assert.Error,
			wantErrExec: errors.New("test error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(AnyTime{}, reqUser.Name, reqUser.Email, reqUser.Password, reqUser.ID, reqUser.ID).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.wantErrExec)
			mock.ExpectCommit()

			r := &repo.UserRepository{
				DB: tt.fields.DB,
			}

			res, err := r.Update(tt.args.ctx, &reqUser)
			tt.wantErr(t, err)
			assert.Equalf(t, tt.want, res, "Error on: Update(%v, %v)", tt.args.ctx, &reqUser)
		})
	}
}

func Test_UserRepository_Delete(t *testing.T) {
	db, mock := Setup()

	query := `UPDATE "users" SET "deleted_at"=$1 WHERE id = $2 AND "users"."deleted_at" IS NULL`

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	column := []string{"id"}
	rows := []driver.Value{id}
	reqUser := model.User{
		Model: model.Model{
			ID: id,
		},
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		rowsMock    *sqlmock.Rows
		want        *model.User
		wantErr     assert.ErrorAssertionFunc
		wantErrExec error
	}{
		{
			name: "Delete User Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
				id:  id,
			},
			rowsMock:    sqlmock.NewRows(column).AddRow(rows...),
			want:        &reqUser,
			wantErr:     assert.NoError,
			wantErrExec: nil,
		},
		{
			name: "Delete User Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
				id:  id,
			},
			rowsMock:    sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			want:        &reqUser,
			wantErr:     assert.Error,
			wantErrExec: errors.New("test error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(AnyTime{}, tt.args.id).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.wantErrExec)
			mock.ExpectCommit()

			r := &repo.UserRepository{
				DB: tt.fields.DB,
			}

			err := r.Delete(tt.args.ctx, tt.args.id)
			tt.wantErr(t, err)
		})
	}
}

func Test_UserRepository_GetById(t *testing.T) {
	db, mock := Setup()

	query := `SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`
	column := []string{"id", "createdAt", "updatedAt", "deletedAt", "name", "email", "password"}

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	user := model.User{
		Model: model.Model{
			ID: id,
		},
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	rows := []driver.Value{id, nil, nil, nil, user.Name, user.Email, user.Password}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		rowsMock *sqlmock.Rows
		want     *model.User
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "Get By Id User Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
				id:  id,
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...),
			want:     &user,
			wantErr:  assert.NoError,
		},
		{
			name: "Get By Id User Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
				id:  id,
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			want:     &model.User{},
			wantErr:  assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(query)).
				WithArgs(id).
				WillReturnRows(tt.rowsMock)

			r := &repo.UserRepository{
				DB: tt.fields.DB,
			}

			res, err := r.GetById(tt.args.ctx, tt.args.id)

			tt.wantErr(t, err)
			assert.Equalf(t, tt.want, res, "Error on: GetById(%v, %+v)", tt.args.ctx, tt.args.id)
		})
	}
}

func Test_UserRepository_GetByEmail(t *testing.T) {
	db, mock := Setup()

	query := `SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`
	column := []string{"id", "createdAt", "updatedAt", "deletedAt", "name", "email", "password"}

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	user := model.User{
		Model: model.Model{
			ID: id,
		},
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	rows := []driver.Value{id, nil, nil, nil, user.Name, user.Email, user.Password}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx   context.Context
		email string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		rowsMock *sqlmock.Rows
		want     *model.User
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "Get By Id User Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:   context.TODO(),
				email: user.Email,
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...),
			want:     &user,
			wantErr:  assert.NoError,
		},
		{
			name: "Get By Id User Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:   context.TODO(),
				email: user.Email,
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			want:     &model.User{},
			wantErr:  assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(query)).
				WithArgs(tt.args.email).
				WillReturnRows(tt.rowsMock)

			r := &repo.UserRepository{
				DB: tt.fields.DB,
			}

			res, err := r.GetByEmail(tt.args.ctx, tt.args.email)

			tt.wantErr(t, err)
			assert.Equalf(t, tt.want, res, "Error on: GetByEmail(%v, %+v)", tt.args.ctx, tt.args.email)
		})
	}
}

func Test_UserRepository_GetAll(t *testing.T) {
	db, mock := Setup()

	queryCount := `SELECT count(*) FROM "users"`
	query := ``
	column := []string{"id", "createdAt", "updatedAt", "deletedAt", "name", "email", "password"}
	columnCount := []string{"count"}

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	user := model.User{
		Model: model.Model{
			ID: id,
		},
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	rows := []driver.Value{id, nil, nil, nil, user.Name, user.Email, user.Password}
	rowsCount := []driver.Value{1}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		req payload.GetAllRequest
	}

	tests := []struct {
		name          string
		fields        fields
		args          args
		rowsMock      *sqlmock.Rows
		rowsMockCount *sqlmock.Rows
		want          []*model.User
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "Get All User Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
				req: payload.GetAllRequest{
					Pagination: pagination.Pagination{
						Page:         1,
						ItemsPerPage: 10,
					},
				},
			},
			rowsMock:      sqlmock.NewRows(column).AddRow(rows...),
			rowsMockCount: sqlmock.NewRows(columnCount).AddRow(rowsCount...),
			want:          []*model.User{&user},
			wantErr:       assert.NoError,
		},
		{
			name: "Get All User Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
				req: payload.GetAllRequest{
					Pagination: pagination.Pagination{
						Page:         1,
						ItemsPerPage: 10,
					},
				},
			},
			rowsMock:      sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			rowsMockCount: sqlmock.NewRows(columnCount).AddRow(rowsCount...),
			want:          []*model.User{},
			wantErr:       assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(queryCount)).WillReturnRows(tt.rowsMockCount)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(tt.rowsMock)
		})

		r := &repo.UserRepository{
			DB: tt.fields.DB,
		}

		res, err := r.GetAll(tt.args.ctx, &tt.args.req.Pagination, &tt.args.req)

		tt.wantErr(t, err)
		assert.Equalf(t, tt.want, res, "Error on: GetAll(%v, %+v)", tt.args.ctx, tt.args.req)
	}
}
