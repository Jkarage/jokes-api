package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkarage/joker-api/controllers"
)

func SetJokesRouter(r *gin.Engine) {
	joke := new(controllers.Joke)
	r.GET("/", joke.Home)
}
