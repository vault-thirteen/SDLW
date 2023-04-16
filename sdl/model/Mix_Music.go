package m

// Mix_Music
/*
struct _Mix_Music {
    Mix_MusicInterface *interface;
    void *context;

    SDL_bool playing;
    Mix_Fading fading;
    int fade_step;
    int fade_steps;

    char filename[1024];
};
typedef struct _Mix_Music Mix_Music;
*/
type Mix_Music struct {
	Interface *Mix_MusicInterface
	Context   *byte

	Playing   Bool
	Fading    Mix_Fading
	FadeStep  int
	FadeSteps int

	Filename [1024]byte
}
