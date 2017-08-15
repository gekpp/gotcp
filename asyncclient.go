package gotcp

import (
	"net"
	"sync"
	"time"
)

type AsyncClient struct {
	server *Server
}

// NewServer creates a server
func NewAsyncClient(config *Config, callback ConnCallback, protocol Protocol) *AsyncClient {
	return &AsyncClient{&Server{
		config:    config,
		callback:  callback,
		protocol:  protocol,
		exitChan:  make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}}
}

// Start starts service
func (c *AsyncClient) Start(conn *net.TCPConn, acceptTimeout time.Duration) {

	select {
	case <-c.server.exitChan:
		return

	default:
	}

	newConn(conn, c.server).Do()
}

// Stop stops service
func (s *AsyncClient) Stop() {
	close(s.server.exitChan)
	s.server.waitGroup.Wait()
}
