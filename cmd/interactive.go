package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// greeter ask
func AskInteractive() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	fmt.Println("Hello,", name, "- welcome to interactive mode!")
}
