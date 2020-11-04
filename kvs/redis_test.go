package kvs

import (
	"fmt"
	"session/config"
	"testing"
)

func TestRunRedis(t *testing.T) {
	// 環境変数を設定
	err := config.GetENV()
	if err != nil {
		t.Fatal(err)
	}
	err = RunRedis()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SetandGet(t *testing.T) {
	// 環境変数を設定
	err := config.GetENV()
	if err != nil {
		t.Fatal(err)
	}

	// redisを起動
	err = RunRedis()
	if err != nil {
		t.Fatal(err)
	}

	err = Set("gdweras", "2_4", Conn)
	if err != nil {
		t.Errorf("set error: %v", err)
	}

	res, err := Get("gdweras", Conn)
	if err != nil {
		t.Errorf("get error: %v", err)
	}
	fmt.Printf("res: %v\n", res)
}
