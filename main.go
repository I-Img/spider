package main

import (
	"fmt"
	"os"
)

var key string

func main() {

	initClient()
	if err := getRecentPic(); err != nil {
		panic(err)
	}
}

func initClient() {
	key = os.Getenv("I_IMG_FLICKER_KEY")
	if key == "" {
		panic(fmt.Errorf("I_IMG_FLICKER_KEY invalid"))
	}
}
