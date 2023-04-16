package m

// Mix_Chunk
/*
typedef struct Mix_Chunk {
    int allocated;
    Uint8 *abuf;
    Uint32 alen;
    Uint8 volume;
} Mix_Chunk;
*/
type Mix_Chunk struct {
	Allocated int
	Abuf      *uint8
	Alen      uint32
	Volume    uint8
}
