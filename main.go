// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     reader := bufio.NewReader(os.Stdin)

//     fmt.Print("What is your name? ")

//     name, _ := reader.ReadString('\n')
//     name = strings.TrimSpace(name)

//     fmt.Printf("Hello, %s! Welcome to the Go Greeter CLI.\n", name)
// }


package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const version = "1.0.0"

func main() {
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
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printHelp()
	}
}

func helloCommand(args []string) {
	name := flag.NewFlagSet("hello", flag.ExitOnError)
	user := name.String("name", "friend", "Your name")
	loud := name.Bool("loud", false, "Shout the greeting")

	name.Parse(args)

	greeting := fmt.Sprintf("Hello, %s!", *user)
	if *loud {
		greeting = strings.ToUpper(greeting)
	}
	fmt.Println(greeting)
}

func byeCommand(args []string) {
	name := flag.NewFlagSet("bye", flag.ExitOnError)
	user := name.String("name", "friend", "Your name")

	name.Parse(args)

	fmt.Printf("Goodbye, %s!\n", *user)
}

func printHelp() {
	fmt.Println("Greeter CLI — Advanced Version")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  greeter <command> [--flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  hello     Greet someone")
	fmt.Println("  bye       Say goodbye")
	fmt.Println("  version   Show version info")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  greeter hello --name Brian")
	fmt.Println("  greeter hello --name Brian --loud")
	fmt.Println("  greeter bye --name Brian")
}
