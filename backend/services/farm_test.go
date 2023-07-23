package services

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/mocks"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"

	"gorm.io/gorm"
	"gotest.tools/v3/assert"
)

var commonFarms = []*models.Farm{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Name: "farm 1",
	},
	{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		Name: "farm 2",
	},
}

func TestServices_Farm_GetAll(t *testing.T) {
	tests := []struct {
		name         string
		prepareMocks func(mf *mocks.MockIFarmRepository)
		want         []*models.Farm
		wantErr      bool
	}{
		{
			name: "success",
			prepareMocks: func(mf *mocks.MockIFarmRepository) {
				mf.EXPECT().All().Times(1).Return(commonFarms, nil)
			},
			want:    commonFarms,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mocks.NewMockIFarmRepository(ctrl)
			tt.prepareMocks(repo)

			s := NewFarmService(repo)
			got, err := s.GetAll()

			assert.Equal(t, tt.wantErr, err != nil, "FarmService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			assert.DeepEqual(t, tt.want, got)
		})
	}
}

func TestServices_Farm_GetByID(t *testing.T) {
	tests := []struct {
		name         string
		prepareMocks func(mf *mocks.MockIFarmRepository)
		id           uint
		want         *models.Farm
		wantErr      bool
	}{
		{
			name: "success",
			prepareMocks: func(mf *mocks.MockIFarmRepository) {
				mf.EXPECT().GetByID(uint(1)).Times(1).Return(commonFarms[0], nil)
			},
			id:      uint(1),
			want:    commonFarms[0],
			wantErr: false,
		},
		{
			name: "not found",
			prepareMocks: func(mf *mocks.MockIFarmRepository) {
				mf.EXPECT().GetByID(uint(1)).Times(1).Return(nil, gorm.ErrRecordNotFound)
			},
			id:      uint(1),
			want:    nil,
			wantErr: true,
		},
		{
			name: "other error",
			prepareMocks: func(mf *mocks.MockIFarmRepository) {
				mf.EXPECT().GetByID(uint(1)).Times(1).Return(nil, gorm.ErrInvalidData)
			},
			id:      uint(1),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mocks.NewMockIFarmRepository(ctrl)
			tt.prepareMocks(repo)

			s := NewFarmService(repo)
			got, err := s.GetByID(tt.id)

			assert.Equal(t, tt.wantErr, err != nil, "FarmService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
			assert.DeepEqual(t, tt.want, got)
		})
	}
}
