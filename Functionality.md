## Supported Functionality
<!-- &#9746; = checked checkbox -->
<!-- &#9744;; = un-checked checkbox -->

### Basics
<table>
    <thead>
        <tr>
            <th>Status</th>
            <th>Group</th>
            <th>Function</th>
            <th>Header File</th>
            <th>Notes</th>
        </tr>
    </thead>
    <tbody>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=8 valign="top">Initialization & Shutdown</td><td>SDL_Init</td>
            <td>SDL.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_InitSubSystem</td>
            <td>SDL.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_Quit</td>
            <td>SDL.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_QuitSubSystem</td>
            <td>SDL.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_WasInit</td>
            <td>SDL.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetMainReady</td>
            <td>SDL_main.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_RegisterApp</td>
            <td>SDL_main.h</td>
            <td>Do not call this directly.</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_UnregisterApp</td>
            <td>SDL_main.h</td>
            <td>Do not call this directly.</td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=9 valign="top">Configuration Variables</td><td>SDL_SetHintWithPriority</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetHint</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ResetHint</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ResetHints</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetHint</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetHintBoolean</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_AddHintCallback</td>
            <td>SDL_hints.h</td>
            <td>Callbacks are broken in Golang.</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_DelHintCallback</td>
            <td>SDL_hints.h</td>
            <td>Callbacks are broken in Golang.</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ClearHints</td>
            <td>SDL_hints.h</td>
            <td></td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=4 valign="top">Error Handling</td><td>SDL_ClearError</td>
            <td>SDL_error.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetError</td>
            <td>SDL_error.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetError</td>
            <td>SDL_error.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9744;</td>
            <td>SDL_GetErrorMsg</td>
            <td>SDL_error.h</td>
            <td>Use SDL_GetError.</td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=15 valign="top">Log Handling</td><td>SDL_LogSetAllPriority</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogSetPriority</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogGetPriority</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogResetPriorities</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_Log</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogVerbose</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogDebug</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogInfo</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogWarn</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogError</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogCritical</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogMessage</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogMessageV</td>
            <td>SDL_log.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogGetOutputFunction</td>
            <td>SDL_log.h</td>
            <td>Callbacks are broken in Golang.</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogSetOutputFunction</td>
            <td>SDL_log.h</td>
            <td>Callbacks are broken in Golang.</td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9744;</td>
            <td rowspan=6 valign="top">Assertions</td>
            <td>SDL_assert</td>
            <td>SDL_assert.h</td>
            <td>Macro. Not available in Go.</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetAssertionHandler</td>
            <td>SDL_assert.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetDefaultAssertionHandler</td>
            <td>SDL_assert.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetAssertionHandler</td>
            <td>SDL_assert.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetAssertionReport</td>
            <td>SDL_assert.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ResetAssertionReport</td>
            <td>SDL_assert.h</td>
            <td></td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=3 valign="top">Querying SDL Version</td>
            <td>SDL_GetVersion</td>
            <td>SDL_version.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetRevision</td>
            <td>SDL_version.h</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetRevisionNumber</td>
            <td>SDL_version.h</td>
            <td>Deprecated.</td>
        </tr>
    </tbody>
</table>

### Audio
<table>
    <thead>
        <tr>
            <th>Status</th>
            <th>Group</th>
            <th>Function</th>
            <th>Header File</th>
            <th>Notes</th>
        </tr>
    </thead>
    <tbody>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=38 valign="top">Audio</td>
            <td>SDL_GetNumAudioDrivers</td>
            <td>SDL_audio.h</td>
            <td></td>
        </tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetAudioDriver</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioInit</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioQuit</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetCurrentAudioDriver</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_OpenAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetNumAudioDevices</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetAudioDeviceName</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetAudioDeviceSpec</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetDefaultAudioInfo</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_OpenAudioDevice</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetAudioStatus</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetAudioDeviceStatus</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_PauseAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_PauseAudioDevice</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_LoadWAV_RW</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_FreeWAV</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_BuildAudioCVT</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_ConvertAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_NewAudioStream</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioStreamPut</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioStreamGet</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioStreamAvailable</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioStreamFlush</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_AudioStreamClear</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_FreeAudioStream</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_MixAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_MixAudioFormat</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_QueueAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_DequeueAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_GetQueuedAudioSize</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_ClearQueuedAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_LockAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_LockAudioDevice</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_UnlockAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_UnlockAudioDevice</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_CloseAudio</td><td>SDL_audio.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>SDL_CloseAudioDevice</td><td>SDL_audio.h</td><td></td></tr>
    </tbody>
</table>
