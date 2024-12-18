package controllers

import (
	"fmt"
	"net/http"
)

func HomeHandler(tpl Template, meta PageMetaData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: us os to grab cwd and load site.json

		config, err := LoadSiteConfig()
		if err != nil {
			http.Error(w, "Error loading site config", http.StatusInternalServerError)
		}
		data := HomePageData{
			PageMetaData: meta,
			Config:       config,
		}
		fmt.Println(data)
		tpl.Execute(w, data)
	}
}
