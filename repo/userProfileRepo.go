package repo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userProfileRepo struct {
	db *gorm.DB
}

func NewUserProfileRepo(db *gorm.DB) UserRepo {
	return &userProfileRepo{
		db: db,
	}
}

type UserRepo interface {
	CreateUser(req CreateRequest, ctx *gin.Context) (*CreateResponse, error)
	GetUserById(req GetUserByIdRequest, ctx *gin.Context) (*GetUserByIdResponse, error)
	UpdateUser(req UpdateUserRequest, ctx *gin.Context) (*UpdateUserResponse, error)
	DeleteUser(id string, ctx *gin.Context) error
}

func (u *userProfileRepo) CreateUser(req CreateRequest, ctx *gin.Context) (*CreateResponse, error) {
	tx := u.db.Table(TableName()).Create(req)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &CreateResponse{}, nil
}

func (u *userProfileRepo) GetUserById(req GetUserByIdRequest, ctx *gin.Context) (*GetUserByIdResponse, error) {
	var res *GetUserByIdResponse
	tx := u.db.Table(TableName()).Where("id = ?", req.Id).First(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (u *userProfileRepo) UpdateUser(req UpdateUserRequest, ctx *gin.Context) (*UpdateUserResponse, error) {
	var res *UpdateUserResponse
	tx := u.db.Table(TableName()).Where("id = ?", req.Id).Updates(&req)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (u *userProfileRepo) DeleteUser(id string, ctx *gin.Context) error {
	var model UpdateUserRequest
	model.Id = id
	tx := u.db.Table(TableName()).Delete(&model)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
