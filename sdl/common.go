package sdl

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

const (
	SdlDll            = "SDL2.dll"
	DllFuncNamePrefix = "SDL_"
)

type FuncMapping struct {
	Fn           *uintptr
	FunctionName string
}

var funcs = []FuncMapping{
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

// LoadLibrary loads the DLL library and its functions.
func LoadLibrary(dllFile string) (err error) {
	if len(dllFile) == 0 {
		dllFile = SdlDll
	}

	fmt.Println(fmt.Sprintf("Loading library: %v.", dllFile))
	sdlDll, err = windows.LoadLibrary(dllFile)
	if err != nil {
		return err
	}

	fmt.Printf("Loading functions: ")
	for _, fm := range funcs {
		fmt.Printf("[%s] ", fm.FunctionName)
		*(fm.Fn), err = windows.GetProcAddress(sdlDll, DllFuncNamePrefix+fm.FunctionName)
		if err != nil {
			return err
		}
	}
	fmt.Println()

	return nil
}

// mustBeNoCallError checks for a call error.
// If call error is set, it finds a name of the function which started this
// function and panics about that function's failure.
func mustBeNoCallError(callErr error) {
	if callErr != nil {
		caller, _, _, _ := runtime.Caller(1) // Skip 1 function up in stack: mustBeNoCallError -> caller.
		fn := runtime.FuncForPC(caller)
		panic(fmt.Sprintf("%s syscall failed: %v", fn.Name(), callErr.Error()))
	}
}

// checkError checks error when necessary.
func checkError(ret uintptr) (err error) {
	if int(ret) == 0 {
		return nil
	}

	return GetError()
}
