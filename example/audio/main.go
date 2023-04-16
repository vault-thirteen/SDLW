package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vault-thirteen/SDLW/sdl"
	"github.com/vault-thirteen/SDLW/win32"
)

func main() {
	var err error
	err = win32.LoadLibrary()
	mustBeNoError(err)
	err = sdl.LoadLibrary(sdl.SdlDll)
	mustBeNoError(err)

	err = work()
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func work() (err error) {
	// Initialization and status check.
	err = sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	status := sdl.WasInit(0)
	fmt.Println(fmt.Sprintf("Initialization status: %v.", status))

	// Audio.
	//sdl.LoadWAV_RW() //TODO

	time.Sleep(time.Second * 1)

	return nil
}
