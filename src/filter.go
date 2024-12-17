package lemin

func FilterPaths(paths []Path, totalAnts int) []Path {
	// best combo  of paths
	bestCombo := []Path{}
	//  used to stop finding paths after diturb all ants (no need more paths)
	remainingAnts := totalAnts
	// compare   with number if rooms to decide wich is the bestcombo
	bestComboRemaining := 0
	for i := 0; i < 1; i++ {
		// compare with remainigbest combo...
		numOfrooms := 0
		// tested path (we will compare it with best combo and decice wich one is the best)
		selectedPaths := []Path{}
		// affect the totale ants to remaningants
		remainingAnts = totalAnts
		// comparing with all other path
		path1 := paths[i]
		// we will have many paths in selected paths so we can compare it with best combo
		selectedPaths = append(selectedPaths, path1)
		// remove the start && end rooms + miminize number of ants
		remainingAnts -= len(path1.Path) - 2
		// we will compare numberof romms with bestcombo remaining to choose small one
		numOfrooms += len(path1.Path) - 2
		// we will iterate over paths only if we have enough number of ants (0+)
		if remainingAnts > 0 {
			for j := i; j < len(paths); j++ {
				// to  skip compare with current path
				if j != i {
					path2 := paths[j]
					// conditions verified only if we haven't same room's in paths
					if !PathsInterfear(selectedPaths, path2) {
						selectedPaths = append(selectedPaths, path2)
						remainingAnts -= len(path2.Path) - 2
						numOfrooms += len(path2.Path) - 2
						/*conditions verify the number of ants ,, if all ants are ditutbed
						we wont need other paths otherwise  */
						if remainingAnts <= 0 {
							return selectedPaths
						}
					}
				}
			}
		}
		// entred the condition only if the lentgh of selected is greater or equal
		// [[1 2 3][1 5 3]]       [[1 4 3]]  || [[1 2 3 ]] [[1 4 3 8 ]]
		if len(selectedPaths) >= len(bestCombo) {
			// must be equal
			//  to filter  paths that have less number of rooms
			// [[1 2 3 ]] [[1 4 3 8 ]]  && 1 < 2
			if len(selectedPaths) == len(bestCombo) && numOfrooms < bestComboRemaining {
				// bestco = 1
				bestComboRemaining = numOfrooms
				// bestcombo = [[1 2 3]]
				bestCombo = selectedPaths
				//  to check the largest path and affect it to best combo
				// (because we have lots of ants and we need the largest slice of paths)
			} else if len(selectedPaths) > len(bestCombo) {
				//  [[1 2 3][1 5 3]] --> best combo
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
