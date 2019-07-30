package spider

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/influxdata/toml"
)

type Verifier interface {
	Verify() (bool, error)
}

type CfgModbus struct {
	Type   string
	Device map[string]CfgDevice
	Host   []CfgHost
}
type CfgDevice struct {
	Register map[string]CfgRegister
	Range    []CfgRange
}
type CfgRegister struct {
	Base uint16
	Type string
	Cmd  string
	Tag  string
}

func (reg *CfgRegister) Count() uint16 {
	switch reg.Type {
	case "signed":
		fallthrough
	case "unsigned":
		fallthrough
	case "hex":
		fallthrough
	case "binary":
		return 1
	case "float":
		fallthrough
	case "long":
		return 2
	case "double":
		return 4
	default:
		return 0
	}
}
func (reg *CfgRegister) Verify() (bool, error) {
	if reg.Count() != 0 {
		return true, nil
	} else {
		return false, fmt.Errorf(" invalid type'%s'", reg.Type)
	}

}
func (reg *CfgRegister) Last() uint16 {
	return reg.Base + reg.Count()
}

type CfgRange struct {
	Base  uint16
	Count uint16
}

func (rag *CfgRange) Last() uint16 {
	return rag.Base + rag.Count
}

type CfgAddrTcp struct {
	Port uint16
}

type CfgAddrRtu struct {
	BaudRate int
	// Data bits: 5, 6, 7 or 8 (default 8)
	DataBits int
	// Stop bits: 1 or 2 (default 1)
	StopBits int
	// Parity: N - None, E - Even, O - Odd (default E)
	// (The use of no parity requires 2 stop bits.)
	Parity string
}
type CfgHost struct {
	Name     string
	Interval uint32
	Address  string
	CfgAddrTcp
	CfgAddrRtu
	Slave []CfgSlave
}

type CfgSlave struct {
	SlaveId uint8
	Device  string
	Name    string
	Collect []string
}

func LoadCfgModbus(path string) (*CfgModbus, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var config CfgModbus
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	return &config, nil

}
