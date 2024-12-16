package lemin

import (
	"fmt"
	"strings"
)


// Moves the ants trought the given paths.
func Sumilation(ants int, paths [][]string) {
	x := make([][]string, ants)
	p1 := make([]path, len(paths))
	for i, v := range paths {
		p1[i] = path{len(v) - 1, v[1:]}
	}
	for i := 0; i < ants; i++ {
		curr := 0
		for in, v := range p1 {
			if p1[curr].wait > v.wait {
				curr = in
			}
		}
		for in := len(p1[curr].way); in < p1[curr].wait; in++ {
			x[i] = append(x[i], "")
		}
		for _, v := range p1[curr].way {

			x[i] = append(x[i], fmt.Sprintf("L%d-%s", i+1, v))
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
