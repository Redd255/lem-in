package lemin


func FilterPaths(paths []Path, totalAnts int) []Path {
	bestCombo := []Path{}
	remainingAnts := totalAnts
	bestComboRemaining := 0
	for i := 0; i < len(paths); i++ {
		numOfrooms := 0
		selectedPaths := []Path{}
		remainingAnts = totalAnts
		path1 := paths[i]
		selectedPaths = append(selectedPaths, path1)
		remainingAnts -=len(path1.Path) - 2 // kan9so number of rooms
		numOfrooms += len(path1.Path) - 2 // kanzido number of rooms
		if remainingAnts > 0 {
			for j := i; j < len(paths); j++ {
				if j != i {
					path2 := paths[j]
					if !PathsInterfear(selectedPaths, path2) {
						selectedPaths = append(selectedPaths, path2)
						remainingAnts -= len(path2.Path) - 2
						numOfrooms += len(path2.Path) - 2
						if remainingAnts <= 0 {
							return selectedPaths
						}
					}
				}
			}
		}
		if len(selectedPaths) >= len(bestCombo) {
			if len(selectedPaths) == len(bestCombo) && numOfrooms < bestComboRemaining {
				bestComboRemaining = numOfrooms
				bestCombo = selectedPaths
			} else if len(selectedPaths) > len(bestCombo) {
				bestCombo = selectedPaths
			}
		}
	}
	return bestCombo
}


func PathsInterfear(paths []Path, path2 Path) bool {
	occupiedRooms := make(map[string]bool)
	for _, path1 := range paths {
		for _, room1 := range path1.Path[1 : len(path1.Path)-1] {
			occupiedRooms[room1] = true
		}
	}
	for _, room2 := range path2.Path[1 : len(path2.Path)-1] {
		if occupiedRooms[room2] {
			return true
		}
	}
	return false
}
