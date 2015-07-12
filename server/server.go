// server.go

package server

import (
  "fmt"
  "git.apache.org/thrift.git/lib/go/thrift"
  "github.com/mdennebaum/pelican/user"
)

type PelicanServer struct {
  host             string
  handler          *UserHandler
  processor        *user.UserSvcProcessor
  transport        *thrift.TServerSocket
  transportFactory thrift.TTransportFactory
  protocolFactory  *thrift.TBinaryProtocolFactory
  server           *thrift.TSimpleServer
}

func NewPelicanServer(host string) *PelicanServer {
    handler := NewUserHandler()
    processor := user.NewUserSvcProcessor(handler)
    transport, err := thrift.NewTServerSocket(host)
    if err != nil {
        panic(err)
    }

    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
    return &PelicanServer{
        host:             host,
        handler:          handler,
        processor:        processor,
        transport:        transport,
        transportFactory: transportFactory,
        protocolFactory:  protocolFactory,
        server:           server,
    }
}

func (ps *PelicanServer) Run() {
    fmt.Printf("server listening on %s\n", ps.host)
    ps.server.Serve()
}

func (ps *PelicanServer) Stop() {
    fmt.Println("stopping server...")
    ps.server.Stop()
}