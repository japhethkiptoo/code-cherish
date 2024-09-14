package utils

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WalkDir(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}

func GetFileExtension(path string) string {
	return filepath.Ext(path)
}
