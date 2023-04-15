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
[List of supported functions](./Functionality.md).

## Adapters
[List of adapters](./Adapters.md).

## SDL 2.0 API by Category
https://wiki.libsdl.org/SDL2/APIByCategory
