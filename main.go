package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tango-Rocker/json-schema-validation/parser"
	"strings"
)

var baseSchema = `{
	"root":{
		"api":{
			"name":"foobar",
			"id": 50,
			"desc?":"zzz"
		},
		"version?": "string"
	}
}`

var example = `{
	"root":{
		"api":{
			"name": "hola",
			"id": 50,
			"desc": "hola"
		}		
	}
}`

func main() {
	schemaMap := make(map[string]interface{}, 0)
	exampleMap := make(map[string]interface{}, 0)

	d := json.NewDecoder(strings.NewReader(baseSchema))
	//sin esto explota, el default de json es decodear los numeros a float64
	d.UseNumber()
	fmt.Println(d.Decode(&schemaMap))

	c := json.NewDecoder(strings.NewReader(example))
	c.UseNumber()
	fmt.Println(c.Decode(&exampleMap))

	//Instanciamos el parser y generamos el arbol semantico
	p := parser.New()
	fields := p.Generate(schemaMap)

	//validamos el ejemplo contra el arbol semantico del schema
	err := fields.Validate(exampleMap)
	if err != nil {
		panic(err)
	}
}
