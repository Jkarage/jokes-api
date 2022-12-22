package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Joke struct{}

func (Joke) Home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to this jokes api,For more documentaion please visit the /docs endpoint.")
}
