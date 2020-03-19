package log

import (
	"io"
	"log"
	"log/syslog"
	"sync"
	"fmt"
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

type Logger struct {
	mu    sync.Mutex
	level int
	outs  []io.Writer
}

func New(level int,out ...io.Writer) *Logger {
	outs := make([]io.Writer,0)
	for _, o := range out {
		outs = append(outs, o)
	}
	
	return &Logger{level: level, outs: outs}
}

func (l *Logger) output(level int, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i:=0;i<len(l.outs);i++{
		lv:=levels[level]
		l.outs[i].Write([]byte(fmt.Sprintf("[%s] +%s",lv,msg)))
	}
}

func (l *Logger) Info(msg interface{}) {
	if l.level<LOG_INFO{
		l.output(l.level,fmt.Sprint(msg))
	}
}
