package mqs

import (
	"context"
	"encoding/json"
	"fim_server/models/log_models"
	"fim_server/service/api/log/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type LogEvent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *LogEvent {
	return &LogEvent{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Request struct {
	Type    int8   `json:"type"`
	IP      string `json:"ip"`
	UserID  uint   `json:"user_id"`
	Level   string `json:"level"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Service string `json:"service"`
}

func (l *LogEvent) Consume(ctx context.Context, key, val string) error {
	var req Request
	err := json.Unmarshal([]byte(val), &req)
	if err != nil {
		logx.Errorf("json解析错误  %s   %s", err.Error(), val)
		return nil
	}
	fmt.Println(req.IP)
	// 查ip对应的地址
	// 调用户基础方法，获取用户昵称
	var info = log_model.LogModel{
		Type:   req.Type,
		IP:     req.IP,
		UserID: req.UserID,
		// Addr:    addr.GetAddr(req.IP),
		Level:   req.Level,
		Title:   req.Title,
		Content: req.Content,
		Service: req.Service,
	}
	if req.UserID != 0 {
		baseInfo, err1 := l.svcCtx.UserRpc.UserBaseInfo(l.ctx, &user_rpc.UserBaseInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err1 == nil {
			info.Name = baseInfo.Name
			info.Avatar = baseInfo.Avatar
		}
	}
	// 判断是不是运行日志
	if info.Type == 3 {
		// 运行日志
		// 先查一下 今天这个服务有没有日志  有没有，有的话就更新，没有再创建
		mutex := sync.Mutex{}
		
		mutex.Lock()
		var logModel log_model.LogModel
		err = l.svcCtx.DB.Take(&logModel, "log_type = ? and service = ? and to_days(created_at) = to_days(now())", 3, info.Service).Error
		mutex.Unlock()
		if err == nil {
			// 找到了
			l.svcCtx.DB.Model(&logModel).Update("content", logModel.Content+"\n"+info.Content)
			logx.Infof("运行日志 %s 更新成功", req.Title)
			return nil
		}
	}
	mutex := sync.Mutex{}
	
	mutex.Lock()
	err = l.svcCtx.DB.Create(&info).Error
	mutex.Unlock()
	if err != nil {
		logx.Error(err)
		return nil
	}
	logx.Infof("日志 %s 保存成功", req.Title)
	
	return nil
}
