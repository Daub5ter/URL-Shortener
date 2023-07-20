package random

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		got := NewRandomString(4)
		if got == "" {
			t.Error()
		} else {
			fmt.Println(got)
		}
	}
}
