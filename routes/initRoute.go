package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkarage/jokes-api/middlewares"
)

func InitRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger())
	SetJokesRouter(r)
	return r
}

func CheckNil(err error) {
	if err != nil {
		panic(err.Error())
	}
}
