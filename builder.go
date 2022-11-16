package main

type Builder struct {
}

func NewBuilder() *Builder {
	return new(Builder)
}

func (p *Builder) Build(input map[string]interface{}) *Field {
	root := NewField("", true, TypeFactory(input))
	return p.build(root, input)
}

func (p *Builder) build(root *Field, input map[string]interface{}) *Field {
	var f *Field
	for key, inner := range input {
		name := key

		required := key[len(key)-1] != '?'

		if !required {
			//le quitamos el '?'
			name = key[:len(key)-1]
		}

		f = NewField(name, required, TypeFactory(input[key]))
		if f.Type == StructType {
			p.build(f, inner.(map[string]interface{}))
		}

		root.AddChild(f)
	}

	return f
}
