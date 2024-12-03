package app

import (
	"github.com/Osedisc/Project1/infras"
	"github.com/Osedisc/Project1/repo"
	"github.com/Osedisc/Project1/service/createUser"
	"github.com/Osedisc/Project1/service/deleteUser"
	"github.com/Osedisc/Project1/service/getUserById"
	"github.com/Osedisc/Project1/service/updateUser"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	bindCreateUser(r)
	bindGeteUserById(r)
	bindUpdateUser(r)
	bindDeleteUser(r)
	r.Run()
}

func bindCreateUser(r *gin.Engine) gin.IRoutes {
	db := infras.DB
	repo := repo.NewUserProfileRepo(db)
	servicer := createUser.NewService(repo)
	handler := createUser.NewHandler(servicer)
	return r.POST("/User", handler.CreateUser)
}

func bindGeteUserById(r *gin.Engine) gin.IRoutes {
	db := infras.DB
	cache := infras.Rdb
	repo := repo.NewUserProfileRepo(db)
	servicer := getUserById.NewService(repo, cache)
	handler := getUserById.NewHandler(servicer)
	return r.GET("/User/:id", handler.GetUserById)
}

func bindUpdateUser(r *gin.Engine) gin.IRoutes {
	db := infras.DB
	cache := infras.Rdb
	repo := repo.NewUserProfileRepo(db)
	servicer := updateUser.NewService(repo, cache)
	handler := updateUser.NewHandler(servicer)
	return r.PATCH("/User/:id", handler.UpdateUser)
}

func bindDeleteUser(r *gin.Engine) gin.IRoutes {
	db := infras.DB
	cache := infras.Rdb
	repo := repo.NewUserProfileRepo(db)
	servicer := deleteUser.NewService(repo, cache)
	handler := deleteUser.NewHandler(servicer)
	return r.DELETE("/User/:id", handler.DeleteUser)
}
