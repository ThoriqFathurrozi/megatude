package datasource

type GempaTerkini struct {
	Gempa
	Potensi string `json:"potensi"`
}

type GempaTerkiniResponse struct {
	GempaTerkini []GempaTerkini `json:"gempa_terkini"`
}
