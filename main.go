package main

import "github.com/jkarage/jokes-api/routes"

func main() {
	r := routes.InitRoutes()
	r.Run()
}
