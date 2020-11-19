package store

import "github.com/iavealokin/MoneyDrive/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}