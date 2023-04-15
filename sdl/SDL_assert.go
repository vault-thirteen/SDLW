package sdl

import "syscall"

type AssertState int

const (
	ASSERTION_RETRY         = AssertState(0)
	ASSERTION_BREAK         = AssertState(1)
	ASSERTION_ABORT         = AssertState(2)
	ASSERTION_IGNORE        = AssertState(3)
	ASSERTION_ALWAYS_IGNORE = AssertState(4)
)

// AssertData
/*
typedef struct SDL_AssertData
{
	int always_ignore;
	unsigned int trigger_count;
	const char *condition;
	const char *filename;
	int linenum;
	const char *function;
	const struct SDL_AssertData *next;
} SDL_AssertData;
*/
// SDL_assert.h
type AssertData struct {
	alwaysIgnore int
	triggerCount uint
	condition    *byte
	filename     *byte
	lineNum      int
	function     *byte
	next         *AssertData
}

// AssertionHandler
// SDL_AssertState (SDLCALL *SDL_AssertionHandler)(const SDL_AssertData* data, void* userdata);
// SDL_assert.h
type AssertionHandler func(data uintptr, userdata uintptr)

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
