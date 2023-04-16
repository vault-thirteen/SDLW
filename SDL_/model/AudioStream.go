package m

import bt "github.com/vault-thirteen/auxie/BasicTypes"

// AudioStream
/*
struct _SDL_AudioStream

	{
		SDL_AudioCVT cvt_before_resampling;
		SDL_AudioCVT cvt_after_resampling;
		SDL_DataQueue *queue;
		SDL_bool first_run;
		Uint8 *staging_buffer;
		int staging_buffer_size;
		int staging_buffer_filled;
		Uint8 *work_buffer_base;
		int work_buffer_len;
		int src_sample_frame_size;
		SDL_AudioFormat src_format;
		Uint8 src_channels;
		int src_rate;
		int dst_sample_frame_size;
		SDL_AudioFormat dst_format;
		Uint8 dst_channels;
		int dst_rate;
		double rate_incr;
		Uint8 pre_resample_channels;
		int packetlen;
		int resampler_padding_samples;
		float *resampler_padding;
		void *resampler_state;
		SDL_ResampleAudioStreamFunc resampler_func;
		SDL_ResetAudioStreamResamplerFunc reset_resampler_func;
		SDL_CleanupAudioStreamResamplerFunc cleanup_resampler_func;
	};
*/
// typedef struct _SDL_AudioStream SDL_AudioStream;
type AudioStream struct {
	CvtBeforeResampling     AudioCVT
	CvtAfterResampling      AudioCVT
	Queue                   *DataQueue
	FirstRun                Bool
	StagingBuffer           *uint8
	StagingBufferSize       int
	StagingBufferFilled     int
	WorkBufferBase          *uint8
	WorkBufferLen           int
	SrcSampleFrameSize      int
	SrcFormat               AudioFormat
	SrcChannels             uint8
	SrcRate                 int
	DstSampleFrameSize      int
	DstFormat               AudioFormat
	DstChannels             uint8
	DstRate                 int
	RateIncr                bt.Double
	PreResampleChannels     uint8
	Packetlen               int
	ResamplerPaddingSamples int
	ResamplerPadding        *float32
	ResamplerState          *byte
	ResamplerFunc           ResampleAudioStreamFunc
	ResetResamplerFunc      ResetAudioStreamResamplerFunc
	CleanupResamplerFunc    CleanupAudioStreamResamplerFunc
}
