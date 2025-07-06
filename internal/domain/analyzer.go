package domain

type Analyzer interface {
	Name() string
	SupportedFileExtensions() []string
	Analyze(line string) bool
}