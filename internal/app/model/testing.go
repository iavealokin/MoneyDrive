package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
	Email: "user@user.org",
	Password: "password",
}
}