package errors

import (
	"errors"
	"strings"
	"testing"
)

// Test_CreateGenericError tests CreateGenericError
func Test_CreateGenericError(t *testing.T) {
	tCode := errorCode(12)
	tMsg := "this is an error"
	err := errorCode(12).CreateGenericError(tMsg)
	if !strings.Contains(err.Error(), tCode.String()) ||
		!strings.Contains(err.Error(), tMsg) {
		t.Errorf("expected error %#vn got different error", err)
	}
}

// Test_CreateInternalError tests CreateInternalError
func Test_CreateInternalError(t *testing.T) {
	tMsg := errors.New("this is an error")
	err := CreateInternalError(tMsg)
	if !strings.Contains(err.Error(), tMsg.Error()) {
		t.Errorf("expected error %#vn got different error", err)
	}
}

// Test_CreateCustomError tests CreateCustomError
func Test_CreateCustomError(t *testing.T) {
	tCode := errorCode(12)
	tMsg := "this is an error"
	tError := errors.New("this is an error")
	res := tCode.CreateCustomError(tMsg)
	err := res(tError)
	if !strings.Contains(err.Error(), tCode.String()) ||
		!strings.Contains(err.Error(), tMsg) {
		t.Errorf("expected error %#vn got different error", err)
	}
}
