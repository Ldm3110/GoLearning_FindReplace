package utilities

import (
	"bufio"
	"os"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, lineNumber int, err error) {
	// Open the base folder
	file, errOpen := os.Open(src)
	if errOpen != nil {
		return 0, nil, 0, errOpen
	}
	defer file.Close()

	// Create output folder
	outputFile, errCreation := os.Create("newText.txt")
	if errCreation != nil {
		return 0, nil, 0, errCreation
	}
	defer outputFile.Close()

	// Analyse the base folder
	lineNumber = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // read the folder line by line
		found, res, occLine := ProcessLine(line, old, new)
		// Write the line in the new folder line by line
		_, errWrite := outputFile.WriteString(res + "\n")
		if errWrite != nil {
			return 0, nil, 0, errWrite
		}
		if found == true {
			lines = append(lines, lineNumber)
			occ += occLine
		}
		lineNumber++
	}
	return occ, lines, lineNumber, err
}
