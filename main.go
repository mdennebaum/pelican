// main.go

package main

import "github.com/mdennebaum/pelican/server"

func main() {
    host := "localhost:9090"
    server := server.NewPelicanServer(host)
    server.Run()
}