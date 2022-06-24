package util

import (
	"fmt"
	"testing"
)

func TestAllocId(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(index int) {
			id := AllocId()
			fmt.Printf("%d->%d \n", index, id)
		}(i)
	}
}
