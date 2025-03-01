// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: setting.proto

package server

import (
	"context"

	"fim_server/service/rpc/setting/internal/logic/setting"
	"fim_server/service/rpc/setting/internal/svc"
	"fim_server/service/rpc/setting/setting_rpc"
)

type SettingServer struct {
	svcCtx *svc.ServiceContext
	setting_rpc.UnimplementedSettingServer
}

func NewSettingServer(svcCtx *svc.ServiceContext) *SettingServer {
	return &SettingServer{
		svcCtx: svcCtx,
	}
}

func (s *SettingServer) SettingInfo(ctx context.Context, in *setting_rpc.Empty) (*setting_rpc.SettingInfoResponse, error) {
	l := settinglogic.NewSettingInfoLogic(ctx, s.svcCtx)
	return l.SettingInfo(in)
}
