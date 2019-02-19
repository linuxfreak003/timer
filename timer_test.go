package timer_test

import (
	"testing"
	"time"

	"github.com/linuxfreak003/timer"
)

func TestTimer(t *testing.T) {
	nt := timer.NewTimer(time.Minute*5 + time.Second*2)
	time.Sleep(2 * time.Second)

	// Reset
	nt.Reset(time.Minute*5 + time.Second*2)
	if nt.Remaining() < time.Minute*5 {
		t.Errorf("Reset failed, remaining time %s too small", nt.Remaining())
	}
	t.Logf("Remaining time: %s", nt.Remaining())

	// Add
	nt.Add(time.Minute * 5)
	if nt.Remaining() < time.Minute*10 {
		t.Errorf("Add failed: remaining time %s too small", nt.Remaining())
	}
	t.Logf("Remaining time: %s", nt.Remaining())

	// Subtract
	nt.Subtract(time.Minute * 5)
	if nt.Remaining() > time.Minute*10 || nt.Remaining() < time.Minute*5 {
		t.Errorf("Add failed: remaining time %s incorrect", nt.Remaining())
	}
	t.Logf("Remaining time: %s", nt.Remaining())
}
