package parser

type Parser interface {
	Parse(content []byte) (interface{}, error)
}

func GetParser(language string) Parser {
	switch language {
	case "js":
		return &JsParser{}
	default:
		return nil
	}
}
