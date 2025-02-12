package utils

import "log"

func InitLogger() {
	// Set up a custom logger if needed
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
