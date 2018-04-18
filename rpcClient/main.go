package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpcprocs"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	multiplyArgs := &rpcprocs.Args{7, 8}
	var multiplyReply int
	err = client.Call("Arith.Multiply", multiplyArgs, &multiplyReply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d", multiplyArgs.A, multiplyArgs.B, multiplyReply)

	divideArgs := &rpcprocs.Args{10, 5}
	var divideReply rpcprocs.Quotient
	err = client.Call("Arith.Divide", divideArgs, &divideReply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("\n Quo: %d/%d = %d \n Rem: %d %% %d = %d", divideArgs.A, divideArgs.B, divideReply.Quo, divideArgs.A, divideArgs.B, divideReply.Rem)
}
