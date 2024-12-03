package getUserById

import (
	"encoding/json"
	"fmt"
	"time"

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
	var res *Response
	err := s.GetByKey(ctx, key, &res)
	if err != nil {
		res, err := s.userRepo.GetUserById(repo.GetUserByIdRequest{
			Id: req.Id,
		}, ctx)
		if err != nil {
			return nil, err
		}
		_ = s.SetByKey(ctx, key, res, time.Hour)
		return &Response{
			Id:          res.Id,
			Name:        res.Name,
			SurName:     res.SurName,
			DateOfBirth: res.DateOfBirth,
		}, nil
	}
	return res, nil
}

func (s *service) GetByKey(ctx *gin.Context, key string, output any) error {
	data, err := s.cache.Get(ctx, key).Bytes()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(data, output)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *service) SetByKey(ctx *gin.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = s.cache.Set(ctx, key, data, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
