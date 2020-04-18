package validators

import (
	mdl "github.com/Ofla/TODO/models"
	"strings"
)


func IsStringEmpty(val string)  bool {
	return len(strings.TrimSpace(val)) == 0
}

func IsNilOrEmpty(i interface{}) bool {
	switch obj := i.(type) {
	case *mdl.Todo:
		return obj == nil || IsStringEmpty(obj.Hash)
	}
	return false
}
