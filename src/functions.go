package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

// Check the the data before add it
func CheckData(lines []string) error {
	//check lines
	if len(lines) < MinLineLength {
		return fmt.Errorf("invalid number of lines")
	}
	// check ants
	numberOfAnts, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil {
		return fmt.Errorf("invalid number of ants: %s", lines[0])
	}
	if numberOfAnts <= 0 || numberOfAnts > MaxAnts {
		return fmt.Errorf("invalid number of ants:%s", lines[0])
	}

	// check start and end
	start := 0
	end := 0
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if i != len(lines)-1 && line == "##start" {
			start++
			if i < len(lines)-1 && (lines[i+1] == "##end" || lines[i-1] == "##end") {
				return fmt.Errorf("error asahbi")
			}
		}
		if i != len(lines)-1 && line == "##end" {
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
	for i, r := range name {
		if i == 0 && (r == '#' || r == 'L') {
			return false
		}
	}
	return true
}

func PathContainsRoom(path []string, roomName string) bool {
	for _, room := range path {
		if room == roomName {
			return true
		}
	}
	return false
}
