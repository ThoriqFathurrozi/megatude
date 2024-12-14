package datasource

type Gempa struct {
	Tanggal     string `json:"Tanggal"`
	Jam         string `json:"Jam"`
	DateTime    string `json:"DateTime"`
	Coordinates string `json:"Coordinates"`
	Lintang     string `json:"Lintang"`
	Bujur       string `json:"Bujur"`
	Magnitude   string `json:"Magnitude"`
	Kedalaman   string `json:"Kedalaman"`
	Wilayah     string `json:"Wilayah"`
}
