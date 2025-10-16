package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	generateNullableTypes()
}

func generateNullableTypes() {
	type NullableTemplateParams struct {
		TypeName, GoType, WriterMethod, LexerMethod string
	}

	var types = []NullableTemplateParams{
		{"Bool", "bool", "Bool", "Bool"},

		{"Int", "int", "Int", "Int"},
		{"Int8", "int8", "Int8", "Int8"},
		{"Int16", "int16", "Int16", "Int16"},
		{"Int32", "int32", "Int32", "Int32"},
		{"Int64", "int64", "Int64", "Int64"},

		{"UInt", "uint", "Uint", "Uint"},
		{"UInt8", "uint8", "Uint8", "Uint8"},
		{"UInt16", "uint16", "Uint16", "Uint16"},
		{"UInt32", "uint32", "Uint32", "Uint32"},
		{"UInt64", "uint64", "Uint64", "Uint64"},

		{"Float32", "float32", "Float32", "Float32"},
		{"Float64", "float64", "Float64", "Float64"},

		{"String", "string", "String", "String"},
	}

	tmpl := template.Must(template.ParseFiles("templates/nullable.tmpl"))

	for _, t := range types {
		typeName := strings.ToLower(t.TypeName)
		file, _ := os.Create("nullable/" + typeName + ".go")
		if err := tmpl.Execute(file, t); err != nil {
			panic(err)
		}
		_ = file.Close()
	}
}
