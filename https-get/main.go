package main // import "github.com/ContainerSolutions/lean-go-containers/https-get"

import (
  "io"
  "log"
  "net/http"
  "os"
)

func main() {
  resp, err := http.Get("https://www.ietf.org/rfc/rfc2795.txt")
  if err != nil {
    log.Fatal(err)
  } else {
    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
  }
}
