package sdl

import "C"

func checkError(cmdOutput C.int) (err error) {
	if int(cmdOutput) == 0 {
		return nil
	}
	return GetError()
}
