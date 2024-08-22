package main

import (
	"flag"
	"fmt"
	"isis-net"
)

func main() {
	const nsel = "00"

	ap := flag.String("area-prefix", "49", "area prefix")
	area := flag.String("area", "0001", "area")
	ip := flag.String("ip", "172.31.255.1", "ip")
	flag.Parse()

	if *ap == "" || *area == "" || *ip == "" {
		flag.Usage()
	}

	ipAddr := isis.Validate(*ip)
	net := isis.ConvertToNET(ipAddr)

	fmt.Printf("NET: %s.%s.%s.%s", *ap, *area, net, nsel)
}