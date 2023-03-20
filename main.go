package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Maxlee136/Funtemps/conv"
)

func main() {
	inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputFile.WriteString("time;temperature\n")

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) == 2 {
			celsius, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				log.Printf("Error parsing temperature value in line: %s", line)
				continue
			}
			fahrenheit := conv.CelsiusToFahrenheit(celsius)
			outputFile.WriteString(fmt.Sprintf("%s;%.2f\n", fields[0], fahrenheit))
		}
	}
}
