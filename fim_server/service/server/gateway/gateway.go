package main

import (
	"encoding/json"
	"fim_server/utils/src/etcd"
	"fim_server/utils/stores/logs"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type Proxy struct{}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		UserId uint `json:"user_id"`
		Role   int  `json:"role"`
	} `json:"data"`
}

func WriteJson(res http.ResponseWriter, message string) {
	byteData, _ := json.Marshal(Data{
		Code:    7,
		Message: message,
	})
	res.Write(byteData)
}

func auth(req *http.Request) error {
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)

	// 认证请求
	authRequest, _ := http.NewRequest("POST", authUrl, nil)
	authRequest.Header = req.Header

	// 获取查询参数token
	token := req.URL.Query().Get("token")
	if token != "" {
		authRequest.Header.Set("Token", token)
	}

	authRequest.Header.Set("ValidPath", req.URL.Path)
	authResult, err := http.DefaultClient.Do(authRequest)
	if err != nil {
		return logs.Error("认证服务错误" + err.Error())
	}

	var authResponse Data
	byteData, _ := io.ReadAll(authResult.Body)
	err = json.Unmarshal(byteData, &authResponse)
	if err != nil {
		return logs.Error("json.Unmarshal" + err.Error())
	}
	if authResponse.Code != 0 {
		return logs.Error("认证不通过", string(byteData))
	}

	// 设置请求头
	if authResponse.Data != nil {
		req.Header.Set("User-Id", fmt.Sprint(authResponse.Data.UserId))
		req.Header.Set("Role", fmt.Sprint(authResponse.Data.Role))
	}

	return nil
}

func (Proxy) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// 匹配请求前缀  /api/user/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("err"))
		return
	}
	service := addrList[1]

	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		WriteJson(res, logs.Error("不匹配的服务", service).Error())
		return
	}

	proxyUrl := fmt.Sprintf("http://%s", addr) // 请求认证服务地址
	logs.Info(fmt.Sprintf("%s %s -> %s%s ", req.RemoteAddr, service, proxyUrl, req.URL.String()))

	// 请求认证服务地址
	err := auth(req)
	if err != nil {
		WriteJson(res, err.Error())
		return
	}

	// 反向代理
	remote, _ := url.Parse(proxyUrl)
	reverseProxy := httputil.NewSingleHostReverseProxy(remote)
	reverseProxy.ServeHTTP(res, req)
}

var configFile = flag.String("f", "gateway.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
}

var config Config

func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &config)
	logs.Info("gateway running", config.Addr)
	http.ListenAndServe(config.Addr, Proxy{})
}
