# SDLW

A wrapper of SDL library for Go language. 

**Notes**

This project will be frozen while Google is continuing to block all the 
attempts to support such cross-platform technologies as LLVM and CLang C 
compiler.

## Building

* Update all the header files in the `include` folder.
* Update the library files in the root folder:
  * `SDL2.lib`, `SDL2.dll`.
* Install the latest version of `LLVM` containing `CLang` compiler:
  * https://github.com/llvm/llvm-project/releases
* Run the build script – `build.bat`.
