# SDLW

A wrapper of SDL library for Go language on Windows O.S.

This library also provides various helper-functions to make usage of the SDL 
library much easier.

**Notes**

* This project is a work in progress.  
  * This means, it is not yet finished and periodically gets updates.


* This project uses system calls.
  * CGO does not work properly on Windows O.S.;
  * CGO is slower than system calls in Golang.


* This library uses a mixed approach for type handling.
  * Simple string types, such as `*char`, are converted into string;
  * Simple integers are converted into appropriate integer types;
  * Buffer pointers are converted into Go's `*uint8` or `*byte`;
  * Some complex objects are converted into Go's objects;
  * Some types which can not be converted into Go, are returned as `uintptr`;
  * Note that Go code must not change C pointers, otherwise the program will 
crash.

## Supported Functionality
[List of supported functions](./Functionality.md).

## SDL 2.0 API by Category
https://wiki.libsdl.org/SDL2/APIByCategory  
Note that the official documentation is very poor.

## Examples
Code examples are available in the `example` folder.
