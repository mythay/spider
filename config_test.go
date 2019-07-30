package spider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTomlConfig(t *testing.T) {
	assert := assert.New(t)

	config, err := LoadCfgModbus("spider.default.toml")
	assert.Empty(err)

	assert.Equal("modbus", config.Type)
	assert.Equal(1, len(config.Device))

	assert.Equal(3, len(config.Device["em3250"].Register))
	assert.Equal("voltage", config.Device["em3250"].Register["input-1"].Tag)
	assert.Equal(2, len(config.Device["em3250"].Range))

}
