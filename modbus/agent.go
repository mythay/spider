package modbus

import (
	"time"

	"github.com/goburrow/modbus"
)

type regReq struct {
	fc    uint8
	base  uint16
	count uint16
}

type Client struct {
	mbClient modbus.Client
	mbReq    []regReq
}

func NewClient(host string, regs []regReq) *Client {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 1 * time.Second
	handler.SlaveId = 1
	// handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		return nil
	}
	return &Client{
		mbClient: modbus.NewClient(handler),
		mbReq:    regs,
	}
}

func (c *Client) Once() error {
	for _, regdef := range c.mbReq {
		switch regdef.fc {
		case modbus.FuncCodeReadHoldingRegisters:
			_, err := c.mbClient.ReadHoldingRegisters(regdef.base, regdef.count)
			if err != nil {
				return err
			} else {
				// fmt.Printf("got %v", resp)
				return nil
			}
		}
	}
	return nil
}
