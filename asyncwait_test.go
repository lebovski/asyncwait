package asyncwait

import (
	"testing"
	"time"
)

// TestPositiveAsyncWait check positive scenario when TestFunc duration less than timeout
func TestPositiveAsyncWait(t *testing.T) {
	ws, fs := testTemplate(20*time.Millisecond, 60*time.Millisecond, 20*time.Millisecond)
	if !ws {
		t.Errorf("AsyncWaitFunc is not true")
	}
	if !fs {
		t.Errorf("TestFunc is not true")
	}
}

// TestNegativeAsyncWait check negative scenario when TestFunc duration more than timeout
func TestNegativeAsyncWait(t *testing.T) {
	ws, fs := testTemplate(80*time.Millisecond, 60*time.Millisecond, 20*time.Millisecond)
	if ws {
		t.Errorf("AsyncWaitFunc is not false")
	}
	if fs {
		t.Errorf("TestFunc is not false")
	}
}

func testTemplate(funcTimeout, timeout, pollInterval time.Duration) (waitStatus, funcStatus bool) {
	funcStatus = false
	go func() {
		time.Sleep(funcTimeout)
		funcStatus = true
	}()
	waitStatus = NewAsyncWait(timeout, pollInterval).Wait(func() bool { return funcStatus })
	return
}
