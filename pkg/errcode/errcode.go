package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	errcode int      `json:"errcode"`
	errmsg  string   `json:"errmsg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{errcode: code, errmsg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.errcode
}

func (e *Error) Msg() string {
	return e.errmsg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.errmsg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
