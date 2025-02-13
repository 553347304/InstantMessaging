package logs

import (
	"fim_server/utils/stores/https"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type request struct {
	Error  string
	Method string
	URL    string
	Body   json.RawMessage // 原始数据
	Header map[string][]string
}

func GinRequest(c *gin.Context, err string) {
	r := request{
		Error:  err,
		Method: c.Request.Method,
		URL:    c.Request.URL.String(),
		Body:   nil,
		Header: c.Request.Header,
	}

	data, _ := c.GetRawData()
	if len(data) != 0 {
		r.Body = data // json 数据
	}

	https.Response(c).Error(r)
	//marshal, _ := json.MarshalIndent(r, "", "\t") // 转格式化后json
	//string(marshal)
}
