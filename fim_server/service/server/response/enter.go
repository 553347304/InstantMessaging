package response

import (
	"fim_server/utils/stores/logs"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := &Body{
			Code:    0,
			Message: "ok",
			Data:    resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	logs.Info("-------------------------------")

	// 错误返回
	httpx.WriteJson(w, http.StatusOK, &Body{
		Code:    7,
		Message: err.Error(),
		Data:    nil,
	})

}
