package utils

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"net"

	"github.com/Bogdante/dhcpstarver/networking"
)

func GenerateRandomMac() ([6]byte, error) {
	var mac [6]byte

	_, err := rand.Read(mac[:])

	if err != nil {
		return [6]byte{}, err
	}

	mac[0] = (mac[0] | 0x02) & 0xfe

	return mac, nil
}

func GenerateRandomTransactionId() (uint32, error) {

	var id [4]byte

	_, err := rand.Read(id[:])

	if err != nil {
		return 0, err
	}

	value := binary.BigEndian.Uint32(id[:])

	return value, nil
}

func StringIPtoBytes(ip string) (*networking.IpAddress, error) {

	netIp := net.ParseIP(ip)

	if netIp == nil || netIp.To4() == nil {
		return &networking.IpAddress{Addr: [4]byte{}}, errors.New("wrong argument ip")
	}

	var ipBytes [4]byte
	copy(ipBytes[:], netIp.To4())

	return &networking.IpAddress{Addr: ipBytes}, nil
}
