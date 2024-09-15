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
		pack, err := dhcp.CreateDhcpDiscoverPackage(byte(i))

		if err != nil {
			log.Printf("Error sending %d package ...", i)
		}

		err = client.SendBuffer(pack)

		if err != nil {
			log.Printf("Error sending buffer %d package ...", i)
		}

		time.Sleep(config.TIMEOUT_BETWEEN_PACKAGES * time.Millisecond)
	}
}
