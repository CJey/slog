package slog

import (
	"io"
	"os"
)

var timeOff bool
var fileOff bool
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

var _srcroot_ string

func init() {
	w = os.Stderr
	priority = LOG_INFO
	callDepth = 3
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

func SetFileOff(off bool) {
	fileOff = off
}

func SetCallDepth(cd int) {
	callDepth = cd
}
