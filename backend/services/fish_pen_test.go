package services

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/mocks"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

func TestServices_FishPen_Create(t *testing.T) {
	tests := []struct {
		name         string
		prepareMocks func(mf *mocks.MockIFishPenRepository, fishPen *models.FishPen)
		wantErr      error
	}{
		{
			name: "ok",
			prepareMocks: func(mf *mocks.MockIFishPenRepository, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByNameForFarm(fishPen.FarmID, fishPen.Name).Times(1).Return(false, nil)
				mf.EXPECT().Create(fishPen).Times(1).DoAndReturn(func(fishPen *models.FishPen) (*models.FishPen, error) {
					// Auto-Generated fields by db / gorm
					fishPen.ID = 3
					fishPen.CreatedAt = time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC)
					fishPen.UpdatedAt = time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC)
					return fishPen, nil
				})
			},
			wantErr: nil,
		},
		{
			name: "err duplicate",
			prepareMocks: func(mf *mocks.MockIFishPenRepository, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByNameForFarm(fishPen.FarmID, fishPen.Name).Times(1).Return(true, nil)
				mf.EXPECT().Create(fishPen).Times(0)
			},
			wantErr: types.ErrDuplicate,
		},
		{
			name: "err unknown exists by name",
			prepareMocks: func(mf *mocks.MockIFishPenRepository, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByNameForFarm(fishPen.FarmID, fishPen.Name).Times(1).Return(false, gorm.ErrInvalidDB)
				mf.EXPECT().Create(fishPen).Times(0)
			},
			wantErr: gorm.ErrInvalidDB,
		},
		{
			name: "err invalid create",
			prepareMocks: func(mf *mocks.MockIFishPenRepository, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByNameForFarm(fishPen.FarmID, fishPen.Name).Times(1).Return(false, nil)
				mf.EXPECT().Create(fishPen).Times(1).Return(nil, types.ErrInvalid)
			},
			wantErr: types.ErrInvalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			fishPen := &models.FishPen{
				Model: gorm.Model{
					ID: 1,
				},
				FarmID:   1,
				Category: types.FishPenCategoryFixed,
				Name:     "fish pen 1",
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			farmRepo := mocks.NewMockIFarmRepository(ctrl)
			fishPenRepo := mocks.NewMockIFishPenRepository(ctrl)

			tt.prepareMocks(fishPenRepo, fishPen)

			// call
			s := NewFishPenService(farmRepo, fishPenRepo)
			_, err := s.Create(fishPen.FarmID, fishPen)

			// check
			if err != tt.wantErr {
				t.Errorf("FishPenService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestServices_FishPen_GetAllForFarm(t *testing.T) {
	tests := []struct {
		name         string
		farmID       uint
		prepareMocks func(
			mf *mocks.MockIFarmRepository,
			mfp *mocks.MockIFishPenRepository,
			farmID uint,
			fishPens []*models.FishPen)
		want    []*models.FishPen
		wantErr error
	}{
		{
			name: "ok",
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, fishPens []*models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetAllForFarm(farmID).Times(1).Return(fishPens, nil)
			},
			wantErr: nil,
		},
		{
			name: "err farm exists not found",
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, _ []*models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(false, nil)
				mfp.EXPECT().GetAllForFarm(farmID).Times(0)
			},
			wantErr: types.ErrNotFound,
		},
		{
			name: "err farm exists unknown",
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, _ []*models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(false, gorm.ErrInvalidDB)
				mfp.EXPECT().GetAllForFarm(farmID).Times(0)
			},
			wantErr: gorm.ErrInvalidDB,
		},
		{
			name: "err unknown",
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, _ []*models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetAllForFarm(farmID).Times(1).Return(nil, gorm.ErrInvalidDB)
			},
			wantErr: gorm.ErrInvalidDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			fishPens := []*models.FishPen{
				{
					Model: gorm.Model{
						ID: 1,
					},
					FarmID:   1,
					Category: types.FishPenCategoryFixed,
					Name:     "fish pen 1",
				},
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			farmRepo := mocks.NewMockIFarmRepository(ctrl)
			fishPenRepo := mocks.NewMockIFishPenRepository(ctrl)

			tt.prepareMocks(farmRepo, fishPenRepo, tt.farmID, fishPens)

			// call
			s := NewFishPenService(farmRepo, fishPenRepo)
			got, err := s.GetAllForFarm(tt.farmID)

			// check
			if tt.wantErr == nil {
				assert.DeepEqual(t, got, fishPens)
			} else {
				if err != tt.wantErr {
					t.Errorf("FishPenService.GetAllForFarm() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestServices_FishPen_GetSingle(t *testing.T) {
	tests := []struct {
		name         string
		farmID       uint
		fishPenID    uint
		prepareMocks func(
			mf *mocks.MockIFarmRepository,
			mfp *mocks.MockIFishPenRepository,
			farmID, fishPenID uint,
			fishPen *models.FishPen)
		want    *models.FishPen
		wantErr error
	}{
		{
			name:      "ok",
			farmID:    1,
			fishPenID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID, fishPenID uint,
				fishPen *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetByID(fishPenID).Times(1).Return(fishPen, nil)
			},
			wantErr: nil,
		},
		{
			name:      "err farm exists not found",
			farmID:    999,
			fishPenID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID, fishPenID uint, _ *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(false, nil)
				mfp.EXPECT().GetByID(fishPenID).Times(0)
			},
			want:    nil,
			wantErr: types.ErrNotFound,
		},
		{
			name:      "err farm exists unknown",
			farmID:    1,
			fishPenID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID, fishPenID uint, _ *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(false, gorm.ErrInvalidDB)
				mfp.EXPECT().GetByID(fishPenID).Times(0)
			},
			want:    nil,
			wantErr: gorm.ErrInvalidDB,
		},
		{
			name:      "err fish pen exists not found",
			farmID:    1,
			fishPenID: 999,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID, fishPenID uint, _ *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetByID(fishPenID).Times(1).Return(nil, types.ErrNotFound)
			},
			want:    nil,
			wantErr: types.ErrNotFound,
		},
		{
			name:      "err fish pen exists unknown",
			farmID:    1,
			fishPenID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID, fishPenID uint, _ *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetByID(fishPenID).Times(1).Return(nil, gorm.ErrInvalidDB)
			},
			want:    nil,
			wantErr: gorm.ErrInvalidDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			farmRepo := mocks.NewMockIFarmRepository(ctrl)
			fishPenRepo := mocks.NewMockIFishPenRepository(ctrl)

			fishPen := &models.FishPen{
				Model: gorm.Model{
					ID: 1,
				},
				FarmID:   1,
				Category: types.FishPenCategoryFixed,
				Name:     "fish pen 1",
			}

			tt.prepareMocks(farmRepo, fishPenRepo, tt.farmID, tt.fishPenID, fishPen)

			// call
			s := NewFishPenService(farmRepo, fishPenRepo)
			got, err := s.GetSingle(tt.farmID, tt.fishPenID)

			// check
			if tt.wantErr == nil {
				assert.DeepEqual(t, got, fishPen)
			} else {
				if err != tt.wantErr {
					t.Errorf("FishPenService.GetSingle() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestServices_FishPen_Update(t *testing.T) {
	tests := []struct {
		name         string
		farmID       uint
		prepareMocks func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, fishPen *models.FishPen)
		wantErr      error
	}{
		{
			name:   "ok no existing fish pen by name",
			farmID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetByNameForFarm(farmID, fishPen.Name).Times(1).Return(nil, types.ErrNotFound)
				mfp.EXPECT().Update(fishPen).Times(1).DoAndReturn(
					func(fp *models.FishPen) (*models.FishPen, error) {
						fp.UpdatedAt = time.Now()
						return fp, nil
					})
			},
			wantErr: nil,
		},
		{
			name:   "ok existing fish pen by name is self",
			farmID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)
				mfp.EXPECT().GetByNameForFarm(farmID, fishPen.Name).Times(1).Return(fishPen, nil)
				mfp.EXPECT().Update(fishPen).Times(1).DoAndReturn(
					func(fp *models.FishPen) (*models.FishPen, error) {
						fp.UpdatedAt = time.Now()
						return fp, nil
					})
			},
			wantErr: nil,
		},
		{
			name:   "err existing fish pen by name is not self",
			farmID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(true, nil)

				// Same name but ID is different.
				mfp.EXPECT().GetByNameForFarm(farmID, fishPen.Name).Times(1).Return(&models.FishPen{
					Model: gorm.Model{
						ID: 999,
					},
					FarmID:   1,
					Category: types.FishPenCategoryFixed,
					Name:     fishPen.Name,
				}, nil)

				// Update not called
				mfp.EXPECT().Update(fishPen).Times(0)
			},
			wantErr: types.ErrDuplicate,
		},
		{
			name:   "err farm not found",
			farmID: 1,
			prepareMocks: func(mf *mocks.MockIFarmRepository, mfp *mocks.MockIFishPenRepository, farmID uint, fishPen *models.FishPen) {
				mf.EXPECT().ExistsByID(farmID).Times(1).Return(false, nil)
				mfp.EXPECT().GetByNameForFarm(farmID, fishPen.Name).Times(0)
				mfp.EXPECT().Update(fishPen).Times(0)
			},
			wantErr: types.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			fishPen := &models.FishPen{
				Model: gorm.Model{
					ID: 1,
				},
				FarmID:   1,
				Category: types.FishPenCategoryFixed,
				Name:     "fish pen 1",
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			farmRepo := mocks.NewMockIFarmRepository(ctrl)
			fishPenRepo := mocks.NewMockIFishPenRepository(ctrl)

			tt.prepareMocks(farmRepo, fishPenRepo, tt.farmID, fishPen)

			// call
			s := NewFishPenService(farmRepo, fishPenRepo)
			_, err := s.Update(tt.farmID, fishPen)

			// check
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("FishPenService.Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
