package web

import (
	"log"
	"net/http"
	"short_url/internal"
)

type routeWithName struct {
	Name string `json:"name"`
	*internal.Route
}

type msg struct {
	Ok bool `json:"ok"`
}

type msgErr struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type msgRoute struct {
	Ok    bool           `json:"ok"`
	Route *routeWithName `json:"route"`
}

type msgRoutes struct {
	Ok     bool             `json:"ok"`
	Routes []*routeWithName `json:"routes"`
	Next   string           `json:"next"`
}

func writeJson(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := internal.Json.NewEncoder(w).Encode(data); err != nil {
		log.Panic(err)
	}
}

func writeJsonOk(w http.ResponseWriter) {
	writeJson(w, &msg{Ok: true}, http.StatusOK)
}

func writeJsonError(w http.ResponseWriter, err string, status int) {
	writeJson(w, &msgErr{Ok: false, Error: err}, status)
}

func writeJsonStorageError(w http.ResponseWriter, err string) {
	log.Printf("[Storage error] %s", err)
	writeJsonError(w, "Storage error", http.StatusInternalServerError)
}

func writeJsonRoute(w http.ResponseWriter, name string, rt *internal.Route) {
	writeJson(w, &msgRoute{
		Ok: true,
		Route: &routeWithName{
			Name:  name,
			Route: rt,
		},
	}, http.StatusOK)
}
