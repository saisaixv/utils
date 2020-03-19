package log

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"sync"
)

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)

const (
	LOG_EMERG   = int(syslog.LOG_EMERG)
	LOG_ALERT   = int(syslog.LOG_ALERT)
	LOG_CRIT    = int(syslog.LOG_CRIT)
	LOG_ERR     = int(syslog.LOG_ERR)
	LOG_WARNING = int(syslog.LOG_WARNING)
	LOG_NOTICE  = int(syslog.LOG_NOTICE)
	LOG_INFO    = int(syslog.LOG_INFO)
	LOG_DEBUG   = int(syslog.LOG_DEBUG)
)

var (
	levels = map[int]string{
		LOG_EMERG:   "EMERG",
		LOG_ALERT:   "ALERT",
		LOG_CRIT:    "CRITICAL",
		LOG_ERR:     "ERROR",
		LOG_WARNING: "WARNING",
		LOG_NOTICE:  "NOTICE",
		LOG_INFO:    "INFO",
		LOG_DEBUG:   "DEBUG",
	}
)

const (
	namePrefix = "LEVEL"
)

func LevelName(level int) string {
	if name, ok := levels[level]; ok {
		return name
	}
	return namePrefix + strconv.Itoa(level)
}

type Muxer interface {
	Output(level int, s string) error
}

type Logger struct {
	mu    sync.Mutex
	level int
	outs  []Muxer
}

func New(level int, out ...Muxer) *Logger {
	outs := make([]Muxer, 0)
	for _, o := range out {
		outs = append(outs, o)
	}

	return &Logger{level: level, outs: outs}
}

func (l *Logger) output(level int, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i := 0; i < len(l.outs); i++ {
		l.outs[i].Output(level, msg)
	}
}

func (l *Logger) Output(level int, a ...interface{}) error {
	if level <= l.level {
		return l.output(level, fmt.Sprint(a...))
	}
	return nil
}

var std = New(LOG_DEBUG, NewLogMux(os.Stderr, "", LstdFlags|Lshortfile))

func Info(msg ...interface{}) {
	std.Output(LOG_INFO, msg...)
}
