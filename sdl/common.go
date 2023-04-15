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
	{&fnLogSetAllPriority, "LogSetAllPriority"},
	{&fnLogSetPriority, "LogSetPriority"},
	{&fnLogGetPriority, "LogGetPriority"},
	{&fnLogResetPriorities, "LogResetPriorities"},
	{&fnLog, "Log"},
	{&fnLogVerbose, "LogVerbose"},
	{&fnLogDebug, "LogDebug"},
	{&fnLogInfo, "LogInfo"},
	{&fnLogWarn, "LogWarn"},
	{&fnLogError, "LogError"},
	{&fnLogCritical, "LogCritical"},
	{&fnLogMessage, "LogMessage"},
	{&fnLogMessageV, "LogMessageV"},
	{&fnLogGetOutputFunction, "LogGetOutputFunction"},
	{&fnLogSetOutputFunction, "LogSetOutputFunction"},
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
	fnLogSetAllPriority    uintptr
	fnLogSetPriority       uintptr
	fnLogGetPriority       uintptr
	fnLogResetPriorities   uintptr
	fnLog                  uintptr
	fnLogVerbose           uintptr
	fnLogDebug             uintptr
	fnLogInfo              uintptr
	fnLogWarn              uintptr
	fnLogError             uintptr
	fnLogCritical          uintptr
	fnLogMessage           uintptr
	fnLogMessageV          uintptr
	fnLogGetOutputFunction uintptr
	fnLogSetOutputFunction uintptr
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

// BytePtrFromStringP converts a Go string into a C string.
// If something goes wrong, it panics.
func BytePtrFromStringP(s string) (cpS *byte) {
	var err error
	cpS, err = windows.BytePtrFromString(s)
	mustBeNoError(err)
	return cpS
}
