package main

import (
	"github.com/Osedisc/Project1/app"
	"github.com/Osedisc/Project1/infras"
)

func main() {
	infras.InitConfig()
	infras.InitRedis()
	infras.InitDB()
	app.Run()
}
