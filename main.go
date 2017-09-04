package main

import (
	"log"

	"github.com/munnerz/metaldata/pkg/authorisation/ip"
	"github.com/munnerz/metaldata/pkg/registry/memory"
	"github.com/munnerz/metaldata/pkg/serve"
)

func main() {
	l := serve.NewListener(
		ip.NewIPAuthorisation(),
		memory.NewMemory(nil),
	)

	if err := l.Serve(); err != nil {
		log.Fatalf("error serving: %s", err.Error())
	}
}
