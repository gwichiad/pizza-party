package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var requestTimeout = time.Second * 5

func (api *api) ingestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := readAndDecode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	if err := api.config.store.insert(ctx, data); err != nil {
		fmt.Printf("insert failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
