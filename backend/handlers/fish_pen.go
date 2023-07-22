package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/services"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

type FishPenHandler struct {
	fishPenService services.IFishPenService
}

func NewFishPenHandler(
	fishPenService services.IFishPenService,
) *FishPenHandler {
	return &FishPenHandler{
		fishPenService: fishPenService,
	}
}

type FishPenResponse struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	MakerModelName *string               `json:"makerModelName"`
	Description    *string               `json:"description"`
	Material       *string               `json:"material"`
	NetMaterial    *string               `json:"netMaterial"`
	Category       types.FishPenCategory `json:"category"`
	WidthCM        uint                  `json:"widthCm"`
	LengthCM       uint                  `json:"lengthCm"`
	HeightCM       uint                  `json:"heightCm"`
}

func fishPenResponseFromModel(fp *models.FishPen) *FishPenResponse {
	return &FishPenResponse{
		ID:             fp.ID,
		Name:           fp.Name,
		MakerModelName: fp.MakerModelName,
		Description:    fp.Description,
		Material:       fp.Material,
		NetMaterial:    fp.NetMaterial,
		Category:       fp.Category,
		WidthCM:        fp.WidthCM,
		LengthCM:       fp.LengthCM,
		HeightCM:       fp.HeightCM,
	}
}

type FishPenRequest struct {
	Name           string                `json:"name" validate:"required,min=1,max=255"`
	MakerModelName *string               `json:"makerModelName" validate:"omitempty,min=0,max=255"`
	Description    *string               `json:"description"`
	Material       *string               `json:"material"`
	NetMaterial    *string               `json:"netMaterial"`
	Category       types.FishPenCategory `json:"category" validate:"oneof=fixed floating submersible submersed other"`
	WidthCM        uint                  `json:"widthCm" validate:"gte=0"`
	LengthCM       uint                  `json:"lengthCm" validate:"gte=0"`
	HeightCM       uint                  `json:"heightCm" validate:"gte=0"`
}

func decodeFishPenRequest(r *http.Request) (*FishPenRequest, error) {
	fpr := &FishPenRequest{}
	if err := json.NewDecoder(r.Body).Decode(&fpr); err != nil {
		return nil, err
	}

	// Clean up string values
	fpr.Name = strings.TrimSpace(fpr.Name)
	fpr.MakerModelName = safeTrim(fpr.MakerModelName)
	fpr.Description = safeTrim(fpr.Description)
	fpr.Material = safeTrim(fpr.Material)
	fpr.NetMaterial = safeTrim(fpr.NetMaterial)

	return fpr, nil
}

func (fpr *FishPenRequest) validate() error {
	v := validator.New()
	return v.Struct(fpr)
}

func (fpr *FishPenRequest) toModel(farmID uint) *models.FishPen {
	return &models.FishPen{
		Name:           fpr.Name,
		FarmID:         farmID,
		MakerModelName: fpr.MakerModelName,
		Description:    fpr.Description,
		Material:       fpr.Material,
		NetMaterial:    fpr.NetMaterial,
		Category:       fpr.Category,
		WidthCM:        fpr.WidthCM,
		LengthCM:       fpr.LengthCM,
		HeightCM:       fpr.HeightCM,
	}
}

func (h *FishPenHandler) Create(w http.ResponseWriter, r *http.Request) {
	farmID, err := parsePathID(r, "farmID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}

	fpr, err := decodeFishPenRequest(r)
	if err != nil {
		log.Printf("Error decoding fishpen request: %v", err)
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}
	if err := fpr.validate(); err != nil {
		respondJSON(w, http.StatusBadRequest,
			map[string]string{"error": fmt.Sprintf(
				"%s: %s", types.BadRequestErrorMessage, err.Error(),
			)},
		)
		return
	}

	fp := fpr.toModel(farmID)
	fp, err = h.fishPenService.Create(fp.FarmID, fp)
	if err != nil {
		if errors.Is(err, types.ErrDuplicate) {
			respondJSON(w, http.StatusConflict, map[string]string{"error": fmt.Sprintf("FishPen already exists")})
			return
		}
		if errors.Is(err, types.ErrInvalid) {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Invalid FishPen: %v", err)})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	respondCreated(w, fishPenResponseFromModel(fp))
}

func (h *FishPenHandler) GetAllForFarm(w http.ResponseWriter, r *http.Request) {
	farmID, err := parsePathID(r, "farmID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}

	fishPens, err := h.fishPenService.GetAllForFarm(farmID)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("Farm Not Found: %d", farmID)})
			return
		}
		log.Printf("Error getting fishpens for farm: %v", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	res := []*FishPenResponse{}
	for _, fp := range fishPens {
		res = append(res, fishPenResponseFromModel(fp))
	}

	respondOK(w, res)
}

func (h *FishPenHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	farmID, err := parsePathID(r, "farmID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}
	fishPenID, err := parsePathID(r, "fishPenID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}

	fishpen, err := h.fishPenService.GetSingle(farmID, fishPenID)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("FishPen Not Found: %d", fishPenID)})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	respondOK(w, fishPenResponseFromModel(fishpen))
}

func (h *FishPenHandler) Update(w http.ResponseWriter, r *http.Request) {
	farmID, err := parsePathID(r, "farmID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}
	fishPenID, err := parsePathID(r, "fishPenID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}

	fpr, err := decodeFishPenRequest(r)
	if err != nil {
		log.Printf("Error decoding fishpen request: %v", err)
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}
	if err := fpr.validate(); err != nil {
		respondJSON(w, http.StatusBadRequest,
			map[string]string{"error": fmt.Sprintf(
				"%s: %s", types.BadRequestErrorMessage, err.Error(),
			)},
		)
		return
	}

	// prepare update model
	fp := fpr.toModel(farmID)
	fp.ID = fishPenID

	fp, err = h.fishPenService.Update(fp.FarmID, fp)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("FishPen Not Found: %d", fishPenID)})
			return
		}
		if errors.Is(err, types.ErrDuplicate) {
			respondJSON(w, http.StatusConflict, map[string]string{"error": fmt.Sprintf("FishPen already exists")})
			return
		}
		if errors.Is(err, types.ErrInvalid) {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Invalid FishPen: %v", err)})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	respondOK(w, fishPenResponseFromModel(fp))
}

func (h *FishPenHandler) Delete(w http.ResponseWriter, r *http.Request) {
	farmID, err := parsePathID(r, "farmID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}
	fishPenID, err := parsePathID(r, "fishPenID")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": types.BadRequestErrorMessage})
		return
	}

	err = h.fishPenService.Delete(farmID, fishPenID)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("FishPen Not Found: %d", fishPenID)})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": types.ServerErrorMessage})
		return
	}

	respondOK(w, nil)
}
