// Code generated by goctl. DO NOT EDIT!
// Source: add.proto

package server

import (
	"context"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/logic"
	"bookstore/rpc/add/internal/svc"
)

type AdderServer struct {
	svcCtx *svc.ServiceContext
}

func NewAdderServer(svcCtx *svc.ServiceContext) *AdderServer {
	return &AdderServer{
		svcCtx: svcCtx,
	}
}

func (s *AdderServer) Ping(ctx context.Context, in *add.Request) (*add.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *AdderServer) Add(ctx context.Context, in *add.AddReq) (*add.AddResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}
