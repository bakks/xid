// +build js

package xid

func readPlatformMachineID() (string, error) {
	return "js", nil
}
