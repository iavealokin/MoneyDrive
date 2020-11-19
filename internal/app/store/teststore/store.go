package teststore

import (
	"github.com/iavealokin/MoneyDrive/internal/app/model"
	"github.com/iavealokin/MoneyDrive/internal/app/store"
)

//Store ..
type Store struct {
	UserRepository *UserRepository
}

func New() *Store {

}

func (s *Store) User() store.UserRepository{
	if s.UserRepository!= nil{
		return s.UserRepository
	}

	s.UserRepository=&UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
		return s.UserRepository
}