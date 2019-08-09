package structs

import (
	"fmt"
	"io"
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
func ExtractFields(arg interface{}, filters []string) ([]Field, error) {
	fields := make([]Field, 0)
	v := reflect.ValueOf(arg)
	t := reflect.TypeOf(arg)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fields, fmt.Errorf("%T is expected to be of type Struct or *Struct", arg)
	}

	if len(filters) > 0 {
		for _, n := range filters {
			ft, found := t.FieldByName(n)
			if !found {
				fields = append(fields, Field{})
				continue
			}
			fields = append(fields, Field{
				Value: v.FieldByName(n),
				Field: ft,
			})
		}
		return fields, nil
	}

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, Field{
			Value: v.Field(i),
			Field: t.Field(i),
		})
	}
	return fields, nil
}

// UnmarshalViper initializes config struct with viper values,
// the struct fields are expected to have `viper` tags
func UnmarshalViper(arg interface{}) {
	fields, _ := ExtractFields(arg, nil)

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

// Print writes the config fields and values to the
func Print(arg interface{}, tag string, writer io.Writer) {
	fields, _ := ExtractFields(arg, nil)

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		if f.Tag("print") == "-" {
			continue
		}
		t := f.Tag(tag)
		if t == "" {
			t = f.Field.Name
		}
		switch f.Value.Kind() {
		case reflect.Int:
			writer.Write([]byte(fmt.Sprintf("%s=%d\n", t, f.Value.Int()))) // #nosec
		case reflect.String:
			if f.Tag("print") == "mask" {
				writer.Write([]byte(fmt.Sprintf("%s=%q\n", t, strings.Mask(f.Value.String())))) // #nosec
			} else {
				writer.Write([]byte(fmt.Sprintf("%s=%q\n", t, f.Value.String()))) // #nosec
			}
		case reflect.Bool:
			writer.Write([]byte(fmt.Sprintf("%s=%t\n", t, f.Value.Bool()))) // #nosec
		}
	}
}

// Titles return printable struct's field titles
func Titles(arg interface{}, tag string, filters []string) ([]string, error) {
	s := []string{}
	fields, err := ExtractFields(arg, filters)
	if err != nil {
		return s, err
	}
	if len(fields) == 0 {
		return s, nil
	}

	for _, f := range fields {
		title := f.Field.Name
		if tag != "" {
			if t := f.Tag(tag); t != "" {
				title = t
			}
		}
		s = append(s, title)
	}

	return s, nil
}
