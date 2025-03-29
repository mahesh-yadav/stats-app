package logger

import (
	"io"
	"log"
	"os"
)

func ConfigureLogger() *os.File {
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	LstdFlags := log.Ldate | log.Ltime | log.Lshortfile
	w := io.MultiWriter(f, os.Stdout)
	// Redirect Go's default logger to use iLogger
	log.SetOutput(w)
	log.SetFlags(LstdFlags)
	log.SetPrefix("StatsApp: ")
	return f
}
