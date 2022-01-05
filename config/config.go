package config

import "github.com/joho/godotenv"

func Update() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
