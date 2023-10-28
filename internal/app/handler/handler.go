package handler

import (
	"net/http"
)

type Handler struct {
}

func (h *Handler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}

}
