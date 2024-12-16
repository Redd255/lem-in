package lemin


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
	}
	path struct {
		wait int
		way  []string
	}
)
