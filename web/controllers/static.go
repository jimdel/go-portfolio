package controllers

import (
	"net/http"
)

type PageMetaData struct {
	Title string
}

func StaticHandler(tpl Template, pageData any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, pageData)
	}
}
