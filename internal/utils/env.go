package utils

import (
	"log"
	"os"
)

func GetEvnVariable(key string) string {
	value, isPresent := os.LookupEnv(key)
	if !isPresent {
		log.Fatalf("Missing %v env variable", key)
	}
	return value
}
