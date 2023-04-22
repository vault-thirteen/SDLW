# SDLW

A wrapper of SDL library for Go language on Windows O.S.

This library also provides various helper-functions to make usage of the SDL 
library much easier.

**Notes**

* This project is a work in progress.  
  * The code is not finished and may periodically get updates.


* This project uses system calls.
  * CGO does not work properly on Windows O.S.;
  * CGO is slower than system calls in Golang.

## SDL 2.0 API by Category
https://wiki.libsdl.org/SDL2/APIByCategory  
Note that the official documentation is very poor.

## Examples
Code examples are available in the `example` folder.
