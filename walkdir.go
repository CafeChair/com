package com
import (
	"os"
	"strings"
	"path/filepath"
)

func WalkDir(dirpath,suffix string) (files []string, err error) {
	files = make([]string,0,30)
	suffix = strings.ToLower(suffix)
	filepath.Walk(dirpath, func(filename string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		} if strings.HasSuffix(strings.ToLower(f.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
		})
	return files,err
}
