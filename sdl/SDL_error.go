package sdl

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/vault-thirteen/SDLW/dll"
	"golang.org/x/sys/windows"
)

// ClearError
// void SDL_ClearError(void);
// https://wiki.libsdl.org/SDL2/SDL_ClearError
func ClearError() {
	_, _, callErr := syscall.SyscallN(fnClearError)
	dll.MustBeNoCallError(callErr)
}

// GetError
// const char* SDL_GetError(void);
// https://wiki.libsdl.org/SDL2/SDL_GetError
func GetError() (err error) {
	cpErrText, _, callErr := syscall.SyscallN(fnGetError)
	dll.MustBeNoCallError(callErr)

	errText := windows.BytePtrToString((*byte)(unsafe.Pointer(cpErrText)))
	if len(errText) == 0 {
		return nil
	}

	return errors.New(errText)
}

// SetError
// int SDL_SetError(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
// https://wiki.libsdl.org/SDL2/SDL_SetError
func SetError(errFormat string, args ...uintptr) (err error) {
	var cpErrFormat *byte
	cpErrFormat, err = windows.BytePtrFromString(errFormat)
	if err != nil {
		return err
	}

	// Golang wants all the args to be in a single array. Let it be so.
	argz := make([]uintptr, 0, len(args)+1)
	argz = append(argz, uintptr(unsafe.Pointer(cpErrFormat)))
	for _, arg := range args {
		argz = append(argz, arg)
	}

	_, _, callErr := syscall.SyscallN(fnSetError, argz...)
	dll.MustBeNoCallError(callErr)

	return nil
}
