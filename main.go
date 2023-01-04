package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jkarage/jokes-api/routes"
	"github.com/jkarage/jokes-api/utils"
)

func main() {
	SetLogOutput()
	r := routes.InitRoutes()
	r.Run()
}

func SetLogOutput() {
	f, err := os.OpenFile("log/runtime.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	utils.CheckNilErr(err)
	gin.DefaultWriter = f
}
