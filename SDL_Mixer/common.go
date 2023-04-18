package sdlm

import (
	"errors"

	"github.com/vault-thirteen/SDLW/dll"
	"golang.org/x/sys/windows"
)

const (
	SdlMixDll         = "SDL2_mixer.dll"
	DllFuncNamePrefix = "Mix_"
)

const (
	ErrExternalFunctionsNotSet = "external functions are not set"
)

var funcs = []dll.FuncMapping{
	{&fnLinked_Version, "Linked_Version"},
	{&fnInit, "Init"},
	{&fnQuit, "Quit"},
	{&fnOpenAudio, "OpenAudio"},
	{&fnOpenAudioDevice, "OpenAudioDevice"},
	//TODO
	{&fnCloseAudio, "CloseAudio"},
}

var (
	sdlMixDll windows.Handle

	fnLinked_Version         uintptr
	fnInit                   uintptr
	fnQuit                   uintptr
	fnOpenAudio              uintptr
	fnOpenAudioDevice        uintptr
	fnQuerySpec              uintptr
	fnAllocateChannels       uintptr
	fnLoadWAV_RW             uintptr
	fnLoadWAV                uintptr
	fnLoadMUS                uintptr
	fnLoadMUS_RW             uintptr
	fnLoadMUSType_RW         uintptr
	fnQuickLoad_WAV          uintptr
	fnQuickLoad_RAW          uintptr
	fnFreeChunk              uintptr
	fnFreeMusic              uintptr
	fnGetNumChunkDecoders    uintptr
	fnGetChunkDecoder        uintptr
	fnHasChunkDecoder        uintptr
	fnGetNumMusicDecoders    uintptr
	fnGetMusicDecoder        uintptr
	fnHasMusicDecoder        uintptr
	fnGetMusicType           uintptr
	fnGetMusicTitle          uintptr
	fnGetMusicTitleTag       uintptr
	fnGetMusicArtistTag      uintptr
	fnGetMusicAlbumTag       uintptr
	fnGetMusicCopyrightTag   uintptr
	fnGetMusicHookData       uintptr
	fnRegisterEffect         uintptr
	fnUnregisterEffect       uintptr
	fnUnregisterAllEffects   uintptr
	fnSetPanning             uintptr
	fnSetPosition            uintptr
	fnSetDistance            uintptr
	fnSetReverseStereo       uintptr
	fnReserveChannels        uintptr
	fnGroupChannel           uintptr
	fnGroupChannels          uintptr
	fnGroupAvailable         uintptr
	fnGroupCount             uintptr
	fnGroupOldest            uintptr
	fnGroupNewer             uintptr
	fnPlayChannel            uintptr
	fnPlayChannelTimed       uintptr
	fnPlayMusic              uintptr
	fnFadeInMusic            uintptr
	fnFadeInMusicPos         uintptr
	fnFadeInChannel          uintptr
	fnFadeInChannelTimed     uintptr
	fnVolume                 uintptr
	fnVolumeChunk            uintptr
	fnVolumeMusic            uintptr
	fnGetMusicVolume         uintptr
	fnMasterVolume           uintptr
	fnHaltChannel            uintptr
	fnHaltGroup              uintptr
	fnHaltMusic              uintptr
	fnExpireChannel          uintptr
	fnFadeOutChannel         uintptr
	fnFadeOutGroup           uintptr
	fnFadeOutMusic           uintptr
	fnFadingMusic            uintptr
	fnFadingChannel          uintptr
	fnPause                  uintptr
	fnResume                 uintptr
	fnPaused                 uintptr
	fnPauseMusic             uintptr
	fnResumeMusic            uintptr
	fnRewindMusic            uintptr
	fnPausedMusic            uintptr
	fnModMusicJumpToOrder    uintptr
	fnSetMusicPosition       uintptr
	fnGetMusicPosition       uintptr
	fnMusicDuration          uintptr
	fnGetMusicLoopStartTime  uintptr
	fnGetMusicLoopEndTime    uintptr
	fnGetMusicLoopLengthTime uintptr
	fnPlaying                uintptr
	fnPlayingMusic           uintptr
	fnSetMusicCMD            uintptr
	fnSetSynchroValue        uintptr
	fnGetSynchroValue        uintptr
	fnSetSoundFonts          uintptr
	fnGetSoundFonts          uintptr
	fnSetTimidityCfg         uintptr
	fnGetTimidityCfg         uintptr
	fnGetChunk               uintptr
	fnCloseAudio             uintptr
)

var (
	extFnGetError uintptr
)

type ExternalFunctions struct {
	ExtFnGetError uintptr
}

// LoadLibrary loads the library and its functions.
func LoadLibrary(dllFile string, extFuncs *ExternalFunctions) (err error) {
	err = saveExternalFunctions(extFuncs)
	if err != nil {
		return err
	}

	err = dll.LoadLibrary(dllFile, &sdlMixDll, funcs, DllFuncNamePrefix)
	if err != nil {
		return err
	}

	return nil
}

// saveExternalFunctions saves external functions.
func saveExternalFunctions(extFuncs *ExternalFunctions) (err error) {
	if extFuncs == nil {
		return errors.New(ErrExternalFunctionsNotSet)
	}

	extFnGetError = extFuncs.ExtFnGetError

	return nil
}

// checkError checks error when necessary.
func checkError(ret uintptr) (err error) {
	if int(ret) == 0 {
		return nil
	}

	return GetError()
}

// mustBeNoError panics if an error occurs.
func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}

// BytePtrFromStringP converts a Go string into a C string.
// If something goes wrong, it panics.
func BytePtrFromStringP(s string) (cpS *byte) {
	var err error
	cpS, err = windows.BytePtrFromString(s)
	mustBeNoError(err)
	return cpS
}
