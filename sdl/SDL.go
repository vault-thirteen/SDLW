package sdl

import (
	"syscall"
)

const (
	INIT_TIMER          uint32 = 0x00000001
	INIT_AUDIO          uint32 = 0x00000010
	INIT_VIDEO          uint32 = 0x00000020
	INIT_JOYSTICK       uint32 = 0x00000200
	INIT_HAPTIC         uint32 = 0x00001000
	INIT_GAMECONTROLLER uint32 = 0x00002000
	INIT_EVENTS         uint32 = 0x00004000
	INIT_SENSOR         uint32 = 0x00008000
	INIT_NOPARACHUTE    uint32 = 0x00100000
	INIT_EVERYTHING     uint32 = INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_EVENTS |
		INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER | INIT_SENSOR
)

// SetMainReady
// void SDL_SetMainReady(void);
// https://wiki.libsdl.org/SDL2/SDL_SetMainReady
func SetMainReady() {
	_, _, _ = syscall.SyscallN(fnSetMainReady)
}

// Init
// int SDL_Init(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_Init
func Init(flags uint32) error {
	ret, _, _ := syscall.SyscallN(fnInit, uintptr(flags))
	return checkError(ret)
}

// InitSubSystem
// int SDL_InitSubSystem(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_InitSubSystem
func InitSubSystem(flags uint32) error {
	ret, _, _ := syscall.SyscallN(fnInitSubSystem, uintptr(flags))
	return checkError(ret)
}

// QuitSubSystem
// void SDL_QuitSubSystem(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_QuitSubSystem
func QuitSubSystem(flags uint32) {
	_, _, _ = syscall.SyscallN(fnQuitSubSystem, uintptr(flags))
}

// WasInit
// Uint32 SDL_WasInit(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_WasInit
func WasInit(flags uint32) (status uint32) {
	ret, _, _ := syscall.SyscallN(fnWasInit, uintptr(flags))
	return uint32(ret)
}

// Quit
// void SDL_Quit(void);
// https://wiki.libsdl.org/SDL2/SDL_Quit
func Quit() {
	_, _, _ = syscall.SyscallN(fnQuit)
}
