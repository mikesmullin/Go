package main

import (
  "fmt"
//  "io"
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
      scanner := bufio.NewScanner(c)
      for scanner.Scan() {
        fmt.Println("cmd: ", scanner.Text());
      }
      if err := scanner.Err(); err != nil {
        fmt.Println("reading input:", err)
      }
      //// echo all incoming data.
      //io.Copy(c, c)
      // shut down the connection
      fmt.Println("closing socket.")
      c.Close()
    }(conn)
  }
}

