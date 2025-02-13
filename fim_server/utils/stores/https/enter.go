package https

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type httpResponse struct {
	Response *http.Response
	Body     []byte // 解析json:json.Unmarshal(r.Body, &map[string]any{}) | 路径参数:url.ParseQuery(string(r.Body))
	Error    interface{}
}

type setHttp struct {
	Method string
	URL    string
	Body   io.Reader
	Header map[string]string
}

func Form(form map[string]any) string {
	// 设置查询参数
	var value string
	for key, v := range form {
		value += fmt.Sprintf("&%s=%v", key, v)
	}
	return value[1:]
}
func newRequest(h setHttp) (*http.Response, []byte, interface{}) {

	// 设置请求体
	request, err1 := http.NewRequest(h.Method, h.URL, h.Body)
	if err1 != nil {
		return nil, nil, "设置请求体错误 ->" + err1.Error()
	}

	// 设置请求头
	if h.Header != nil {
		for key, v := range h.Header {
			request.Header.Set(key, v)
		}
	}

	// 执行请求
	response, err2 := http.DefaultClient.Do(request)
	if err2 != nil {
		return nil, nil, "请求错误 -> " + err2.Error()
	}
	if response.StatusCode != 200 {
		return nil, nil, fmt.Sprint("状态码错误 -> ", response.StatusCode)
	}

	// 读取数据
	read, err3 := io.ReadAll(response.Body)
	if err3 != nil {
		return nil, nil, "读取错误 -> " + err3.Error()
	}

	response.Body = io.NopCloser(bytes.NewBuffer(read)) // 重新写入数据

	return response, read, nil
}
