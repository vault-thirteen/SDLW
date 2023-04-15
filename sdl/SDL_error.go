package sdl

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ClearError
// void SDL_ClearError(void);
// https://wiki.libsdl.org/SDL2/SDL_ClearError
func ClearError() {
	_, _, _ = syscall.SyscallN(fnClearError)
}

// GetError
// const char* SDL_GetError(void);
// https://wiki.libsdl.org/SDL2/SDL_GetError
func GetError() (err error) {
	cpErrText, _, _ := syscall.SyscallN(fnGetError)

	errText := windows.BytePtrToString((*byte)(unsafe.Pointer(cpErrText)))
	if len(errText) == 0 {
		return nil
	}

	return errors.New(errText)
}

// SetError
// int SDL_SetError(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
// https://wiki.libsdl.org/SDL2/SDL_SetError
func SetError(errFormat string, args ...uintptr) {
	// While this function always returns -1, the return is useless.

	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+1)
	argz = append(argz, uintptr(unsafe.Pointer(BytePtrFromStringP(errFormat))))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, _ = syscall.SyscallN(fnSetError, argz...)
}
