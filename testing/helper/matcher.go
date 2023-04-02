package helper

import (
	"fmt"
	"log"
	"reflect"

	"github.com/golang/mock/gomock"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// bcrypt matcher
type bcryptMatcher struct {
	password []byte
}

func MatchBCrypt(password []byte) gomock.Matcher {
	return &bcryptMatcher{password}
}

func (m *bcryptMatcher) Matches(x interface{}) bool {
	if hpw, ok := x.([]byte); ok {
		return bcrypt.CompareHashAndPassword(hpw, m.password) == nil
	}
	return false
}

func (m *bcryptMatcher) String() string {
	log.Println(string(m.password))
	return fmt.Sprintf("matches to %s with bcrypt", m.password)
}

// create user matcher
type createUsertMatcher struct {
	Name     string
	Email    string
	Password []byte
}

func MatchCreateUser(user model.User) gomock.Matcher {
	return &createUsertMatcher{
		Name:     user.Name,
		Email:    user.Email,
		Password: []byte(user.Password),
	}
}

func (m *createUsertMatcher) Matches(x interface{}) bool {
	reflectedValue := reflect.ValueOf(x).Elem()

	user := x.(*model.User)
	hpw := []byte(user.Password)

	if m.Name != reflectedValue.FieldByName("Name").String() {
		return false
	} else if m.Email != reflectedValue.FieldByName("Email").String() {
		return false
	}
	return bcrypt.CompareHashAndPassword(hpw, m.Password) == nil
}

func (m *createUsertMatcher) String() string {
	user := &model.User{
		Name:     m.Name,
		Email:    m.Email,
		Password: string(m.Password),
	}
	return fmt.Sprintf("is equal to %v (%s)", user, reflect.TypeOf(user))
}
