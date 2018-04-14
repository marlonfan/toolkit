package util

import (
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	i := 3
	Retry(3, time.Second*1, func() error {
		i--
		return errors.New("")
	})

	if i != 0 {
		t.Logf("current i is %d", i)
		t.Error("retry function error")
	}

	Retry(3, time.Second*1, func() error {
		i--
		return nil
	})

	if _, ok := Retry(3, time.Second*1, func() error {
		i--
		return nil
	}).(error); ok {
		t.Logf("current i is %d", i)
		t.Error("retry function error")
	}
}
