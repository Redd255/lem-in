package lemin

import (
	"fmt"
	"strings"
)

// Moves the ants trought the given paths.
func Sumilation(ants int, paths [][]string) {
	// 2d slice to track the movements of ants   each time
	x := make([][]string, ants)
	//  store wait and way
	p1 := make([]path, len(paths))
	for i, v := range paths {
		// exeample path = [1 2 3 ]
		// p1[0] =   wait : 2 , way : [2 3 ]
		p1[i] = path{len(v) - 1, v[1:]}
	}
	for i := 0; i < ants; i++ {
		// the paths already sorted so we take the one the one with less wait time wich always the first one
		curr := 0
		// in > index , v > wait time with way , p1 all paths
		for in, v := range p1 {
			//  check the wait of all rooms and store the index of the one with less wait time 
			//  curr var will change each time  the condition find path with less wait time 
			if p1[curr].wait > v.wait {
				curr = in
			}
		}

		for in := len(p1[curr].way); in < p1[curr].wait; in++ {
			x[i] = append(x[i], "")

		}

		for _, v := range p1[curr].way {
			x[i] = append(x[i], fmt.Sprintf("L%d-%s", i+1, v))
			fmt.Println(x[i])

		}
		p1[curr].wait += 1
	}

	DisplayResult(x)

}

// DisplayResult outputs the results of the path calculation.
func DisplayResult(result [][]string) {
	for i := 0; i < len(result[len(result)-1]); i++ {
		x := []string{}
		for _, v := range result {
			if len(v) > i && v[i] != "" {
				x = append(x, v[i])
			}
		}
		fmt.Println(strings.Join(x, " "))
	}
}
