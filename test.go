package main

import (
	"errors"
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
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	status := sdl.WasInit(0)
	fmt.Println(fmt.Sprintf("Initialization status: %v.", status))

	// Set, get and clear an error.
	err = sdl.SetError("Test: %d.", 123)
	if err != nil {
		return err
	}
	manualErr := sdl.GetError()
	//fmt.Println(manualErr)
	if manualErr.Error() != "Test: 123." {
		return errors.New("errors do not work")
	}
	sdl.ClearError()
	manualErr = sdl.GetError()
	//fmt.Println(manualErr)
	if manualErr != nil {
		return errors.New("errors do not work")
	}

	time.Sleep(time.Second * 1)

	return nil
}
