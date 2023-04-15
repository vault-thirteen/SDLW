# SDLW

A wrapper of SDL library for Go language on Windows O.S. 

**Notes**

* This project is a work in progress.  
  * This means, it is not yet finished and periodically gets updates.


* This project uses system calls.
  * CGO does not work properly on Windows O.S.;
  * CGO is slower than system calls in Golang.


* This library uses a modular approach.  
  * Complex values (objects) returned by C library have simple Go types (`uintptr`) 
  to speed up the work. 
  * If you need to get inside the objects, you should use 
  adapters which convert types from `uintptr` into object types and reversed.
  * Note that Go code must not change C pointers, otherwise the program will 
crash.

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
            <td rowspan=6 valign="top">Initialization and Shutdown</td><td>SDL_Init</td>
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
            <td align="center">&#9744;</td>
            <td rowspan=4 valign="top">Configuration Variables</td><td>SDL_HintPriority</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9744;</td>
            <td>SDL_SetHintWithPriority</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9744;</td>
            <td>SDL_GetHint</td>
            <td></td>
        </tr>
        <tr>
            <td align="center">&#9744;</td>
            <td>SDL_SetHint</td>
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
    </tbody>
</table>

## SDL 2.0 API by Category
https://wiki.libsdl.org/SDL2/APIByCategory
