package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Gib die Zahl a ein: ")
	aString, _ := reader.ReadString('\n')
	a, err := strconv.ParseFloat(strings.TrimSpace(aString), 64)
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Print("Gib die Zahl b ein: ")
	bString, _ := reader.ReadString('\n')
	b, err := strconv.ParseFloat(strings.TrimSpace(bString), 64)
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Println("Summe von a und b:", a+b)
	fmt.Println("Differenz von a und b:", a-b)
	fmt.Println("Produkt von a und b:", a*b)
	fmt.Println("Quotient von a und b:", a/b)

}
