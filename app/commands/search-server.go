package commands

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func SearchServer(ctx *cli.Context) {
	host := ctx.String("host")

	nameservers, err := net.LookupNS(host)

	if err != nil {
		log.Fatalf("Error => %v", err)
		return
	}

	fmt.Printf("HOST => '%s'\n", host)

	for _, nameserver := range nameservers {
		fmt.Printf("NAME SERVER => %s\n", nameserver.Host)
	}

}
