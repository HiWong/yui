package common

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	root string // 当前的工作目录
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	root = dir
}

func GetAbsPath(filename string) (absPath string) {
	absPath = path.Join(root, filename)
	return
}

func ReadFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func WriteFile(filePath string, data []byte) {
	ioutil.WriteFile(filePath, data, 0644)
}
