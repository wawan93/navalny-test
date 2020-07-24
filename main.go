package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gopkg.in/webdeskltd/dadata.v2"
)

func main() {
	mux := new(http.ServeMux)
	mux.Handle("/suggest", http.HandlerFunc(handler))

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatalf("can't start listen: %s", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if len(query) <= 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("query must be more than 3 symbols"))
		return
	}

	result, err := suggest(query)
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

func suggest(query string) ([]string, error) {
	token := os.Getenv("API_KEY")
	api := dadata.NewDaData(token, "")
	params := dadata.SuggestRequestParams{Query: query}
	response, err := api.SuggestAddresses(params)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, res := range response {
		result = append(result, res.Value)
	}
	return result, nil
}
