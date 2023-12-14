package util

import "sync"

var lock = &sync.Mutex{}
var reqQue int64 = 1
var waitCount int64 = 0

func AllocId() int64 {
	var result int64 = 1

	waitCount++
	lock.Lock()

	waitCount--

	result = reqQue

	if waitCount == 0 {
		reqQue = 1
	} else {
		reqQue++
	}

	lock.Unlock()
	return result
}
