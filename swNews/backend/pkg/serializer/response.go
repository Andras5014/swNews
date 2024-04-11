package serializer

type Response struct {
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
