package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vault-thirteen/SDLW/audio"
	"github.com/vault-thirteen/SDLW/sdl"
	m "github.com/vault-thirteen/SDLW/sdl/model"
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
	var audioInfo *audio.Info
	audioInfo, err = audio.GetInfo()
	if err != nil {
		return err
	}
	device := audioInfo.PlaybackDevices[0]
	fmt.Println(fmt.Sprintf("Using device: %s.", device.Name))

	// Load a WAV file.
	filePath := `C:\Windows\Media\chimes.wav`
	var wavSpec m.AudioSpec
	var wavBuffer *uint8
	var wavLength uint32
	_ = sdl.LoadWAV(filePath, &wavSpec, &wavBuffer, &wavLength)
	defer func() {
		sdl.FreeWAV(wavBuffer)
	}()

	// Open the device.
	var s2 m.AudioSpec
	deviceId := sdl.OpenAudioDevice(device.Name, 0, &wavSpec, &s2, 0)
	if deviceId == 0 {
		return sdl.GetError()
	}
	defer sdl.CloseAudioDevice(deviceId)
	fmt.Println("Device ID:", deviceId)

	// Play audio.
	ret := sdl.QueueAudio(deviceId, wavBuffer, wavLength)
	if ret != 0 {
		return sdl.GetError()
	}
	sdl.PauseAudioDevice(deviceId, 0)

	time.Sleep(time.Second * 3) //TODO:How to get length of audio ?

	return nil
}
