# SDLW

A wrapper of SDL library for Go language on Windows O.S. 

**Notes**

This project uses system calls because:
* CGO does not work normally on Windows O.S.;
* CGO is slower than system calls in Golang.

## SDL 2.0 API by Category
https://wiki.libsdl.org/SDL2/APIByCategory
