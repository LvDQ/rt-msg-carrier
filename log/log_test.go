package log

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	logger.Debug("This is debug message.")
	logger.Error("This is error message.")
	time.Sleep(2 * time.Second)
	logger.Debug("This is debug message.")
	logger.Error("This is error message.")
	t.Log("complete")
}
