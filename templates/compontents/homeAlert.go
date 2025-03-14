package compontents

import (
	"github.com/ckwcfm/learn-go/rss/templates"
)

type HomeAlertTemplate struct {
	File string
}

var HomeAlert = &HomeAlertTemplate{
	File: "views/compontents/alerts/homeAlert.html",
}

type HomeAlertData struct {
	Message string
	Action  string
}

func (t *HomeAlertTemplate) Alert(data HomeAlertData) templates.Template[HomeAlertData] {
	return templates.CreateTemplate(t.File, "home-alert", data)
}
