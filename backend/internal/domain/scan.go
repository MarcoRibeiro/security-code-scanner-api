package domain

type Scan struct {
	Id            string
	Path          string
	Configuration *Configuration
	Findings      []Finding
	Done          bool
	Err           string
}
