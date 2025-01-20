package log_service

import (
	"bytes"
	"context"
	"fim_server/models/log_models"
	"fim_server/models/user_models"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type PusherServerInterface interface {
	Response(http.ResponseWriter, *http.Request, []byte) string
	Info(context.Context, string, string)
}
type pusherServer struct {
	UserId  string `json:"user_id"`
	IP      string `json:"ip"`
	Type    string `json:"type"`
	Level   string `json:"level"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Service string `json:"service"`
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewPusher(serviceName, _type string) PusherServerInterface {
	return &pusherServer{
		Service: serviceName,
		Type:    _type,
	}
}

func (p *pusherServer) Response(w http.ResponseWriter, r *http.Request, body []byte) string {
	_method := r.Method
	url := r.URL.String()
	header := string(conv.Json().Marshal(r.Header))
	request, _ := io.ReadAll(r.Body)
	response := string(body)
	r.Body = io.NopCloser(bytes.NewBuffer(request)) // 重新写入数据
	return fmt.Sprintf(""+
		`<div class="code">
			<div class="request">
				<span class="method">%s</span>
				<span class="url">%s</span>
				<pre class="header">%s</pre>
				<pre class="body">%s</pre>
			</div>
			<div class="response">
				<pre class="body">%s</pre>
			</div>
		</div>`, _method, url, header, request, response)
}
func (p *pusherServer) Info(ctx context.Context, title, content string) {
	p.Save(ctx, "info", title, content)
}
func (p *pusherServer) Error(ctx context.Context, title, content string) {
	p.Save(ctx, "error", title, content)
}
func (p *pusherServer) Warn(ctx context.Context, title, content string) {
	p.Save(ctx, "warn", title, content)
}
func (p *pusherServer) Save(ctx context.Context, level, title, content string) {
	p.IP = ctx.Value("ip").(string)
	p.UserId = ctx.Value("user_id").(string)
	p.Level = level
	p.Title = title
	p.Content = content
	addr := method.Addr().GetAddr(p.IP)
	userId := conv.Type(p.UserId).Uint()
	
	mutex := sync.Mutex{}
	mutex.Lock()
	var user user_models.UserModel
	logs.Info(userId)
	DB.Take(&user,  userId)
	DB.Create(log_models.LogModel{
		UserId:  userId,
		Name:    user.Name,
		Avatar:  user.Avatar,
		IP:      p.IP,
		Addr:    addr,
		Service: p.Service,
		Type:    p.Type,
		Level:   p.Level,
		Title:   p.Title,
		Content: p.Content,
	})
	mutex.Unlock()
}
