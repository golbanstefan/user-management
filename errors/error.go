package errors

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

const VALIDATION_ERROR = "Error on validating"

type ErrorJson struct {
	Message string
	Data    interface{}
}

func (e ErrorJson) Error() string {
	b, err := json.Marshal(&e)
	CheckError(err)
	return string(b)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type FieldError struct {
	Err validator.FieldError
}

type ValidationError struct {
	Field     string
	Condition string
	Value     interface{}
	Message   string
}

func (q FieldError) ToString() string {
	var sb strings.Builder
	sb.WriteString("Validation failed on field '" + q.Err.Field() + "'")
	sb.WriteString(", Condition: " + q.Err.ActualTag())
	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.Err.Param() != "" {
		sb.WriteString(" { " + q.Err.Param() + " }")
	}
	if q.Err.Value() != nil && q.Err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", Actual: %v", q.Err.Value()))
	}
	return sb.String()
}

func (q FieldError) Json() ValidationError {
	return ValidationError{
		Field:     q.Err.Field(),
		Condition: q.Err.ActualTag(),
		Value:     q.Err.Value(),
		Message:   q.ToString(),
	}
}

// FirebaseError is an error type containing an error code string.
type FirebaseError struct {
	Error struct {
		Code    int64
		Message string
	}
}

func (f FirebaseError) Normalize(err error) []byte {
	return []byte(strings.Trim(err.Error(), "http error status: 400;body:"))
}

func ErrToJson(err error) error {
	var e ErrorJson
	switch t := reflect.ValueOf(err).Type().String(); t {
	case "*internal.FirebaseError":
		var f FirebaseError
		f.Normalize(err)
		json.Unmarshal(f.Normalize(err), &f)
		e.Message = f.Error.Message
		e.Data = f.Error
	case "validator.ValidationErrors":
		var errorsAr []ValidationError
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errorsAr = append(errorsAr, FieldError{fieldErr}.Json())
		}
		e.Message = VALIDATION_ERROR
		e.Data = errorsAr
	//todo add more cases, for different errors
	default:
		e.Message = err.Error()
		e.Data = err
	}
	return e
}
