package main

import (
	"github.com/jkarage/jokes-api/log"
	"github.com/jkarage/jokes-api/routes"
)

func main() {
	log.SetLogOutput()
	r := routes.InitRoutes()
	r.Run()
}
