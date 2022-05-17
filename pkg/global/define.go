package global

type HttpResponseCode int

const (
	OK HttpResponseCode = iota
	InputError
	JsonError
	DbError
	UnknownError
)

type HttpResponseMsg string

const (
	OKMsg         HttpResponseMsg = "ok"
	InputErrorMsg                 = "input error"
)

type HttpResponse struct {
	Code HttpResponseCode `json:"code"`
	Msg  HttpResponseMsg  `json:"msg"`
}
