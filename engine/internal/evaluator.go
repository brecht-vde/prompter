package engine

func Eval(template Template, data map[string]interface{}) string {
	for _, s := range template.s {
		switch t := s.(type) {
		case *Variable:
			value, ok := data[t.Identifier]

			if !ok {
				panic("data missing for template")
			}

			t.Identifier = value.(string)
		}
	}

	return template.String()
}
