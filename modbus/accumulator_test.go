package modbus

import (
	"reflect"
	"testing"

	"github.com/mythay/spider"
)

func Test_accumulate_no_range(t *testing.T) {
	device := &spider.MbDevice{
		Register: map[string]spider.MbRegister{"value-1": spider.MbRegister{Base: 0, Type: "float"}, "value-2": spider.MbRegister{Base: 4, Type: "signed"}},
		Range:    nil,
	}
	type args struct {
		collect []string
		device  *spider.MbDevice
	}
	tests := []struct {
		name    string
		args    args
		want    []regReq
		wantErr bool
	}{
		// TODO: Add test cases.
		{"only one float reg", args{[]string{"value-1"}, device}, []regReq{regReq{0, 0, 2}}, false},
		{"only one singed reg", args{[]string{"value-2"}, device}, []regReq{regReq{0, 4, 1}}, false},
		{"two regs", args{[]string{"value-1", "value-2"}, device}, []regReq{regReq{0, 0, 2}, regReq{0, 4, 1}}, false},
		{"invalid value", args{[]string{"value-notexist"}, device}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accumulate(tt.args.collect, tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("accumulate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accumulate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accumulate_with_range(t *testing.T) {
	device := &spider.MbDevice{
		Register: map[string]spider.MbRegister{"value-1": spider.MbRegister{Base: 0, Type: "float"}, "value-2": spider.MbRegister{Base: 4, Type: "signed"}},
		Range:    []spider.MbRange{spider.MbRange{0, 10}},
	}
	type args struct {
		collect []string
		device  *spider.MbDevice
	}
	tests := []struct {
		name    string
		args    args
		want    []regReq
		wantErr bool
	}{
		// TODO: Add test cases.
		{"only one float reg", args{[]string{"value-1"}, device}, []regReq{regReq{0, 0, 2}}, false},
		{"only one singed reg", args{[]string{"value-2"}, device}, []regReq{regReq{0, 4, 1}}, false},
		{"two regs", args{[]string{"value-1", "value-2"}, device}, []regReq{regReq{0, 0, 5}}, false},
		{"invalid value", args{[]string{"value-notexist"}, device}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accumulate(tt.args.collect, tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("accumulate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accumulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
