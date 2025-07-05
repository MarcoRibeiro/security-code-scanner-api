package analyzers

import (
	"regexp"
)

type CrossSiteScriptingAnalyzer struct{
    scriptTagRegex *regexp.Regexp
}

func NewCrossSiteScriptingAnalyzer() *CrossSiteScriptingAnalyzer {
	scriptTagRegex := regexp.MustCompile(`(?i).*Alert(.*).*`)
	return &CrossSiteScriptingAnalyzer{scriptTagRegex: scriptTagRegex}
}

func (a *CrossSiteScriptingAnalyzer) Name() string {
	return "Cross-Site Scripting Analyzer"
}

func (a *CrossSiteScriptingAnalyzer) SupportedFileExtensions() []string {
	return []string{".html", ".js"}
}

func (a *CrossSiteScriptingAnalyzer) Analyze(data string) (bool, error) {
	matches := a.scriptTagRegex.MatchString(data)
	return matches, nil
}