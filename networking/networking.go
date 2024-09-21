package networking

import (
	"encoding/binary"
	"net"
)

type IpAddress struct {
	Addr [4]byte
}

type Client struct {
	connection net.Conn
}

func CreateNewClient() (*Client, error) {
	connection, err := net.Dial("udp", "255.255.255.255:67")

	if err != nil {
		return nil, err
	}

	return &Client{connection: connection}, nil
}

func (c *Client) SendBuffer(buffer []byte) error {
	_, err := c.connection.Write(buffer)

	return err
}

func (c *Client) CloseConnection() error {
	return c.connection.Close()
}

func (inPoolAddr *IpAddress) IsLessOrEqual(poolAddr *IpAddress) bool {
	inPoolAddrUint := binary.BigEndian.Uint32(inPoolAddr.Addr[:])
	poolAddrUint := binary.BigEndian.Uint32(poolAddr.Addr[:])

	return inPoolAddrUint <= poolAddrUint
}

func (addr *IpAddress) Next() {
	addrUint := binary.BigEndian.Uint32(addr.Addr[:])
	addrUint++
	binary.BigEndian.PutUint32(addr.Addr[:], addrUint)
}
