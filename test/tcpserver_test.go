package test

import (
	"learn-go/network"
	"testing"
)

func TestTcpServer(t *testing.T) {
	network.TcpServer("0.0.0.0:8081")
}
