typedef void (SDLCALL *SDL_LogOutputFunction)(void *userdata, int category, SDL_LogPriority priority, const char *message);
extern DECLSPEC void SDLCALL SDL_LogSetAllPriority(SDL_LogPriority priority);
extern DECLSPEC void SDLCALL SDL_LogSetPriority(int category, SDL_LogPriority priority);
extern DECLSPEC SDL_LogPriority SDLCALL SDL_LogGetPriority(int category);
extern DECLSPEC void SDLCALL SDL_LogResetPriorities(void);
extern DECLSPEC void SDLCALL SDL_Log(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
extern DECLSPEC void SDLCALL SDL_LogVerbose(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
extern DECLSPEC void SDLCALL SDL_LogDebug(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
extern DECLSPEC void SDLCALL SDL_LogInfo(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
extern DECLSPEC void SDLCALL SDL_LogWarn(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
extern DECLSPEC void SDLCALL SDL_LogError(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
extern DECLSPEC void SDLCALL SDL_LogCritical(int category, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(2);
extern DECLSPEC void SDLCALL SDL_LogMessage(int category, SDL_LogPriority priority, SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(3);
extern DECLSPEC void SDLCALL SDL_LogMessageV(int category, SDL_LogPriority priority, const char *fmt, va_list ap);
extern DECLSPEC void SDLCALL SDL_LogGetOutputFunction(SDL_LogOutputFunction *callback, void **userdata);
extern DECLSPEC void SDLCALL SDL_LogSetOutputFunction(SDL_LogOutputFunction callback, void *userdata);
