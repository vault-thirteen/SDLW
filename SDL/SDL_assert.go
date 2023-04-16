package sdl

// SDL_assert.h.

import "syscall"

// SetAssertionHandler
/*
void SDL_SetAssertionHandler(
                    SDL_AssertionHandler handler,
                    void *userdata);
*/
// https://wiki.libsdl.org/SDL2/SDL_SetAssertionHandler
func SetAssertionHandler(handler uintptr, userdata uintptr) {
	_, _, _ = syscall.SyscallN(fnSetAssertionHandler, handler, userdata)
}

// GetDefaultAssertionHandler
// SDL_AssertionHandler SDL_GetDefaultAssertionHandler(void);
// https://wiki.libsdl.org/SDL2/SDL_GetDefaultAssertionHandler
func GetDefaultAssertionHandler() uintptr {
	ret, _, _ := syscall.SyscallN(fnGetDefaultAssertionHandler)
	return ret
}

// GetAssertionHandler
// SDL_AssertionHandler SDL_GetAssertionHandler(void **puserdata);
// https://wiki.libsdl.org/SDL2/SDL_GetAssertionHandler
func GetAssertionHandler(puserdata uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(fnGetAssertionHandler, puserdata)
	return ret
}

// GetAssertionReport
// const SDL_AssertData * SDL_GetAssertionReport(void);
// https://wiki.libsdl.org/SDL2/SDL_GetAssertionReport
func GetAssertionReport() uintptr {
	ret, _, _ := syscall.SyscallN(fnGetAssertionReport)
	return ret
}

// ResetAssertionReport
// void SDL_ResetAssertionReport(void);
// https://wiki.libsdl.org/SDL2/SDL_ResetAssertionReport
func ResetAssertionReport() {
	_, _, _ = syscall.SyscallN(fnResetAssertionReport)
}
