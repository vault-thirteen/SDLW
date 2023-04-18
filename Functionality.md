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

### Various Functions
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
            <td rowspan=1 valign="top">I/O</td>
            <td>SDL_RWFromFile</td>
            <td>SDL_rwops.h</td>
            <td></td>
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
            <td rowspan=39 valign="top">Audio</td>
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
        <tr><td align="center">&#9746;</td><td>SDL_LoadWAV</td><td>SDL_audio.h</td><td>Reimplemented macro.</td></tr>
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


### SDL Mixer
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
            <td rowspan=90 valign="top">Audio</td>
            <td>Mix_Linked_Version</td>
            <td>SDL_mixer.h</td>
            <td></td>
        </tr>
        <tr><td align="center">&#9746;</td><td>Mix_Init</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>Mix_Quit</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>Mix_OpenAudio</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>Mix_OpenAudioDevice</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_QuerySpec</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_AllocateChannels</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_LoadWAV_RW</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_LoadWAV</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_LoadMUS</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_LoadMUS_RW</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_LoadMUSType_RW</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_QuickLoad_WAV</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_QuickLoad_RAW</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FreeChunk</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FreeMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetNumChunkDecoders</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetChunkDecoder</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_HasChunkDecoder</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetNumMusicDecoders</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicDecoder</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_HasMusicDecoder</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicType</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicTitle</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicTitleTag</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicArtistTag</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicAlbumTag</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicCopyrightTag</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicHookData</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_RegisterEffect</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_UnregisterEffect</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_UnregisterAllEffects</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetPanning</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetPosition</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetDistance</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetReverseStereo</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_ReserveChannels</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GroupChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GroupChannels</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GroupAvailable</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GroupCount</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GroupOldest</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GroupNewer</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_PlayChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_PlayChannelTimed</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_PlayMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeInMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeInMusicPos</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeInChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeInChannelTimed</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_Volume</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_VolumeChunk</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_VolumeMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicVolume</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_MasterVolume</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_HaltChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_HaltGroup</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_HaltMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_ExpireChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeOutChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeOutGroup</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadeOutMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadingMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_FadingChannel</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_Pause</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_Resume</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_Paused</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_PauseMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_ResumeMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_RewindMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_PausedMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_ModMusicJumpToOrder</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetMusicPosition</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicPosition</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_MusicDuration</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicLoopStartTime</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicLoopEndTime</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetMusicLoopLengthTime</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_Playing</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_PlayingMusic</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetMusicCMD</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetSynchroValue</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetSynchroValue</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetSoundFonts</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetSoundFonts</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_SetTimidityCfg</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetTimidityCfg</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9744;</td><td>Mix_GetChunk</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>Mix_CloseAudio</td><td>SDL_mixer.h</td><td></td></tr>
        <tr><td align="center">&#9746;</td><td>Mix_GetError</td><td>SDL_mixer.h</td><td>Calls SDL_GetError.</td></tr>
    </tbody>
</table>
