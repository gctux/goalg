package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Gib deinen Namen ein: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Hallo", name, "! Du hast einen schönen Namen!")
}
