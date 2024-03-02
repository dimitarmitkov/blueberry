package main

import (
	"math/rand"
	"time"
)

func GenerateRandomText(length int) string {
	// Define the characters you want to include in the random text
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate the random text
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}

	return string(result)
}
