package main

import (
	"log"
	"os"
)

const (
	infoLevel = 1
	errLevel  = 2
)

type Log struct {
	level int
}

func NewLogger(logLevel string) *Log {
	if logLevel == "err" {
		return &Log{1}
	}
	if logLevel == "info" {
		return &Log{2}
	}

	return &Log{0}
}

func (l *Log) Infof(format string, args ...interface{}) {
	if l.level >= infoLevel {
		log.SetPrefix("INFO: ")
		log.SetOutput(os.Stdout)
		log.Printf(format, args...)
	}
}

func (l *Log) Errorf(format string, args ...interface{}) {
	if l.level >= errLevel {
		log.SetPrefix("ERROR: ")
		log.SetOutput(os.Stderr)
		log.Printf(format, args...)
	}
}

func (l *Log) Fatalf(format string, args ...interface{}) {
	log.SetPrefix("FATAL: ")
	log.SetOutput(os.Stderr)
	log.Fatalf(format, args...)
}
