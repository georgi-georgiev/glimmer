package main

type LoggerWriter interface {
	Write(entry *logEntry) error
	Sync() error
	Close() error
}
