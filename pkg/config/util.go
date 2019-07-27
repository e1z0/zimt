package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"

	"github.com/radiohive/zimt/pkg/strings"
	"github.com/radiohive/zimt/pkg/structs"
)

// unmarshal initializes config struct with viper values,
// the struct fields are expected to have `viper` tags
func unmarshal(arg interface{}) {
	fields := structs.ExtractFields(arg)

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		tag := f.Tag("viper")
		if !f.Value.CanSet() || tag == "" {
			continue
		}
		switch f.Value.Kind() {
		case reflect.Int:
			f.Value.SetInt(viper.GetInt64(tag))
		case reflect.String:
			f.Value.SetString(viper.GetString(tag))
		case reflect.Bool:
			f.Value.SetBool(viper.GetBool(tag))
		}
	}
}

// report prints the config fields and values to the standard output
func report(arg interface{}) {
	fields := structs.ExtractFields(arg)

	if len(fields) == 0 {
		return
	}

	for _, f := range fields {
		tag := f.Tag("viper")
		if tag == "" {
			continue
		}
		switch f.Value.Kind() {
		case reflect.Int:
			fmt.Printf("%s=%d\n", tag, f.Value.Int())
		case reflect.String:
			if f.Tag("print") == "mask" {
				fmt.Printf("%s=%q\n", tag, strings.Mask(f.Value.String()))
			} else {
				fmt.Printf("%s=%q\n", tag, f.Value.String())
			}
		case reflect.Bool:
			fmt.Printf("%s=%t\n", tag, f.Value.Bool())
		}
	}
}
