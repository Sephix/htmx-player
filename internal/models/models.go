package models

type Nav struct {
	Title    string
	Link     string
	IsActive bool
}

type Search struct {
	Field string
	Value string
	Url   string
}
