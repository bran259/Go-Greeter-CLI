package cmd

import (
	"flag"
	"fmt"

	"go-greeter/utils"
)

// greeter lang --lang sw --name Brian
func LangCommand(args []string) {
	fs := flag.NewFlagSet("lang", flag.ExitOnError)

	lang := fs.String("lang", "en", "Language code")
	name := fs.String("name", utils.LoadDefaultName(), "Your name")

	fs.Parse(args)

	greet := utils.LangGreetings[*lang]
	if greet == "" {
		fmt.Println("Unknown language. Try: en, es, fr, de, sw, it")
		return
	}

	fmt.Println(greet, *name+"!")
}
