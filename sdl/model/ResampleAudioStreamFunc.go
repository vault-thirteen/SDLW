package m

// ResampleAudioStreamFunc
// typedef int (*SDL_ResampleAudioStreamFunc)(SDL_AudioStream *stream, const void *inbuf, const int inbuflen, void *outbuf, const int outbuflen);
type ResampleAudioStreamFunc func(stream uintptr, inbuf uintptr, inbuflen uintptr, outbuf uintptr, outbuflen uintptr) uintptr
