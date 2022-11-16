package main

import "fmt"

type Field struct {
	DataType
	name       string
	children   []*Field
	isRequired bool
}

func NewField(name string, required bool, t DataType) *Field {
	f := new(Field)
	f.name = name
	f.DataType = t
	f.isRequired = required
	f.children = make([]*Field, 0)

	return f
}

const (
	missingMsg = "missing field %s"
	//TODO: unir estos 2 mensajes en uno, imprimiendo el expected y el actual
	structExpectedMSg = "field %s expected as struct, but found %v"
	typeXExpectedMsg  = "field %s must be of type %d but found %v"
)

func (f *Field) Validate(tree map[string]interface{}) error {
	source := tree[f.name]

	if source == nil {
		if f.isRequired {
			return fmt.Errorf(missingMsg, f.name)
		}
		return nil
	}

	if f.Type == StructType {
		if child, ok := source.(map[string]interface{}); ok {
			for i := 0; i < len(f.children); i++ {
				err := f.children[i].Validate(child)
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf(structExpectedMSg, f.name, source)
		}
	} else {
		if !f.DataType.IsValid(source) {
			return fmt.Errorf(typeXExpectedMsg, f.name, f.DataType.Type, source)
		}
	}

	return nil
}

func (f *Field) AddChild(child *Field) {
	f.children = append(f.children, child)
}
