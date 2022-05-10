package commands

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func SearchIps(ctx *cli.Context) {
	host := ctx.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatalf("Error => %v", err)
		return
	}

	fmt.Printf("HOST => '%s'\n", host)

	for _, ip := range ips {
		fmt.Printf("IP => %s\n", ip)
	}
}
