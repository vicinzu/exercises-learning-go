package main

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams []Team
	Wins  map[string]uint
}

func main() {
	// Nothing
}
