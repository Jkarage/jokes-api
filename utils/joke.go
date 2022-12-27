package utils

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"time"
)

type Jokes struct {
	Jokes []Joke `json:"jokes"`
}

type Joke struct {
	Text     string   `json:"text"`
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Tags     []string `json:"tags"`
}

var jokes Jokes

func init() {
	jsonFile, err := os.Open("assets/jokes.json")
	CheckNilErr(err)

	// fmt.Println("Successfully Opened jokes.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &jokes)
	CheckNilErr(err)
}

func RandomJoke() string {
	max := len(jokes.Jokes)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	index := r.Intn(max)
	return jokes.Jokes[index].Text
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
