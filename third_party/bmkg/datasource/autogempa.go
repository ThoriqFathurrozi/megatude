package datasource

type AutoGempa struct {
	Gempa
	Potensi   string `json:"potensi"`
	Dirasakan string `json:"dirasakan"`
	ShakeMap  string `json:"shakemap"`
}
