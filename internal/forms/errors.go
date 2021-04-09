package forms

type errors map[string][]string

// Add adds error message for given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns first error message
func (e errors) Get(field string) string {
	errString := e[field]
	if len(errString) == 0 {
		return ""
	}

	return errString[0]
}
