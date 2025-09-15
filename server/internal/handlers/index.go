package handlers

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://old.kochen.app")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
