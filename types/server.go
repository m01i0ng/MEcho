package types

type Server interface {
  Start()
  Stop()
  Serve()
}
