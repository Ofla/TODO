package helpers_test

import (
	hlp "github.com/Ofla/TODO/helpers"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := hlp.GetLogger()
	if logger == nil {
		t.Errorf("expected a logger got nil")
	}
}
