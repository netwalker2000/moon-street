package dao

import "moon-street/internal/model"

type UserDataAccessXormImpl struct {
}

func (s *UserDataAccessXormImpl) GetById(id int64) (model.User, error) {
	return model.User{}, nil
}

func (s *UserDataAccessXormImpl) GetByName(name string) (model.User, error) {
	return model.User{}, nil
}

func (s *UserDataAccessXormImpl) Save(user model.User) (int64, error) {
	return 0, nil
}
