package analyzer

import "code-cherish/core/models"

type Analyzer interface {
	Analyze(ast interface{}) []models.Issue
}

func getAnalyzer(language string) Analyzer {
	switch language {
	case "js":
		return &JsAnalyzer{}
	default:
		return nil
	}
}
