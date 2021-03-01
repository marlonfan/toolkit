package util

import (
	"time"
)

// Retry 简单的重试方法, 指定固定的重试次数和时间,成功就返回
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Retry(attempts, sleep, fn)
		}
		return err
	}
	return nil
}
