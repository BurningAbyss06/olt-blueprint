package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/metalpoch/olt-blueprint/common/constants"
	"github.com/metalpoch/olt-blueprint/common/model"
	"github.com/metalpoch/olt-blueprint/common/pkg/tracking"
	"github.com/metalpoch/olt-blueprint/traffic/repository"
	"gorm.io/gorm"
)

type feedUsecase struct {
	repo     repository.FeedRepository
	telegram tracking.Telegram
}

func NewFeedUsecase(db *gorm.DB, telegram tracking.Telegram) *feedUsecase {
	return &feedUsecase{repository.NewFeedRepository(db), telegram}
}

func (use feedUsecase) GetDevice(id uint) (*model.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := use.repo.GetDevice(ctx, id)
	if err != nil {
		go use.telegram.Notification(
			constants.MODULE_UPDATE,
			constants.CATEGORY_DATABASE,
			fmt.Sprintf("(feedUsecase).GetDevice - use.repo.GetDevice(ctx, %d)", id),
			err,
		)
		return nil, err
	}

	return (*model.Device)(res), err
}

func (use feedUsecase) GetAllDevice() ([]*model.DeviceLite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := use.repo.GetAllDevice(ctx)
	if err != nil {
		go use.telegram.Notification(
			constants.MODULE_UPDATE,
			constants.CATEGORY_DATABASE,
			"(feedUsecase).GetAllDevice - use.repo.GetAllDevice(ctx, %d)",
			err,
		)
		return nil, err
	}

	var devices []*model.DeviceLite
	for _, d := range res {
		devices = append(devices, &model.DeviceLite{
			ID:          d.ID,
			IP:          d.IP,
			SysName:     d.SysName,
			SysLocation: d.SysLocation,
			IsAlive:     d.IsAlive,
			LastCheck:   d.LastCheck,
			CreatedAt:   d.CreatedAt,
			UpdatedAt:   d.UpdatedAt,
		})
	}

	return devices, err
}

func (use feedUsecase) GetInterface(id uint) (*model.Interface, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := use.repo.GetInterface(ctx, id)
	if err != nil {
		go use.telegram.Notification(
			constants.MODULE_UPDATE,
			constants.CATEGORY_DATABASE,
			fmt.Sprintf("(feedUsecase).GetInterface - use.repo.GetInterface(ctx, %d)", id),
			err,
		)
		return nil, err
	}

	return (*model.Interface)(res), err
}
