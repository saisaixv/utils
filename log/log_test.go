package log

import(
	"testing"
	"os"
)

func Info_test(t *testing.T)  {
	logger:=New(LOG_INFO,os.Stderr)
	logger.Info("hello world")
}