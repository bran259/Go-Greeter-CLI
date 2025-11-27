package main

import (
	"fmt"
	"os"

	"go-greeter/cmd"
	"go-greeter/utils"
)

func main() {
	utils.ShowBanner()

	if len(os.Args) < 2 {
		cmd.PrintHelp()
		return
	}

	switch os.Args[1] {
	case "hello":
		cmd.HelloCommand(os.Args[2:])
	case "bye":
		cmd.ByeCommand(os.Args[2:])
	case "time":
		cmd.TimeCommand(os.Args[2:])
	case "lang":
		cmd.LangCommand(os.Args[2:])
	case "random":
		cmd.RandomCommand()
	case "ask":
		cmd.AskInteractive()
	case "version":
		fmt.Println("Greeter CLI v3.0.0")
	default:
		fmt.Println("Unknown command:", os.Args[1])
		cmd.PrintHelp()
	}
}
