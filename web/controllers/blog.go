package controllers

import (
	"net/http"
)

func BlogHandler(tpl Template, meta PageMetaData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := BlogPageData{
			PageMetaData: meta,
		}
		tpl.Execute(w, data)
	}
}
