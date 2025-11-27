package cmd

import (
	"flag"
	"fmt"
	"strings"

	"go-greeter/utils"
)

// Handles: greeter hello --name Brian --loud
func HelloCommand(args []string) {
	fs := flag.NewFlagSet("hello", flag.ExitOnError)
	name := fs.String("name", utils.LoadDefaultName(), "Your name")
	loud := fs.Bool("loud", false, "Uppercase greeting")

	fs.Parse(args)

	message := fmt.Sprintf("Hello, %s!", *name)

	if *loud {
		message = strings.ToUpper(message)
	}

	fmt.Println(utils.ColorGreen, message, utils.ColorReset)
}
