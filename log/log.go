package log

import (
	"fmt"
	dfc "github.com/athlum/pkg/dateFormatconv"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

const format = "%s %s gosp[%s]: %s\n"

const (
	LEVEL_INFO  = "INFO"
	LEVEL_DEBUG = "DEBUG"
	LEVEL_WARN  = "WARN"
	LEVEL_ERROR = "ERROR"
	LEVEL_FATAL = "FATAL"
)

var (
	Default        = &Log{os.Stdout, 2}
	_dateFormat = ""
)

func init() {
	_dateFormat, _ = dfc.Format("yyyy-MM-dd HH:mm:ss.SSS")
}

type Log struct {
	writer    io.Writer
	calldepth int
}

func (l *Log) Printf(level, msg string) {
	_, file, line, ok := runtime.Caller(l.calldepth)
	if !ok {
		file = "???"
		line = 1
	}
	fileSep := fmt.Sprintf("%s:L%v", file[strings.LastIndex(file, "/")+1:], line)
	l.writer.Write([]byte(fmt.Sprintf(format, time.Now().Format(_dateFormat), level, fileSep, msg)))
}

func (l *Log) Info(v ...interface{}) {
	l.Printf(LEVEL_INFO, fmt.Sprint(v...))
}

func (l *Log) Infof(format string, v ...interface{}) {
	l.Printf(LEVEL_INFO, fmt.Sprintf(format, v...))
}

func (l *Log) Debug(v ...interface{}) {
	l.Printf(LEVEL_DEBUG, fmt.Sprint(v...))
}

func (l *Log) Debugf(format string, v ...interface{}) {
	l.Printf(LEVEL_DEBUG, fmt.Sprintf(format, v...))
}

func (l *Log) Warn(v ...interface{}) {
	l.Printf(LEVEL_WARN, fmt.Sprint(v...))
}

func (l *Log) Warnf(format string, v ...interface{}) {
	l.Printf(LEVEL_WARN, fmt.Sprintf(format, v...))
}

func (l *Log) Error(v ...interface{}) {
	l.Printf(LEVEL_ERROR, fmt.Sprint(v...))
}

func (l *Log) Errorf(format string, v ...interface{}) {
	l.Printf(LEVEL_ERROR, fmt.Sprintf(format, v...))
}

func (l *Log) Fatal(v ...interface{}) {
	l.Printf(LEVEL_FATAL, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Log) Fatalf(format string, v ...interface{}) {
	l.Printf(LEVEL_FATAL, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Info(v ...interface{}) {
	Default.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Default.Infof(format, v...)
}

func Debug(v ...interface{}) {
	Default.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	Default.Debugf(format, v...)
}

func Warn(v ...interface{}) {
	Default.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	Default.Warnf(format, v...)
}

func Error(v ...interface{}) {
	Default.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Default.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	Default.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	Default.Fatalf(format, v...)
}

