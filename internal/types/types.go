package types

import "github.com/gorilla/websocket"

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	MaxConn int    `yaml:"maxConnections"`
}

type Client struct {
	Srv *Server

	// Websocket connection
	Conn *websocket.Conn

	// Channel for outbound messages
	Send chan []byte
}

type Server struct {
	// Registered clients
	Clients map[*Client]bool

	Address string

	// Capacity of concurrent client connections
	MaxConn int

	// Broadcast channel
	Broadcast chan []byte
}
