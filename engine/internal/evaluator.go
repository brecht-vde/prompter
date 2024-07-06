package internal

import (
	"fmt"
	"reflect"
	"strings"
)

func Eval(template Template, data map[string]interface{}) (string, error) {
	for _, s := range template.s {
		switch t := s.(type) {
		case *Variable:
			value, ok := data[t.Identifier]

			if !ok {
				return "", fmt.Errorf("missing or invalid data for variable '%s'", t.Identifier)
			}

			t.Value = fmt.Sprintf("%v", value)
		case *Join:
			value, ok := data[t.Identifier]

			if !ok {
				return "", fmt.Errorf("missing or invalid data for join '%s'", t.Identifier)
			}

			cvs := reflect.ValueOf(value)

			if cvs.Kind() != reflect.Slice {
				return "", fmt.Errorf("invalid type for join '%s'. you must provide a slice type", t.Identifier)
			}

			v := make([]string, cvs.Len())

			for i := 0; i < len(v); i++ {
				v[i] = fmt.Sprintf("%v", cvs.Index(i).Interface())
			}

			t.Value = strings.Join(v, t.Separator)
		}
	}

	return template.String(), nil
}
