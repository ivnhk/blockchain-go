package main

import (
	"fmt"
	"github.com/ivnhk/blockchain-go/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}
	s := network.NewServer(opts)
	fmt.Println(s)
}
