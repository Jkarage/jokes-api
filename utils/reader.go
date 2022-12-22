package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadPdf() {
	f, err := os.Open("/assets/jokes.pdf")
	if err != nil {
		panic(err.Error())
	}
	input := bufio.NewReader(f)
	read, err := input.ReadBytes(byte(38))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(read)
}
