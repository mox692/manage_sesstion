package config

import (
	"os"
	"testing"
)

func Test_GetENV(t *testing.T) {
	err := GetENV()
	if err != nil {
		t.Errorf("error!: %s\n", err)
	}

	cases := []string{
		"REDIS_HOST",
		"REDIS_PORT",
	}

	want := []string{
		"127.0.0.1",
		"6379",
	}

	for i, v := range cases {
		if os.Getenv(v) != want[i] {
			t.Errorf("GetENV error: want %s, but get %s\n", want[i], os.Getenv(v))
		}
	}
}
