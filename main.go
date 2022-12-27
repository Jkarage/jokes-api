package main

import (
	"fmt"

	"github.com/jkarage/jokes-api/utils"
)

func main() {
	rjoke := utils.RandomJoke()
	fmt.Println(rjoke)
}
