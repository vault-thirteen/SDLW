package sdl

/*
#include "include/SDL.h"
*/
import "C"

const (
	INIT_TIMER          = C.SDL_INIT_TIMER
	INIT_AUDIO          = C.SDL_INIT_AUDIO
	INIT_VIDEO          = C.SDL_INIT_VIDEO
	INIT_JOYSTICK       = C.SDL_INIT_JOYSTICK
	INIT_HAPTIC         = C.SDL_INIT_HAPTIC
	INIT_GAMECONTROLLER = C.SDL_INIT_GAMECONTROLLER
	INIT_EVENTS         = C.SDL_INIT_EVENTS
	INIT_SENSOR         = C.SDL_INIT_SENSOR
	INIT_NOPARACHUTE    = C.SDL_INIT_NOPARACHUTE
	INIT_EVERYTHING     = C.SDL_INIT_EVERYTHING
)

// SetMainReady
// void SDL_SetMainReady(void);
// https://wiki.libsdl.org/SDL2/SDL_SetMainReady
func SetMainReady() {
	C.SDL_SetMainReady()
}

// Init
// int SDL_Init(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_Init
func Init(flags uint32) error {
	return checkError(C.SDL_Init(C.Uint32(flags)))
}

// InitSubSystem
// int SDL_InitSubSystem(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_InitSubSystem
func InitSubSystem(flags uint32) error {
	return checkError(C.SDL_InitSubSystem(C.Uint32(flags)))
}

// QuitSubSystem
// void SDL_QuitSubSystem(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_QuitSubSystem
func QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

// WasInit
// Uint32 SDL_WasInit(Uint32 flags);
// https://wiki.libsdl.org/SDL2/SDL_WasInit
func WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(C.Uint32(flags)))
}

// Quit
// void SDL_Quit(void);
// https://wiki.libsdl.org/SDL2/SDL_Quit
func Quit() {
	C.SDL_Quit()
}
