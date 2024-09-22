package main

import (
	"code-cherish/core/analyzer"
	"fmt"
)

func analyzeCode(code string, lang string) {
	analyzer, err := analyzer.GetAnalyzer(lang)

	if err != nil {
		fmt.Println(err)
		return
	}

	issues := analyzer.Analyze(code)
	fmt.Println(issues)
}

func main() {
	analyzeCode("console.log('Hello, World!')", "js")
	fmt.Println("AST")
}
