package redis

import (
	"testing"
)

func TestDoSet(t *testing.T) {
	InitPool("192.168.10.168", 6379, "caton", 1, 30, 30, 30, 30)

}
