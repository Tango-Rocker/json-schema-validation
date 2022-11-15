package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	StructType = iota
	NumberType
	StringType
)

type ValidatorFunc func(value interface{}) bool
type DataType struct {
	Type int
	f    ValidatorFunc
}

func ValidateString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func ValidateNumber(value interface{}) bool {
	_, ok := value.(json.Number)
	return ok
}

func (dt *DataType) IsValid(value interface{}) bool {
	return dt.f(value)
}

var (
	structType = DataType{
		Type: StructType,
		f:    nil,
	}

	numberType = DataType{
		Type: NumberType,
		f:    ValidateNumber,
	}

	stringType = DataType{
		Type: StringType,
		f:    ValidateString,
	}
)

func TypeFactory(t interface{}) DataType {
	switch t.(type) {
	case json.Number:
		return numberType
	case string:
		return stringType
	case map[string]interface{}:
		return structType
	default:
		msg := fmt.Sprintf("data type %v not implemented yet, sorry! ", reflect.TypeOf(t).String())
		panic(msg)
	}
}
