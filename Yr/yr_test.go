package yr_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Testfunksjonen tar inn en filbane og utfører testene som er beskrevet i oppgaven.
func Testfunksjonen(filbane string) {
	// Åpne filen for lesing.
	fil, err := os.Open(filbane)
	if err != nil {
		fmt.Println("Kunne ikke åpne filen:", err)
		return
	}
	defer fil.Close()

	// Les filen linje for linje og utfør testene.
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		linje := scanner.Text()
		deler := strings.Split(linje, ";")

		switch len(deler) {
		case 4:
			// Test 1: Antall linjer i filen er 16756.
			// Bruk en teller for å telle antall linjer i filen.
			// Her antar vi at hver linje er en værmåling.
			// Du kan endre dette hvis linjene i filen har en annen struktur.
			teller++
		case 5:
			// Test 2-4: Sjekk om de ønskede verdiene stemmer overens med forventede verdier.
			if deler[3] == "6" {
				if deler[4] != "42,8" {
					fmt.Println("Test feilet: ", linje)
				}
			} else if deler[3] == "0" {
				if deler[4] != "32" {
					fmt.Println("Test feilet: ", linje)
				}
			} else if deler[3] == "-11" {
				if deler[4] != "12,2" {
					fmt.Println("Test feilet: ", linje)
				}
			}
		case 1:
			// Test 5: Endre teksten i linjen og legg til studentens navn.
			if strings.Contains(linje, "Data er gyldig per") && strings.Contains(linje, "Meteorologisk institutt (MET)") {
				nyLinje := strings.Replace(linje, "Data er gyldig per", "Data er basert på gyldig data (per "+dato+")", 1)
				nyLinje += "endringen er gjort av STUDENTENS_NAVN"
				fmt.Println(nyLinje)
			}
		}
	}

	// Test 6: Gjennomsnittstemperatur for august 2022 i grader Celsius.
	// Du kan bruke samme teller som i Test 1 for å telle antall målinger for august 2022.
	// Deretter kan du bruke en variabel for å holde summen av temperaturene.
	// Når du har summen og antall målinger, kan du regne ut gjennomsnittet.
	// Husk å sjekke at antall målinger er større enn null før du deler på det.
}

func main() {
	filbane := "værdata.txt"

	Testfunksjonen(filbane)
}
