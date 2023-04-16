package sdl

// SDL_log.h.

import (
	"syscall"
	"unsafe"

	m "github.com/vault-thirteen/SDLW/SDL/model"
)

// LogSetAllPriority
// void SDL_LogSetAllPriority(SDL_LogPriority priority);
// https://wiki.libsdl.org/SDL2/SDL_LogSetAllPriority
func LogSetAllPriority(priority m.LogPriority) {
	_, _, _ = syscall.SyscallN(fnLogSetAllPriority, uintptr(priority))
}

// LogSetPriority
/*
void SDL_LogSetPriority(int category,
                        SDL_LogPriority priority);
*/
// https://wiki.libsdl.org/SDL2/SDL_LogSetPriority
func LogSetPriority(category int, priority m.LogPriority) {
	_, _, _ = syscall.SyscallN(fnLogSetPriority, uintptr(category), uintptr(priority))
}

// LogGetPriority
// SDL_LogPriority SDL_LogGetPriority(int category);
// https://wiki.libsdl.org/SDL2/SDL_LogGetPriority
func LogGetPriority(category int) m.LogPriority {
	ret, _, _ := syscall.SyscallN(fnLogGetPriority, uintptr(category))
	return m.LogPriority(ret)
}

// LogResetPriorities
// void SDL_LogResetPriorities(void);
// https://wiki.libsdl.org/SDL2/SDL_LogResetPriorities
func LogResetPriorities() {
	_, _, _ = syscall.SyscallN(fnLogResetPriorities)
}

// Log
// void SDL_Log(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
// https://wiki.libsdl.org/SDL2/SDL_Log
func Log(fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+1)
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLog, argz...)
}

// LogVerbose
// void SDL_LogVerbose(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
// https://wiki.libsdl.org/SDL2/SDL_LogVerbose
func LogVerbose(category int, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+2)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogVerbose, argz...)
}

// LogDebug
// void SDL_LogDebug(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
// https://wiki.libsdl.org/SDL2/SDL_LogDebug
func LogDebug(category int, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+2)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogDebug, argz...)
}

// LogInfo
// void SDL_LogInfo(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
// https://wiki.libsdl.org/SDL2/SDL_LogInfo
func LogInfo(category int, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+2)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogInfo, argz...)
}

// LogWarn
// void SDL_LogWarn(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
// https://wiki.libsdl.org/SDL2/SDL_LogWarn
func LogWarn(category int, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+2)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogWarn, argz...)
}

// LogError
// void SDL_LogError(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
// https://wiki.libsdl.org/SDL2/SDL_LogError
func LogError(category int, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+2)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogError, argz...)
}

// LogCritical
// void SDL_LogCritical(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
// https://wiki.libsdl.org/SDL2/SDL_LogCritical
func LogCritical(category int, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+2)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogCritical, argz...)
}

// LogMessage
/*
void SDL_LogMessage(int category,
                    SDL_LogPriority priority,
                    SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(3);
*/
// https://wiki.libsdl.org/SDL2/SDL_LogMessage
func LogMessage(category int, priority m.LogPriority, fmt string, args ...uintptr) {
	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+3)
	argz = append(argz, uintptr(category))
	argz = append(argz, uintptr(priority))
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnLogMessage, argz...)
}

// LogMessageV
/*
void SDL_LogMessageV(int category,
                     SDL_LogPriority priority,
                     const char *fmt, va_list ap);
*/
// https://wiki.libsdl.org/SDL2/SDL_LogMessageV
func LogMessageV(category int, priority m.LogPriority, fmt string, ap string) {
	_, _, _ = syscall.SyscallN(fnLogMessageV, uintptr(category), uintptr(priority), uintptr(unsafe.Pointer(BytePtrFromStringP(fmt))), uintptr(unsafe.Pointer(BytePtrFromStringP(ap))))
}

// LogGetOutputFunction
// void SDL_LogGetOutputFunction(SDL_LogOutputFunction *callback, void **userdata);
// https://wiki.libsdl.org/SDL2/SDL_LogGetOutputFunction
// TODO: Test this when callbacks are fixed in Golang.
func LogGetOutputFunction(callback uintptr, userdata uintptr) {
	_, _, _ = syscall.SyscallN(fnLogGetOutputFunction, callback, userdata)
}

// LogSetOutputFunction
// void SDL_LogSetOutputFunction(SDL_LogOutputFunction callback, void *userdata);
// https://wiki.libsdl.org/SDL2/SDL_LogSetOutputFunction
// TODO: Test this when callbacks are fixed in Golang.
func LogSetOutputFunction(callback uintptr, userdata uintptr) {
	_, _, _ = syscall.SyscallN(fnLogSetOutputFunction, callback, userdata)
}
