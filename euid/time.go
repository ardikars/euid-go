package euid

import "time"

func currentTimestamp() uint64 {
	var nano = time.Now().UnixNano()
	return uint64(nano / 1e6)
}
