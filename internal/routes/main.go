package routes

import (
	"encoding/json"
	"net/http"
)

type cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

type provider interface {
	Suggest(query string) ([]string, error)
}

type MainHandler struct {
	Cache    cache
	Provider provider
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if len(query) <= 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("query must be more than 3 symbols"))
		return
	}

	result, err := h.Provider.Suggest(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
