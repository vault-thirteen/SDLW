package m

// SDL_VARIABLE_LENGTH_ARRAY
/*
// This is for a variable-length array at the end of a struct:
// struct x { int y; char z[SDL_VARIABLE_LENGTH_ARRAY]; };
// Use this because GCC 2 needs different magic than other compilers.
#if (defined(__GNUC__) && (__GNUC__ <= 2)) || defined(__CC_ARM) || defined(__cplusplus)
#define SDL_VARIABLE_LENGTH_ARRAY 1
#else
#define SDL_VARIABLE_LENGTH_ARRAY
#endif
*/
const SDL_VARIABLE_LENGTH_ARRAY = 1 //TODO: ???
