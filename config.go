package spider

type Config struct {
	Type   string
	Device map[string]Device
	Host   []Host
}
type Device struct {
	Register []Register
	Range    []Range
}
type Register struct {
	Base uint32
	Type string
	Cmd  string
	Tag  string
	Name string
}

type Range struct {
	Base  uint32
	Count uint32
}

type Host struct {
	Name     string
	Interval uint32
	Address  string
	Port     uint32
	Serial   string
	Baudrate uint32
	Slave    []Slave
}

type Slave struct {
	SlaveId uint8
	Device  string
	Name    string
	Collect []string
}
