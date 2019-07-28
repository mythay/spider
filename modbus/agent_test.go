package modbus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectModbusTcp(t *testing.T) {
	assert := assert.New(t)

	regs := []regReq{{3, 0, 2}}
	x := NewClient("127.0.0.1:502", regs)
	assert.NotNil(x)
	assert.Nil(x.Once())
}

func BenchmarkTemplateParallel(b *testing.B) {
	b.StopTimer()
	regs := []regReq{{3, 1, 2}}
	x := NewClient("127.0.0.1:502", regs)
	// fmt.Printf("hi %d\n", b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x.Once()
	}
}
