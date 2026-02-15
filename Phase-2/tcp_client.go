package main

import (
	"log"
	"log/slog"
	"net"
)

// peer

type Peer struct {
	conn net.Conn
}

func NewPree(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (p *Peer) readLoop() {
	for {

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
	peer := NewPree(conn)
	s.addPeerCh <- peer
	slog.Info("new peer connected", "remoteAddr", conn.RemoteAddr())
	peer.readLoop()
}

func main() {
	server := NewServer(Config{})
	log.Fatal(server.Start())
}
