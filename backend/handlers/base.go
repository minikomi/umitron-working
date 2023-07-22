package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondOK(w http.ResponseWriter, payload interface{}) {
	respondJSON(w, http.StatusOK, payload)
}

func respondCreated(w http.ResponseWriter, payload interface{}) {
	respondJSON(w, http.StatusCreated, payload)
}

func respondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func safeTrim(s *string) *string {
	if s == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*s)

	return &trimmed
}

func parsePathID(r *http.Request, key string) (uint, error) {
	vars := mux.Vars(r)
	valStr, ok := vars[key]
	if !ok {
		return 0, fmt.Errorf("missing key in path %s", key)
	}
	val, err := strconv.ParseUint(valStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}
