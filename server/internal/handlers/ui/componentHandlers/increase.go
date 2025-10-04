package componentHandlers

import (
	"net/http"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/ui/components"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils"
)

var Count = 0

func PostIncrease(w http.ResponseWriter, r *http.Request) {
	Count++
	w.Header().Set("Content-Type", "text/html")

	component := components.Counter(Count)
	utils.Render(w, r, component)
}
