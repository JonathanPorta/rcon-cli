package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/kidoman/go-steam"
)

func main() {
	debug := flag.Bool("debug", false, "debug")
	flag.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		fmt.Println("Please set ADDR.")
		return
	}
	server, err := steam.Connect(addr)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	ping, err := server.Ping()
	if err != nil {
		fmt.Printf("steam: could not ping %v: %v\n", addr, err)
		return
	}
	fmt.Printf("steam: ping to %v: %v\n", addr, ping)
	info, err := server.Info()
	if err != nil {
		fmt.Printf("steam: could not get server info from %v: %v\n", addr, err)
		return
	}
	fmt.Printf("steam: info of %v: %v\n", addr, info)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
