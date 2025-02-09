package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/Binject/debug/pe"
	"golang.org/x/sys/windows"
)

func SetHiddenAttribute(filePath string) error {
	utf16Path, err := windows.UTF16PtrFromString(filePath)
	if err != nil {
		return err
	}
	return windows.SetFileAttributes(utf16Path, windows.FILE_ATTRIBUTE_HIDDEN)
}

// 获取 COFF 符号表基址，条件是导出表中 main.PATHS 的值必须为 "UNABLE0"
func getCOFFSymbolBaseForSpecifiedFile(filePath string) (uint32, error) {
	// 读取 PE 头
	peFile, err := pe.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer peFile.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	offset := bytes.Index(data, []byte{'U', 'N', 'A', 'B', 'L', 'E', '0'})
	if offset == -1 {
		return 0, fmt.Errorf("未找到 UNABLE0")
	}

	for _, pe := range peFile.Symbols {
		if strings.Contains(pe.Name, "main.PATHS") {
			return uint32(offset) - pe.Value, nil
		}
	}
	return 0, fmt.Errorf("未找到 main.PATHS")
}

func findSymbol(filePath, symbolName string) (uint64, error) {
	symbolName = "main." + symbolName

	peFile, err := pe.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件 %v: %v", filePath, err)
	}
	defer peFile.Close()

	for _, pe := range peFile.Symbols {
		if pe.Name == symbolName {
			COFFSymbolsBase, err := getCOFFSymbolBaseForSpecifiedFile(filePath)
			if err != nil {
				return 0, err
			}
			return uint64(pe.Value) + uint64(COFFSymbolsBase), nil
		}
	}
	return 0, fmt.Errorf("未找到指定 Symbols")
}

func readFile(filePath string) (int, []byte) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
	}

	return len(data), data
}

func writeFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0666)
	return err
}

func isPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
