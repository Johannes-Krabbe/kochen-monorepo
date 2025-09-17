package pageHandlers

import (
	"net/http"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/ui/pages"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	component := pages.Homepage()
	utils.Render(w, r, component)
}
