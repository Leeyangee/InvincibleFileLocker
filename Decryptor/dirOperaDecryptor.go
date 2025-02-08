package main

import (
	"bytes"
	"fmt"
	"sort"
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

func decryptSubDirByBFS(dir string) {
	docs := getAllDocs(dir)
	sort.Sort(dirElementSortBySize(docs))

	for i, doc := range docs {
		if IS_DEBUG {
			fmt.Println(i, doc.depth, doc.size, doc.path)
		}
		_, docData := readFile(doc.path)
		docHeader := docData[:17]
		docData = docData[17:]

		if bytes.Equal(docHeader, RSA1_HEADER) {
			ENC_FILE_FIND++
			data, err := rsa_decrypt(docData)
			if err != nil {
				fmt.Println(err)
				continue
			}
			writeFile(doc.path, data)
			ENC_FILE_DECRYPTED++
		} else if bytes.Equal(docHeader, AES1_HEADER) {
			ENC_FILE_FIND++
			data, err := aes_decrypt(docData)
			if err != nil {
				fmt.Println(err)
				continue
			}
			writeFile(doc.path, data)
			ENC_FILE_DECRYPTED++
		}
	}
}
