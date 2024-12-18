package main

import (
	"fmt"
	"net/http"

	"portfolio/web/controllers"
	"portfolio/web/templates"
	"portfolio/web/views"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const PORT = ":8080"

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Not Found!", http.StatusNotFound)
	})

	type Route struct {
		Path         string
		Name         string
		Meta         string
		ViewName     string
		DisplayInNav bool
	}
	var routes = []Route{
		{Name: "Home", Path: "/", Meta: "jimdel | Home", ViewName: "home.gohtml", DisplayInNav: false},
		{Name: "About", Path: "/about", Meta: "About", ViewName: "home.gohtml", DisplayInNav: true},
		{Name: "Projects", Path: "/projects", Meta: "Projects", ViewName: "home.gohtml", DisplayInNav: true},
		{Name: "Resume", Path: "/resume", Meta: "Resume", ViewName: "home.gohtml", DisplayInNav: true},
	}

	type pageData struct {
		Routes      []Route
		CurrentPage Route
	}

	for _, route := range routes {
		pageData := pageData{routes, route}
		r.Get(route.Path, controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout.gohtml", route.ViewName)), pageData))
	}

	fmt.Printf("Server listening on port %v\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Printf("<< SERVER ERROR >>")
		panic(err)
	}
}
