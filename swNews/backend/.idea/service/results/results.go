package results

import (
	"backend/pkg/cache"
	"backend/pkg/conf"
	"backend/pkg/serializer"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type ResultItemJSON struct {
	Id             string `json:"Id"`
	DisciplineName string `json:"DisciplineName"`
}

type EventJSON struct {
	Sport []SportJSON `json:"Sports"`
}
type SportJSON struct {
	ResultList []ResultItemJSON `json:"DisciplineList"`
}
type ResultItem struct {
	Id             string `json:"id"`
	DisciplineName string `json:"disciplineName"`
}

var eventJsonFile = "/event.json"

type ResultService struct {
}

func (service *ResultService) GetResults() serializer.Response {
	res, ok := cache.Get("results")
	if ok {
		return serializer.Response{
			Data: res,
		}
	}
	var event EventJSON
	bytes, err := os.ReadFile(conf.AppConf.StaticPath + eventJsonFile)
	if err != nil {
		return serializer.AppError("Failed to read event.json", err)
	}
	err = json.Unmarshal(bytes, &event)
	if err != nil {
		return serializer.AppError("Failed to parse event.json", err)
	}
	resultRes := make([]ResultItem, 0)
	for _, result := range event.Sport[0].ResultList {
		resultRes = append(resultRes, ResultItem{
			Id:             result.Id,
			DisciplineName: result.DisciplineName,
		})
	}
	_ = cache.Set("results", resultRes, 0)
	return serializer.Response{
		Data: resultRes,
	}
}

type DetailResultJSON struct {
	Heats []HeatJSON `json:"Heats"`
}
type HeatJSON struct {
	UtcTime  string        `json:"UtcDateTime"`
	Name     string        `json:"Name"`
	Athletes []AthleteJSON `json:"Results"`
}

type AthleteJSON struct {
	CountryName string           `json:"NAT"`
	TotalPoints string           `json:"TotalPoints"`
	Rank        int              `json:"Rank"`
	FirstName   string           `json:"FirstName"`
	LastName    string           `json:"LastName"`
	Age         int              `json:"AthleteResultAge"`
	PointBehind string           `json:"PointsBehind"`
	Competitors []CompetitorJSON `json:"Competitors"`
}

type CompetitorJSON struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Country   string `json:"NAT"`
	Age       int    `json:"AthleteResultAge"`
}

type DetailService struct {
	Id string `uri:"id" binding:"required,min=1"`
}

type DetailResult struct {
	Name          string         `json:"name"`
	UtcTime       string         `json:"utcTime"`
	AthleteGroups []AthleteGroup `json:"athleteGroups"`
}

type AthleteGroup struct {
	Country     string          `json:"countryName"`
	Rank        int             `json:"rank"`
	TotalPoints string          `json:"points"`
	PointBehind string          `json:"pointBehind"`
	Athletes    []SimpleAthlete `json:"athletes"`
}

type SimpleAthlete struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func init() {
	gob.Register([]DetailResult{})
	gob.Register([]ResultItem{})

}

func (service *DetailService) GetDetail() serializer.Response {
	if res, ok := cache.Get("result-" + service.Id); ok {
		return serializer.Response{
			Data: res,
		}
	}
	var detail DetailResultJSON
	jsonFile := "/results/" + service.Id + ".json"
	bytes, err := os.ReadFile(conf.AppConf.StaticPath + jsonFile)
	if err != nil {
		return serializer.AppError(fmt.Sprintf("Failed to read %s", jsonFile), err)
	}
	err = json.Unmarshal(bytes, &detail)
	if err != nil {
		return serializer.AppError(fmt.Sprintf("Failed to parse %s", jsonFile), err)
	}
	resultRes := make([]DetailResult, 0)
	for _, heat := range detail.Heats {
		resItem := DetailResult{
			Name:    heat.Name,
			UtcTime: heat.UtcTime,
		}
		resItem.AthleteGroups = make([]AthleteGroup, 0)
		for _, athlete := range heat.Athletes {
			athleteItem := AthleteGroup{
				TotalPoints: athlete.TotalPoints,
				PointBehind: athlete.PointBehind,
				Country:     athlete.CountryName,
				Rank:        athlete.Rank,
			}
			if athlete.Competitors != nil {
				//	双人项目
				athleteItem.Athletes = make([]SimpleAthlete, 2)
				athleteItem.Athletes[0] = SimpleAthlete{
					FirstName: athlete.Competitors[0].FirstName,
					LastName:  athlete.Competitors[0].LastName,
					Age:       athlete.Competitors[0].Age,
				}
				athleteItem.Athletes[1] = SimpleAthlete{
					FirstName: athlete.Competitors[1].FirstName,
					LastName:  athlete.Competitors[1].LastName,
					Age:       athlete.Competitors[1].Age,
				}
			} else {
				//	单人项目
				athleteItem.Athletes = make([]SimpleAthlete, 1)
				athleteItem.Athletes[0] = SimpleAthlete{
					FirstName: athlete.FirstName,
					LastName:  athlete.LastName,
					Age:       athlete.Age,
				}
			}
			resItem.AthleteGroups = append(resItem.AthleteGroups, athleteItem)
		}
		resultRes = append(resultRes, resItem)
	}
	_ = cache.Set("result-"+service.Id, resultRes, 0)
	return serializer.Response{
		Data: resultRes,
	}
}
