package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin/render"
)

// Render :
type Render struct {
	Templates   map[string]*template.Template
	TemplateDir string
	Debug       bool
}

// NewRender :
func NewRender(dir string) Render {
	r := Render{

		Templates:   map[string]*template.Template{},
		Debug:       false,
		TemplateDir: dir,
	}

	return r
}

// AddSingle :
func (r Render) AddSingle(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}

	r.Templates[name] = tmpl

	if r.Debug {
		log.Printf("Adding template %v\n", name)
	}
}

// AddDirectory :
func (r Render) AddDirectory(dir string) {
	err := filepath.Walk(fmt.Sprintf("%s/%s", r.TemplateDir, dir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error walking directory %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && info.Name() != "layout.html" {
			r.AddSingle(r.formatTemplateName(path), template.Must(template.ParseFiles(path)))
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

// AddDirectoryWithLayout :
func (r Render) AddDirectoryWithLayout(dir, layout string) {
	err := filepath.Walk(fmt.Sprintf("%s/%s", r.TemplateDir, dir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error walking directory %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && info.Name() != "layout.html" {
			r.AddSingle(r.formatTemplateName(path), template.Must(template.ParseFiles(layout, path)))
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

// Instance :
func (r Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.Templates[name],
		Data:     data,
	}
}

func (r Render) formatTemplateName(raw string) string {
	output := strings.Replace(raw, r.TemplateDir, "", 1)
	output = strings.Replace(output, filepath.Ext(output), "", 1)
	return output[1:]
}
