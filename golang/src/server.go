package main

import "rpcServer"
import "log"
import "net/rpc"
import "net/http"
import "net"

func main() {
    arith := new(rpcServer.Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", "127.0.0.1:1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}
