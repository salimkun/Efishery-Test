package util

import (
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

func GeneratePassword() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().Unix())

	length := 4

	ran_str := make([]rune, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = letters[rand.Intn(len(letters))]
	}

	// Displaying the random string
	str := string(ran_str)

	return str
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
