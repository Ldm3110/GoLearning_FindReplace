package utilities

import "fmt"

func PrintResults(lineNumber, occ int, lines []int) {
	// Print results
	fmt.Println("======  Results  ======")
	defer fmt.Println("==== End of Results ====")
	fmt.Println("Number of lines in folder : ", lineNumber)
	fmt.Println("Number of occurrences : ", occ)
	fmt.Println("Lines concerned : ", lines)
}
