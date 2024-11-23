package lemin

const (
	MaxAnts       = 100000
	MinLineLength = 3
)

type (
	Colony struct {
		Rooms          []*Room
		currentSpecial string
		Start          string
		End            string
		Ants           int
	}
	Room struct {
		Name        string
		Coordinates [2]int
		Tunnel      []*Room
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
)
