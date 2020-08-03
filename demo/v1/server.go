package main

import "github.com/m01i0ng/MEcho"

func main() {
  server := MEcho.NewServer("v1")
  server.Serve()
}
