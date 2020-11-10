package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
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
	statusID := request.StatusID
	userID := request.UserID

	err := kvs.Set(userID, statusID, kvs.Conn)
	if err != nil {
		return &session.SetStatus{SetStatusCode: STATUS_BAD}, xerrors.Errorf("kvs.Set err :%w", err)
	}
	ss.success(ctx, request)
	return &session.SetStatus{SetStatusCode: STATUS_OK}, nil
}

func (ss *sessionServer) GetSession(ctx context.Context, request *session.SessionRequest) (*session.GetStatus, error) {
	userID := request.UserID
	status, err := kvs.Get(userID, kvs.Conn)

	if err != nil {
		return &session.GetStatus{GetStatusCode: STATUS_BAD}, xerrors.Errorf("kvs.Get err :%w", err)
	}
	ss.success(ctx, request)
	return &session.GetStatus{GetStatusCode: status}, nil
}

func (ss *sessionServer) success(ctx context.Context, request *session.SessionRequest) {
	// *******************Todo: リクエスト時のログ処理
	log.Println("get request.")
}

func main() {

	// 環境変数のセット
	err := config.GetENV()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	// redisの起動処理
	err = kvs.RunRedis()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server runnnig!(port: %s)\n", os.Getenv("REDIS_PORT"))
	grpcServer := grpc.NewServer()
	session.RegisterSessionServer(grpcServer, &sessionServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
