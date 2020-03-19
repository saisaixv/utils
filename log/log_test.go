package log

import (
	"testing"

	"github.com/labstack/gommon/log"
)

func Info_test(t *testing.T) {
	log.Info("hello world")
}
