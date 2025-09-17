package componentHandlers

import (
	"net/http"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/ui/components"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils"
)

var count = 0

func PostIncrease(w http.ResponseWriter, r *http.Request) {
	count++
	w.Header().Set("Content-Type", "text/html")

	component := components.Counter(count)
	utils.Render(w, r, component)
}
