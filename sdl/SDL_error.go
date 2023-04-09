package sdl

/*
#include "include/SDL.h"
*/
import "C"
import "errors"

// GetError
// const char* SDL_GetError(void);
// https://wiki.libsdl.org/SDL2/SDL_GetError
func GetError() (err error) {
	e := C.SDL_GetError()
	if err != nil {
		s := C.GoString(e)
		if len(s) > 0 {
			return errors.New(s)
		}
	}

	return nil
}
