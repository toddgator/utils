package utils

import (
	"log"
	"os"
)

func SetLogFile(filepath string) {
	// Attempting to open the log file supplied
	logFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Clean up when the done
	defer logFile.Close()

	// Set our file as the output destination for logging.
	log.SetOutput(logFile)
}
