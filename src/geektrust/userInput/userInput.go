package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FormattedInput() [][]string {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return make([][]string, 0)
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")
		return make([][]string, 0)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text [][]string

	for scanner.Scan() {
		text = append(text, strings.Split(strings.ReplaceAll(scanner.Text(), " ", "/"), "/"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
