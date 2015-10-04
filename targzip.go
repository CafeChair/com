package com

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func handleErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func tarFile(path string, tw *tar.Writer, fi os.FileInfo) {
	fr, err := os.Open(path)
	handleErr(err)
	defer fr.Close()

	hdr := new(tar.Header)
	hdr.Name = path
	hdr.Size = fi.Size()
	hdr.Mode = int64(fi.Mode())
	hdr.ModTime = fi.ModTime()

	err = tw.WriteHeader(hdr)
	handleErr(err)

	_, err = io.Copy(tw, fr)
	handleErr(err)
	return
}

func tarDir(dirPath string, tw *tar.Writer) {
	dir, err := os.Open(dirPath)
	handleErr(err)
	defer dir.Close()
	fis, err := dir.Readdir(0)
	handleErr(err)

	for _, fi := range fis {
		curPath := dirPath + "/" + fi.Name()
		if fi.IsDir() {
			tarDir(curPath, tw)
		} else {
			tarFile(curPath, tw, fi)
		}
	}
	return
}

func TarGz(tarFile string, srcDir string) {
	fw, err := os.Create(tarFile)
	handleErr(err)
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	tarDir(srcDir, tw)
	return
}
