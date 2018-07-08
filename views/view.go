package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is the relative location of all view layout templates.
	LayoutDir = "views/layouts/"

	// TemplateDir is the relative location of all view templates.
	TemplateDir = "views/"

	// TemplateExt is the file extension for all view templates.
	TemplateExt = ".gohtml"
)

// NewView creates a new View from the given HTML template files.
func NewView(layout string, files ...string) *View {
	// Process the input template files
	addTemplateExt(files)
	addTemplatePath(files)

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

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// Render executes the Layout template, writing the results to the io.Writer.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// View represents a view defined by HTML templates.
type View struct {
	Template *template.Template
	Layout   string
}

// addTemplatePath takes a slice of template file path strings and prepends the
// TemplateDir directory to each string in the slice.
//
// If the TemplateDir is "views/", then the input of {"home"} would modify the
// values in the slice to be {"views/home"}.
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes a slice of template file path strings and appends the
// TemplateExt extension to each string in the slice.
//
// If the TemplateExt is ".gohtml", then the input of {"home"} would modify the
// values in the slice to be {"home.gohtml"}.
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
