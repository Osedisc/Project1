package createUser

import (
	"github.com/Osedisc/Project1/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) *service {
	return &service{userRepo: userRepo}
}

func (s *service) Execute(ctx *gin.Context, req Request) (*Response, error) {
	id := uuid.New().String()
	_, err := s.userRepo.CreateUser(repo.CreateRequest{
		Id:          id,
		Name:        req.Name,
		SurName:     req.SurName,
		DateOfBirth: req.DateOfBirth,
	}, ctx)
	if err != nil {
		return nil, err
	}
	return &Response{
		Id:          id,
		Name:        req.Name,
		SurName:     req.SurName,
		DateOfBirth: req.DateOfBirth,
	}, nil
}
