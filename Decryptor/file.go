package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func getAllDocs(dir string) []dirElement {

	dir = filepath.Clean(dir)

	var result = []dirElement{}

	glob_result, err := filepath.Glob(dir)
	if err != nil {
		fmt.Println("在处理通配符时出错，检查通配符是否正确", err)
		return nil
	}

	for _, dir_1 := range glob_result {

		err := filepath.Walk(dir_1, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				fmt.Println("filepath.Walk 游走文件出错:", err)
			}

			depth := len(strings.Split(path, string(os.PathSeparator)))

			if !info.IsDir() {
				dir_ele := &dirElement{
					path:  path,
					size:  info.Size(),
					depth: int64(depth),
				}
				result = append(result, *dir_ele)
			}
			return nil
		})

		if err != nil {
			fmt.Println("filepath.Walk 遍历目录时出错:", err)
		}
	}

	return result
}

func getSubDirs(dir string) []string {
	var result []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("读取文件时出错:", err)
	}

	for _, file := range files {
		if file.IsDir() {
			result = append(result, filepath.Join(dir, file.Name()))
		}
	}
	return result
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
