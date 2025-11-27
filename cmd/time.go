package cmd

import (
	"fmt"
	"time"

	"go-greeter/utils"
)

// greeter time
func TimeCommand(args []string) {
	name := utils.LoadDefaultName()
	hour := time.Now().Hour()

	var greeting string
	switch {
	case hour < 12:
		greeting = "Good morning"
	case hour < 18:
		greeting = "Good afternoon"
	default:
		greeting = "Good evening"
	}

	fmt.Println(greeting+",", name+"!")
}
