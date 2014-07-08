envar
=====

Golang flag-style environment variable parser


Example Usage
-------------

```
package main

import (
  "fmt"
  "io"
  "log"
  "net"

  "github.com/partkyle/envar"
)

var (
  host string
  port int
)

func init() {
  envar.StringVar(&host, "ECHO_HOST", "localhost", "host to bind to")
  envar.IntVar(&port, "ECHO_PORT", 9000, "port to bind to")

  envar.Parse()
}

func main() {
  listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
  if err != nil {
    log.Fatalf("err listening: %s")
  }

  log.Printf("listening on %s", listener.Addr())

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatalf("err accept: %s", err)
    }

    go func(conn net.Conn) {
      io.Copy(conn, conn)
    }(conn)
  }
}

```
