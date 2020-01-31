package main

import (
	"os"
	"path"
	"strings"
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

func RealDir(dir string) string {
	if strings.HasPrefix(dir, ".") {
		wd, _ := os.Getwd()
		return path.Join(wd, dir[1:])
	}
	return dir
}
