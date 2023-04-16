package audio

import m "github.com/vault-thirteen/SDLW/sdl/model"

type Device struct {
	Index int
	Name  string
	Spec  *m.AudioSpec
}
