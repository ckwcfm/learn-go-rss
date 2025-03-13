package pages

import (
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/templates"
)

// * Template Names
type templateName struct {
	Form string
	List string
	Oob  string
}

// * Template
type postTemplate struct {
	File         string
	templateName templateName
}

// * Post Template
var Post = &postTemplate{
	File: "views/pages/post.html",
	templateName: templateName{
		Form: "post-form",
		List: "post-list",
		Oob:  "oob-post-item",
	},
}

// * Data Types
type FormData struct {
	Title   string
	Content string
}

type PostFormData struct {
	Form  FormData
	Error string
}

type PostListData struct {
	Posts []models.Post
}

type OobPostItemData = models.Post

// * Template Functions
func (f *postTemplate) Form(data PostFormData) templates.Template[PostFormData] {
	return templates.CreateTemplate(f.File, f.templateName.Form, data)
}

func (f *postTemplate) List(data PostListData) templates.Template[PostListData] {
	return templates.CreateTemplate(f.File, f.templateName.List, data)
}

func (f *postTemplate) OobPostItem(data OobPostItemData) templates.Template[OobPostItemData] {
	return templates.CreateTemplate(f.File, f.templateName.Oob, data)
}
