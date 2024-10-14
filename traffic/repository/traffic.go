package repository

import (
	"context"

	"github.com/metalpoch/olt-blueprint/common/entity"
	"github.com/metalpoch/olt-blueprint/common/model"
	"gorm.io/gorm"
)

type trafficRepository struct {
	db *gorm.DB
}

func NewTrafficRepository(db *gorm.DB) *trafficRepository {
	return &trafficRepository{db}
}

func (repo trafficRepository) GetTrafficByInterface(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Where("interface_id = ? AND date BETWEEN ? AND ?", id, date.InitDate, date.EndDate).
		Order("date").
		Find(&traffic).
		Error
	return traffic, err
}

func (repo trafficRepository) GetTrafficByDevice(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Model(&entity.Traffic{}).
		Select("date, SUM(\"in\") AS \"in\", SUM(out) AS out, SUM(bandwidth) as bandwidth").
		Joins("JOIN interfaces ON interfaces.id = traffics.interface_id").
		Where("interfaces.device_id = ? AND traffics.date BETWEEN ? AND ?", id, date.InitDate, date.EndDate).
		Group("date").
		Order("date").
		Find(&traffic).
		Error
	return traffic, err
}

func (repo trafficRepository) GetTrafficByFat(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Select("date, \"in\", out, bandwidth").
		Where("date BETWEEN ? AND ?", date.InitDate, date.EndDate).
		Joins("JOIN fats ON fats.interface_id = traffics.interface_id").
		Where("fats.id = ?", id).
		Order("date").
		Find(&traffic).
		Error
	return traffic, err
}

func (repo trafficRepository) GetTrafficByState(ctx context.Context, state string, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Model(&entity.Traffic{}).
		Select("date, SUM(\"in\") AS \"in\", SUM(out) AS out, SUM(bandwidth) as bandwidth").
		Joins("JOIN interfaces as i ON i.id = traffics.interface_id").
		Joins("JOIN fats as f ON f.interface_id = i.id").
		Joins("JOIN locations as l ON l.id = f.location_id").
		Where("l.state = ? AND traffics.date BETWEEN ? AND ?", state, date.InitDate, date.EndDate).
		Group("date").
		Order("date").
		Find(&traffic).
		Error

	return traffic, err
}

func (repo trafficRepository) GetTrafficByCounty(ctx context.Context, state, county string, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Model(&entity.Traffic{}).
		Select("date, SUM(\"in\") AS \"in\", SUM(out) AS out, SUM(bandwidth) as bandwidth").
		Joins("JOIN interfaces as i ON i.id = traffics.interface_id").
		Joins("JOIN fats as f ON f.interface_id = i.id").
		Joins("JOIN locations as l ON l.id = f.location_id").
		Where("l.state = ? AND l.county = ? AND traffics.date BETWEEN ? AND ?", state, county, date.InitDate, date.EndDate).
		Group("date").
		Order("date").
		Find(&traffic).
		Error

	return traffic, err
}

// By Municipality
func (repo trafficRepository) GetTrafficByLocationID(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Model(&entity.Traffic{}).
		Select("date, SUM(\"in\") AS \"in\", SUM(out) AS out, SUM(bandwidth) as bandwidth").
		Joins("JOIN interfaces as i ON i.id = traffics.interface_id").
		Joins("JOIN fats as f ON f.interface_id = i.id").
		Where("f.location_id = ? AND traffics.date BETWEEN ? AND ?", id, date.InitDate, date.EndDate).
		Group("date").
		Order("date").
		Find(&traffic).
		Error

	return traffic, err
}
