package handlers

import (
	"net/http"
	"personal_website/components"
)

func TestStyle(w http.ResponseWriter, r *http.Request) {
	component := components.TestStyle()
	component.Render(r.Context(), w)	
}
