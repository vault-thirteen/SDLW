@ECHO OFF

go env CC
SET CC="D:\Program Files\LLVM\bin\clang.exe"
go env CC

go env CGO_ENABLED
SET CGO_ENABLED=1
go env CGO_ENABLED

go build .
