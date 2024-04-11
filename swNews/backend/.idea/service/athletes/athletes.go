package athletes

import (
	"backend/pkg/cache"
	"backend/pkg/conf"
	"backend/pkg/serializer"
	"encoding/gob"
	"encoding/json"
	"os"
)

var athleteJsonFile = "/athletes.json"

type Athlete struct {
	Country     string `json:"country"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	Dob         string `json:"dob"`
	CountryCode string `json:"countryCode"`
}

type AthleteService struct {
}

type CountryJSON struct {
	CountryName string        `json:"CountryName"`
	CountryCode string        `json:"CountryCode"`
	AthleteList []AthleteJSON `json:"Participations"`
}

type AthleteJSON struct {
	Gender    int    `json:"Gender"`
	LastName  string `json:"PreferredLastName"`
	FirstName string `json:"PreferredFirstName"`
	Dob       string `json:"DOB"`
}

func init() {
	gob.Register([]Athlete{})
}

func (service *AthleteService) GetAthletes() serializer.Response {
	res, ok := cache.Get("athletes")
	if ok {
		return serializer.Response{
			Data: res,
		}
	}
	countryList := make([]CountryJSON, 0)
	bytes, err := os.ReadFile(conf.AppConf.StaticPath + athleteJsonFile)
	if err != nil {
		return serializer.AppError("Failed to read athletes.json", err)
	}
	err = json.Unmarshal(bytes, &countryList)
	if err != nil {
		return serializer.AppError("Failed to parse athletes.json", err)
	}
	athleteRes := make([]Athlete, 0)
	for _, country := range countryList {
		for _, athlete := range country.AthleteList {
			athleteRes = append(athleteRes, Athlete{
				Country:     country.CountryName,
				CountryCode: country.CountryCode,
				FirstName:   athlete.FirstName,
				LastName:    athlete.LastName,
				Gender:      parseGender(athlete.Gender),
				Dob:         athlete.Dob,
			})
		}
	}
	_ = cache.Set("athletes", athleteRes, 0)
	return serializer.Response{
		Data: athleteRes,
	}
}

func parseGender(code int) string {
	if code == 1 {
		return "Female"
	} else {
		return "Male"
	}

}
