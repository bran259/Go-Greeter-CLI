// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const version = "1.0.0"

// func main() {
// 	if len(os.Args) < 2 {
// 		printHelp()
// 		return
// 	}

// 	command := os.Args[1]

// 	switch command {
// 	case "hello":
// 		helloCommand(os.Args[2:])
// 	case "bye":
// 		byeCommand(os.Args[2:])
// 	case "version":
// 		fmt.Println("Greeter CLI version", version)
// 	default:
// 		fmt.Printf("Unknown command: %s\n\n", command)
// 		printHelp()
// 	}
// }

// func helloCommand(args []string) {
// 	name := flag.NewFlagSet("hello", flag.ExitOnError)
// 	user := name.String("name", "friend", "Your name")
// 	loud := name.Bool("loud", false, "Shout the greeting")

// 	name.Parse(args)

// 	greeting := fmt.Sprintf("Hello, %s!", *user)
// 	if *loud {
// 		greeting = strings.ToUpper(greeting)
// 	}
// 	fmt.Println(greeting)
// }

// func byeCommand(args []string) {
// 	name := flag.NewFlagSet("bye", flag.ExitOnError)
// 	user := name.String("name", "friend", "Your name")

// 	name.Parse(args)

// 	fmt.Printf("Goodbye, %s!\n", *user)
// }

// func printHelp() {
// 	fmt.Println("Greeter CLI — Advanced Version")
// 	fmt.Println()
// 	fmt.Println("Usage:")
// 	fmt.Println("  greeter <command> [--flags]")
// 	fmt.Println()
// 	fmt.Println("Commands:")
// 	fmt.Println("  hello     Greet someone")
// 	fmt.Println("  bye       Say goodbye")
// 	fmt.Println("  version   Show version info")
// 	fmt.Println()
// 	fmt.Println("Examples:")
// 	fmt.Println("  greeter hello --name Brian")
// 	fmt.Println("  greeter hello --name Brian --loud")
// 	fmt.Println("  greeter bye --name Brian")
// }

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"strings"
	"time"
)

/* -----------------------------------------------------------
   GLOBALS & CONSTANTS
----------------------------------------------------------- */

const version = "2.5.0"

var colors = map[string]string{
	"reset": "\033[0m",
	"blue":  "\033[34m",
	"green": "\033[32m",
	"yellow": "\033[33m",
	"purple": "\033[35m",
}

// Multilingual greetings
var langGreetings = map[string]string{
	"en": "Hello",
	"es": "¡Hola",
	"fr": "Bonjour",
	"sw": "Jambo",
	"de": "Hallo",
	"it": "Ciao",
}

/* -----------------------------------------------------------
   ENTRY POINT
----------------------------------------------------------- */

func main() {
	showBanner()

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "hello":
		helloCommand(os.Args[2:])
	case "bye":
		byeCommand(os.Args[2:])
	case "version":
		fmt.Println("Greeter CLI version", version)
	case "time":
		timeGreetCommand(os.Args[2:])
	case "lang":
		langCommand(os.Args[2:])
	case "random":
		randomGreet()
	case "setname":
		saveDefaultName(os.Args[2:])
	case "ask":
		askInteractive()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printHelp()
	}
}

/* -----------------------------------------------------------
   COMMAND: hello
----------------------------------------------------------- */

func helloCommand(args []string) {
	fs := flag.NewFlagSet("hello", flag.ExitOnError)
	user := fs.String("name", loadDefaultName(), "Your name")
	loud := fs.Bool("loud", false, "Uppercase greeting")
	quiet := fs.Bool("quiet", false, "Disable colored output")

	fs.Parse(args)

	g := fmt.Sprintf("Hello, %s!", *user)
	if *loud {
		g = strings.ToUpper(g)
	}
	if *quiet {
		fmt.Println(g)
	} else {
		fmt.Println(colors["green"], g, colors["reset"])
	}
}

/* -----------------------------------------------------------
   COMMAND: bye
----------------------------------------------------------- */

func byeCommand(args []string) {
	fs := flag.NewFlagSet("bye", flag.ExitOnError)
	user := fs.String("name", loadDefaultName(), "Your name")
	quiet := fs.Bool("quiet", false, "Disable colored output")

	fs.Parse(args)

	msg := fmt.Sprintf("Goodbye, %s!", *user)
	if *quiet {
		fmt.Println(msg)
	} else {
		fmt.Println(colors["blue"], msg, colors["reset"])
	}
}

/* -----------------------------------------------------------
   COMMAND: time
----------------------------------------------------------- */

func timeGreetCommand(args []string) {
	name := loadDefaultName()
	hour := time.Now().Hour()

	var message string
	switch {
	case hour < 12:
		message = "Good morning"
	case hour < 18:
		message = "Good afternoon"
	default:
		message = "Good evening"
	}

	fmt.Printf("%s, %s!\n", message, name)
}

/* -----------------------------------------------------------
   COMMAND: lang
----------------------------------------------------------- */

func langCommand(args []string) {
	fs := flag.NewFlagSet("lang", flag.ExitOnError)
	lang := fs.String("lang", "en", "Language code (en, es, fr, sw, de, it)")
	name := fs.String("name", loadDefaultName(), "Your name")

	fs.Parse(args)

	greeting, ok := langGreetings[*lang]
	if !ok {
		fmt.Println("Unknown language code. Try: en, es, fr, sw, de, it")
		return
	}

	fmt.Printf("%s %s!\n", greeting, *name)
}

/* -----------------------------------------------------------
   COMMAND: random
----------------------------------------------------------- */

func randomGreet() {
	greetings := []string{
		"You're awesome",
		"Keep going",
		"Have a great day",
		"Stay positive",
		"You're doing great",
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(greetings[rand.Intn(len(greetings))] + "!")
}

/* -----------------------------------------------------------
   COMMAND: setname (writes config file)
----------------------------------------------------------- */

func saveDefaultName(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: greeter setname <your_name>")
		return
	}

	name := args[0]
	usr, _ := user.Current()
	file := usr.HomeDir + "/.greeterconfig"

	os.WriteFile(file, []byte(name), 0644)
	fmt.Println("Default name saved:", name)
}

/* -----------------------------------------------------------
   CONFIG HELPER
----------------------------------------------------------- */

func loadDefaultName() string {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.greeterconfig"

	data, err := os.ReadFile(file)
	if err != nil {
		return "friend"
	}
	return strings.TrimSpace(string(data))
}

/* -----------------------------------------------------------
   COMMAND: ask (interactive mode)
----------------------------------------------------------- */

func askInteractive() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Hello, %s! Welcome to the Greeter interactive mode.\n", name)
}

/* -----------------------------------------------------------
   HELP
----------------------------------------------------------- */

func printHelp() {
	fmt.Println("Greeter CLI — Enhanced Edition")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  greeter <command> [--flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  hello       Standard greeting")
	fmt.Println("  bye         Say goodbye")
	fmt.Println("  time        Greet based on current time")
	fmt.Println("  lang        Multilingual greeting")
	fmt.Println("  random      Fun random greeting")
	fmt.Println("  ask         Interactive mode")
	fmt.Println("  setname     Save default name")
	fmt.Println("  version     Show version info")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  greeter hello --name Brian")
	fmt.Println("  greeter time")
	fmt.Println("  greeter lang --lang sw --name Brian")
	fmt.Println("  greeter random")
}

func showBanner() {
	fmt.Println(colors["purple"])
	fmt.Println("=====================================")
	fmt.Println("         GREETER CLI v2.5.0          ")
	fmt.Println("=====================================")
	fmt.Println(colors["reset"])
}


