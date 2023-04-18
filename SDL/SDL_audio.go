package sdl

// SDL_audio.h.

import (
	"syscall"
	"unsafe"

	"github.com/vault-thirteen/SDLW/SDL/model"
	"golang.org/x/sys/windows"
)

// Audio flags.
const (
	AUDIO_MASK_BITSIZE  = m.Uint16(0xFF)
	AUDIO_MASK_DATATYPE = m.Uint16(1 << 8)
	AUDIO_MASK_ENDIAN   = m.Uint16(1 << 12)
	AUDIO_MASK_SIGNED   = m.Uint16(1 << 15)
)

// Audio format flags.
const (
	AUDIO_U8     = 0x0008
	AUDIO_S8     = 0x8008
	AUDIO_U16LSB = 0x0010
	AUDIO_S16LSB = 0x8010
	AUDIO_U16MSB = 0x1010
	AUDIO_S16MSB = 0x9010
	AUDIO_S32LSB = 0x8020
	AUDIO_S32MSB = 0x9020
	AUDIO_F32LSB = 0x8120
	AUDIO_F32MSB = 0x9120
)

// Which audio format changes are allowed when opening a device.
const (
	AUDIO_ALLOW_FREQUENCY_CHANGE = 0x00000001
	AUDIO_ALLOW_FORMAT_CHANGE    = 0x00000002
	AUDIO_ALLOW_CHANNELS_CHANGE  = 0x00000004
	AUDIO_ALLOW_SAMPLES_CHANGE   = 0x00000008
	AUDIO_ALLOW_ANY_CHANGE       = AUDIO_ALLOW_FREQUENCY_CHANGE | AUDIO_ALLOW_FORMAT_CHANGE | AUDIO_ALLOW_CHANNELS_CHANGE | AUDIO_ALLOW_SAMPLES_CHANGE
)

func AUDIO_BITSIZE(x m.Uint16) m.Uint16 {
	return x & AUDIO_MASK_BITSIZE
}
func AUDIO_ISFLOAT(x m.Uint16) m.Uint16 {
	return x & AUDIO_MASK_DATATYPE
}
func AUDIO_ISBIGENDIAN(x m.Uint16) m.Uint16 {
	return x & AUDIO_MASK_ENDIAN
}
func AUDIO_ISSIGNED(x m.Uint16) m.Uint16 {
	return x & AUDIO_MASK_SIGNED
}

// GetNumAudioDrivers
// int SDL_GetNumAudioDrivers(void);
// https://wiki.libsdl.org/SDL2/SDL_GetNumAudioDrivers
func GetNumAudioDrivers() m.Int {
	ret, _, _ := syscall.SyscallN(fnGetNumAudioDrivers)
	return m.Int(ret)
}

// GetAudioDriver
// const char* SDL_GetAudioDriver(int index);
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDriver
func GetAudioDriver(index m.Int) string {
	ret, _, _ := syscall.SyscallN(fnGetAudioDriver, uintptr(index))
	return windows.BytePtrToString((*byte)(unsafe.Pointer(ret)))
}

// AudioInit
// int SDL_AudioInit(const char *driver_name);
// https://wiki.libsdl.org/SDL2/SDL_AudioInit
func AudioInit(driverName string) m.Int {
	var err error
	var cpDriverName *byte
	cpDriverName, err = windows.BytePtrFromString(driverName)
	mustBeNoError(err)

	ret, _, _ := syscall.SyscallN(fnAudioInit, uintptr(unsafe.Pointer(cpDriverName)))
	return m.Int(ret)
}

// AudioQuit
// void SDL_AudioQuit(void);
// https://wiki.libsdl.org/SDL2/SDL_AudioQuit
func AudioQuit() {
	_, _, _ = syscall.SyscallN(fnAudioQuit)
}

