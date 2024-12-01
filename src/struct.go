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
	Ant struct {
		Id   int
		Path []string
		Next int
	}

	Path struct {
		Path   []string
		AntsIn int
	}
	Tunnels struct {
		Romms [2]string
	}
)
