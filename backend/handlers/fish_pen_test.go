package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/mocks"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

func TestHandlers_FishPenRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		fpr     *FishPenRequest
		wantErr bool
	}{
		{
			name: "valid",
			fpr: &FishPenRequest{
				Name:     "test",
				Category: "fixed",
			},
			wantErr: false,
		},
		{
			name: "name empty",
			fpr: &FishPenRequest{
				Name:     "",
				Category: "fixed",
			},
			wantErr: true,
		},
		{
			name: "name too long",
			fpr: &FishPenRequest{
				Name:     fmt.Sprintf("%256s", "a"),
				Category: "fixed",
			},
			wantErr: true,
		},
		{
			name: "category unknown",
			fpr: &FishPenRequest{
				Name:     "test",
				Category: "banana",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fpr.validate(); (err != nil) != tt.wantErr {
				t.Errorf("FishPenRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

	}
}

type fishPenMatcher struct {
	expected *models.FishPen
}

func (f fishPenMatcher) Matches(x any) bool {
	fpr, ok := x.(*models.FishPen)
	if !ok {
		return false
	}
	return f.expected.Name == fpr.Name && f.expected.Category == fpr.Category
}
func (m fishPenMatcher) String() string {
	return fmt.Sprintf("is equal to %v", m.expected)
}

func TestHandlers_FishPen_Create(t *testing.T) {
	tests := []struct {
		name           string
		farmID         string
		fishPenRequest *FishPenRequest
		prepareMocks   func(
			fpr *FishPenRequest,
			fps *mocks.MockIFishPenService,
		)
		wantCode int
	}{
		{
			name:   "ok",
			farmID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "fixed",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				farmID := uint(1)
				fp := fpr.toModel(farmID)

				fps.EXPECT().
					Create(farmID, fishPenMatcher{expected: fp}).
					Return(fp, nil)
			},
			wantCode: http.StatusCreated,
		},
		{
			name:   "invalid farm ID",
			farmID: "aaaa",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "fixed",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name:   "invalid request body",
			farmID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "",
				Category: "fixed",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name:   "service error",
			farmID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "fixed",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				farmID := uint(1)
				fp := fpr.toModel(farmID)

				fps.EXPECT().Create(farmID, fishPenMatcher{expected: fp}).Return(nil, errors.New("test error"))
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// preapre
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fps := mocks.NewMockIFishPenService(ctrl)
			tt.prepareMocks(tt.fishPenRequest, fps)

			h := NewFishPenHandler(fps)

			body, err := json.Marshal(tt.fishPenRequest)
			if err != nil {
				t.Fatal(err)
			}
			r := httptest.NewRequest(
				http.MethodPost,
				fmt.Sprintf("/api/farms/%s/fishpens", tt.farmID),
				bytes.NewBuffer(body),
			)
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r = mux.SetURLVars(r, map[string]string{
				"farmID": fmt.Sprint(tt.farmID),
			})

			// call
			h.Create(w, r)

			// check
			if w.Code != tt.wantCode {
				t.Errorf("FishPenHandler.Create() code = %v, wantCode %v", w.Code, tt.wantCode)
			}
		})
	}
}

func TestHandlers_FishPen_GetAllForFarm(t *testing.T) {
	tests := []struct {
		name         string
		farmID       string
		prepareMocks func(
			farmID string,
			fps *mocks.MockIFishPenService,
		)
		wantCode int
	}{
		{
			name:   "ok",
			farmID: "1",
			prepareMocks: func(
				farmID string,
				fps *mocks.MockIFishPenService,
			) {
				farmIDInt := uint(1)

				fps.EXPECT().
					GetAllForFarm(farmIDInt).
					Return([]*models.FishPen{}, nil)
			},
			wantCode: http.StatusOK,
		},
		{
			name:   "invalid farm ID",
			farmID: "aaaa",
			prepareMocks: func(
				farmID string,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().GetAllForFarm(gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fps := mocks.NewMockIFishPenService(ctrl)
			tt.prepareMocks(tt.farmID, fps)

			h := NewFishPenHandler(fps)

			r := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/api/farms/%s/fishpens", tt.farmID),
				nil,
			)

			w := httptest.NewRecorder()
			r = mux.SetURLVars(r, map[string]string{
				"farmID": tt.farmID,
			})

			// call
			h.GetAllForFarm(w, r)

			// check
			if w.Code != tt.wantCode {
				t.Errorf("Handlers.FishPen.GetAllForFarm() code = %v, wantCode %v", w.Code, tt.wantCode)
			}
		})
	}
}

func TestHandlers_FishPen_GetByID(t *testing.T) {
	tests := []struct {
		name         string
		farmID       uint
		fishPenID    uint
		prepareMocks func(
			farmID, fishPenID uint,
			fps *mocks.MockIFishPenService,
		)
		wantCode int
	}{
		{
			name:      "ok",
			farmID:    1,
			fishPenID: 1,
			prepareMocks: func(
				farmID, fishPenID uint,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().GetSingle(farmID, fishPenID).Return(&models.FishPen{
					Model: gorm.Model{
						ID: 1,
					},
				}, nil)
			},
			wantCode: http.StatusOK,
		},
		{
			name:      "not found",
			farmID:    1,
			fishPenID: 1,
			prepareMocks: func(
				farmID, fishPenID uint,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().GetSingle(farmID, fishPenID).Return(nil, types.ErrNotFound)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name:      "Unknown err",
			farmID:    1,
			fishPenID: 1,
			prepareMocks: func(
				farmID, fishPenID uint,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().GetSingle(farmID, fishPenID).Return(nil, errors.New("some err"))
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fps := mocks.NewMockIFishPenService(ctrl)
			tt.prepareMocks(tt.farmID, tt.fishPenID, fps)

			h := NewFishPenHandler(fps)

			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/farms/%d/fishpens/%d", tt.fishPenID, tt.farmID), nil)
			w := httptest.NewRecorder()
			r = mux.SetURLVars(r, map[string]string{
				"farmID":    fmt.Sprint(tt.farmID),
				"fishPenID": fmt.Sprint(tt.fishPenID),
			})

			// call
			h.GetByID(w, r)

			// check
			if w.Code != tt.wantCode {
				t.Errorf("FishPenHandler.GetByID() code = %v, wantCode %v", w.Code, tt.wantCode)
			}
		})
	}
}

func TestHandlers_FishPen_Update(t *testing.T) {
	tests := []struct {
		name           string
		farmID         string
		fishPenID      string
		fishPenRequest *FishPenRequest
		prepareMocks   func(
			fpr *FishPenRequest,
			fps *mocks.MockIFishPenService,
		)
		wantCode int
	}{
		{
			name:      "ok",
			farmID:    "1",
			fishPenID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "other",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				farmID := uint(1)
				fishPenID := uint(1)

				fp := fpr.toModel(farmID)
				fp.FarmID = farmID
				fp.ID = fishPenID

				fps.EXPECT().
					Update(uint(farmID), fishPenMatcher{expected: fp}).
					Return(fp, nil)
			},
			wantCode: http.StatusOK,
		},
		{
			name:      "validation error",
			farmID:    "1",
			fishPenID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "banana",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name:      "not found",
			farmID:    "1",
			fishPenID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "other",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().
					Update(uint(1), gomock.Any()).
					Return(nil, types.ErrNotFound)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name:      "duplicate",
			farmID:    "1",
			fishPenID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "other",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().
					Update(uint(1), gomock.Any()).
					Return(nil, types.ErrDuplicate)
			},
			wantCode: http.StatusConflict,
		},
		{
			name:      "unknown error",
			farmID:    "1",
			fishPenID: "1",
			fishPenRequest: &FishPenRequest{
				Name:     "test",
				Category: "other",
			},
			prepareMocks: func(
				fpr *FishPenRequest,
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().
					Update(uint(1), gomock.Any()).
					Return(nil, errors.New("some err"))
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fps := mocks.NewMockIFishPenService(ctrl)
			tt.prepareMocks(tt.fishPenRequest, fps)

			h := NewFishPenHandler(fps)

			body, err := json.Marshal(tt.fishPenRequest)
			if err != nil {
				t.Fatal(err)
			}
			r := httptest.NewRequest(
				http.MethodPut,
				fmt.Sprintf("/api/farms/%s/fishpens/%s", tt.fishPenID, tt.farmID),
				bytes.NewBuffer(body),
			)
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r = mux.SetURLVars(r, map[string]string{
				"farmID":    fmt.Sprint(tt.farmID),
				"fishPenID": fmt.Sprint(tt.fishPenID),
			})

			// call
			h.Update(w, r)

			// check
			if w.Code != tt.wantCode {
				t.Errorf("FishPenHandler.Update() code = %v, wantCode %v", w.Code, tt.wantCode)
			}
		})
	}
}

