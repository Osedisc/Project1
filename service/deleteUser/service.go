package deleteUser

import (
	"github.com/Osedisc/Project1/repo"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type service struct {
	userRepo repo.UserRepo
	cache    *redis.Client
}

func NewService(userRepo repo.UserRepo, cache *redis.Client) *service {
	return &service{
		userRepo: userRepo,
		cache:    cache}
}

func (s *service) Execute(ctx *gin.Context, req Request) (*Response, error) {
	key := "user::" + req.Id
	err := s.userRepo.DeleteUser(req.Id, ctx)
	if err != nil {
		return nil, err
	}
	_ = s.cache.Del(ctx, key)
	return &Response{}, nil
}
