package bmkg

import (
	autogempa "github.com/ThoriqFathurrozi/megatude/third_party/bmkg/datasource"
)

type BMKGResponse struct {
	InfoGempa InfoGempa `json:"Infogempa"`
}

type InfoGempa struct {
	Gampa autogempa.AutoGempa `json:"gempa"`
}
