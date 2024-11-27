package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(1000000) // Generates a random 6-digit number
	return fmt.Sprintf("%06d", otp)
}
