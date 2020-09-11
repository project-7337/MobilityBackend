package tests

import (
	"backend/config"
	"testing"
)

func TestConfig(t *testing.T) {
	var configs config.Conf
	configs.ReadConf()

	if configs.Test != "successfull" {
		t.Errorf("Configs Read Failed, expected %v, got %v", "successfull", configs.Test)
	}
}
