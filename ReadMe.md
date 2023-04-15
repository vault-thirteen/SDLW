# SDLW

A wrapper of SDL library for Go language. 

**Notes**

This project uses system calls because:
* CGO does not work normally;
* CGO is slower than system calls in Golang.
