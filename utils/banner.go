package utils

import "fmt"

func ShowBanner() {
	fmt.Println(ColorBlue)
	fmt.Println("=====================================")
	fmt.Println("         GREETER CLI v3.0.0          ")
	fmt.Println("=====================================")
	fmt.Println(ColorReset)
}
