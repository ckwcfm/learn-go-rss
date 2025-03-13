package contents

import (
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/templates"
)

type bookTemplate struct {
	File         string
	templateName bookTemplateName
}
type bookTemplateName struct {
	Content     string
	Form        string
	List        string
	OobListItem string
}

var Book = &bookTemplate{
	File: "views/pages/book.html",
	templateName: bookTemplateName{
		Content:     "content",
		Form:        "book-form",
		List:        "book-list",
		OobListItem: "oob-book-item",
	},
}

type BookFormData struct {
	Title  string `required:"true"`
	Author string `required:"true"`
	Error  string
}

type BookListData struct {
	Books []models.Book
}

type OobBookListItemData = models.Book

type BookContentData struct {
	BookFormData BookFormData `required:"true"`
	BookListData BookListData `required:"true"`
}

func (b *bookTemplate) Form(data BookFormData) templates.Template[BookFormData] {
	return templates.CreateTemplate(b.File, b.templateName.Form, data)
}

func (b *bookTemplate) List(data BookListData) templates.Template[BookListData] {
	return templates.CreateTemplate(b.File, b.templateName.List, data)
}

func (b *bookTemplate) OobListItem(data OobBookListItemData) templates.Template[OobBookListItemData] {
	return templates.CreateTemplate(b.File, b.templateName.OobListItem, data)
}

func (b *bookTemplate) Content(data BookContentData) templates.Template[BookContentData] {
	return templates.CreateTemplate(b.File, b.templateName.Content, data)
}
