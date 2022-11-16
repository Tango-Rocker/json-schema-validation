package main

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

const errorMsg = "data type %s not implemented yet, sorry! "

func TypeFactory(t interface{}) DataType {
	switch t.(type) {
	case json.Number:
		return numberType
	case string:
		return stringType
	case map[string]interface{}:
		return structType
	default:
		panic(fmt.Sprintf(errorMsg, reflect.TypeOf(t).String()))
	}
}
