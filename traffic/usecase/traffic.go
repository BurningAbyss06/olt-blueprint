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

type trafficUsecase struct {
	repo     repository.TrafficRepository
	telegram tracking.Telegram
}

func NewTrafficUsecase(db *gorm.DB, telegram tracking.Telegram) *trafficUsecase {
	return &trafficUsecase{repository.NewTrafficRepository(db), telegram}
}

func (use trafficUsecase) GetTrafficByInterface(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := use.repo.GetTrafficByInterface(ctx, id, date)
	if err != nil {
		go use.telegram.Notification(
			constants.MODULE_UPDATE,
			constants.CATEGORY_DATABASE,
			fmt.Sprintf("(trafficUsecase).GetTrafficByInterface - use.repo.GetTrafficByInterface(ctx, %d, %v)", id, date),
			err,
		)
		return nil, err
	}

	traffics := []*model.TrafficResponse{}
	for _, e := range res {
		traffics = append(traffics, &model.TrafficResponse{
			Date:      e.Date,
			Bandwidth: e.Bandwidth,
			In:        e.In,
			Out:       e.Out,
		})
	}

	return traffics, err
}

func (use trafficUsecase) GetTrafficByDevice(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := use.repo.GetTrafficByDevice(ctx, id, date)
	if err != nil {
		go use.telegram.Notification(
			constants.MODULE_UPDATE,
			constants.CATEGORY_DATABASE,
			fmt.Sprintf("(trafficUsecase).GetTrafficByDevice - use.repo.GetTrafficByDevice(ctx, %d, %v)", id, date),
			err,
		)
		return nil, err
	}

	traffics := []*model.TrafficResponse{}
	for _, e := range res {
		traffics = append(traffics, &model.TrafficResponse{
			Date:      e.Date,
			Bandwidth: e.Bandwidth,
			In:        e.In,
			Out:       e.Out,
		})
	}

	return traffics, err
}

func (use trafficUsecase) GetTrafficByFat(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := use.repo.GetTrafficByFat(ctx, id, date)
	if err != nil {
		go use.telegram.Notification(
			constants.MODULE_UPDATE,
			constants.CATEGORY_DATABASE,
			fmt.Sprintf("(trafficUsecase).GetTrafficByFat - use.repo.GetTrafficByFat(ctx, %d, %v)", id, date),
			err,
		)
		return nil, err
	}

	traffics := []*model.TrafficResponse{}
	for _, e := range res {
		traffics = append(traffics, &model.TrafficResponse{
			Date:      e.Date,
			Bandwidth: e.Bandwidth,
			In:        e.In,
			Out:       e.Out,
		})
	}

	return traffics, err
}
