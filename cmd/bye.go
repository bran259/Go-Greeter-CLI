package cmd

import (
	"flag"
	"fmt"

	"go-greeter/utils"
)

// greeter bye --name Brian
func ByeCommand(args []string) {
	fs := flag.NewFlagSet("bye", flag.ExitOnError)
	name := fs.String("name", utils.LoadDefaultName(), "Your name")
	fs.Parse(args)

	fmt.Println(utils.ColorBlue, "Goodbye,", *name+"!", utils.ColorReset)
}
