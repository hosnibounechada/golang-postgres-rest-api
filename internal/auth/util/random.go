package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomUsername(firstName, lastName string) string {
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)

	rand.Seed(time.Now().UnixNano())
	randomDigits := fmt.Sprintf("%04d", rand.Intn(10000))

	username := firstName + lastName + randomDigits

	return username
}
