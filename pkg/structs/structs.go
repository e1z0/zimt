package structs

import (
	"reflect"
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
