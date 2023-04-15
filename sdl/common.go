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
	{&fnSetHintWithPriority, "SetHintWithPriority"},
	{&fnSetHint, "SetHint"},
	{&fnResetHint, "ResetHint"},
	{&fnResetHints, "ResetHints"},
	{&fnGetHint, "GetHint"},
	{&fnGetHintBoolean, "GetHintBoolean"},
	{&fnAddHintCallback, "AddHintCallback"},
	{&fnDelHintCallback, "DelHintCallback"},
	{&fnClearHints, "ClearHints"},

	// Error Handling.
	{&fnClearError, "ClearError"},
	{&fnGetError, "GetError"},
	{&fnSetError, "SetError"},

	// Log Handling.
	//TODO
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
	fnSetHintWithPriority uintptr
	fnSetHint             uintptr
	fnResetHint           uintptr
	fnResetHints          uintptr
	fnGetHint             uintptr
	fnGetHintBoolean      uintptr
	fnAddHintCallback     uintptr
	fnDelHintCallback     uintptr
	fnClearHints          uintptr

	// Error Handling.
	fnClearError uintptr
	fnGetError   uintptr
	fnSetError   uintptr

	// Log Handling.
	//TODO
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

// mustBeNoError panics if an error occurs.
func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}
