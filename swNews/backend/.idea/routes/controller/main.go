package controller

import (
	"backend/pkg/serializer"
)

func ErrorResponse(err error) serializer.Response {

	return serializer.ParamErr("Invalid parameters.", err)
}
