package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Error: Illegal number of arguments")
	}

	domain := os.Args[1]
	cname, addrs, err := net.LookupSRV("", "", domain)

	if err != nil {
		return
	}

	s := ""
	for _, addr := range addrs {
		s += fmt.Sprintf(
			"%s 0 IN SRV %d %d %d %s\n",
			cname, addr.Priority, addr.Weight, addr.Port, addr.Target)
	}

	fmt.Print(s)
}
