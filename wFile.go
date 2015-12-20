package com

import (
	"os"
	"time"
)

type writeFile struct {
	result  chan string
	success chan bool
}

func WriteFile(filename, str string, w writeFile) {
	timeNow := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05")
	go func() {
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			w.result <- "[Error]"
			w.success <- false
		}
		defer file.Close()
		_, err = file.WriteString("[" + timeNow + "]" + "\t" + str + "\n")
		if err != nil {
			w.result <- "[Error]"
			w.success <- false
		} else {
			w.result <- str
			w.success <- true
		}
	}()
}

