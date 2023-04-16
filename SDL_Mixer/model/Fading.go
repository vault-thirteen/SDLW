package mm

// Fading
/*
	typedef enum {
	    MIX_NO_FADING,
	    MIX_FADING_OUT,
	    MIX_FADING_IN
	} Mix_Fading;
*/
type Fading int

const (
	MIX_NO_FADING  = Fading(0)
	MIX_FADING_OUT = Fading(1)
	MIX_FADING_IN  = Fading(2)
)
