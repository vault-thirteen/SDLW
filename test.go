package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vault-thirteen/SDLW/sdl"
)

func main() {
	var err error
	err = work()
	if err != nil {
		log.Fatal(err)
	}
}

func work() (err error) {
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	fmt.Println("Test")
	time.Sleep(time.Second)

	return nil
}