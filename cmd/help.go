package cmd

import "fmt"

// PrintHelp shows all available commands and usage details.
func PrintHelp() {
	fmt.Println("=====================================")
	fmt.Println("            GREETER CLI v3.0         ")
	fmt.Println("=====================================")
	fmt.Println("Usage:")
	fmt.Println("  go run main.go <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  hello     Greet someone")
	fmt.Println("  bye       Say goodbye")
	fmt.Println("  time      Time-based greeting")
	fmt.Println("  lang      Set or view language")
	fmt.Println("  random    Random fun greeting")
	fmt.Println("  ask       Interactive mode")
	fmt.Println("  version   Show version info")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go hello --name Brian")
	fmt.Println("  go run main.go hello --name Brian --loud")
	fmt.Println("  go run main.go bye --name Brian")
	fmt.Println("  go run main.go time")
	fmt.Println("  go run main.go lang --set spanish")
	fmt.Println("  go run main.go random")
	fmt.Println("  go run main.go ask")
}
