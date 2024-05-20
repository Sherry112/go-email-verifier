package utils

import "log"

func ErrorLogger(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
