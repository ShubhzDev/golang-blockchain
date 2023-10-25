package main

import "fmt"
import "github.com/shubzdev/golang-blockchain/network"

func main(){
	trLocal := network.NewLocalTransport("LOCAL")
	opts := network.ServerOpts{
		Transposrts: []network.Transport(trLocal)
	}

	s:= network.NewServer(opts)
	s.Start()
}