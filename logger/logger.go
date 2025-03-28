package logger

import (
	"io"
	"log"
	"os"
)

func SetupLogger() (*log.Logger, *os.File) {
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	LstdFlags := log.Ldate | log.Ltime | log.Lshortfile
	w := io.MultiWriter(f, os.Stdout)
	iLogger := log.New(w, "iLogger: ", LstdFlags)
	return iLogger, f
}
