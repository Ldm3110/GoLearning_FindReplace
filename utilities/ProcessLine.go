package utilities

import "strings"

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	// Search if "Go" is in the line
	found = strings.Contains(line, old)
	if found == true {
		occ = strings.Count(line, "Go")
		res = strings.Replace(line, old, new, -1)
	} else {
		res = line
	}

	return found, res, occ
}
