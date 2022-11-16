package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// definicion de como es la estructura de los payloads que vamos a analizar
// el caracter '?' es un flag para indicar que el campo es opcional..
// seguramente nadie use el caracter '?' en un payload real verdad? ... verdad?
var jsonRules = `{
	"root":{
		"api":{
			"name":"foobar",
			"id": 50,
			"desc?":"zzz"
		},
		"version?": "string"
	}
}`

// ejmplo
var jsonExample = `{
	"root":{
		"api":{
			"name": "hola",
			"id": 50		
		}	
	}
}`

func main() {
	//levantamos la data de los jsons dentro en un Schema -> map[string]interface{}
	rules, err := getSchema(jsonRules)
	if err != nil {
		panic(err)
	}

	example, err2 := getSchema(jsonExample)
	if err2 != nil {
		panic(err)
	}

	//Instanciamos el builder y generamos el arbol estructural
	ruleBuilder := NewBuilder()
	tree := ruleBuilder.Build(rules)

	//validamos el ejemplo contra el arbol del schema
	err = tree.Validate(example)
	if err != nil {
		panic(err)
	}

}

func getSchema(payload string) (map[string]interface{}, error) {
	schema := make(map[string]interface{}, 0)
	d := json.NewDecoder(strings.NewReader(payload))
	//sin esto explota, el default de json es decodear los numeros a float64
	d.UseNumber()
	fmt.Println()

	return schema, d.Decode(&schema)
}
