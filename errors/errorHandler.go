package errors

import (
	"encoding/json"
	"strconv"
)

// Handler describes the behaviour of error
type Handler interface {
	CreateGenericError(msg string) error
	CreateCustomError(msg string) error
}

type errorCode uint16

// ErrorModel represents the model of error
type ErrorModel struct {
	Code       errorCode
	GenericMsg string
	ErrorMsg   string
	Details    interface{} // contains any useful info to debug the error
}

// Error serializes the error
func (err ErrorModel) Error() string {
	errMarshalled, _ := json.Marshal(&err)
	return string(errMarshalled)
}

// String converts the Code to string
func (code errorCode) String() string {
	return strconv.Itoa(int(code))
}

// CreateGenericError creates generic error with correlationID
func (code errorCode) CreateGenericError(msg string) error {
	return ErrorModel{
		Code:       code,
		GenericMsg: msg,
	}

}

// CreateCustomError creates custom error with correlationID and extra information
func (code errorCode) CreateCustomError(msg string) func(err error, details ...interface{}) error {
	createCustomError := func(err error, details ...interface{}) error {
		errModel := ErrorModel{
			Code:       code,
			GenericMsg: msg,
			ErrorMsg:   err.Error(),
			Details:    details,
		}
		return errModel
	}
	return createCustomError
}

// CreateInternalError creates internal error with detailed error
// used to create errors for logging
func CreateInternalError(err error, errorDescription ...interface{}) error {
	errorModel := ErrorModel{
		ErrorMsg: err.Error(),
		Details:  errorDescription,
	}
	return errorModel
}
var (
	ContextDeadlineExceededError = errorCode(1).CreateGenericError("timeout exceeded")
	InvalidContext               = errorCode(2).CreateGenericError("invalid timeout")
)

var (
	DBGetError          = errorCode(100).CreateGenericError("Error on Get Todo")
	DBUpdateError       = errorCode(101).CreateGenericError("Error on Update Todo")
	DBUpdateStatusError = errorCode(102).CreateGenericError("Error on Update Status of Todo")
	DBCreateError       = errorCode(103).CreateGenericError("Error on Create Todo")
	DBDeleteError       = errorCode(103).CreateGenericError("Error on Delete Todo")
	DBGetAllError       = errorCode(104).CreateGenericError("Error on Get allTodos")
)

var (
	InvalidInputError      = errorCode(400).CreateGenericError("Invalid input")
	InvalidIDError         = errorCode(401).CreateGenericError("Invalid ID")
	InvalidDBTypeError     = errorCode(403).CreateGenericError("Invalid DB type")
	InvalidServerTypeError = errorCode(406).CreateGenericError("Invalid server type")
	MarshallError          = errorCode(407).CreateCustomError("Error on Marshalling object")
	UnmarshallError        = errorCode(408).CreateCustomError("Error on Unmarshalling object")
	NotFoundError          = errorCode(411).CreateGenericError("element Not found")
)
