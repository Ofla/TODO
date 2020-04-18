package validators

import (
	mdl "github.com/Ofla/TODO/models"
	"testing"
)

func TestIsNilOrEmpty(t *testing.T) {
	obj := mdl.Todo{}
	if ok :=  IsNilOrEmpty(&obj); !ok {
		t.Errorf("expected empty/nil obj: %v, got %v", ok, !ok)
	}
}

func TestIsStringEmpty(t *testing.T) {
	obj := "   "
	if ok := IsStringEmpty(obj); !ok {
		t.Errorf("expected empty string: %v, got %v", ok, !ok)
	}
}