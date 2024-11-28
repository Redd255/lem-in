package lemin

import (
	"errors"
	"fmt"
	"os"
	"sort"
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
	c.Ants, _ = strconv.Atoi(lines[0])

	//Add the rooms
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		switch line {
		case "##start":
			c.currentSpecial = "start"
		case "##end":
			c.currentSpecial = "end"
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
		switch c.currentSpecial {
		case "start":
			c.Start = name
		case "end":
			c.End = name
		}
		c.currentSpecial = ""

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
	sourceRoom.Tunnel = append(sourceRoom.Tunnel, destinationRoom)
	destinationRoom.Tunnel = append(destinationRoom.Tunnel, sourceRoom)
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
	// start room
	start := c.GetRoom(c.Start)

	var paths []Path
	queue := [][]*Room{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current.Name == c.End {
			temp := []string{}
			for _, room := range path {
				temp = append(temp, room.Name)
			}
			paths = append(paths, Path{Path: temp})
			continue
		}

		for _, adj := range current.Tunnel {
			if !PathContainsRoom(path, adj) {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, adj)
				queue = append(queue, newPath)
			}
		}
	}

	if len(paths) == 0 {
		return nil, errors.New("there is no path from start to end")
	}

	// Sort paths by length (shortest to longest)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i].Path) < len(paths[j].Path)
	})

	return paths, nil
}
