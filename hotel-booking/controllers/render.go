package controllers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func resolveViewsDir() string {
	candidates := []string{"views", filepath.Join("hotel-booking", "views")}
	for _, dir := range candidates {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir
		}
	}
	return "views"
}

func renderTemplate(w http.ResponseWriter, name string, data any) {
	viewsDir := resolveViewsDir()
	layout := filepath.Join(viewsDir, "layout.html")
	target := filepath.Join(viewsDir, name)
	tmpl, err := template.ParseFiles(layout, target)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "render error: "+err.Error(), http.StatusInternalServerError)
	}
}
