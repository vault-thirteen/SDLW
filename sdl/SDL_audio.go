package sdl

import (
	"syscall"
	"unsafe"

	bt "github.com/vault-thirteen/TIFFer/models/basic-types"
	"golang.org/x/sys/windows"
)

// SDL_audio.h.

// typedef Uint16 SDL_AudioFormat;
type AudioFormat uint16

// Audio flags.
const (
	AUDIO_MASK_BITSIZE  = 0xFF
	AUDIO_MASK_DATATYPE = 1 << 8
	AUDIO_MASK_ENDIAN   = 1 << 12
	AUDIO_MASK_SIGNED   = 1 << 15
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

// AudioCallback
// typedef void (SDLCALL * SDL_AudioCallback) (void *userdata, Uint8 * stream, int len);
type AudioCallback func(userdata uintptr, stream uintptr, len uintptr)

// AudioSpec
/*
typedef struct SDL_AudioSpec
{
	int freq;
	SDL_AudioFormat format;
	Uint8 channels;
	Uint8 silence;
	Uint16 samples;
	Uint16 padding;
	Uint32 size;
	SDL_AudioCallback callback;
	void *userdata;
} SDL_AudioSpec;
*/
type AudioSpec struct {
	Freq     int
	Format   AudioFormat
	Channels uint8
	Silence  uint8
	Samples  uint16
	Padding  uint16
	Size     uint32
	Callback AudioCallback
	Userdata uintptr
}

// AudioFilter
// typedef void (SDLCALL * SDL_AudioFilter) (struct SDL_AudioCVT * cvt, SDL_AudioFormat format);
type AudioFilter func(cvt uintptr, format uintptr)

const AUDIOCVT_MAX_FILTERS = 9

// AudioCVT
/*
typedef struct SDL_AudioCVT
{
	int needed;
	SDL_AudioFormat src_format;
	SDL_AudioFormat dst_format;
	double rate_incr;
	Uint8 *buf;
	int len;
	int len_cvt;
	int len_mult;
	double len_ratio;
	SDL_AudioFilter filters[SDL_AUDIOCVT_MAX_FILTERS + 1];
	int filter_index;
} SDL_AUDIOCVT_PACKED SDL_AudioCVT;
*/
type AudioCVT struct {
	Needed      int
	SrcFormat   AudioFormat
	DstFormat   AudioFormat
	RateIncr    bt.Double
	Buf         *byte
	Len         int
	LenCvt      int
	LenMult     int
	LenRatio    bt.Double
	Filters     [AUDIOCVT_MAX_FILTERS + 1]AudioFilter
	FilterIndex int
}

// GetNumAudioDrivers
// int SDL_GetNumAudioDrivers(void);
// https://wiki.libsdl.org/SDL2/SDL_GetNumAudioDrivers
func GetNumAudioDrivers() int {
	ret, _, _ := syscall.SyscallN(fnGetNumAudioDrivers)
	return int(ret)
}

// GetAudioDriver
// const char* SDL_GetAudioDriver(int index);
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDriver
func GetAudioDriver(index int) string {
	ret, _, _ := syscall.SyscallN(fnGetAudioDriver, uintptr(index))
	return windows.BytePtrToString((*byte)(unsafe.Pointer(ret)))
}

// AudioInit
// int SDL_AudioInit(const char *driver_name);
// https://wiki.libsdl.org/SDL2/SDL_AudioInit
func AudioInit(driverName string) int {
	var err error
	var cpDriverName *byte
	cpDriverName, err = windows.BytePtrFromString(driverName)
	mustBeNoError(err)

	ret, _, _ := syscall.SyscallN(fnAudioInit, uintptr(unsafe.Pointer(cpDriverName)))
	return int(ret)
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
func OpenAudio(desired uintptr, obtained uintptr) int {
	ret, _, _ := syscall.SyscallN(fnOpenAudio, desired, obtained)
	return int(ret)
}

// typedef Uint32 SDL_AudioDeviceID;
type AudioDeviceID uint32

// GetNumAudioDevices
// int SDL_GetNumAudioDevices(int iscapture);
// https://wiki.libsdl.org/SDL2/SDL_GetNumAudioDevices
func GetNumAudioDevices(isCapture int) int {
	ret, _, _ := syscall.SyscallN(fnGetNumAudioDevices, uintptr(isCapture))
	return int(ret)
}

// GetAudioDeviceName
/*
const char* SDL_GetAudioDeviceName(int index,
                                   int iscapture);
*/
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDeviceName
func GetAudioDeviceName(index int, isCapture int) string {
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
func GetAudioDeviceSpec(index int, isCapture int, spec uintptr) int {
	ret, _, _ := syscall.SyscallN(fnGetAudioDeviceSpec, uintptr(index), uintptr(isCapture), spec)
	return int(ret)
}

// GetDefaultAudioInfo
/*
int SDL_GetDefaultAudioInfo(char **name,
                            SDL_AudioSpec *spec,
                            int iscapture);
*/
// https://wiki.libsdl.org/SDL2/SDL_GetDefaultAudioInfo
func GetDefaultAudioInfo(name uintptr, spec uintptr, isCapture int) int {
	ret, _, _ := syscall.SyscallN(fnGetDefaultAudioInfo, name, spec, uintptr(isCapture))
	return int(ret)
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
func OpenAudioDevice(device string, isCapture int, desired uintptr, obtained uintptr, allowedChanges int) AudioDeviceID {
	var err error
	var cpDevice *byte
	cpDevice, err = windows.BytePtrFromString(device)
	mustBeNoError(err)

	ret, _, _ := syscall.SyscallN(fnOpenAudioDevice, uintptr(unsafe.Pointer(cpDevice)), uintptr(isCapture), desired, obtained, uintptr(allowedChanges))
	return AudioDeviceID(ret)
}

type AudioStatus int

const (
	AUDIO_STOPPED = AudioStatus(0)
	AUDIO_PLAYING = AudioStatus(1)
	AUDIO_PAUSED  = AudioStatus(2)
)

// GetAudioStatus
// SDL_AudioStatus SDL_GetAudioStatus(void);
// https://wiki.libsdl.org/SDL2/SDL_GetAudioStatus
func GetAudioStatus() AudioStatus {
	ret, _, _ := syscall.SyscallN(fnGetAudioStatus)
	return AudioStatus(ret)
}

// GetAudioDeviceStatus
// SDL_AudioStatus SDL_GetAudioDeviceStatus(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_GetAudioDeviceStatus
func GetAudioDeviceStatus(dev uintptr) AudioStatus {
	ret, _, _ := syscall.SyscallN(fnGetAudioDeviceStatus, dev)
	return AudioStatus(ret)
}

// PauseAudio
// void SDL_PauseAudio(int pause_on);
// https://wiki.libsdl.org/SDL2/SDL_PauseAudio
func PauseAudio(pauseOn int) {
	_, _, _ = syscall.SyscallN(fnPauseAudio, uintptr(pauseOn))
}

// PauseAudioDevice
/*
void SDL_PauseAudioDevice(SDL_AudioDeviceID dev,
                          int pause_on);
*/
// https://wiki.libsdl.org/SDL2/SDL_PauseAudioDevice
func PauseAudioDevice(dev AudioDeviceID, pauseOn int) {
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
func LoadWAV_RW(src uintptr, freeSrc int, spec uintptr, audioBuf uintptr, audioLen uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(fnLoadWAV_RW, src, uintptr(freeSrc), spec, audioBuf, audioLen)
	return ret
}

// FreeWAV
// void SDL_FreeWAV(Uint8 * audio_buf);
// https://wiki.libsdl.org/SDL2/SDL_FreeWAV
func FreeWAV(audioBuf uintptr) {
	_, _, _ = syscall.SyscallN(fnFreeWAV, audioBuf)
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
func BuildAudioCVT(cvt uintptr, srcFormat AudioFormat, srcChannels uint8, srcRate int, dstFormat AudioFormat, dstChannels uint8, dstRate int) int {
	ret, _, _ := syscall.SyscallN(fnBuildAudioCVT, cvt, uintptr(srcFormat), uintptr(srcChannels), uintptr(srcRate), uintptr(dstFormat), uintptr(dstChannels), uintptr(dstRate))
	return int(ret)
}

// ConvertAudio
// int SDL_ConvertAudio(SDL_AudioCVT * cvt);
// https://wiki.libsdl.org/SDL2/SDL_ConvertAudio
func ConvertAudio(cvt uintptr) int {
	ret, _, _ := syscall.SyscallN(fnConvertAudio, cvt)
	return int(ret)
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
func NewAudioStream(srcFormat AudioFormat, srcChannels uint8, srcRate int, dstFormat AudioFormat, dstChannels uint8, dstRate int) uintptr {
	ret, _, _ := syscall.SyscallN(fnNewAudioStream, uintptr(srcFormat), uintptr(srcChannels), uintptr(srcRate), uintptr(dstFormat), uintptr(dstChannels), uintptr(dstRate))
	return ret
}

// AudioStreamPut
// int SDL_AudioStreamPut(SDL_AudioStream *stream, const void *buf, int len);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamPut
func AudioStreamPut(stream uintptr, buf uintptr, len int) int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamPut, stream, buf, uintptr(len))
	return int(ret)
}

// AudioStreamGet
// int SDL_AudioStreamGet(SDL_AudioStream *stream, void *buf, int len);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamGet
func AudioStreamGet(stream uintptr, buf uintptr, len int) int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamGet, stream, buf, uintptr(len))
	return int(ret)
}

// AudioStreamAvailable
// int SDL_AudioStreamAvailable(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamAvailable
func AudioStreamAvailable(stream uintptr) int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamAvailable, stream)
	return int(ret)
}

// AudioStreamFlush
// int SDL_AudioStreamFlush(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamFlush
func AudioStreamFlush(stream uintptr) int {
	ret, _, _ := syscall.SyscallN(fnAudioStreamFlush, stream)
	return int(ret)
}

// AudioStreamClear
// void SDL_AudioStreamClear(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_AudioStreamClear
func AudioStreamClear(stream uintptr) {
	_, _, _ = syscall.SyscallN(fnAudioStreamClear, stream)
}

// FreeAudioStream
// void SDL_FreeAudioStream(SDL_AudioStream *stream);
// https://wiki.libsdl.org/SDL2/SDL_FreeAudioStream
func FreeAudioStream(stream uintptr) {
	_, _, _ = syscall.SyscallN(fnFreeAudioStream, stream)
}

const MIX_MAXVOLUME = 128

// MixAudio
/*
void SDL_MixAudio(Uint8 * dst, const Uint8 * src,
                  Uint32 len, int volume);
*/
// https://wiki.libsdl.org/SDL2/SDL_MixAudio
func MixAudio(dst uintptr, src uintptr, len uint32, volume int) {
	_, _, _ = syscall.SyscallN(fnMixAudio, dst, src, uintptr(len), uintptr(volume))
}

// MixAudioFormat
/*
void SDL_MixAudioFormat(Uint8 * dst,
                        const Uint8 * src,
                        SDL_AudioFormat format,
                        Uint32 len, int volume);
*/
// https://wiki.libsdl.org/SDL2/SDL_MixAudioFormat
func MixAudioFormat(dst uintptr, src uintptr, format AudioFormat, len uint32, volume int) {
	_, _, _ = syscall.SyscallN(fnMixAudioFormat, dst, src, uintptr(format), uintptr(len), uintptr(volume))
}

// QueueAudio
// int SDL_QueueAudio(SDL_AudioDeviceID dev, const void *data, Uint32 len);
// https://wiki.libsdl.org/SDL2/SDL_QueueAudio
func QueueAudio(dev AudioDeviceID, data uintptr, len uint32) int {
	ret, _, _ := syscall.SyscallN(fnQueueAudio, uintptr(dev), data, uintptr(len))
	return int(ret)
}

// DequeueAudio
// Uint32 SDL_DequeueAudio(SDL_AudioDeviceID dev, void *data, Uint32 len);
// https://wiki.libsdl.org/SDL2/SDL_DequeueAudio
func DequeueAudio(dev AudioDeviceID, data uintptr, len uint32) uint32 {
	ret, _, _ := syscall.SyscallN(fnDequeueAudio, uintptr(dev), data, uintptr(len))
	return uint32(ret)
}

// GetQueuedAudioSize
// Uint32 SDL_GetQueuedAudioSize(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_GetQueuedAudioSize
func GetQueuedAudioSize(dev AudioDeviceID) uint32 {
	ret, _, _ := syscall.SyscallN(fnGetQueuedAudioSize, uintptr(dev))
	return uint32(ret)
}

// ClearQueuedAudio
// void SDL_ClearQueuedAudio(SDL_AudioDeviceID dev);
// https://wiki.libsdl.org/SDL2/SDL_ClearQueuedAudio
func ClearQueuedAudio(dev AudioDeviceID) {
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
func LockAudioDevice(dev AudioDeviceID) {
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
func UnlockAudioDevice(dev AudioDeviceID) {
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
func CloseAudioDevice(dev AudioDeviceID) {
	_, _, _ = syscall.SyscallN(fnCloseAudioDevice, uintptr(dev))
}