// GetCurrentAudioDriver
// const char* SDL_GetCurrentAudioDriver(void);
// https://wiki.libsdl.org/SDL2/SDL_GetCurrentAudioDriver
func GetCurrentAudioDriver() string {
	ret, _, _ := syscall.SyscallN(fnGetCurrentAudioDriver)
	return windows.BytePtrToString((*byte)(unsafe.Pointer(ret)))
}

// OpenAudio
/*
int SDL_OpenAudio(SDL_AudioSpec * desired,
                  SDL_AudioSpec * obtained);
*/
// https://wiki.libsdl.org/SDL2/SDL_OpenAudio
func OpenAudio(desired *m.AudioSpec, obtained *m.AudioSpec) m.Int {
	ret, _, _ := syscall.SyscallN(fnOpenAudio, uintptr(unsafe.Pointer(desired)), uintptr(unsafe.Pointer(obtained)))
	return m.Int(ret)
}

// GetNumAudioDevices
// int SDL_GetNumAudioDevices(int iscapture);
// https://wiki.libsdl.org/SDL2/SDL_GetNumAudioDevices
func GetNumAudioDevices(isCapture m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnGetNumAudioDevices, uintptr(isCapture))
	return m.Int(ret)
}

// GetAudioDeviceName
/*
const char* SDL_GetAudioDeviceName(int index,
                                   int iscapture);
*/
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDeviceName
func GetAudioDeviceName(index m.Int, isCapture m.Int) string {
	ret, _, _ := syscall.SyscallN(fnGetAudioDeviceName, uintptr(index), uintptr(isCapture))
	return windows.BytePtrToString((*byte)(unsafe.Pointer(ret)))
}

// GetAudioDeviceSpec
/*
int SDL_GetAudioDeviceSpec(int index,
                           int iscapture,
                           SDL_AudioSpec *spec);
*/
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDeviceSpec
func GetAudioDeviceSpec(index m.Int, isCapture m.Int, spec *m.AudioSpec) m.Int {
	ret, _, _ := syscall.SyscallN(fnGetAudioDeviceSpec, uintptr(index), uintptr(isCapture), uintptr(unsafe.Pointer(spec)))
	return m.Int(ret)
}

// GetDefaultAudioInfo
/*
int SDL_GetDefaultAudioInfo(char **name,
                            SDL_AudioSpec *spec,
                            int iscapture);
*/
// https://wiki.libsdl.org/SDL2/SDL_GetDefaultAudioInfo
func GetDefaultAudioInfo(name **m.Char, spec *m.AudioSpec, isCapture m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnGetDefaultAudioInfo, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(spec)), uintptr(isCapture))
	return m.Int(ret)
}

// OpenAudioDevice
/*
SDL_AudioDeviceID SDL_OpenAudioDevice(
                          const char *device,
                          int iscapture,
                          const SDL_AudioSpec *desired,
                          SDL_AudioSpec *obtained,
                          int allowed_changes);
*/
// https://wiki.libsdl.org/SDL2/SDL_OpenAudioDevice
func OpenAudioDevice(device string, isCapture m.Int, desired *m.AudioSpec, obtained *m.AudioSpec, allowedChanges m.Int) m.AudioDeviceID {
	var err error
	var cpDevice *byte
	cpDevice, err = windows.BytePtrFromString(device)
	mustBeNoError(err)

	ret, _, _ := syscall.SyscallN(fnOpenAudioDevice, uintptr(unsafe.Pointer(cpDevice)), uintptr(isCapture), uintptr(unsafe.Pointer(desired)), uintptr(unsafe.Pointer(obtained)), uintptr(allowedChanges))
	return m.AudioDeviceID(ret)
}

// GetAudioStatus
// SDL_AudioStatus SDL_GetAudioStatus(void);
// https://wiki.libsdl.org/SDL2/SDL_GetAudioStatus
func GetAudioStatus() m.AudioStatus {
	ret, _, _ := syscall.SyscallN(fnGetAudioStatus)
	return m.AudioStatus(ret)
}

