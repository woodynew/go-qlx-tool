package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetAPPRootPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}

func GetFileRootPath(path string) string {
	return GetAPPRootPath() + "/" + path
}
