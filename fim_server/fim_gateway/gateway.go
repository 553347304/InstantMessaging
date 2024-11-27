package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fim_server/common/etcd"
	"fim_server/utils/stores/logs"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
}

var config Config

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		UserId uint `json:"userId"`
		Role   int  `json:"role"`
	} `json:"data"`
}

func Json(res http.ResponseWriter, data Data) {
	logs.Error(data.Message)
	byteData, _ := json.Marshal(data)
	res.Write(byteData)
}

func auth(req *http.Request) error {
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	authRequest, _ := http.NewRequest("POST", authUrl, nil)
	authRequest.Header = req.Header
	authRequest.Header.Set("ValidPath", req.URL.Path)
	authResult, err := http.DefaultClient.Do(authRequest)
	if err != nil {
		return logs.Error("认证服务错误" + err.Error())
	}

	var authResponse Data
	byteData, _ := io.ReadAll(authResult.Body)
	err = json.Unmarshal(byteData, &authResponse)
	if err != nil {
		return logs.Error("解析失败" + err.Error())
	}
	if authResponse.Code != 0 { // 认证失败
		return logs.Error("认证失败")
	}

	if authResponse.Data != nil {
		req.Header.Set("User-Id", fmt.Sprint(authResponse.Data.UserId))
		req.Header.Set("Role", fmt.Sprint(authResponse.Data.Role))
	}

	return nil
}

func proxy(url string, body io.Reader, res http.ResponseWriter, req *http.Request, ser string) error {

	proxyReq, err := http.NewRequest(req.Method, url, body)
	if err != nil {
		return errors.New(err.Error())
	}

	proxyReq.Header = req.Header
	proxyReq.Header.Del("ValidPath")

	proxyResponse, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return errors.New("服务异常" + err.Error())
	}
	_, err = io.Copy(res, proxyResponse.Body)
	if err != nil {
		return errors.New("io copy 失败" + err.Error())
	}

	logs.Info(ser, proxyReq.Header)
	return nil
}

func gateway(res http.ResponseWriter, req *http.Request) {
	newBody := io.NopCloser(req.Body)     // 防止req.Body被读取两次
	reqByteData, _ := io.ReadAll(newBody) // 读取请求体
	body := bytes.NewBuffer(reqByteData)  // 将请求体重新写入

	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		Json(res, Data{Code: 7, Message: "服务错误"})
		return
	}
	var service = addrList[1]

	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		Json(res, Data{Code: 7, Message: "不匹配的服务:" + service})
		return
	}

	remoteAddr := strings.Split(req.RemoteAddr, ":")
	if len(remoteAddr) != 2 {
		Json(res, Data{Code: 7, Message: "服务错误"})
		return
	}

	url := fmt.Sprintf("http://%s%s", addr, req.URL.String())
	ser := fmt.Sprintf("%s %s -> %s ", remoteAddr[0], service, url)

	// 请求认证服务地址
	err := auth(req)
	if err != nil {
		Json(res, Data{Code: 7, Message: err.Error()})
		return
	}

	// 转发到实际服务上
	err = proxy(url, body, res, req, ser)
	if err != nil {
		Json(res, Data{Code: 7, Message: err.Error()})
		return
	}

}

// cd fim_gateway
// go run gateway.go
func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &config)

	http.HandleFunc("/", gateway)

	fmt.Printf("gateway running http://%s\n", config.Addr)

	http.ListenAndServe(config.Addr, nil)

}