// GetAudioDeviceStatus
// SDL_AudioStatus SDL_GetAudioDeviceStatus(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDeviceStatus
func GetAudioDeviceStatus(dev m.AudioDeviceID) m.AudioStatus {
	ret, _, _ := syscall.SyscallN(fnGetAudioDeviceStatus, uintptr(dev))
	return m.AudioStatus(ret)
}

// PauseAudio
// void SDL_PauseAudio(int pause_on);
// https://wiki.libsdl.org/SDL2/SDL_PauseAudio
func PauseAudio(pauseOn m.Int) {
	_, _, _ = syscall.SyscallN(fnPauseAudio, uintptr(pauseOn))
}

// PauseAudioDevice
/*
void SDL_PauseAudioDevice(SDL_AudioDeviceID dev,
                          int pause_on);
*/
// https://wiki.libsdl.org/SDL2/SDL_PauseAudioDevice
func PauseAudioDevice(dev m.AudioDeviceID, pauseOn m.Int) {
	_, _, _ = syscall.SyscallN(fnPauseAudioDevice, uintptr(dev), uintptr(pauseOn))
}

// LoadWAV_RW
/*
SDL_AudioSpec* SDL_LoadWAV_RW(SDL_RWops * src,
                              int freesrc,
                              SDL_AudioSpec * spec,
                              Uint8 ** audio_buf,
                              Uint32 * audio_len);
*/
// https://wiki.libsdl.org/SDL2/SDL_LoadWAV_RW
func LoadWAV_RW(src uintptr, freeSrc m.Int, spec *m.AudioSpec, audioBuf **m.Uint8, audioLen *m.Uint32) *m.AudioSpec {
	ret, _, _ := syscall.SyscallN(fnLoadWAV_RW, src, uintptr(freeSrc), uintptr(unsafe.Pointer(spec)), uintptr(unsafe.Pointer(audioBuf)), uintptr(unsafe.Pointer(audioLen)))
	return (*m.AudioSpec)(unsafe.Pointer(ret))
}

// LoadWAV
/*
SDL_AudioSpec* SDL_LoadWAV(const char*    file,
                           SDL_AudioSpec* spec,
                           Uint8**        audio_buf,
                           Uint32*        audio_len)
*/
// https://wiki.libsdl.org/SDL2/SDL_LoadWAV
func LoadWAV(file string, spec *m.AudioSpec, audioBuf **m.Uint8, audioLen *m.Uint32) *m.AudioSpec {
	// #define SDL_LoadWAV(file, spec, audio_buf, audio_len) SDL_LoadWAV_RW(SDL_RWFromFile(file, "rb"),1, spec,audio_buf,audio_len)
	// This is a C macro.
	// This means, it is not directly accessible via the DLL file !

	ops := RWFromFile(file, "rb")
	if ops == 0 {
		panic(ops)
	}

	ret, _, _ := syscall.SyscallN(fnLoadWAV_RW, ops, 0, uintptr(unsafe.Pointer(spec)), uintptr(unsafe.Pointer(audioBuf)), uintptr(unsafe.Pointer(audioLen)))
	return (*m.AudioSpec)(unsafe.Pointer(ret))
}

// FreeWAV
// void SDL_FreeWAV(Uint8 * audio_buf);
// https://wiki.libsdl.org/SDL2/SDL_FreeWAV
func FreeWAV(audioBuf *m.Uint8) {
	_, _, _ = syscall.SyscallN(fnFreeWAV, uintptr(unsafe.Pointer(audioBuf)))
}

