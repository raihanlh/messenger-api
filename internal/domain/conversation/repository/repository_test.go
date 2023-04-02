package repository_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	repo "gitlab.com/raihanlh/messenger-api/internal/domain/conversation/repository"
	"gitlab.com/raihanlh/messenger-api/internal/model"
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

	query := `INSERT INTO "conversations" ("created_at","updated_at","deleted_at","id") VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING RETURNING "id"`
	queryUser := `INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email","password","id") VALUES ($1,$2,$3,$4,$5,$6,$7),($8,$9,$10,$11,$12,$13,$14) ON CONFLICT DO NOTHING RETURNING "id"`
	queryParticipant := `INSERT INTO "user_participants" ("conversation_id","user_id") VALUES ($1,$2),($3,$4) ON CONFLICT DO NOTHING RETURNING "conversation_id","user_id`

	id := "6fd33930-d76e-401a-a3ba-7a03352812c2"
	column := []string{"id"}
	columnUser := []string{"id"}
	columnConv := []string{"conversation_id", "user_id"}

	sender := &model.User{
		Model: model.Model{
			ID: "34251esd-d76e-401a-a3ba-7a03352812c2",
		},
		Name:     "Test Sender",
		Email:    "sender@example.com",
		Password: "password",
	}
	receiver := &model.User{
		Model: model.Model{
			ID: "47dsga9t-d76e-401a-a3ba-7a03352812c2",
		},
		Name:     "Test Receiver",
		Email:    "receiver@example.com",
		Password: "password",
	}
	conversation := model.Conversation{
		Model: model.Model{
			ID: id,
		},
		Users: []*model.User{sender, receiver},
	}
	rows := []driver.Value{id}
	rowUser1 := []driver.Value{sender.ID}
	rowUser2 := []driver.Value{receiver.ID}
	rowConv1 := []driver.Value{id, sender.ID}
	rowConv2 := []driver.Value{id, receiver.ID}

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		rowsMock     *sqlmock.Rows
		rowsUserMock *sqlmock.Rows
		rowsConvMock *sqlmock.Rows
		want         *model.Conversation
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name: "Create Conversation Success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
			},
			rowsMock:     sqlmock.NewRows(column).AddRow(rows...),
			rowsUserMock: sqlmock.NewRows(columnUser).AddRow(rowUser1...).AddRow(rowUser2...),
			rowsConvMock: sqlmock.NewRows(columnConv).AddRow(rowConv1...).AddRow(rowConv2...),
			want:         &conversation,
			wantErr:      assert.NoError,
		},
		{
			name: "Create Conversation Failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.TODO(),
			},
			rowsMock: sqlmock.NewRows(column).AddRow(rows...).RowError(0, errors.New("test error")),
			want:     &conversation,
			wantErr:  assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := conversation.Users

			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(query)).
				WithArgs(AnyTime{}, AnyTime{}, nil, sqlmock.AnyArg()).
				WillReturnRows(tt.rowsMock)
			mock.ExpectQuery(regexp.QuoteMeta(queryUser)).
				WithArgs(AnyTime{}, AnyTime{}, nil, users[0].Name, users[0].Email, users[0].Password, sqlmock.AnyArg(),
					AnyTime{}, AnyTime{}, nil, users[1].Name, users[1].Email, users[1].Password, sqlmock.AnyArg()).
				WillReturnRows(tt.rowsMock)
			mock.ExpectQuery(regexp.QuoteMeta(queryParticipant)).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnRows(tt.rowsMock)

			mock.ExpectCommit()

			r := &repo.ConversationRepository{
				DB: tt.fields.DB,
			}

			res, err := r.Create(tt.args.ctx, &conversation)
			tt.want.ID = res.ID
			tt.wantErr(t, err)
			assert.Equalf(t, tt.want, res, "Error on: Create(%v, %v)", tt.args.ctx, &conversation)
		})
	}
}
