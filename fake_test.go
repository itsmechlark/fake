package fake

import (
	"testing"
)

func TestSetLang(t *testing.T) {
	err := SetLang("en")
	if err != nil {
		t.Error("SetLang should successfully set lang")
	}

	err = SetLang("sd")
	if err == nil {
		t.Error("SetLang with nonexistent lang should return error")
	}
}

// TestConcurrentSafety runs fake methods in multiple go routines concurrently.
// This test should be run with the race detector enabled.
func TestConcurrentSafety(t *testing.T) {
	workerCount := 10
	doneChan := make(chan struct{})

	for i := 0; i < workerCount; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				FirstName()
				LastName()
				Gender()
				FullName()
				Day()
				Country()
				Company()
				Industry()
				Street()
			}
			doneChan <- struct{}{}
		}()
	}

	for i := 0; i < workerCount; i++ {
		<-doneChan
	}
}
