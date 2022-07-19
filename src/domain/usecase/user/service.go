package user

import (
	"clean-go/src/infra/repos"
)

type Service struct {
	 repo repos.UserRepository
}

func NewService(r repos.UserRepository) *Service {
	return &Service{
		 repo: r,
	}
}



func (s *Service) GetUser() string {
	
	return s.repo.GetUser()
}
