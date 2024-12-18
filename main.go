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

	// Routes
	r.Get("/", controllers.HomeHandler(views.Must(views.ParseFS(templates.FS, "layout.gohtml", "home.gohtml")), controllers.PageMetaData{Title: "Home"}))
	r.Get("/blog", controllers.BlogHandler(views.Must(views.ParseFS(templates.FS, "layout.gohtml", "blog.gohtml")), controllers.PageMetaData{Title: "Blog"}))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Not Found!", http.StatusNotFound)
	})

	fmt.Printf("Server listening on port %v\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Printf("<< SERVER ERROR >>")
		panic(err)
	}
}
