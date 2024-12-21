package datasource

type GempaDirasakan struct {
	Gempa
	Dirasakan string `json:"dirasakan"`
}
