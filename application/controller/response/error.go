package response

import (
	"encoding/json"
	"net/http"
)

type errRes struct {
	Code    int    `json:"code"`
	Result  string `json:"result"`
	Message string `json:"message"`
}

/*
ServerErrorRender はサーバエラーを返す.

Example Response:
	{
		"code": 503,
		"result": "internal server error",
	}
*/
func ServerErrorRender() (int, []byte) {
	code := http.StatusServiceUnavailable // 503
	errRes := &errRes{
		Code:   code,
		Result: "internal server error",
	}
	res, _ := json.Marshal(errRes)
	return code, res
}

/*
ValidationErrorRender はバリデーションのエラーを返す.

Example Response:
	{
		"code": 422,
		"result": "validation error",
		"message": "商品IDは必須です"
	}
*/
func ValidationErrorRender(m string) (int, []byte) {
	code := http.StatusUnprocessableEntity // 422
	errRes := &errRes{
		Code:    code,
		Message: "validation error",
		Result:  m,
	}
	res, err := json.Marshal(errRes)
	if err != nil {
		// FIXME: ログを出す
		return ServerErrorRender()
	}
	return code, res
}

/*
NoRecordErrorRender はレコードがない場合のエラーを返す.

Example Response:
	{
		"code": 404,
		"result": "no record error",
		"message": "商品が存在しません"
	}
*/
func NoRecordErrorRender(m string) (int, []byte) {
	code := http.StatusNotFound // 404
	errRes := &errRes{
		Code:    code,
		Message: "no record error",
		Result:  m,
	}
	res, err := json.Marshal(errRes)
	if err != nil {
		// FIXME: ログを出す
		return ServerErrorRender()
	}
	return code, res
}
