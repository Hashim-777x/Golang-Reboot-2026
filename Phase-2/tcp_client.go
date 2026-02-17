package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

// peer

type Peer struct {
	conn  net.Conn
	msgCh chan []byte
}

func NewPree(conn net.Conn, msgCh chan []byte) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
	}
}

//	func (p *Peer) readLoop() {
//		for {
//			}
//		}
func (p *Peer) readLoop() {
	// 1. Create a buffer (a byte array slice)
	// We allocate 2048 bytes (2KB) of memory to hold incoming data.
	buf := make([]byte, 2048)

	for {
		// 2. Read from the connection into the buffer
		// This is a "BLOCKING" call. The Goroutine pauses here
		// until data arrives. It uses 0% CPU while waiting.
		n, err := p.conn.Read(buf)
		if err != nil {
			return // Exit the loop and close the goroutine if the peer disconnects
		}

		// 3. Process the data
		// We take a 'slice' of the buffer from index 0 to n (actual bytes read)
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgCh <- msgBuf
		// Log what we received
		//slog.Info("received message", "msg", string(msg), "bytes", n)
	}
}

// peer end
const defaultListnerAddress = ":5005"

type Config struct {
	ListnerAddress string
}

type Server struct {
	Config
	peer      map[*Peer]bool
	ln        net.Listener
	addPeerCh chan *Peer
	quickCh   chan struct{}
	msgCh     chan []byte
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListnerAddress) == 0 {
		cfg.ListnerAddress = defaultListnerAddress
	}
	return &Server{
		Config:    cfg,
		peer:      make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
		quickCh:   make(chan struct{}),
		msgCh:     make(chan []byte),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListnerAddress)
	if err != nil {
		return err
	}
	s.ln = ln

	go s.loop()
	slog.Info("server running ", "lsitenAddress", s.ListnerAddress)

	return s.acceptLoop()

}

func (s *Server) loop() {
	for {
		select {
		case rawMsg := <-s.msgCh:
			fmt.Println(rawMsg)

		case <-s.quickCh:
			return
		case peer := <-s.addPeerCh:
			s.peer[peer] = true
		}
	}
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPree(conn, s.msgCh)
	s.addPeerCh <- peer
	slog.Info("new peer connected", "remoteAddr", conn.RemoteAddr())
	peer.readLoop()
}

func main() {
	server := NewServer(Config{})
	log.Fatal(server.Start())
}
