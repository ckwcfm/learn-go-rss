package layouts

import (
	"github.com/ckwcfm/learn-go/rss/templates"
)

type mainLayoutTemplate struct {
	File string
}

var MainLayout = &mainLayoutTemplate{
	File: "views/layouts/main.html",
}

func (t *mainLayoutTemplate) Layout() templates.Template[any] {
	return templates.CreateTemplate[any](t.File, "", struct{}{})
}
