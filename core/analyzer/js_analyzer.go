package analyzer

import "code-cherish/core/models"

type JsAnalyzer struct{}

func (j *JsAnalyzer) Analyze(ast interface{}) []models.Issue {
	return []models.Issue{}
}
