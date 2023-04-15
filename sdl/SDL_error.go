package sdl

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/vault-thirteen/SDLW/dll"
	"golang.org/x/sys/windows"
)

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
