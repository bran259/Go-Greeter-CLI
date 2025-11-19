package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("What is your name? ")

    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Printf("Hello, %s! Welcome to the Go Greeter CLI.\n", name)
}
