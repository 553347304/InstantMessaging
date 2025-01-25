package method

import (
	"fim_server/utils/stores/logs"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpResponse struct {
	Response *http.Response
	Body     []byte // 解析json:json.Unmarshal(r.Body, &map[string]any{}) | 路径参数:url.ParseQuery(string(r.Body))
	Error    error
}

type httpServerInterface interface {
	Get(map[string]any) httpResponse
	Post(map[string]any) httpResponse
}
type httpServer struct {
	url string
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Http(url string) httpServerInterface {
	return &httpServer{url: url}
}

func (l *httpServer) param(params map[string]any) string {
	if len(params) == 0 {
		return ""
	}
	var value string
	for key, v := range params {
		value += fmt.Sprintf("&%s=%v", key, v)
	}
	return value[1:]
}

func (l *httpServer) Get(params map[string]any) httpResponse {
	response, err := http.Get(l.url + "?" + l.param(params))
	if err != nil {
		return httpResponse{Error: logs.Error(err)}
	}
	defer response.Body.Close()
	read, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return httpResponse{Error: logs.Error(err)}
	}
	return httpResponse{Response: response, Body: read, Error: nil}
}

func (l *httpServer) Post(params map[string]any) httpResponse {
	response, err := http.Post(l.url + "?" + l.param(params),"application/json",strings.NewReader(""))
	if err != nil {
		return httpResponse{Error: logs.Error(err)}
	}
	defer response.Body.Close()
	read, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return httpResponse{Error: logs.Error(err)}
	}
	return httpResponse{Response: response, Body: read, Error: nil}
}
