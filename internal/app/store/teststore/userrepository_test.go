package teststore_test

import (
	"testing"

	"github.com/iavealokin/MoneyDrive/internal/app/model"
	"github.com/iavealokin/MoneyDrive/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)


func TestUserRepository_Create(t *testing.T){
	s:= teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}


func TestUserRepository_FindByEmail (t *testing.T){
	s:= teststore.New()


	email :="user@user.org"
	_, err :=s.User().FindByEmail(email)
	assert.Error(t, err)	

	u := model.TestUser(t)
	u.Email= email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}