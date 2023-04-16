package m

// Mix_Fading
/*
	typedef enum {
	    MIX_NO_FADING,
	    MIX_FADING_OUT,
	    MIX_FADING_IN
	} Mix_Fading;
*/
type Mix_Fading int

const (
	MIX_NO_FADING  = Mix_Fading(0)
	MIX_FADING_OUT = Mix_Fading(1)
	MIX_FADING_IN  = Mix_Fading(2)
)
