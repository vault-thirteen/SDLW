package m

// Mix_MusicInterface
/*
typedef struct
{
    const char *tag;
    Mix_MusicAPI api;
    Mix_MusicType type;
    SDL_bool loaded;
    SDL_bool opened;

	// A lot of function pointers ...
	int (*Load)(void);
	int (*Open)(const SDL_AudioSpec *spec);
	void *(*CreateFromRW)(SDL_RWops *src, int freesrc);
	void *(*CreateFromFile)(const char *file);
	void (*SetVolume)(void *music, int volume);
	int (*GetVolume)(void *music);
	int (*Play)(void *music, int play_count);
	SDL_bool (*IsPlaying)(void *music);
	int (*GetAudio)(void *music, void *data, int bytes);
	int (*Jump)(void *music, int order);
	int (*Seek)(void *music, double position);
	double (*Tell)(void *music);
	double (*Duration)(void *music);
	double (*LoopStart)(void *music);
	double (*LoopEnd)(void *music);
	double (*LoopLength)(void *music);
	const char* (*GetMetaTag)(void *music, Mix_MusicMetaTag tag_type);
	void (*Pause)(void *music);
	void (*Resume)(void *music);
	void (*Stop)(void *music);
	void (*Delete)(void *music);
	void (*Close)(void);
	void (*Unload)(void);
} Mix_MusicInterface;
*/
type Mix_MusicInterface struct {
	Tag    *byte
	Api    Mix_MusicAPI
	Type   Mix_MusicType
	Loaded Bool
	Opened Bool

	// A lot of function pointers ...
}
