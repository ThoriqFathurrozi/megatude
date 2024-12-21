package bmkg

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

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

func (b *BMKG) GetSourceData() (BMKGResponseAuto, BMKGResponseTerkini, BMKGResponseDirasakan) {
	var wg sync.WaitGroup

	wg.Add(3)

	var BMKGResponseAuto BMKGResponseAuto
	var BMKGResponseTerkini BMKGResponseTerkini
	var BMKGResponseDirasakan BMKGResponseDirasakan

	go b.requestData("/autogempa.json", &BMKGResponseAuto, &wg)
	go b.requestData("/gempaterkini.json", &BMKGResponseTerkini, &wg)
	go b.requestData("/gempadirasakan.json", &BMKGResponseDirasakan, &wg)

	wg.Wait()

	return BMKGResponseAuto, BMKGResponseTerkini, BMKGResponseDirasakan

}

func (b *BMKG) requestData(urlTarget string, Data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	url := b.cfg.Resource.Url

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	urlsParse := fmt.Sprintf("%s%s", url, urlTarget)

	client := &http.Client{Transport: tr}

	response, err := client.Get(urlsParse)

	if err != nil {
		fmt.Println(err)
		return
	}
	resData, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(resData, &Data)

	defer response.Body.Close()

}
