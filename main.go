package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	passwordSize = 24

	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{};:'\",.<>/?"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	password := PassGen()

	fmt.Println(password)

	fmt.Println(passwordSize)
	fmt.Println(uppercase)
	fmt.Println(lowercase)
	fmt.Println(numbers)
	fmt.Println(symbols)
}

func PassGen() string {
	password := strings.Builder{}
	password.Grow(passwordSize)

	for range passwordSize {
		switch rand.Intn(4) {
		case 0:
			password.WriteByte(uppercase[rand.Intn(len(uppercase))])
		case 1:
			password.WriteByte(lowercase[rand.Intn(len(lowercase))])
		case 2:
			password.WriteByte(numbers[rand.Intn(len(numbers))])
		case 3:
			password.WriteByte(symbols[rand.Intn(len(symbols))])
		}
	}

	return password.String()
}
