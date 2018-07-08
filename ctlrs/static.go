package ctlrs

import (
	"github.com/joberly/imgs/views"
)

// Static is a controller which contains static GET only views.
type Static struct {
	Home     *views.View
	Contact  *views.View
	NotFound *views.View
}

// NewStatic creates and initializes a new Static controller.
func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("bootstrap", "static/home"),
		Contact:  views.NewView("bootstrap", "static/contact"),
		NotFound: views.NewView("bootstrap", "static/notfound"),
	}
}
