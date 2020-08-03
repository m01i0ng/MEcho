package MEcho

import (
  "fmt"
  "net"

  "github.com/kataras/golog"
  "github.com/m01i0ng/MEcho/types"
)

type Server struct {
  Name      string
  IPVersion string
  IP        string
  Port      int
}

func (s *Server) Start() {
  golog.Infof("[start] server listening at tcp://%s:%d", s.IP, s.Port)

  go func() {
    addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
    if err != nil {
      golog.Errorf("[start] resolve tcp addr error: %v", err)
      return
    }

    l, err := net.ListenTCP(s.IPVersion, addr)
    if err != nil {
      golog.Errorf("[start] listen error: %v", err)
      return
    }

    golog.Infof("[start] start server %s ok", s.Name)

    for {
      conn, err := l.AcceptTCP()
      if err != nil {
        golog.Errorf("[start] accept error: %v", err)
        continue
      }

      go func() {
        for {
          buf := make([]byte, 512)
          n, err := conn.Read(buf)
          if err != nil {
            golog.Errorf("[start] recv buf error: %v", err)
            continue
          }

          if _, err := conn.Write(buf[:n]); err != nil {
            golog.Errorf("[start] write back error: %v", err)
            continue
          }
        }
      }()
    }
  }()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
  s.Start()

  select {}
}

func NewServer(name string) types.Server {
  s := &Server{
    Name:      name,
    IPVersion: "tcp4",
    IP:        "0.0.0.0",
    Port:      8888,
  }

  return s
}
