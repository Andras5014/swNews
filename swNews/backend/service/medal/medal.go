package medal

import (
	"backend/pkg/cache"
	"backend/pkg/serializer"
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
)

type WrapJSON struct {
	M Medals `json:"Medals"`
}

type Medals struct {
	SportsMedals []SportsMedalJson `json:"SportMedals"`
}
type SportsMedalJson struct {
	Countries []CountryJSON `json:"Countries"`
}
type CountryJSON struct {
	CountryName string    `json:"CountryName"`
	CountryCode string    `json:"CountryCode"`
	Rank        int       `json:"Rank"`
	TotalCount  int       `json:"TotalCount"`
	Gold        CountJSON `json:"Gold"`
	Silver      CountJSON `json:"Silver"`
	Bronze      CountJSON `json:"Bronze"`
}

type CountJSON struct {
	Count int `json:"Count"`
}

type Item struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Rank        int    `json:"rank"`
	TotalCount  int    `json:"totalCount"`
	Gold        int    `json:"gold"`
	Silver      int    `json:"silver"`
	Bronze      int    `json:"bronze"`
}

func init() {
	gob.Register([]Item{})
}

type Service struct {
}

var url = "https://api.worldaquatics.com/fina/competitions/3337/medals"

func (service *Service) GetMedal() serializer.Response {
	res, ok := cache.Get("medals")
	if ok {
		return serializer.Response{
			Data: res,
		}
	}

	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return serializer.AppError("Failed to get medal data", err)
	}

	reader := response.Body
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		return serializer.AppError("Failed to read medal data", err)
	}
	var wrap WrapJSON
	err = json.Unmarshal(data, &wrap)
	if err != nil {
		return serializer.AppError("Failed to parse medal data", err)
	}
	medalRes := make([]Item, 0)
	for _, country := range wrap.M.SportsMedals[0].Countries {
		medalRes = append(medalRes, Item{
			Country:     country.CountryName,
			CountryCode: country.CountryCode,
			Rank:        country.Rank,
			TotalCount:  country.TotalCount,
			Gold:        country.Gold.Count,
			Silver:      country.Silver.Count,
			Bronze:      country.Bronze.Count,
		})
	}
	_ = cache.Set("medals", medalRes, 0)
	return serializer.Response{
		Data: medalRes,
	}
}
