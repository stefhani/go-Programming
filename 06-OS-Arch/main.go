package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

/*
>go run main.go
>go build  main.go
>./main
>go install
> bin 06-OS-Arch
*/
