package ctlrs

import (
	"fmt"
	"net/http"

	"github.com/joberly/imgs/views"
)

// NewUsers creates a new Users controller.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

// Users is a controller for registered users.
type Users struct {
	NewView *views.View
}

// New renders the form where a user can create a new account.
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// SignupForm contains data from a submitted user account signup
// form processed by Create.
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create processes the signup form when a user requests creation
// of a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Printf("new signup: email %s password %s\n", form.Email, form.Password)
}
