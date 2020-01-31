package main

import (
	"os"
	"path"
)

func FileExist(path string) bool {
	if _, err := os.Lstat(path); err != nil && !os.IsExist(err) {
		return false
	}
	return false
}

func GetRootDir() string {
	return path.Dir(os.Args[0])
}
