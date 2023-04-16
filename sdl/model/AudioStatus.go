package m

type AudioStatus int

const (
	AUDIO_STOPPED = AudioStatus(0)
	AUDIO_PLAYING = AudioStatus(1)
	AUDIO_PAUSED  = AudioStatus(2)
)
