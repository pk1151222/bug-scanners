package utils

import (
	"log"
)

// LogError is a helper function to log errors
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
