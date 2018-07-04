package views

import (
	"html/template"
	"io"
	"path/filepath"
)

// NewView creates a new View from the given HTML template files.
func NewView(layout string, files ...string) *View {
	matches, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	files = append(files, matches...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Execute executes the template, writing the results to the io.Writer.
func (v *View) Execute(wr io.Writer, data interface{}) error {
	return v.Template.ExecuteTemplate(wr, v.Layout, data)
}

// View represents a view defined by HTML templates.
type View struct {
	Template *template.Template
	Layout   string
}
