package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kidoman/go-steam"
)

// CLI Flags
var debug bool
var command string

// ENV Vars
var addr string
var pass string

func init() {
	flag.Usage = func() {
	    fmt.Fprintf(os.Stderr, "RCON CLI usage: rcon [-debug] command\n")
	    flag.PrintDefaults()
	    os.Exit(2)
	}

	flag.BoolVar(&debug, "debug", false, "Show debug output. Useful for debugging.")
	flag.Parse()

	if debug {
		steam.SetLog(log.New())
	}

	command = flag.Arg(0)

	if command == "" {
		command = "status"
	}

	addr = os.Getenv("ADDR")
	pass = os.Getenv("RCON_PASSWORD")

	if addr == "" || pass == "" {
		log.Fatal("Please set ADDR & RCON_PASSWORD.")
	}
}

func main() {
	var retries = 0
	var max_retries = 5

	for {
		o := &steam.ConnectOptions{RCONPassword: pass}
		rcon, err := steam.Connect(addr, o)
		if err != nil {
			if retries > max_retries {
				log.Fatal("Maximum retry limit exceeded. Please verify your configuration and try again later.")
			}
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			retries++
			continue
		}

		defer rcon.Close()

		for {
			resp, err := rcon.Send(command)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println(resp)
			time.Sleep(5 * time.Second)
		}
	}
}
