package analyzers

import (
	"regexp"
)

type SQLInjectionAnalyzer struct{
	pattern *regexp.Regexp
}

func NewSQLInjectionAnalyzer() *SQLInjectionAnalyzer {
	pattern := regexp.MustCompile(`(?is)"[^"]*SELECT[^"]*WHERE[^"]*%s[^"]*"`)
	return &SQLInjectionAnalyzer{pattern: pattern}
}

func (a *SQLInjectionAnalyzer) Name() string {
	return "SQL Injection Analyzer"
}

func (a *SQLInjectionAnalyzer) SupportedFileExtensions() []string {
	return []string{".cs", ".go", ".java", ".js", ".php", ".py", ".rb", ".ts"}
}

func (a *SQLInjectionAnalyzer) Analyze(line string) (bool, error) {	
	matches := a.pattern.MatchString(line)
	return matches, nil
}