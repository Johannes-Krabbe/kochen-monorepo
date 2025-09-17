package uiHandlers

import (
	"net/http"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/ui/pages"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	component := pages.Homepage()
	component.Render(r.Context(), w)
}
