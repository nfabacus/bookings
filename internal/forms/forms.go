package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}

// valid returns true if no errors, otherwise returns false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initialises a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}
