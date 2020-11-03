package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"session/session"

	"google.golang.org/grpc"
)

type sessionServer struct {
	session.UnimplementedSessionServer
}

func (ss *sessionServer) SetSession(ctx context.Context, request *session.SessionRequest) (*session.SetStatus, error) {
	// redisへのセッション登録処理
}
func (ss *sessionServer) GetSession(ctx context.Context, request *session.SessionRequest) (*session.GetStatus, error) {
	// redisからのセッション取得処理
}

func newServer() {

}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("err: %w", err)
	}
	fmt.Println("runnnig!")

	// redisの起動処理

	grpcServer := grpc.NewServer()
	session.RegisterSessionServer(grpcServer, &sessionServer{})
	grpcServer.Serve(lis)

}