// BuildAudioCVT
/*
int SDL_BuildAudioCVT(SDL_AudioCVT * cvt,
                      SDL_AudioFormat src_format,
                      Uint8 src_channels,
                      int src_rate,
                      SDL_AudioFormat dst_format,
                      Uint8 dst_channels,
                      int dst_rate);
*/
// https://wiki.libsdl.org/SDL2/SDL_BuildAudioCVT
func BuildAudioCVT(cvt *m.AudioCVT, srcFormat m.AudioFormat, srcChannels m.Uint8, srcRate m.Int, dstFormat m.AudioFormat, dstChannels m.Uint8, dstRate m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnBuildAudioCVT, uintptr(unsafe.Pointer(cvt)), uintptr(srcFormat), uintptr(srcChannels), uintptr(srcRate), uintptr(dstFormat), uintptr(dstChannels), uintptr(dstRate))
	return m.Int(ret)
}

// ConvertAudio
// int SDL_ConvertAudio(SDL_AudioCVT * cvt);
// https://wiki.libsdl.org/SDL2/SDL_ConvertAudio
func ConvertAudio(cvt *m.AudioCVT) m.Int {
	ret, _, _ := syscall.SyscallN(fnConvertAudio, uintptr(unsafe.Pointer(cvt)))
	return m.Int(ret)
}

// NewAudioStream
/*
SDL_AudioStream * SDL_NewAudioStream(const SDL_AudioFormat src_format,
                   const Uint8 src_channels,
                   const int src_rate,
                   const SDL_AudioFormat dst_format,
                   const Uint8 dst_channels,
                   const int dst_rate);
*/
// https://wiki.libsdl.org/SDL2/SDL_NewAudioStream
func NewAudioStream(srcFormat m.AudioFormat, srcChannels m.Uint8, srcRate m.Int, dstFormat m.AudioFormat, dstChannels m.Uint8, dstRate m.Int) *m.AudioStream {
	ret, _, _ := syscall.SyscallN(fnNewAudioStream, uintptr(srcFormat), uintptr(srcChannels), uintptr(srcRate), uintptr(dstFormat), uintptr(dstChannels), uintptr(dstRate))
	return (*m.AudioStream)(unsafe.Pointer(ret))
}

// AudioStreamPut
// int SDL_AudioStreamPut(SDL_AudioStream *stream, const void *buf, int len);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamPut
func AudioStreamPut(stream *m.AudioStream, buf *m.Void, len m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamPut, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(buf)), uintptr(len))
	return m.Int(ret)
}

// AudioStreamGet
// int SDL_AudioStreamGet(SDL_AudioStream *stream, void *buf, int len);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamGet
func AudioStreamGet(stream *m.AudioStream, buf *m.Void, len m.Int) m.Int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamGet, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(buf)), uintptr(len))
	return m.Int(ret)
}

// AudioStreamAvailable
// int SDL_AudioStreamAvailable(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamAvailable
func AudioStreamAvailable(stream *m.AudioStream) m.Int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamAvailable, uintptr(unsafe.Pointer(stream)))
	return m.Int(ret)
}

// AudioStreamFlush
// int SDL_AudioStreamFlush(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamFlush
func AudioStreamFlush(stream *m.AudioStream) m.Int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamFlush, uintptr(unsafe.Pointer(stream)))
	return m.Int(ret)
}

// AudioStreamClear
// void SDL_AudioStreamClear(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamClear
func AudioStreamClear(stream *m.AudioStream) {
	_, _, _ = syscall.SyscallN(fnAudioStreamClear, uintptr(unsafe.Pointer(stream)))
}

// FreeAudioStream
// void SDL_FreeAudioStream(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_FreeAudioStream
func FreeAudioStream(stream *m.AudioStream) {
	_, _, _ = syscall.SyscallN(fnFreeAudioStream, uintptr(unsafe.Pointer(stream)))
}

const MIX_MAXVOLUME = 128

// MixAudio
/*
void SDL_MixAudio(Uint8 * dst, const Uint8 * src,
                  Uint32 len, int volume);
*/
// https://wiki.libsdl.org/SDL2/SDL_MixAudio
func MixAudio(dst *m.Uint8, src *m.Uint8, len m.Uint32, volume m.Int) {
	_, _, _ = syscall.SyscallN(fnMixAudio, uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(src)), uintptr(len), uintptr(volume))
}

