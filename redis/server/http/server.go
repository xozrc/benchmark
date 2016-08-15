package http

import (
	"net/http"

	"github.com/xozrc/benchmark/redis/server/http/api"
)

func Listen(addr string) error {
	// router := &mux.Router{}
	// s := router.PathPrefix("/api").Subrouter()
	// s.Handle("/set", http.HandlerFunc(api.Set)).Methods("POST")

	return http.ListenAndServe(addr, http.HandlerFunc(api.Set))
}
