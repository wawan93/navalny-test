package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
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

	if val, err := h.Cache.Get(query); err == nil {
		log.Printf("from cache: query: %s", query)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
		return
	}

	result, err := h.Provider.Suggest(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.Cache.Set(query, buf.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("from provider: query: %s", query)

	_, err = io.Copy(w, &buf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
