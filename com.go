package com
import (
	"os"
	"strings"
	"io/ioutil"
	"crypto/md5"
	"path/filepath"
)

//遍历目录以及子目录下的文件,可以指定文件后缀名
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

// 遍历目录下所有文件的md5值
func MD5Files(root string) (map[string][md5.Size]byte, error) {
	mp := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(filename string, f os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		data,err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		mp[filename] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil,err
	}
	return mp,nil
}

//指定文件符是否是文件
func IsFile(filename string) bool {
	file,err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

//指定文件或者目录是否存在
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//逐行读取文件返回列表
func ParseFile(filename string) ([]string, error) {
	str := make([]string, 0)
	file,err := os.Open(filename)
	if err != nil {
		return nil,err
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line,err := buf.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			str = append(str, line)
		}
	}
	return str,nil
}