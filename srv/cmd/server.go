package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mrlreable/console-chat-app/cfg"
	"github.com/mrlreable/console-chat-app/internal/types"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Server running\n")
}

func newServer(maxConn int, address string) *types.Server {
	return &types.Server{
		MaxConn:   maxConn,
		Address:   address,
		Clients:   make(map[*types.Client]bool),
		Broadcast: make(chan []byte),
	}
}

func serve(server *types.Server) error {
	http.HandleFunc("/ws", handler)

	err := http.ListenAndServe(server.Address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

func main() {
	var config types.Config

	cfg.SetConfigName("config")
	cfg.SetConfigType(cfg.Yaml)

	err := cfg.NewConfig(&config)
	if err != nil {
		log.Fatal("NewConfig: ", err)
	}

	fmt.Printf("Read in config: %+v\n", config)

	host := config.Server.Host + ":" + strconv.Itoa(config.Server.Port)
	srv := newServer(config.Server.MaxConn, host)

	err = serve(srv)
	if err != nil {
		log.Fatal("serve: ", err)
	}
}
