package main

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/form"
	"github.com/webdevfuel/go-and-htmx-validation/template"
	"github.com/webdevfuel/go-and-htmx-validation/user"
)

var decoder *form.Decoder

type FormData struct {
	ID   string `form:"id"`
	Name string `form:"name"`
}

func (f FormData) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.ID, validation.Required, validation.Length(8, 0)),
		validation.Field(&f.Name, validation.Required),
	)
}

func main() {
	decoder = form.NewDecoder()
	r := http.NewServeMux()
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users := user.ListUsers()
		component := template.Users(users)
		component.Render(r.Context(), w)
	})
	r.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		// parse form
		err := r.ParseForm()
		if err != nil {
			renderFormComponent(w, r)
			return
		}
		// decode form into struct
		var formData FormData
		err = decoder.Decode(&formData, r.PostForm)
		if err != nil {
			renderFormComponent(w, r)
			return
		}
		// validate struct
		err = formData.Validate()
		if err != nil {
			renderFormComponent(w, r)
			return
		}
		// construct user
		u := user.User{
			ID:   formData.ID,
			Name: formData.Name,
		}
		// update slice of users
		user.AddUser(u)
		// render form and user
		renderFormComponent(w, r)
		renderUserComponent(w, r, u)
	})
	http.ListenAndServe("localhost:3000", r)
}

func renderFormComponent(w http.ResponseWriter, r *http.Request) {
	formComponent := template.Form()
	formComponent.Render(r.Context(), w)
}

func renderUserComponent(w http.ResponseWriter, r *http.Request, u user.User) {
	userComponent := template.User(u)
	userComponent.Render(r.Context(), w)
}
