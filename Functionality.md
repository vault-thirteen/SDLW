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
            <th>Notes</th>
        </tr>
    </thead>
    <tbody>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=6 valign="top">Initialization & Shutdown</td><td>SDL_Init</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_InitSubSystem</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_Quit</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_QuitSubSystem</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetMainReady</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_WasInit</td>
            <td></td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=9 valign="top">Configuration Variables</td><td>SDL_SetHintWithPriority</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetHint</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ResetHint</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ResetHints</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetHint</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetHintBoolean</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_AddHintCallback</td>
            <td>Callbacks are broken in Golang</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_DelHintCallback</td>
            <td>Callbacks are broken in Golang</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_ClearHints</td>
            <td></td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=3 valign="top">Error Handling</td><td>SDL_ClearError</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_GetError</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_SetError</td>
            <td></td>
        </tr>
<!-- -->
        <tr>
            <td align="center">&#9746;</td>
            <td rowspan=15 valign="top">Log Handling</td><td>SDL_LogSetAllPriority</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogSetPriority</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogGetPriority</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogResetPriorities</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_Log</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogVerbose</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogDebug</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogInfo</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogWarn</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogError</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogCritical</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogMessage</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogMessageV</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogGetOutputFunction</td>
            <td>Callbacks are broken in Golang</td>
        </tr>
        <tr>
            <td align="center">&#9746;</td>
            <td>SDL_LogSetOutputFunction</td>
            <td>Callbacks are broken in Golang</td>
        </tr>
    </tbody>
</table>
