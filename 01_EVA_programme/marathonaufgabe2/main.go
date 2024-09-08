/*
Aufgabe 3

Ein Marathonläufer (42,195 km) möchte eine bestimmte Endzeit laufen (z.B. 3 Stunden, 35 Minuten, 30 Sekunden). Sein Trainer möchte ein Programm schreiben, dass ihm nach einer bestimmten Teilstrecke (z.B. 10 km) die benötigte Zwischenzeit ausgibt. Hilf dem Trainer und schreibe ein Programm, dass das entsprechende leistet. Erstelle vorher mit dem Struktogrammeditor ein Struktogramm und binde es als Bild ein!

Möglicher Programmablauf:

Eingabe der Endzeit: 3:35:10
Teilstrecke: 10

Zwischenzeit: 0:51:04
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
	fmt.Print("Eingabe der Endzeit (s:mm:ss): ")
	eZeitStr, _ := reader.ReadString('\n')
	var eZeitList []string = strings.Split(eZeitStr, ":")

	if len(eZeitList) != 3 {
		fmt.Println("Falsches Zeitformat!")
		return
	}

	eStunden, err := strconv.Atoi(strings.TrimSpace(eZeitList[0]))
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	eMinuten, err := strconv.Atoi(strings.TrimSpace(eZeitList[1]))
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	eSekunden, err := strconv.Atoi(strings.TrimSpace(eZeitList[2]))
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

	Zwischenzeit := strconv.Itoa(aStunden)
	if aMinuten < 10 {
		Zwischenzeit = Zwischenzeit + ":0" + strconv.Itoa(aMinuten)
	} else {
		Zwischenzeit = Zwischenzeit + ":" + strconv.Itoa(aMinuten)
	}
	if aSekunden < 10 {
		Zwischenzeit = Zwischenzeit + ":0" + strconv.Itoa(aSekunden)
	} else {
		Zwischenzeit = Zwischenzeit + ":" + strconv.Itoa(aSekunden)
	}

	println("Zwischenzeit:", Zwischenzeit)

}
