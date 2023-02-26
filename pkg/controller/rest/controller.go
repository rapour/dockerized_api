package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dockerized/api/pkg/repository"
	tools "dockerized/api/tools/http"
)

type restHandler struct {
	config   tools.Config
	database repository.Repository
}

func NewRestHandler(c tools.Config, db repository.Repository) *restHandler {
	return &restHandler{
		config:   c,
		database: db,
	}
}

func (h *restHandler) Serve() error {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", h.Ping)
	mux.HandleFunc("/view", h.View)

	server := http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%d", h.config.Port),
	}
	fmt.Printf("starting a REST API at %s\n", server.Addr)

	return server.ListenAndServe()
}

func (h *restHandler) Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s", req.Host)
}

func (h *restHandler) View(w http.ResponseWriter, req *http.Request) {

	records, err := h.database.GetRecords()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	serialized, err := json.Marshal(records)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.Write(serialized)
}
