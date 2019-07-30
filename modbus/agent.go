package modbus

import (
	"fmt"
	"net"
	"time"

	mb "github.com/goburrow/modbus"
	"github.com/mythay/spider"
)

type mbRange struct {
	spider.CfgRange
	resp []byte
	err  error
	ts   time.Time
}

type client struct {
	host       spider.CfgHost
	device     map[string]spider.CfgDevice
	conn       mb.Client
	tcpHandler *mb.TCPClientHandler
	rtuHandler *mb.RTUClientHandler
	slave      []slave
}

func ParseRegister(client *client) error {
	var slaves []slave
	for _, cfgSlave := range client.host.Slave {
		ranges, err := accumulate(cfgSlave.Collect, client.device[cfgSlave.Device])
		if err != nil {
			return err
		}
		var xrange []mbRange
		for _, item := range ranges {
			xrange = append(xrange, mbRange{CfgRange: item})
		}
		slaves = append(slaves, slave{CfgSlave: cfgSlave, parent: client, req: xrange})
	}
	client.slave = slaves
	return nil
}

func CreateConnection(client *client) error {
	if net.ParseIP(client.host.Address) != nil { // valid ip address
		port := client.host.Port
		if port == 0 {
			port = 502
		}
		handler := mb.NewTCPClientHandler(fmt.Sprintf("%s:%d", client.host.Address, port))
		handler.Timeout = 1 * time.Second
		client.conn = mb.NewClient(handler)
		client.tcpHandler = handler
	} else {
		handler := mb.NewRTUClientHandler(client.host.Address)
		client.conn = mb.NewClient(handler)
		client.rtuHandler = handler

	}
	return nil
}

func newClient(host spider.CfgHost, device map[string]spider.CfgDevice) (*client, error) {
	client := &client{host: host, device: device}

	err := ParseRegister(client)
	if err != nil {
		return nil, err
	}
	err = CreateConnection(client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type slave struct {
	spider.CfgSlave
	parent *client
	req    []mbRange
}
