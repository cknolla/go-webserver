package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, else false
func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// Has ensures that a form field is in Post and is not empty
func (form *Form) Has(field string, r *http.Request) bool {
	value := r.Form.Get(field)
	if value == "" {
		return false
	}
	return true
}


// MinLength checks for string minimum length
func (form *Form) MinLength(field string, length int, r *http.Request) bool {
	value := r.Form.Get(field)
	if len(value) < length {
		form.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}


// IsEmail checks for valid email address
func (form *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(form.Get(field)) {
		form.Errors.Add(field, "Invalid email address")
		return false
	}
	return true
}