// MixAudioFormat
/*
void SDL_MixAudioFormat(Uint8 * dst,
                        const Uint8 * src,
                        SDL_AudioFormat format,
                        Uint32 len, int volume);
*/
// https://wiki.libsdl.org/SDL2/SDL_MixAudioFormat
func MixAudioFormat(dst *m.Uint8, src *m.Uint8, format m.AudioFormat, len m.Uint32, volume m.Int) {
	_, _, _ = syscall.SyscallN(fnMixAudioFormat, uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(src)), uintptr(format), uintptr(len), uintptr(volume))
}

// QueueAudio
// int SDL_QueueAudio(SDL_AudioDeviceID dev, const void *data, Uint32 len);
// https://wiki.libsdl.org/SDL2/SDL_QueueAudio
func QueueAudio(dev m.AudioDeviceID, data *m.Uint8, len m.Uint32) m.Int {
	ret, _, _ := syscall.SyscallN(fnQueueAudio, uintptr(dev), uintptr(unsafe.Pointer(data)), uintptr(len))
	return m.Int(ret)
}

// DequeueAudio
// Uint32 SDL_DequeueAudio(SDL_AudioDeviceID dev, void *data, Uint32 len);
// https://wiki.libsdl.org/SDL2/SDL_DequeueAudio
func DequeueAudio(dev m.AudioDeviceID, data *m.Void, len m.Uint32) m.Uint32 {
	ret, _, _ := syscall.SyscallN(fnDequeueAudio, uintptr(dev), uintptr(unsafe.Pointer(data)), uintptr(len))
	return m.Uint32(ret)
}

// GetQueuedAudioSize
// Uint32 SDL_GetQueuedAudioSize(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_GetQueuedAudioSize
func GetQueuedAudioSize(dev m.AudioDeviceID) m.Uint32 {
	ret, _, _ := syscall.SyscallN(fnGetQueuedAudioSize, uintptr(dev))
	return m.Uint32(ret)
}

// ClearQueuedAudio
// void SDL_ClearQueuedAudio(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_ClearQueuedAudio
func ClearQueuedAudio(dev m.AudioDeviceID) {
	_, _, _ = syscall.SyscallN(fnClearQueuedAudio, uintptr(dev))
}

// LockAudio
// void SDL_LockAudio(void);
// https://wiki.libsdl.org/SDL2/SDL_LockAudio
func LockAudio() {
	_, _, _ = syscall.SyscallN(fnLockAudio)
}

// LockAudioDevice
// void SDL_LockAudioDevice(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_LockAudioDevice
func LockAudioDevice(dev m.AudioDeviceID) {
	_, _, _ = syscall.SyscallN(fnLockAudioDevice, uintptr(dev))
}

// UnlockAudio
// void SDL_UnlockAudio(void);
// https://wiki.libsdl.org/SDL2/SDL_UnlockAudio
func UnlockAudio() {
	_, _, _ = syscall.SyscallN(fnUnlockAudio)
}

// UnlockAudioDevice
// void SDL_UnlockAudioDevice(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_UnlockAudioDevice
func UnlockAudioDevice(dev m.AudioDeviceID) {
	_, _, _ = syscall.SyscallN(fnUnlockAudioDevice, uintptr(dev))
}

// CloseAudio
// void SDL_CloseAudio(void);
// https://wiki.libsdl.org/SDL2/SDL_CloseAudio
func CloseAudio() {
	_, _, _ = syscall.SyscallN(fnCloseAudio)
}

// CloseAudioDevice
// void SDL_CloseAudioDevice(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_CloseAudioDevice
func CloseAudioDevice(dev m.AudioDeviceID) {
	_, _, _ = syscall.SyscallN(fnCloseAudioDevice, uintptr(dev))
}
