package main

import (
	"address-suggester/internal/cache"
	"address-suggester/internal/providers"
	"address-suggester/internal/routes"

	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("API_KEY")
	provider := &providers.Dadata{
		Token: token,
	}

	c := cache.NewInMemoryCache()

	mux := new(http.ServeMux)
	mux.Handle("/suggest", &routes.MainHandler{
		Provider: provider,
		Cache:    c,
	})

	log.Println("start listen :80")

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatalf("can't start listen: %s", err)
	}
}
