package config

import "testing"

func TestReadConfigReturnsConfig(t *testing.T) {
	var config Config
	config = Read("config_test.json")
	if config.ColorCubeBitSize != 4 {
		t.Fail()
	}
	if config.OutputFilename != "out.png" {
		t.Fail()
	}

}
