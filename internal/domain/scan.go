package domain

type Scan struct {
	Id       string
	Path     string
	Exclude  []string
	Findings []Finding
	Done     bool
	Err      string
}
