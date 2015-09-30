package com

import "log"

const (	
	CLR_R = "\x1b[31;1m"
	CLR_G = "\x1b[32;1m"
	CLR_B = "\x1b[34;1m"
)

func ColorLog(color string, info interface{}) {
	var outputType string
	if color == CLR_B {
		outputType = "Loginfo"
	} else if color == CLR_R {
		outputType = "Errinfo"
	} else if color == CLR_G {
		outputType = "Runinfo"
	}
	log.Printf("%s: %s%s%s\n", outputType, color,info, "\x1b[0m")
}
