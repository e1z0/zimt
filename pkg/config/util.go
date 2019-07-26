package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type structField struct {
	value reflect.Value
	field reflect.StructField
	tag   string
}

func extractStructFields(arg interface{}, tag string) []structField {
	var fields []structField

	v := reflect.ValueOf(arg)
	t := reflect.TypeOf(arg)

	if v.Kind() != reflect.Ptr {
		return fields
	}

	sv := v.Elem()
	st := t.Elem()
	if sv.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < sv.NumField(); i++ {
		fv := sv.Field(i)
		ft := st.Field(i)
		t := ft.Tag.Get(tag)

		if !fv.IsValid() || tag == "" {
			continue
		}

		fields = append(fields, structField{value: fv, field: ft, tag: t})
	}

	return fields
}

// unmarshal initializes config struct with viper values,
// the struct fields are expected to have `viper` tags
func unmarshal(arg interface{}) {
	fields := extractStructFields(arg, "viper")

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		if !f.value.CanSet() {
			continue
		}
		switch f.value.Kind() {
		case reflect.Int:
			f.value.SetInt(viper.GetInt64(f.tag))
		case reflect.String:
			f.value.SetString(viper.GetString(f.tag))
		case reflect.Bool:
			f.value.SetBool(viper.GetBool(f.tag))
		}
	}
}

// report prints the config fields and values to the standard output
func report(arg interface{}) {
	fields := extractStructFields(arg, "viper")

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		switch f.value.Kind() {
		case reflect.Int:
			fmt.Printf("%s=%d\n", f.tag, f.value.Int())
		case reflect.String:
			fmt.Printf("%s=%q\n", f.tag, f.value.String())
		case reflect.Bool:
			fmt.Printf("%s=%t\n", f.tag, f.value.Bool())
		}
	}
}
