package m

// AssertionHandler
// SDL_AssertState (SDLCALL *SDL_AssertionHandler)(const SDL_AssertData* data, void* userdata);
// SDL_assert.h
type AssertionHandler func(data uintptr, userdata uintptr) uintptr
