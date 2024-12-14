package datasource

type GempaDirasakan struct {
	Gempa
	Dirasakan string `json:"dirasakan"`
}

type GempaDirasakanResponse struct {
	GempaDirasakan []GempaDirasakan `json:"gempa_dirasakan"`
}
