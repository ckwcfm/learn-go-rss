package partials

import (
	"github.com/ckwcfm/learn-go/rss/templates"
)

type sidenavTemplate struct {
	File string
}

var Sidenav = &sidenavTemplate{
	File: "views/partials/sidenav.html",
}

func (t *sidenavTemplate) Sidenav() templates.Template[any] {
	return templates.CreateTemplate[any](t.File, "sidebar", struct{}{})
}
