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
	{&fnSetMainReady, "SetMainReady"},
	{&fnInit, "Init"},
	{&fnInitSubSystem, "InitSubSystem"},
	{&fnQuitSubSystem, "QuitSubSystem"},
	{&fnWasInit, "WasInit"},
	{&fnQuit, "Quit"},

	// Error Handling.
	{&fnClearError, "ClearError"},
	{&fnGetError, "GetError"},
	{&fnSetError, "SetError"},
}

var (
	sdlDll windows.Handle

	// Initialization and Shutdown.
	fnSetMainReady  uintptr
	fnInit          uintptr
	fnInitSubSystem uintptr
	fnQuitSubSystem uintptr
	fnWasInit       uintptr
	fnQuit          uintptr

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
