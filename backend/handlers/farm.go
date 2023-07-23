package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/services"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

type FarmHandler struct {
	farmService services.IFarmService
}

func NewFarmHandler(
	farmService services.IFarmService,
) *FarmHandler {
	return &FarmHandler{
		farmService: farmService,
	}
}

type FarmResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (h *FarmHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	fs, err := h.farmService.GetAll()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	res := []*FarmResponse{}
	for _, f := range fs {
		res = append(res, &FarmResponse{
			ID:   f.ID,
			Name: f.Name,
		})
	}

	respondOK(w, res)
}

func (h *FarmHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	farmID, err := parsePathID(r, "farmID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		return
	}

	f, err := h.farmService.GetByID(uint(farmID))
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("Farm not found: %d", farmID)})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	respondOK(w, &FarmResponse{
		ID:   f.ID,
		Name: f.Name,
	})
}
