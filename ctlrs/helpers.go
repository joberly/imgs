package ctlrs

import (
	"net/http"

	"github.com/gorilla/schema"
)

// Common code for parsing the user account signup form.
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
