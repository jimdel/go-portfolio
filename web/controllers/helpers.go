package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

const SITECONFIG_PATH = "/web/static/site.config.json"

type Test struct {
	String string `json:"string"`
}
type SiteConfig struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Hobbies     []string      `json:"hobbies"`
	Bio         template.HTML `json:"bio"`
	Email       string        `json:"email"`
}

type HomePageData struct {
	PageMetaData PageMetaData
	Config       SiteConfig
}

type BlogPageData struct {
	PageMetaData PageMetaData
	Config       SiteConfig
}

func LoadSiteConfig() (SiteConfig, error) {
	var config SiteConfig
	path, err := os.Getwd()
	if err != nil {
		return config, err
	}
	fmt.Println(path)
	path = filepath.Join(path, SITECONFIG_PATH)
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}

	fmt.Println(file)

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return config, err
	}
	fmt.Println(decoder)
	return config, nil
}
