package main

type LoggerFormatter interface {
	Format(*logEntry, *[]byte) error
}
