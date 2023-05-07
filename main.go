package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name     string
	Age      uint8
	Document string `required:"true"`
}

func validateFields(stc any) error {
	t := reflect.TypeOf(stc)
	v := reflect.ValueOf(stc)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)

		required := fieldType.Tag.Get("required")

		if required == "" || required == "false" {
			continue
		}

		vf := v.Field(i)

		switch vf.Kind() {
		case reflect.String:
			if vf.String() == "" {
				return fmt.Errorf("o campo %s é obrigatório", fieldType.Name)
			}
		case reflect.Uint:
			if vf.Uint() == 0 {
				return fmt.Errorf("o campo %s é obrigatório", fieldType.Name)
			}
		}

	}

	return nil
}

func main() {
	p := Person{
		Name:     "Israel Souza",
		Age:      23,
		Document: "000000",
	}

	if err := validateFields(p); err != nil {
		panic(err)
	}
}
