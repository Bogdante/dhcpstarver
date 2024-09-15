package networking

import "net"

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
