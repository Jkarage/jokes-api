package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkarage/jokes-api/controllers"
)

func SetJokesRouter(r *gin.Engine) {
	joke := new(controllers.Joke)
	r.GET("/", joke.Home)
	r.GET("/joke", joke.Joke)
}

func SetNaughtyJokesRouter(r *gin.Engine) {

}
