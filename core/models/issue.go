package models

type Issue struct {
	Type        string
	Description string
	Line        int
	Column      int
	FilePath    string
}
