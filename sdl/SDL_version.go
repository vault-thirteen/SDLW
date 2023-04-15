package sdl

// SDL_version.h.

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Version
/*
typedef struct SDL_version
{
	Uint8 major;
	Uint8 minor;
	Uint8 patch;
} SDL_version;
*/
// SDL_version.h
type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

// GetVersion
// void SDL_GetVersion(SDL_version * ver);
// https://wiki.libsdl.org/SDL2/SDL_GetVersion
func GetVersion() *Version {
	var ver = new(Version)

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
