package config

import (
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	_, err := LoadConfiguration()
	if err != nil {
		t.Errorf("expected loading config, got error: %v", err)
	}
}
