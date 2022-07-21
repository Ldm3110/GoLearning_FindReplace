package main

import (
	"FindReplace/utilities"
	"fmt"
)

func main() {
	// init values
	srcFile := "text.txt"
	oldText := "Go "
	newText := "Golang "

	// Working
	occ, lines, lineNumber, err := utilities.FindReplaceFile(srcFile, oldText, newText)
	if err != nil {
		fmt.Println("Error -", err)
		return
	}

	utilities.PrintResults(lineNumber, occ, lines)
}
