package dhcp

import (
	"bytes"
	"encoding/binary"

	"github.com/Bogdante/dhcpstarver/utils"
)

type DhcpHeader struct {
	op      byte
	htype   byte
	hlen    byte
	hops    byte
	xid     uint32
	secs    uint16
	flags   uint16
	ciaddr  uint32
	yiaddr  uint32
	siaddr  uint32
	giaddr  uint32
	chaddr  [16]byte
	sname   [64]byte
	file    [128]byte
	options [64]byte
}

func CreateDhcpRequestPackage(lastRequestIpByte byte) ([]byte, error) {

	var dhcpPack DhcpHeader
	dhcpPack.op = 1
	dhcpPack.htype = 1
	dhcpPack.hlen = 6
	dhcpPack.hops = 0
	dhcpPack.secs = 0
	dhcpPack.flags = 0x8000

	fakeMac, err := utils.GenerateRandomMac()

	if err != nil {
		panic("Can't generate fake MAC-address")
	}

	fakeXid, err := utils.GenerateRandomTransactionId()

	if err != nil {
		panic("Can't generate fake XID")
	}

	copy(dhcpPack.chaddr[:6], fakeMac[:])

	dhcpPack.xid = fakeXid

	customOptions := make([]byte, 0, 64)
	customOptions = append(customOptions, 0x63, 0x82, 0x53, 0x63)                  // Magic COOKIE: DHCP
	customOptions = append(customOptions, 0x32, 4, 192, 168, 0, lastRequestIpByte) // Request Address
	customOptions = append(customOptions, 0x35, 0x01, 0x03)                        // TYPE Message Request

	customOptions = append(customOptions, 0xff) // END
	copy(dhcpPack.options[:], customOptions)

	var buffer bytes.Buffer

	err = binary.Write(&buffer, binary.BigEndian, &dhcpPack)

	if err != nil {
		panic("Can't write to the BUFFER")
	}

	return buffer.Bytes(), nil
}
