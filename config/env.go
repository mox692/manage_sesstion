package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetENV() error {

	// デプロイの際にENVは指定しておく。
	if os.Getenv("ENV") == "production" {
		err := godotenv.Load("../prod.env")
		if err != nil {
			fmt.Printf("can't read .envfile")
			return err
		}
	} else {
		err := godotenv.Load("../local.env")
		if err != nil {
			fmt.Printf("can't read .envfile")
			return err
		}
	}

	return nil
}
