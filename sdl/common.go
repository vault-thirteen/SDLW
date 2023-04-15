package sdl

import (
	"github.com/vault-thirteen/SDLW/dll"
	"golang.org/x/sys/windows"
)

const (
	SdlDll            = "SDL2.dll"
	DllFuncNamePrefix = "SDL_"
)

var funcs = []dll.FuncMapping{
	// Initialization and Shutdown.
	{&fnInit, "Init"},
	{&fnInitSubSystem, "InitSubSystem"},
	{&fnQuit, "Quit"},
	{&fnQuitSubSystem, "QuitSubSystem"},
	{&fnSetMainReady, "SetMainReady"},
	{&fnWasInit, "WasInit"},

	// Configuration Variables.
	// TODO

	// Error Handling.
	{&fnClearError, "ClearError"},
	{&fnGetError, "GetError"},
	{&fnSetError, "SetError"},
}

var (
	sdlDll windows.Handle

	// Initialization and Shutdown.
	fnInit          uintptr
	fnInitSubSystem uintptr
	fnQuit          uintptr
	fnQuitSubSystem uintptr
	fnSetMainReady  uintptr
	fnWasInit       uintptr

	// Configuration Variables.
	// TODO
	// https://wiki.libsdl.org/SDL2/SDL_HintPriority
	// https://wiki.libsdl.org/SDL2/SDL_SetHintWithPriority
	// https://wiki.libsdl.org/SDL2/SDL_GetHint
	// https://wiki.libsdl.org/SDL2/SDL_SetHint

	// Error Handling.
	fnClearError uintptr
	fnGetError   uintptr
	fnSetError   uintptr
)

// LoadLibrary loads the library and its functions.
func LoadLibrary(dllFile string) (err error) {
	err = dll.LoadLibrary(dllFile, &sdlDll, funcs, DllFuncNamePrefix)
	if err != nil {
		return err
	}

	return nil
}

// checkError checks error when necessary.
func checkError(ret uintptr) (err error) {
	if int(ret) == 0 {
		return nil
	}

	return GetError()
}
