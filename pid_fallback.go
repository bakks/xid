// +build !js

package xid

import "os"

func getPid() int {
	return os.Getpid()
}
