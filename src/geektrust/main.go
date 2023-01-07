package main

import (
	"geektrust/train"
	userinput "geektrust/userInput"
)

func main() {

	text := userinput.FormattedInput()
	train.Train(text)

}