// delete test
func TestHandlers_FishPen_Delete(t *testing.T) {
	tests := []struct {
		name         string
		farmID       string
		fishPenID    string
		prepareMocks func(
			fps *mocks.MockIFishPenService,
		)
		wantCode int
	}{
		{
			name:      "ok",
			farmID:    "1",
			fishPenID: "1",
			prepareMocks: func(
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().
					Delete(uint(1), uint(1)).
					Return(nil)
			},
			wantCode: http.StatusNoContent,
		},
		{
			name:      "not found",
			farmID:    "1",
			fishPenID: "1",
			prepareMocks: func(
				fps *mocks.MockIFishPenService,
			) {
				fps.EXPECT().
					Delete(uint(1), uint(1)).
					Return(types.ErrNotFound)
			},
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fps := mocks.NewMockIFishPenService(ctrl)
			tt.prepareMocks(fps)

			h := NewFishPenHandler(fps)

			r := httptest.NewRequest(
				http.MethodDelete,
				fmt.Sprintf("/api/farms/%s/fishpens/%s", tt.fishPenID, tt.farmID),
				nil,
			)

			w := httptest.NewRecorder()
			r = mux.SetURLVars(r, map[string]string{
				"farmID":    fmt.Sprint(tt.farmID),
				"fishPenID": fmt.Sprint(tt.fishPenID),
			})

			// call
			h.Delete(w, r)

			// check
			if w.Code != tt.wantCode {
				t.Errorf("FishPenHandler.Delete() code = %v, wantCode %v", w.Code, tt.wantCode)
			}
		})
	}
}
