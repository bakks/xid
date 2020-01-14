// +build js

package xid

import (
	"math/rand"
	"strconv"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())
var id = strconv.Itoa(rand.New(seed).Intn(32768))

func readPlatformMachineID() (string, error) {
	return id, nil
}
