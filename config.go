package slog

import (
	"io"
	"os"
	"strings"
)

var timeOff bool
var lineOff bool
var levelOff bool
var callDepth int

type Priority int
type Writer func(lvl Priority, msg string)

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

var w io.Writer
var sw Writer
var priority Priority

var codeRoot string

func init() {
	w = os.Stderr
	priority = LOG_NOTICE
	callDepth = 3
}

func SetCodeRoot(root string) {
	codeRoot = root
}

func SetPriority(p Priority) {
	switch {
	case p > LOG_DEBUG:
		p = LOG_DEBUG
	case p < LOG_EMERG:
		p = LOG_EMERG
	}
	priority = p
}

func SetPriorityString(p string) {
	p = strings.ToUpper(p)
	switch p {
	case "EMERG":
		priority = LOG_EMERG
	case "ALERT":
		priority = LOG_ALERT
	case "CRIT":
		priority = LOG_CRIT
	case "ERR":
		priority = LOG_ERR
	case "WARNING":
		priority = LOG_WARNING
	case "NOTICE":
		priority = LOG_NOTICE
	case "INFO":
		priority = LOG_INFO
	case "DEBUG":
		priority = LOG_DEBUG
	}
}

func SetWriter(wt io.Writer) {
	sw = nil
	w = wt
}

func SetWriterFunc(wtf Writer) {
	sw = wtf
	w = nil
}

func SetLevelOff(off bool) {
	levelOff = off
}

func SetTimeOff(off bool) {
	timeOff = off
}

func SetLineOff(off bool) {
	lineOff = off
}

func SetCallDepth(cd int) {
	callDepth = cd
}
