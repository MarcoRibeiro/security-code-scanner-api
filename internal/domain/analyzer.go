package domain

type Analyzer interface {
	Name() string
	SupportedFileExtensions() []string
	Analyze(data string) (bool, error)
}