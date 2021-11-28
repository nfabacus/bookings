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
		return false
	}
	return true
}
