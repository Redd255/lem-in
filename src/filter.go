package lemin

func SelectBestPaths(paths []Path, totalAnts int) []Path {
	var optimalPaths []Path
	remainingAnts := totalAnts

	for _, path := range paths {
		if !pathConflicts(optimalPaths, path) && calculatePathCapacity(path) <= remainingAnts {
			optimalPaths = append(optimalPaths, path)
			remainingAnts -= calculatePathCapacity(path)
		}
		if remainingAnts <= 0 {
			break
		}
	}
	return optimalPaths
}

func pathConflicts(existingPaths []Path, newPath Path) bool {
	occupiedRooms := make(map[string]bool)

	for _, path := range existingPaths {
		for _, room := range path.Path[1 : len(path.Path)-1] {
			occupiedRooms[room] = true
		}
	}

	for _, room := range newPath.Path[1 : len(newPath.Path)-1] {
		if occupiedRooms[room] {
			return true
		}
	}

	return false
}

func calculatePathCapacity(path Path) int {
	return len(path.Path) - 2
}
