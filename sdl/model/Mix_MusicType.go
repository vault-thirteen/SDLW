package m

// Mix_MusicType
/*
typedef enum {
    MUS_NONE,
    MUS_CMD,
    MUS_WAV,
    MUS_MOD,
    MUS_MID,
    MUS_OGG,
    MUS_MP3,
    MUS_MP3_MAD_UNUSED,
    MUS_FLAC,
    MUS_MODPLUG_UNUSED,
    MUS_OPUS
} Mix_MusicType;
*/
type Mix_MusicType int

const (
	MUS_NONE           = Mix_MusicType(0)
	MUS_CMD            = Mix_MusicType(1)
	MUS_WAV            = Mix_MusicType(2)
	MUS_MOD            = Mix_MusicType(3)
	MUS_MID            = Mix_MusicType(4)
	MUS_OGG            = Mix_MusicType(5)
	MUS_MP3            = Mix_MusicType(6)
	MUS_MP3_MAD_UNUSED = Mix_MusicType(7)
	MUS_FLAC           = Mix_MusicType(8)
	MUS_MODPLUG_UNUSED = Mix_MusicType(9)
	MUS_OPUS           = Mix_MusicType(10)
)
