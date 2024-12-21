package bmkg

import (
	gempa "github.com/ThoriqFathurrozi/megatude/third_party/bmkg/datasource"
)

// auto gempa response
type BMKGResponseAuto struct {
	InfoGempa InfoGempaAuto `json:"Infogempa"`
}

type InfoGempaAuto struct {
	Gampa gempa.AutoGempa `json:"gempa"`
}

// gempa terkini response
type BMKGResponseTerkini struct {
	InfoGempa InfoGempaTerkini `json:"Infogempa"`
}

type InfoGempaTerkini struct {
	GempaTerkiniRes []gempa.GempaTerkini `json:"gempa"`
}

// gempa dirasakan response
type BMKGResponseDirasakan struct {
	InfoGempa InfoGempaDirasakan `json:"Infogempa"`
}

type InfoGempaDirasakan struct {
	GempaDirasakanRes []gempa.GempaDirasakan `json:"gempa"`
}
