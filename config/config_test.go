package config

import "testing"

func TestGetConfig(t *testing.T) {
	c := GetConfig()
	t.Logf("config: %+v\n", c)
}
