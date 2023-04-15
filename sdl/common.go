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
	{&fnGetError, "GetError"},
	{&fnSetMainReady, "SetMainReady"},
	{&fnInit, "Init"},
	{&fnInitSubSystem, "InitSubSystem"},
	{&fnQuitSubSystem, "QuitSubSystem"},
	{&fnWasInit, "WasInit"},
	{&fnQuit, "Quit"},
}

var (
	sdlDll          windows.Handle
	fnGetError      uintptr
	fnSetMainReady  uintptr
	fnInit          uintptr
	fnInitSubSystem uintptr
	fnQuitSubSystem uintptr
	fnWasInit       uintptr
	fnQuit          uintptr
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
