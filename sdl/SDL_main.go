package sdl

// SDL_main.h.

import "syscall"

// SetMainReady
// void SDL_SetMainReady(void);
// https://wiki.libsdl.org/SDL2/SDL_SetMainReady
func SetMainReady() {
	_, _, _ = syscall.SyscallN(fnSetMainReady)
}

// RegisterApp
// Most applications do not need to, and should not, call this directly; SDL
// will call it when initializing the video subsystem.

// UnregisterApp
// Most applications do not need to, and should not, call this directly; SDL
// will call it when deinitializing the video subsystem.
