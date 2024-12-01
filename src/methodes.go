package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Add the data from the file
func (c *Colony) AddData(lines []string) error {
	//Check the the data before add it
	if err := CheckData(lines); err != nil {
		return err
	}

	//Add the number of ants
	Ants, _ = strconv.Atoi(lines[0])

	//Add the rooms
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		switch line {
		case "##start":
			currentSpecial = "start"
		case "##end":
			currentSpecial = "end"
		default:
			if err := c.ProcessLine(line); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Colony) ProcessLine(line string) error {
	RoomVal := strings.Fields(line)
	tunVal := strings.Split(line, "-")

	// for room
	if len(RoomVal) == 3 {
		name := strings.TrimSpace(RoomVal[0])
		if !CheckName(name) {
			return fmt.Errorf("invalid room name: %s", name)
		}

		// Add room to colony
		if err := c.AddRoom(name); err != nil {
			return err
		}

		// Assign start or end based on currentSpecial
		switch currentSpecial {
		case "start":
			Start = name
		case "end":
			End = name
		}
		currentSpecial = ""

		// for tunnels
	} else if len(tunVal) == 2 {
		if err := c.AddTunnels(tunVal[0], tunVal[1]); err != nil {
			return err
		}
	}
	return nil
}

// adds the Rooms
func (c *Colony) AddRoom(name string) error {

	// Check if the room already exists
	for _, room := range c.Rooms {
		if room.Name == name {
			return fmt.Errorf("room %s  already exists! ", name)
		}
	}
	// Add the new room
	c.Rooms = append(c.Rooms, &Room{Name: name})
	return nil
}

// adds the tunnels between rooms
func (c *Colony) AddTunnels(from, to string) error {
	if c.GetRoom(from) == nil || c.GetRoom(to) == nil {
		return fmt.Errorf("room : %s doesent exist to link it with : %s", from, to)
	}
	sourceRoom := c.GetRoom(from)
	destinationRoom := c.GetRoom(to)
	sourceRoom.Tunnel = append(sourceRoom.Tunnel, to)
	destinationRoom.Tunnel = append(destinationRoom.Tunnel, from)
	return nil
}

func (c *Colony) GetRoom(name string) *Room {
	for _, room := range c.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

// Breadth-First Search (BFS) algorithm to find all paths from start to end room
func (c *Colony) FindPaths() ([]Path, error) {

	var paths []Path
	queue := [][]string{{Start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:] // Dequeue
		currentRoomName := path[len(path)-1]

		if currentRoomName == End {
			paths = append(paths, Path{Path: path})
			continue
		}

		// Explore adjacent rooms (tunnels)
		currentRoom := c.GetRoom(currentRoomName)
		for _, adj := range currentRoom.Tunnel {
			if !PathContainsRoom(path, adj) {
				// Avoid cycles by checking if the room is already in the path
				newPath := append([]string{}, path...) // Create a new path
				newPath = append(newPath, adj)
				queue = append(queue, newPath)
			}
		}
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("there is no path from start to end")
	}

	return paths, nil
}
