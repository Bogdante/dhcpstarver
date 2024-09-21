package args

import (
	"flag"

	"github.com/Bogdante/dhcpstarver/networking"
	"github.com/Bogdante/dhcpstarver/utils"
)

func ParseCmdArguments() (bool, *networking.IpAddress, *networking.IpAddress, int64) {
	option := flag.Bool("r", true, "Send DHCP Request Packages")
	poolStart := flag.String("s", "192.168.0.100", "Starting IP-address from DHCP pool")
	poolEnd := flag.String("e", "192.168.0.200", "Final IP-adrees from DHCP pool")
	delay := flag.Int64("d", 1000, "Delay between sent packages")

	flag.Parse()

	poolStartBytes, err := utils.StringIPtoBytes(*poolStart)

	if err != nil {
		panic("Error parsing starting IP address")
	}

	poolEndBytes, err := utils.StringIPtoBytes(*poolEnd)

	if err != nil {
		panic("Errro parsing end IP address")
	}

	return *option, poolStartBytes, poolEndBytes, *delay
}
