package main

import (
	"net/http"

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
		u := user.User{
			ID:   r.FormValue("id"),
			Name: r.FormValue("name"),
		}
		user.AddUser(u)
		component := template.User(u)
		component.Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:3000", r)
}
