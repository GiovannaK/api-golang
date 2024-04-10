package main

import "github.com/GiovannaK/api-golang/configs"

func main() {
	cfg, err := configs.LoadConfig("configs")

	if err != nil {
		panic(err)
	}
}
