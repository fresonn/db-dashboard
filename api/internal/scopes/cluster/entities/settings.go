package entities

type PostmasterSettings struct {
	Settings []Setting `json:"settings"`
}

type Setting struct {
	Name        string `json:"name"`
	Setting     string `json:"setting"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
}
