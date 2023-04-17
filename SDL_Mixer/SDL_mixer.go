package sdlm

// SDL_mixer.h.

import (
	"syscall"
	"unsafe"

	m "github.com/vault-thirteen/SDLW/SDL/model"
	mm "github.com/vault-thirteen/SDLW/SDL_Mixer/model"
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
func Init(flags mm.InitFlags) mm.InitFlags {
	ret, _, _ := syscall.SyscallN(fnInit, uintptr(flags))
	return mm.InitFlags(ret)
}

//extern DECLSPEC void SDLCALL Mix_Quit(void);
func Quit() {
	_, _, _ = syscall.SyscallN(fnQuit)
}

//extern DECLSPEC int SDLCALL Mix_OpenAudio(int frequency, Uint16 format, int channels, int chunksize);
func OpenAudio(frequency int, format uint16, channels int, chunkSize int) int {
	ret, _, _ := syscall.SyscallN(fnOpenAudio, uintptr(frequency), uintptr(format), uintptr(channels), uintptr(chunkSize))
	return int(ret)
}

//extern DECLSPEC int SDLCALL Mix_OpenAudioDevice(int frequency, Uint16 format, int channels, int chunksize, const char* device, int allowed_changes);
func OpenAudioDevice(frequency int, format uint16, channels int, chunkSize int, device string, allowedChanges int) int {
	ret, _, _ := syscall.SyscallN(fnOpenAudioDevice, uintptr(frequency), uintptr(format), uintptr(channels), uintptr(chunkSize), uintptr(unsafe.Pointer(BytePtrFromStringP(device))), uintptr(allowedChanges))
	return int(ret)
}

//TODO

//extern DECLSPEC void SDLCALL Mix_CloseAudio(void);
func CloseAudio() {
	_, _, _ = syscall.SyscallN(fnCloseAudio)
}
