package analyzer

import "code-cherish/core/models"

type JsAnalyzer struct{}

func (j *JsAnalyzer) Analyze(ast interface{}) models.Analysis {
	var analysis models.Analysis

	analysis.Issues = append(analysis.Issues, models.Issue{
		Type:        "js",
		Description: "Configuration: Allow users to configure which rules to apply",
		Location: models.Location{
			Filepath: "test.js",
			Line:     1,
			Column:   1,
		},
	})

	analysis.Suggestions = append(analysis.Suggestions, models.Suggestion{
		Description: "Configuration: Allow users to configure which rules to apply",
		Fix:         "Configuration: Allow users to configure which rules to apply",
		Location: models.Location{
			Filepath: "test.js",
			Line:     1,
			Column:   1,
		},
	})

	return analysis
}
