package sdlm

// SDL_mixer.h.

import (
	"syscall"
	"unsafe"

	m "github.com/vault-thirteen/SDLW/SDL/model"
)

const (
	CHANNELS          = 8
	DEFAULT_FREQUENCY = 44100
	DEFAULT_CHANNELS  = 2
)

//extern DECLSPEC const SDL_version * SDLCALL Mix_Linked_Version(void);
func Linked_Version() *m.Version {
	ret, _, _ := syscall.SyscallN(fnLinked_Version)
	return (*m.Version)(unsafe.Pointer(ret))
}

// Init
// extern DECLSPEC int SDLCALL Mix_Init(int flags);
// "It returns the flags successfully initialized, or 0 on failure."
// https://wiki.libsdl.org/SDL3_mixer/Mix_Init
func Init(flags m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnInit, uintptr(flags))
	return m.Int(ret)
}

//extern DECLSPEC void SDLCALL Mix_Quit(void);
func Quit() {
	_, _, _ = syscall.SyscallN(fnQuit)
}

//extern DECLSPEC int SDLCALL Mix_OpenAudio(int frequency, Uint16 format, int channels, int chunksize);
func OpenAudio(frequency m.Int, format m.Uint16, channels m.Int, chunkSize m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnOpenAudio, uintptr(frequency), uintptr(format), uintptr(channels), uintptr(chunkSize))
	return m.Int(ret)
}

//extern DECLSPEC int SDLCALL Mix_OpenAudioDevice(int frequency, Uint16 format, int channels, int chunksize, const char* device, int allowed_changes);
func OpenAudioDevice(frequency m.Int, format m.Uint16, channels m.Int, chunkSize m.Int, device string, allowedChanges m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnOpenAudioDevice, uintptr(frequency), uintptr(format), uintptr(channels), uintptr(chunkSize), uintptr(unsafe.Pointer(BytePtrFromStringP(device))), uintptr(allowedChanges))
	return m.Int(ret)
}

//TODO

//extern DECLSPEC void SDLCALL Mix_CloseAudio(void);
func CloseAudio() {
	_, _, _ = syscall.SyscallN(fnCloseAudio)
}
