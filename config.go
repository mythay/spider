package spider

import (
	"io/ioutil"
	"os"

	"github.com/influxdata/toml"
)

type MbConfig struct {
	Type   string
	Device map[string]MbDevice
	Host   []MbHost
}
type MbDevice struct {
	Register map[string]MbRegister
	Range    []MbRange
}
type MbRegister struct {
	Base uint16
	Type string
	Cmd  string
	Tag  string
}

func (reg *MbRegister) Count() uint16 {
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
func (reg *MbRegister) Isvalid() bool {
	return reg.Count() != 0
}
func (reg *MbRegister) Last() uint16{
	return reg.Base + reg.Count()
}

type MbRange struct {
	Base  uint16
	Count uint16
}

func(rag * MbRange) Last() uint16{
	return rag.Base + rag.Count
}

type MbHost struct {
	Name     string
	Interval uint32
	Address  string
	Port     uint32
	Serial   string
	Baudrate uint32
	Slave    []MbSlave
}

type MbSlave struct {
	SlaveId uint8
	Device  string
	Name    string
	Collect []string
}

func LoadMbConfig(path string) (*MbConfig, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var config MbConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	return &config, nil

}
