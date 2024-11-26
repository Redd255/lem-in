package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Read the data from the file
func (c *Colony) ParseFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	if err := c.AddData(lines); err != nil {
		return err
	}
	return nil
}
func (c *Colony) checks(lines []string) error {
	start := 0
	end := 0
	for _, line := range lines {
		if line == "##start" {
			start++
		}
		if line == "##end" {
			end++
		}
	}
	if start != 1 || end != 1 {
		return fmt.Errorf("chkatkhwr")
	}
	return nil
}

// store ants number
func (c *Colony) AddAnts(lines []string) error {
	numberOfAnts, err := strconv.Atoi(lines[0])
	if err != nil {
		return err
	}
	if numberOfAnts <= 0 || numberOfAnts > MaxAnts {
		return fmt.Errorf("invalid number of ants:%s", lines[0])
	}
	c.Ants = numberOfAnts
	return nil
}

func (c *Colony) AddData(lines []string) error {
	if err := c.AddAnts(lines); err != nil {
		return err
	}
	if err := c.checks(lines); err != nil {
		return err
	}
	for _, a := range lines[1:] {
		line := strings.TrimSpace(a)
		if len(line) < MinLineLength {
			return fmt.Errorf("invalid data format: %s", line)
		}

		switch line {
		case "##start":
			if c.Start != "" {
				return fmt.Errorf("start has already been defined as %s", c.Start)
			}
			c.currentSpecial = "start"
		case "##end":
			if c.End != "" {
				return fmt.Errorf("end has already been defined as %s", c.End)
			}
			c.currentSpecial = "end"
		default:
			if err := c.processLine(line); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Colony) processLine(line string) error {
	values := strings.Fields(line)

	// Handle Room Definition
	if len(values) == 3 {
		name := strings.TrimSpace(values[0])
		if !checkname(name) {
			return fmt.Errorf("invalid room name: %s", name)
		}

		x, y, err := parseCoordinates(values[1], values[2], name)
		if err != nil {
			return err
		}

		// Add room to colony
		if err := c.AddRoom(name, x, y); err != nil {
			return err
		}

		// Assign start or end based on currentSpecial
		switch c.currentSpecial {
		case "start":
			c.Start = name
		case "end":
			c.End = name
		}
		c.currentSpecial = "" // Reset currentSpecial

		// Handle Tunnel Definition
	} else if tunVal := strings.Split(line, "-"); len(tunVal) == 2 {
		if err := c.AddTunnels(tunVal[0], tunVal[1]); err != nil {
			return err
		}
	}
	return nil
}

// Helper function to parse coordinates
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

func (c *Colony) AddRoom(name string, x int, y int) error {
	cords := [2]int{x, y}

	// Check if the room already exists
	for _, room := range c.Rooms {
		if room.Name == name {
			return fmt.Errorf("room %s %d %d already exists! ", name, cords[0], cords[1])
		}
	}
	// Add the new room
	c.Rooms = append(c.Rooms, &Room{Name: name, Coordinates: cords})
	return nil
}
func checkname(name string) bool {
	if name == "" || name[0] == '#' || name[0] == 'L' {
		return false
	}
	for _, r := range name {
		if unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func (c *Colony) GetRoom(name string) *Room {
	for _, room := range c.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

func (c *Colony) AddTunnels(from, to string) error {
	if c.GetRoom(from) == nil || c.GetRoom(to) == nil {
		return fmt.Errorf("room : %s doesent exist to link it with : %s", from, to)
	}
	sourceRoom := c.GetRoom(from)
	destinationRoom := c.GetRoom(to)
	sourceRoom.Tunnel = append(sourceRoom.Tunnel, destinationRoom)
	destinationRoom.Tunnel = append(destinationRoom.Tunnel, sourceRoom)
	return nil
}
