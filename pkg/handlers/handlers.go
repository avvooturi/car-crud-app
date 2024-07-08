package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/avvooturi/car-crud-app/pkg/storage"
)

type CarHandler struct {
	db storage.DB
}

func (h *CarHandler) HandleFetchOne(w http.ResponseWriter, r *http.Request) error {
	res, err := h.db.FetchOne("123ABC")

	if err != nil {
		return err
	}

	b, err := json.Marshal(res)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.Write(b)
	return nil
}

func New(_db storage.DB) *CarHandler {
	return &CarHandler{db: _db}
}
