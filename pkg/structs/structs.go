package structs

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"

	"github.com/radiohive/zimt/pkg/strings"
)

// Field represents single field of struct
type Field struct {
	Value reflect.Value
	Field reflect.StructField
}

// Tag returns the value associated with key in the tag string
func (f Field) Tag(name string) string {
	return f.Field.Tag.Get(name)
}

// ExtractFields return slice of struct's fields
func ExtractFields(arg interface{}) (fields []Field) {
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
		if !fv.IsValid() {
			continue
		}
		fields = append(fields, Field{
			Value: fv,
			Field: st.Field(i),
		})
	}

	return fields
}

// UnmarshalViper initializes config struct with viper values,
// the struct fields are expected to have `viper` tags
func UnmarshalViper(arg interface{}) {
	fields := ExtractFields(arg)

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		t := f.Tag("viper")
		if !f.Value.CanSet() || t == "" {
			continue
		}
		switch f.Value.Kind() {
		case reflect.Int:
			f.Value.SetInt(viper.GetInt64(t))
		case reflect.String:
			f.Value.SetString(viper.GetString(t))
		case reflect.Bool:
			f.Value.SetBool(viper.GetBool(t))
		}
	}
}

// Print prints the config fields and values to the standard output
func Print(arg interface{}, tag string) {
	fields := ExtractFields(arg)

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		t := f.Tag(tag)
		if t == "" {
			continue
		}
		switch f.Value.Kind() {
		case reflect.Int:
			fmt.Printf("%s=%d\n", t, f.Value.Int())
		case reflect.String:
			if f.Tag("print") == "mask" {
				fmt.Printf("%s=%q\n", t, strings.Mask(f.Value.String()))
			} else {
				fmt.Printf("%s=%q\n", t, f.Value.String())
			}
		case reflect.Bool:
			fmt.Printf("%s=%t\n", t, f.Value.Bool())
		}
	}
}
