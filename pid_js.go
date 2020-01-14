// +build js

package xid

import (
	"math/rand"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())
var id = rand.New(seed).Intn(32768)

func getPid() int {
	return id
}
