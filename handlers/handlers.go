package handlers

import (
	"net/http"
	"personal_website/components"
)

func TestStyle(w http.ResponseWriter, r *http.Request) {
	component := components.Desktop()
	component.Render(r.Context(), w)	
}
