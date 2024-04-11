package schedule

import (
	"backend/pkg/cache"
	"backend/pkg/conf"
	"backend/pkg/serializer"
	"encoding/gob"
	"encoding/json"
	"os"
)

type MatchJSON struct {
	DisciplineName string     `json:"DisciplineName"`
	Id             string     `json:"Id"`
	Heats          []HeatJSON `json:"Heats"`
}

type HeatJSON struct {
	Date         string `json:"Date"`
	Time         string `json:"Time"`
	UtcDateTime  string `json:"UtcDateTime"`
	Name         string `json:"Name"`
	PhaseName    string `json:"PhaseName"`
	ResultStatus string `json:"ResultStatus"`
}

func init() {
	gob.Register([]Item{})
}

type Service struct {
}

var resultFolder = "/results"

type Item struct {
	DisciplineName string `json:"disciplineName"`
	Id             string `json:"id"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	UtcDateTime    string `json:"utcDateTime"`
	Name           string `json:"name"`
	PhaseName      string `json:"phaseName"`
	ResultStatus   string `json:"resultStatus"`
}

func (service *Service) GetSchedule() serializer.Response {
	res, ok := cache.Get("schedule")
	if ok {
		return serializer.Response{
			Data: res,
		}
	}
	var resList []Item
	dir, err := os.ReadDir(conf.AppConf.StaticPath + resultFolder)
	if err != nil {
		return serializer.AppError("Failed to read results folder", err)
	}
	for _, file := range dir {
		var match MatchJSON
		bytes, err := os.ReadFile(conf.AppConf.StaticPath + resultFolder + "/" + file.Name())
		if err != nil {
			return serializer.AppError("Failed to read match file", err)
		}
		err = json.Unmarshal(bytes, &match)
		if err != nil {
			return serializer.AppError("Failed to parse match file", err)
		}
		for _, heat := range match.Heats {
			resList = append(resList, Item{
				DisciplineName: match.DisciplineName,
				Id:             match.Id,
				Date:           heat.Date,
				Time:           heat.Time,
				UtcDateTime:    heat.UtcDateTime,
				Name:           heat.Name,
				PhaseName:      heat.PhaseName,
				ResultStatus:   heat.ResultStatus,
			})
		}
	}
	_ = cache.Set("schedule", resList, 0)
	return serializer.Response{
		Data: resList,
	}
}
