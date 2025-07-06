package mocks

type mockAnalyzer struct {
	name       string
	extensions []string
	matchLine  string
}

func (m *mockAnalyzer) Name() string                      { return m.name }
func (m *mockAnalyzer) SupportedFileExtensions() []string { return m.extensions }
func (m *mockAnalyzer) Analyze(data string) (bool, error) {
	return data == m.matchLine, nil
}

func NewMockAnalyzer(name string, extensions []string, matchLine string) *mockAnalyzer {
	return &mockAnalyzer{
		name:       name,
		extensions: extensions,
		matchLine:  matchLine,
	}
}