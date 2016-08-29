package slog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func format(lvl, msg string) string {
	nl := ""
	if !strings.HasSuffix(msg, "\n") {
		nl = "\n"
	}
	var now, file string
	var line int

	if !timeOff {
		now = time.Now().Format("01-02 15:04:05.000")
	}
	if !lineOff {
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}
		if len(file) > len(codeRoot) && file[:len(codeRoot)] == codeRoot {
			file = file[len(codeRoot):]
		}
	}
	if timeOff {
		if levelOff {
			lvl = ""
		} else {
			lvl += " "
		}
		if lineOff {
			return fmt.Sprintf("%s%s%s", lvl, msg, nl)
		}
		return fmt.Sprintf("%s%s:%d %s%s", lvl, file, line, msg, nl)
	}
	if levelOff {
		lvl = ""
	} else {
		lvl += " "
	}
	if lineOff {
		return fmt.Sprintf("%s %s%s%s", now, lvl, msg, nl)
	}
	return fmt.Sprintf("%s %s%s:%d %s%s", now, lvl, file, line, msg, nl)
}

func write(p Priority, msg string) {
	if p > priority {
		return
	}
	switch p {
	case LOG_EMERG:
		msg = format("EMG", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_EMERG, msg)
		}
	case LOG_ALERT:
		msg = format("ALT", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_ALERT, msg)
		}
	case LOG_CRIT:
		msg = format("CRT", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_CRIT, msg)
		}
	case LOG_ERR:
		msg = format("ERR", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_ERR, msg)
		}
	case LOG_WARNING:
		msg = format("WRN", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_WARNING, msg)
		}
	case LOG_NOTICE:
		msg = format("NTC", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_NOTICE, msg)
		}
	case LOG_INFO:
		msg = format("INF", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_INFO, msg)
		}
	case LOG_DEBUG:
		msg = format("DBG", msg)
		if sw == nil {
			w.Write([]byte(msg))
		} else {
			sw(LOG_DEBUG, msg)
		}
	}
	return
}

func Alert(msg string) {
	write(LOG_ALERT, msg)
}

func Alertf(format string, v ...interface{}) {
	write(LOG_ALERT, fmt.Sprintf(format, v...))
}

func Crit(msg string) {
	write(LOG_CRIT, msg)
}

func Critf(format string, v ...interface{}) {
	write(LOG_CRIT, fmt.Sprintf(format, v...))
}

func Debug(msg string) {
	write(LOG_DEBUG, msg)
}

func Debugf(format string, v ...interface{}) {
	write(LOG_DEBUG, fmt.Sprintf(format, v...))
}

func Emerg(msg string) {
	write(LOG_EMERG, msg)
}

func Emergf(format string, v ...interface{}) {
	write(LOG_EMERG, fmt.Sprintf(format, v...))
}

func Err(msg string) {
	write(LOG_ERR, msg)
}

func Errf(format string, v ...interface{}) {
	write(LOG_ERR, fmt.Sprintf(format, v...))
}

func Info(msg string) {
	write(LOG_INFO, msg)
}

func Infof(format string, v ...interface{}) {
	write(LOG_INFO, fmt.Sprintf(format, v...))
}

func Notice(msg string) {
	write(LOG_NOTICE, msg)
}

func Noticef(format string, v ...interface{}) {
	write(LOG_NOTICE, fmt.Sprintf(format, v...))
}

func Warning(msg string) {
	write(LOG_WARNING, msg)
}

func Warningf(format string, v ...interface{}) {
	write(LOG_WARNING, fmt.Sprintf(format, v...))
}
