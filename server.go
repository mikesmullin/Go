package main

import (
  "fmt"
  "log"
  "net"
  "bufio"
)

func main() {
  l, err := net.Listen("tcp", ":9339")
  if err != nil {
    log.Fatal(err)
  }
  defer l.Close()

  fmt.Println("server listening on localhost tcp/9339...")

  for {
    // wait for a connection
    conn, err := l.Accept()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Print(".")

    // Handle the connection in a new goroutine
    // the loop then returns to accepting, so that
    // multiple connections may be served concurrently.
    go func(c net.Conn) {
      defer func() {
        fmt.Println("closing socket.")
        c.Close()
      }()
      reader := bufio.NewReader(c)
      for {
        bytes, err := reader.ReadBytes(0)
        line := string(bytes)
        fmt.Println("line:", line)
        if err != nil {
          fmt.Println("err:", err)
          break
        }
        switch line[0] {
          case '<': // XML
            fmt.Println("xml")
          case '{': // JSON
            fmt.Println("JSON")
          default:
            panic("unrecognized format")
        }
      }
    }(conn)
  }
}

