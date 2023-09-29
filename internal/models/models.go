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

func GetNav(title, link string, isActive bool) Nav {
	return Nav{
		Title:    title,
		Link:     link,
		IsActive: isActive,
	}
}
func GetSearch(field, value, url string) Search {
	return Search{
		Field: field,
		Value: value,
		Url:   url,
	}
}
