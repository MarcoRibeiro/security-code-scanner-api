package entities

import (
	"github.com/yourusername/yourproject/internal/domain/valueObjects"
)

type Scan struct {
	Id       string
	Path     string
	Exclude  []string
	Findings []valueObjects.Finding
	Done     bool
	Err      string
}
