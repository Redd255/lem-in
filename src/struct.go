package lemin

var (
	Ants           int
	Start          string
	End            string
	currentSpecial string
)

type (
	Colony struct {
		Rooms []*Room // containg room struct wihch is the name and the tunnel  --> exectly point to
	}
	Room struct {
		Name   string
		Tunnel []string
	}
	Path struct {
		Path []string // takes the path by  bfs 
	}
	path struct {
		wait int  // The number of time steps the path has been occupied (ch7al mn nmla fatt 3la dak lpath)
		way  []string // the actual path (excluding the starting node)
	}
)
