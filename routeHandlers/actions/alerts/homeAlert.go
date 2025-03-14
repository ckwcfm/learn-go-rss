package alerts

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/templates/compontents"
)

func ActionHomeAlert(w http.ResponseWriter, r *http.Request) {
	message := compontents.HomeAlertData{
		Message: "Hello, World!",
		Action:  "Close!!!",
	}
	alert := compontents.HomeAlert.Alert(message)
	alert.Render(w, r)
}
