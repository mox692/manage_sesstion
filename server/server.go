package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"session/kvs"
	"session/session"

	"github.com/gomodule/redigo/redis"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

type sessionServer struct {
	session.UnimplementedSessionServer
}

func (ss *sessionServer) SetSession(ctx context.Context, request *session.SessionRequest) (*session.SetStatus, error) {
	log.Printf("get reqest!!\n")
	statusID := request.StatusID
	userID := request.UserID

	err := kvs.Set(userID, statusID, kvs.Conn)
	if err != nil {
		return &session.SetStatus{StatusCode: 500}, xerrors.Errorf("kvs.Set err :%w", err)
	}
	ss.success(ctx, request)
	return &session.SetStatus{StatusCode: 200}, nil
}

func (ss *sessionServer) GetSession(ctx context.Context, request *session.SessionRequest) (*session.GetStatus, error) {
	userID := request.UserID
	data, err := kvs.Get(userID, kvs.Conn)

	// redisのNilReturn エラーだけは別途処理。
	if unwrapErr := errors.Unwrap(err); unwrapErr == redis.ErrNil {
		return &session.GetStatus{StatusCode: 500, ErrMessage: kvs.NilReturn}, nil
	}
	if err != nil {
		return &session.GetStatus{StatusCode: 500, ErrMessage: err.Error()}, err
	}
	ss.success(ctx, request)
	return &session.GetStatus{StatusCode: 200, Data: data}, nil
}

func (ss *sessionServer) success(ctx context.Context, request *session.SessionRequest) {
	// *******************Todo: リクエスト時のログ処理
	log.Println("get request! No err!")
}

func main() {

	// 環境変数のセット
	// err := config.GetENV()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	// redisの起動処理
	err = kvs.RunRedis()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("redis server runnnig!(port: %s)\n", os.Getenv("REDIS_PORT"))
	grpcServer := grpc.NewServer()
	session.RegisterSessionServer(grpcServer, &sessionServer{})

	fmt.Printf("gRPC server runnnig!(port: %s)\n", os.Getenv("SERVER_PORT"))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
