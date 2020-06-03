package users

import "github.com/gonzavz/go-doit-api/src/utils/errors"

type Service interface {
	GetByID(string) (User, errors.RestError)
}

type service struct {
}

func NewService() Service {
	return &service
}

func (s *service) GetByID(string id) (User, errors.RestError) {

}
