package analyzer

import (
	"code-cherish/core/models"
	"errors"
)

type Analyzer interface {
	Analyze(ast interface{}) models.Analysis
}

func GetAnalyzer(language string) (Analyzer, error) {
	switch language {
	case "js":
		return &JsAnalyzer{}, nil
	default:
		return nil, errors.New(`not supported`)
	}
}
