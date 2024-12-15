package lemin

import (
	"fmt"
	"sort"
)

func SimulateAntMovement(totalAnts int, paths []Path, targetRoom string) {
	// Distribute ants across multiple paths based on their capacities.
	distributeAnts(paths, totalAnts)

	// Initialize ants with their assigned paths.
	ants := initializeAnts(totalAnts, paths)

	// Track room occupancy to avoid collisions.
	occupiedRooms := make(map[string]bool)
	var output string

	for len(ants) > 0 {
		usedTunnels := make(map[Tunnel]bool) // Track used tunnels to prevent simultaneous access.
		for i := 0; i < len(ants); i++ {
			ant := ants[i]
			currentRoom := ant.Path[ant.Next-1]
			nextRoom := ant.Path[ant.Next]
			tunnel := Tunnel{Rooms: [2]string{currentRoom, nextRoom}}

			// Check if the tunnel and next room are available.
			if !usedTunnels[tunnel] && !occupiedRooms[nextRoom] {
				output += fmt.Sprintf("L%d-%s ", ant.Id, nextRoom)
				usedTunnels[tunnel] = true

				// Update room occupancy.
				if ant.Next < len(ant.Path)-1 {
					occupiedRooms[nextRoom] = true
				}
				occupiedRooms[currentRoom] = false

				// Move the ant to the next room.
				ants[i].Next++
			}

			// Remove the ant if it reaches the target room.
			if nextRoom == targetRoom {
				ants = append(ants[:i], ants[i+1:]...)
				i-- // Adjust index after removal.
			}
		}
		output += "\n"
	}
	fmt.Print(output)
}

// distributeAnts evenly allocates ants to paths based on their lengths and current loads.
func distributeAnts(paths []Path, totalAnts int) {
	if len(paths) == 1 {
		// Assign all ants to the single available path.
		paths[0].AntsIn = totalAnts
		return
	}
	for ant := 1; ant <= totalAnts; ant++ {
		// Sort paths by their weight: path length + current ant load.
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].Path)+paths[i].AntsIn < len(paths[j].Path)+paths[j].AntsIn
		})
		// Assign the ant to the path with the smallest weight.
		paths[0].AntsIn++
		//fmt.Println(paths[0])
	}
}

// initializeAnts creates a list of ants and assigns them to paths.
func initializeAnts(totalAnts int, paths []Path) []Ant {
	ants := make([]Ant, totalAnts)
	pathIndex := 0

	for i := 0; i < totalAnts; i++ {
		for paths[pathIndex].AntsIn == 0 {
			pathIndex = (pathIndex + 1) % len(paths)
		}
		ants[i] = Ant{
			Id:   i + 1,
			Path: paths[pathIndex].Path,
			Next: 1,
		}
		//fmt.Println(ants[i])
		paths[pathIndex].AntsIn--
	}
	return ants
}
