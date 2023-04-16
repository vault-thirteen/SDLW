package m

import bt "github.com/vault-thirteen/auxie/BasicTypes"

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
