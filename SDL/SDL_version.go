package sdl

// SDL_version.h.

import (
	"syscall"
	"unsafe"

	m "github.com/vault-thirteen/SDLW/sdl/model"
	"golang.org/x/sys/windows"
)

// GetVersion
// void SDL_GetVersion(SDL_version * ver);
// https://wiki.libsdl.org/SDL2/SDL_GetVersion
func GetVersion() *m.Version {
	var ver = new(m.Version)

	_, _, _ = syscall.SyscallN(fnGetVersion, uintptr(unsafe.Pointer(ver)))
	return ver
}

// GetRevision
// const char* SDL_GetRevision(void);
// https://wiki.libsdl.org/SDL2/SDL_GetRevision
func GetRevision() string {
	ret, _, _ := syscall.SyscallN(fnGetRevision)
	return windows.BytePtrToString((*byte)(unsafe.Pointer(ret)))
}

// GetRevisionNumber
// int SDL_GetRevisionNumber(void);
// https://wiki.libsdl.org/SDL2/SDL_GetRevisionNumber
func GetRevisionNumber() int {
	ret, _, _ := syscall.SyscallN(fnGetRevisionNumber)
	return int(ret)
}
