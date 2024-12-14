package bmkg

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ThoriqFathurrozi/megatude/configs"
)

type BMKG struct {
	cfg *configs.Config
}

func NewBMKG() *BMKG {
	cfg := configs.GetConfig()
	return &BMKG{
		cfg: cfg,
	}
}

func (b *BMKG) GetSourceData() BMKGResponse {
	url := b.cfg.Resource.Url

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	autogempaurl := fmt.Sprintf("%s/autogempa.json", url)

	response, err := client.Get(autogempaurl)
	if err != nil {
		return BMKGResponse{}
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return BMKGResponse{}
	}

	var responseObject BMKGResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
