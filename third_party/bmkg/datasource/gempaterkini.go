package datasource

type GempaTerkini struct {
	Gempa
	Potensi string `json:"potensi"`
}
