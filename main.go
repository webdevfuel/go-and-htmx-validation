package main

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/webdevfuel/go-and-htmx-validation/template"
	"github.com/webdevfuel/go-and-htmx-validation/user"
)

func main() {
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
		// put values into map
		values := make(map[string]string)
		for k, v := range r.PostForm {
			values[k] = v[0]
		}
		// validate map
		err = validation.Validate(values,
			validation.Map(
				validation.Key("id", validation.Required, validation.Length(8, 0)),
				validation.Key("name", validation.Required),
			),
		)
		if err != nil {
			renderFormComponent(w, r)
			return
		}
		// construct user
		u := user.User{
			ID:   values["id"],
			Name: values["name"],
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
