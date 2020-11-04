package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"session/config"
	"session/kvs"
	"session/session"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

var (
	STATUS_OK  = "0"
	STATUS_BAD = "-1"
)

type sessionServer struct {
	session.UnimplementedSessionServer
}

func (ss *sessionServer) SetSession(ctx context.Context, request *session.SessionRequest) (*session.SetStatus, error) {

	// redisへのセッション登録処理
	statusID := request.StatusID
	userID := request.UserID

	err := kvs.Set(userID, statusID, kvs.Conn)
	if err != nil {
		return &session.SetStatus{SetStatusCode: STATUS_BAD}, xerrors.Errorf("kvs.Set err :%w", err)
	}
	return &session.SetStatus{SetStatusCode: STATUS_OK}, nil
}

func (ss *sessionServer) GetSession(ctx context.Context, request *session.SessionRequest) (*session.GetStatus, error) {
	// redisからのセッション取得処理
	userID := request.UserID

	status, err := kvs.Get(userID, kvs.Conn)
	if err != nil {
		return &session.GetStatus{GetStatusCode: STATUS_BAD}, xerrors.Errorf("kvs.Get err :%w", err)
	}
	return &session.GetStatus{GetStatusCode: status}, nil
}

func main() {

	// 環境変数のセット
	err := config.GetENV()
	if err != nil {
		log.Fatal("err: %w", err)
	}

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("err: %w", err)
	}
	fmt.Println("runnnig!")

	// redisの起動処理
	err = kvs.RunRedis()
	if err != nil {
		log.Fatal("err: %w", err)
	}

	grpcServer := grpc.NewServer()
	session.RegisterSessionServer(grpcServer, &sessionServer{})
	grpcServer.Serve(lis)
}
