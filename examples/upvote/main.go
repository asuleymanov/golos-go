package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/asuleymanov/golos-go"
)

var (
	voter = ""
	key   = ""
)

func main() {
	cls, err := client.NewClient([]string{"wss://api.golos.cf", "wss://ws.golos.io"}, "work")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer cls.Close()

	cls.SetKeys(&client.Keys{PKey: []string{key}})

	if err := run(cls); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run(cls *client.Client) (err error) {
	flag.Parse()
	// Process args.
	args := flag.Args()

	if len(args) != 2 {
		return errors.New("2 arguments required")
	}
	author, permlink := args[0], args[1]

	fmt.Println(cls.Vote(voter, author, permlink, 10000))

	return nil
}
