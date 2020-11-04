package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"session/config"
	"session/session"

	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

var Conn redis.Conn

type sessionServer struct {
	session.UnimplementedSessionServer
}

func (ss *sessionServer) SetSession(ctx context.Context, request *session.SessionRequest) (*session.SetStatus, error) {
	// redisへのセッション登録処理

	return &session.SetStatus{}, nil
}
func (ss *sessionServer) GetSession(ctx context.Context, request *session.SessionRequest) (*session.GetStatus, error) {
	// redisからのセッション取得処理

	return &session.GetStatus{}, nil
}

func runRedis() error {

	addr := os.Getenv("REDIS_ADDRESS")
	var err error
	Conn, err = redis.Dial("tcp", addr)
	if err != nil {
		return err
	}
	return nil
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
	err = runRedis()
	if err != nil {
		log.Fatal("err: %w", err)
	}

	grpcServer := grpc.NewServer()
	session.RegisterSessionServer(grpcServer, &sessionServer{})
	grpcServer.Serve(lis)

}

func checkENV() bool {
	return os.Getenv("ENV") == "local"
}
