package main

import (
	"context"
	"session/config"
	"session/kvs"
	"session/session"
	"testing"
)

func Test_SetSessionAndGetSession(t *testing.T) {

	server := &sessionServer{}

	// 環境変数を設定
	err := config.GetENV()
	if err != nil {
		t.Fatal(err)
	}

	// redisを起動
	err = kvs.RunRedis()
	if err != nil {
		t.Fatal(err)
	}

	setRequests := []session.SessionRequest{
		{
			StatusID: "fas",
			UserID:   "fasd",
		},
		{
			StatusID: "ffdsas",
			UserID:   "fasadsd",
		},
		{
			StatusID: "2.3",
			UserID:   "motoyuki",
		},
		{
			StatusID: "2.2",
			UserID:   "yuki",
		},
	}
	ctx := context.Background()

	for _, request := range setRequests {
		status, err := server.SetSession(ctx, &request)

		if err != nil {
			t.Errorf("error: %w", err)
		}
		if status.StatusCode != 200 {
			t.Errorf("bad status: want %d, get %d", 200, status.StatusCode)
		}
	}

	getRequests := []session.SessionRequest{
		{
			UserID: "fasd",
		},
		{
			UserID: "fasadsd",
		},
		{
			UserID: "motoyuki",
		},
		{
			UserID: "yuki",
		},
	}

	expects := []string{
		"fas",
		"ffdsas",
		"2.3",
		"2.2",
	}

	for i, request := range getRequests {
		status, err := server.GetSession(ctx, &request)

		if err != nil {
			t.Errorf("err: %v", err)
		}
		if status.Data != expects[i] {
			t.Errorf("GetSession err: expect %s ,but got %s\n", expects[i], status.Data)
		}
		if status.StatusCode != 200 {
			t.Errorf("GetSession err: expect %d ,but got %d\n", 200, status.StatusCode)
		}
	}
}

func Test_GetSessionNilReturn(t *testing.T) {

	server := &sessionServer{}

	// 環境変数を設定
	err := config.GetENV()
	if err != nil {
		t.Fatal(err)
	}

	// redisを起動
	err = kvs.RunRedis()
	if err != nil {
		t.Fatal(err)
	}

	// 存在しないkeyでリクエスト
	request := session.SessionRequest{UserID: "rwqekjlqew43"}

	ctx := context.Background()
	getStatus, err := server.GetSession(ctx, &request)

	if err != nil {
		t.Errorf("%+v", err)
	}
	if getStatus.StatusCode != 500 {
		t.Errorf("expect %d, but got %s", 500, getStatus)
	}
	if getStatus.ErrMessage != kvs.NilReturn {
		t.Errorf("expect %s, but got %s", kvs.NilReturn, getStatus.ErrMessage)
	}
}
