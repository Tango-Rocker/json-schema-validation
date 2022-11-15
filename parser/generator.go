package parser

import (
	"github.com/Tango-Rocker/json-schema-validation/schema"
)

type Parser struct {
}

func New() *Parser {
	return new(Parser)
}

func (p *Parser) Generate(input map[string]interface{}) *schema.Field {
	root := schema.NewField("", true, schema.TypeFactory(input))
	return p.generate(root, input)
}

func (p *Parser) generate(root *schema.Field, input map[string]interface{}) *schema.Field {
	var f *schema.Field
	for key, inner := range input {
		name := key

		required := key[len(key)-1] != '?'

		if !required {
			//le quitamos el '?'
			name = key[:len(key)-1]
		}

		f = schema.NewField(name, required, schema.TypeFactory(input[key]))
		if f.Type == schema.StructType {
			p.generate(f, inner.(map[string]interface{}))
		}

		root.AddChild(f)
	}

	return f
}
