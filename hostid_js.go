// +build linux

package xid

import (
	"math/rand"
	"strconv"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())
var id = strconv.ItoA(rand.New(seed).Intn(32768))

func readPlatformMachineID() (string, error) {
	return id, nil
}
