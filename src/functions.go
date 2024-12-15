package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

// Check the the data before add it
func CheckData(lines []string) error {
	// check start and end
	start := 0
	end := 0
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if i != len(lines)-1 && line == "##start" && strings.TrimSpace(lines[i+1]) != "##end" {
			start++
		}
		if i != len(lines)-1 && line == "##end" && strings.TrimSpace(lines[i-1]) != "##Start" {
			end++
		}

		checkCoor := strings.Split(line, " ")
		if len(checkCoor) == 3 {
			_, err := strconv.Atoi(checkCoor[1])
			if err != nil {
				return fmt.Errorf("there is a problem in coordinations")
			}
			_, err1 := strconv.Atoi(checkCoor[2])
			if err1 != nil {
				return fmt.Errorf("there is a problem in coordinations")
			}
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

// Check if the path contains the given room
func PathContainsRoom(path []string, roomName string) bool {
	for _, room := range path {
		if room == roomName {
			return true
		}
	}
	return false
}
