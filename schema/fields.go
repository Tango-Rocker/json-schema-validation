package schema

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

func (f *Field) Validate(tree map[string]interface{}) error {
	source := tree[f.name]

	if source == nil {
		if f.isRequired {
			return fmt.Errorf("missing field %s", f.name)
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
			return fmt.Errorf("field %s expected as struct, but foun %v", f.name, source)
		}
	} else {
		if !f.DataType.IsValid(source) {
			return fmt.Errorf("field %s must be of type %d but found %v", f.name, f.DataType.Type, source)
		}
	}

	return nil
}

func (f *Field) AddChild(child *Field) {
	f.children = append(f.children, child)
}
