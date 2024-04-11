package comment

import (
	"backend/pkg/cache"
	"backend/pkg/serializer"
	"encoding/gob"
	"time"
)

type Comment struct {
	Content string `json:"content"`
	Time    int64  `json:"time"`
	Name    string `json:"name"`
}

type CreateService struct {
	Content string `json:"content" binding:"required" `
	Name    string `json:"name" binding:"required"`
}

func init() {
	gob.Register([]Comment{})
}

func (service *CreateService) Create() serializer.Response {
	comments := make([]Comment, 0)
	res, ok := cache.Get("comment")
	if ok {
		comments = res.([]Comment)
	}
	comment := Comment{
		Content: service.Content,
		Time:    time.Now().Unix(),
		Name:    service.Name,
	}
	comments = append(comments, comment)
	err := cache.Set("comment", comments, 0)
	if err != nil {
		return serializer.AppError("Failed to save comment", err)
	}
	return serializer.Response{}
}

type ListService struct {
	PageNo int `uri:"pageNo" binding:"required"`
}

type ListResponse struct {
	PageNo    int       `json:"pageNo"`
	PageSize  int       `json:"pageSize"`
	TotalPage int       `json:"totalPage"`
	List      []Comment `json:"list"`
}

func (service *ListService) List() serializer.Response {
	comments := make([]Comment, 0)
	res, ok := cache.Get("comment")
	if ok {
		comments = res.([]Comment)
	}
	pageSize := 10
	totalPage := len(comments) / pageSize
	if len(comments)%pageSize != 0 {
		totalPage++
	}
	start := (service.PageNo - 1) * pageSize
	end := start + pageSize
	if end > len(comments) {
		end = len(comments)
	}
	return serializer.Response{
		Data: ListResponse{
			PageNo:    service.PageNo,
			PageSize:  pageSize,
			TotalPage: totalPage,
			List:      comments[start:end],
		},
	}
}
