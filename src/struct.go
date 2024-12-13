package lemin

const (
	MaxAnts       = 100000
	MinLineLength = 7
)

var (
	Ants           int
	Start          string
	End            string
	currentSpecial string
)

type (
	Colony struct {
		Rooms []*Room
	}
	Room struct {
		Name   string
		Tunnel []string
	}
	Path struct {
		Path   []string
		AntsIn int
	}
	Ant struct {
		Id   int      
		Path []string 
		Next int      
	}
	Tunnel struct {
		Rooms [2]string
	}
)
