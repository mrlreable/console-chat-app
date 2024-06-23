package main

import (
	"fmt"
	"log"

	"github.com/mrlreable/console-chat-app/cfg"
	"github.com/mrlreable/console-chat-app/internal/types"
)

type Client struct {
	Username string
}

func main() {
	var config types.Config

	cfg.SetConfigName("config")
	cfg.SetConfigType(cfg.Yaml)

	err := cfg.NewConfig(&config)
	if err != nil {
		log.Fatalf("NewConfig: %v", err)
	}

	fmt.Printf("%+v", config)
}
