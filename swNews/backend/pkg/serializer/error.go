package serializer

import (
	"backend/pkg/logger"
)

const (
	CodeSuccess      = 0
	CodeError        = 1
	CodeParamErr     = 2
	CodeDBErr        = 3
	CodeTokenErr     = 4
	CodeAuthErr      = 5
	CodeUserNotFound = 6
)

func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "Invalid parameters."
	}
	return Err(CodeParamErr, msg, err)
}

func AppError(msg string, err error) Response {
	return Err(CodeError, msg, err)
}

func Err(code int, msg string, err error) Response {
	logger.Error(msg, err)
	//fmt.Println(err)
	return Response{
		Code: code,
		Msg:  msg,
	}
}
