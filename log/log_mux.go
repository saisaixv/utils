package log

import (
	"io"
	"log"
)

// go log
type LogMux struct {
	l *log.Logger
}

func NewLogMux(out io.Writer, prefix string, flags int) *LogMux {
	m := &LogMux{}
	if prefix != "" {
		prefix = prefix + " "
	}
	m.l = log.New(out, prefix, flags)
	return m
}

func (m *LogMux) Output(level int, s string) error {
	return m.l.Output(4, "["+LevelName(level)+"]"+" "+s)
}
