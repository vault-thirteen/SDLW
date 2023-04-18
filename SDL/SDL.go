package sdl

// SDL.h.

import (
	"syscall"

	m "github.com/vault-thirteen/SDLW/SDL/model"
)

const (
	INIT_TIMER          m.Uint32 = 0x00000001
	INIT_AUDIO          m.Uint32 = 0x00000010
	INIT_VIDEO          m.Uint32 = 0x00000020
	INIT_JOYSTICK       m.Uint32 = 0x00000200
	INIT_HAPTIC         m.Uint32 = 0x00001000
	INIT_GAMECONTROLLER m.Uint32 = 0x00002000
	INIT_EVENTS         m.Uint32 = 0x00004000
	INIT_SENSOR         m.Uint32 = 0x00008000
	INIT_NOPARACHUTE    m.Uint32 = 0x00100000
	INIT_EVERYTHING     m.Uint32 = INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_EVENTS |
		INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER | INIT_SENSOR
)

// Init
// int SDL_Init(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_Init
func Init(flags m.Uint32) error {
	ret, _, _ := syscall.SyscallN(fnInit, uintptr(flags))
	return checkError(ret)
}

// InitSubSystem
// int SDL_InitSubSystem(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_InitSubSystem
func InitSubSystem(flags m.Uint32) error {
	ret, _, _ := syscall.SyscallN(fnInitSubSystem, uintptr(flags))
	return checkError(ret)
}

// QuitSubSystem
// void SDL_QuitSubSystem(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_QuitSubSystem
func QuitSubSystem(flags m.Uint32) {
	_, _, _ = syscall.SyscallN(fnQuitSubSystem, uintptr(flags))
}

// WasInit
// Uint32 SDL_WasInit(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_WasInit
func WasInit(flags m.Uint32) (status m.Uint32) {
	ret, _, _ := syscall.SyscallN(fnWasInit, uintptr(flags))
	return m.Uint32(ret)
}

// Quit
// void SDL_Quit(void);
// https://wiki.libsdl.org/SDL2/SDL_Quit
func Quit() {
	_, _, _ = syscall.SyscallN(fnQuit)
}
