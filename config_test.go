package spider

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"github.com/influxdata/toml"
	"os"
  )

  func TestLoadTomlConfig(t *testing.T) {

	assert := assert.New(t)

 
  
	f, err := os.Open("spider.default.toml")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    buf, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }
    var config Config
    if err := toml.Unmarshal(buf, &config); err != nil {
        panic(err)
	}
	
	assert.Equal("modbus", config.Type )
	assert.Equal(1, len(config.Device))

	assert.Equal(1,1)
  
  }