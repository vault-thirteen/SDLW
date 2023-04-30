# SDLW

A wrapper of SDL library for Go language on Windows O.S.

This library also provides various helper-functions to make usage of the SDL 
library much easier.

**Notes**

* This project is frozen due to bugs in Go language and SDL library.
  * More information can be found in comments to the `Test` function and C tests.


* This project uses system calls.
  * CGO does not work properly on Windows O.S.
  * CGO is slower than system calls in Golang.

## SDL 2.0 API by Category
https://wiki.libsdl.org/SDL2/APIByCategory  
Note that the official documentation is very poor.

## Examples
Code examples are available in the `example` folder.
