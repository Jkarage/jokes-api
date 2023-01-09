package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkarage/jokes-api/utils"
)

type Joke struct{}

func (Joke) Home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to this jokes api,For more documentaion please visit the /docs endpoint.")
}

func (Joke) Joke(c *gin.Context) {
	joke, author := utils.RandomJoke()
	c.JSON(http.StatusOK, gin.H{
		"joke":   joke,
		"author": author,
	})
}
