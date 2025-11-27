package cmd

import (
	"fmt"
	"math/rand"
	"time"
)

// greeter random
func RandomCommand() {
	lines := []string{
		"You're awesome",
		"Keep going!",
		"Have a great day",
		"Stay positive",
		"You're doing great!",
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(lines[rand.Intn(len(lines))])
}
