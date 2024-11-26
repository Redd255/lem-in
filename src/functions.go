package lemin

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Check the the data before add it
func CheckData(lines []string) error {
	//check lines
	if len(lines) < MinLineLength {
		return fmt.Errorf("invalid number of lines")
	}

	// check ants
	numberOfAnts, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("invalid number of ants: %s", lines[0])
	}
	if numberOfAnts <= 0 || numberOfAnts > MaxAnts {
		return fmt.Errorf("invalid number of ants:%s", lines[0])
	}

	// check start and end
	start := 0
	end := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "##start" {
			start++
		}
		if line == "##end" {
			end++
		}
	}
	if start != 1 || end != 1 {
		return fmt.Errorf("there is a problem in start or end")
	}

	return nil
}

// Check room name before add it
func CheckName(name string) bool {
	if name != "" {
		for i, r := range name {
			if i == 0 && (r == '#' || r == 'L') {
				return false
			}
			if unicode.IsSpace(r) {
				return false
			}
		}
	}
	return true
}

// Check room coordinates before add it
func parseCoordinates(xStr, yStr, name string) (int, int, error) {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid x coordinate for room %s: %s", name, xStr)
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid y coordinate for room %s: %s", name, yStr)
	}
	return x, y, nil
}
