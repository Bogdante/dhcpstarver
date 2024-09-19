package main

import (
	"log"
	"time"

	"github.com/Bogdante/dhcpstarver/config"
	"github.com/Bogdante/dhcpstarver/dhcp"
	"github.com/Bogdante/dhcpstarver/networking"
)

func main() {
	client, err := networking.CreateNewClient()
	defer client.CloseConnection()

	if err != nil {
		panic("Cant Create Client")
	}

	for i := config.IP_POOL_START_LAST_BYTE; i <= config.IP_POOL_END_LAST_BYTE; i++ {
		pack, err := dhcp.CreateDhcpRequestPackage(byte(i))

		if err != nil {
			log.Printf("Error creating DHCP REQUEST PACKAGE with %d last byte ...", i)
		}

		err = client.SendBuffer(pack)

		if err != nil {
			log.Printf("Error sending buffer with reserved last byte %d ...", i)
		}

		time.Sleep(config.TIMEOUT_BETWEEN_PACKAGES * time.Millisecond)
	}
}
