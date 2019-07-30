package modbus

import (
	"fmt"

	"github.com/mythay/spider"
)

type innerRange struct {
	org  spider.CfgRange
	calc spider.CfgRange
}

func (rag *innerRange) AdjustRange(reg *spider.CfgRegister) bool {
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

func accumulate(collect []string, device spider.CfgDevice) ([]spider.CfgRange, error) {
	var regs []spider.CfgRange
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
				regs = append(regs, spider.CfgRange{Base: reg.Base, Count: reg.Count()})
			}
		} else {
			return nil, fmt.Errorf("'%s' reg not exist\n", item)
		}
	}
	for _, rag := range rags {
		if rag.calc.Base+rag.calc.Count > 0 {
			regs = append(regs, spider.CfgRange{Base: rag.calc.Base, Count: rag.calc.Count})
		}
	}
	return regs, nil
}
