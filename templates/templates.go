package templates

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateKey struct {
	file         string
	templateName string
}

func NewTemplateKey(file string, templateName string) TemplateKey {
	return TemplateKey{
		file:         file,
		templateName: templateName,
	}
}

func (k *TemplateKey) GetFile() string {
	return k.file
}

func (k *TemplateKey) GetTemplateName() string {
	return k.templateName
}

type Template[T any] struct {
	templ *template.Template
	Data  T
	Key   TemplateKey
}

func (t *Template[T]) Render(w http.ResponseWriter, r *http.Request) {
	t.templ = template.Must(template.ParseFiles(t.Key.file))
	t.templ.ExecuteTemplate(w, t.Key.templateName, t.Data)
}

// Helper function at package level instead of method
func CreateTemplate[T any](file string, templateName string, initialData T) Template[T] {
	log.Println("Creating template: [templates.go CreateTemplate | line 41]")
	return Template[T]{
		Key:  NewTemplateKey(file, templateName),
		Data: initialData,
	}
}

// ExecuteTemplates executes multiple templates with the data from the last template
func ExecuteTemplates(w http.ResponseWriter, data any, templates ...Template[any]) error {
	files := make([]string, len(templates))
	for i, t := range templates {
		files[i] = t.Key.GetFile()
	}

	tmpl := template.Must(template.ParseFiles(files...))
	return tmpl.Execute(w, data)
}
