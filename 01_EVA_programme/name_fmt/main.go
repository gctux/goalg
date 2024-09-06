package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hallo", name, "! Du hast einen schönen Namen!")
}
