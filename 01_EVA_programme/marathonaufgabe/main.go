/*
Aufgabe 3

Ein Marathonläufer (42,195 km) möchte eine bestimmte Endzeit laufen (z.B. 3 Stunden, 35 Minuten, 30 Sekunden). Sein Trainer möchte ein Programm schreiben, dass ihm nach einer bestimmten Teilstrecke (z.B. 10 km) die benötigte Zwischenzeit ausgibt. Hilf dem Trainer und schreibe ein Programm, dass das entsprechende leistet. Erstelle vorher mit dem Struktogrammeditor ein Struktogramm und binde es als Bild ein!

Möglicher Programmablauf:

Eingabe der Endzeit
Stunden: 3
Minuten: 35
Sekunden: 30
Teilstrecke: 10

Zwischenzeit
Stunden: 0
Minuten: 51
Sekunden: 4
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Eingabe der Endzeit\n")
	fmt.Print("Stunden: ")
	eStundenStr, _ := reader.ReadString('\n')
	eStunden, err := strconv.Atoi(strings.TrimSpace(eStundenStr))
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Print("Minuten: ")
	eMinutenStr, _ := reader.ReadString('\n')
	eMinuten, err := strconv.Atoi(strings.TrimSpace(eMinutenStr))
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Print("Sekunden: ")
	eSekundenStr, _ := reader.ReadString('\n')
	eSekunden, err := strconv.Atoi(strings.TrimSpace(eSekundenStr))
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Print("Teilstrecke: ")
	teilstreckeStr, _ := reader.ReadString('\n')
	teilstrecke, err := strconv.ParseFloat(strings.TrimSpace(teilstreckeStr), 64)
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	eZeit := eStunden*60*60 + eMinuten*60 + eSekunden
	zZeit := int(math.Round(teilstrecke / 42.195 * float64(eZeit)))
	aStunden := zZeit / (60 * 60)
	aMinuten := (zZeit % (60 * 60)) / 60
	aSekunden := (zZeit % (60 * 60)) % 60

	println("Zwischenzeit:", aStunden, ":", aMinuten, ":", aSekunden)

}
