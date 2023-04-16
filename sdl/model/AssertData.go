package m

// AssertData
/*
typedef struct SDL_AssertData
{
	int always_ignore;
	unsigned int trigger_count;
	const char *condition;
	const char *filename;
	int linenum;
	const char *function;
	const struct SDL_AssertData *next;
} SDL_AssertData;
*/
// SDL_assert.h
type AssertData struct {
	AlwaysIgnore int
	TriggerCount uint
	Condition    *byte
	Filename     *byte
	LineNum      int
	Function     *byte
	Next         *AssertData
}
