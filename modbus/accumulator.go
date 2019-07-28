package modbus

import (
	"fmt"

	"github.com/mythay/spider"
)

type innerRange struct {
	org  spider.MbRange
	calc spider.MbRange
}

func (rag *innerRange) AdjustRange(reg *spider.MbRegister) bool {
	if reg.Base >= rag.org.Base && reg.Last() <= rag.org.Last() {
		if rag.calc.Base+rag.calc.Count == 0 { // first time, all value are zero
			rag.calc.Base = reg.Base
			rag.calc.Count = reg.Count()
		} else {
			if reg.Base < rag.calc.Base {
				rag.calc.Count += rag.calc.Base - reg.Base
				rag.calc.Base = reg.Base
			}
			if reg.Last() > rag.calc.Last() {
				rag.calc.Count += reg.Last() - rag.calc.Last()
			}
		}

		return true
	}
	return false
}

func accumulate(collect []string, device *spider.MbDevice) ([]regReq, error) {
	var regs []regReq
	rags := make([]innerRange, len(device.Range))
	for i, rag := range device.Range {
		rags[i].org = rag
	}
	for _, item := range collect {
		if reg, ok := device.Register[item]; ok {
			match := false
			for i, _ := range rags {
				if rags[i].AdjustRange(&reg) {
					match = true
					break
				}
			}
			if !match {
				regs = append(regs, regReq{base: reg.Base, count: reg.Count()})
			}
		} else {
			return nil, fmt.Errorf("'%s' reg not exist\n", item)
		}
	}
	for _, rag := range rags {
		if rag.calc.Base+rag.calc.Count > 0 {
			regs = append(regs, regReq{base: rag.calc.Base, count: rag.calc.Count})
		}
	}
	return regs, nil
}
