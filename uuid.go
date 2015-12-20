package com

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func NewUUID() string {
	var x [16]byte
	unixNow := time.Now().UnixNano()
	mrand := rand.New(rand.NewSource(unixNow)).Int63()
	binary.BigEndian.PutUint32(x[0:], uint32(unixNow))
	binary.BigEndian.PutUint32(x[4:], uint32(unixNow))
	binary.BigEndian.PutUint32(x[8:], uint32(mrand))
	binary.BigEndian.PutUint32(x[12:], uint32(mrand))
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", x[:4], x[4:6], x[6:8], x[8:10], x[10:])
}
