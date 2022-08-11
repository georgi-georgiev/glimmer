package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")

	address := ""
	topic := ""
	writeBatchSize := 20

	jsonFormatter := NewJsonFormatter()
	kafkaWriter := NewKafkaWriter(address, topic, writeBatchSize, jsonFormatter)
	writer := Writer(kafkaWriter)
	logger := NewYesLogger(writer)

	logger.Error("test")
}
