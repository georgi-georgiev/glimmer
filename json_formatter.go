package main

import "encoding/json"

type jsonFormatter struct {
}

func NewJsonFormatter() LoggerFormatter {
	return &jsonFormatter{}
}

func (f *jsonFormatter) Format(entry *logEntry, buf *[]byte) (err error) {
	entry.File = toShort(entry.File)
	jsonBuf, err := json.Marshal(entry)
	*buf = append(*buf, jsonBuf...)
	return
}
