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

	progress_bar.SetValue(0)
	for i, doc := range docs {
		progress_tips.SetText("正在解密: " + doc.path)
		progress_bar.SetValue(float64(i+1) / float64(len(docs)))
		if IS_DEBUG {
			fmt.Println(i, doc.depth, doc.size, doc.path)
		}
		_, docData := readFile(doc.path)
		if len(docData) > 17 {
			docHeader := docData[:17]
			docData = docData[17:]

			if bytes.Equal(docHeader, RSA1_HEADER) {
				ENC_FILE_FIND++
				data, err := rsa_decrypt(docData)
				if err != nil {
					addError(fmt.Errorf("%v, %s", err, doc.path))
					continue
				}
				writeFile(doc.path, data)
				ENC_FILE_DECRYPTED++
			} else if bytes.Equal(docHeader, AES1_HEADER) {
				ENC_FILE_FIND++
				data, err := aes_decrypt(docData)
				if err != nil {
					addError(fmt.Errorf("%v, %s", err, doc.path))
					continue
				}
				writeFile(doc.path, data)
				ENC_FILE_DECRYPTED++
			}
		}
	}
}
