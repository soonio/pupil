package logic

import (
	"errors"
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/app/model"

	"github.com/soonio/pupil/pkg/pagination"

	"gorm.io/gorm"
)

var Dict = new(dictLogic)

type dictLogic struct{}

func (i *dictLogic) List(page *pagination.Paginator) error {
	var rows = make([]*model.Dict, 0)
	return Paginator(page, rows, func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.Dict{})
	})
}

func (i *dictLogic) Save(k, v string) error {
	one := &model.Dict{}
	err := app.DB.Where("k", k).Take(one).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return app.DB.Create(&model.Dict{K: k, V: v}).Error
		}
		return err
	}

	return app.DB.Model(&model.Dict{}).Where("k", k).Update("v", v).Error
}

func (i *dictLogic) Get(k string) (v string, err error) {
	one := &model.Dict{}
	err = app.DB.Where("k", k).Take(one).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}
	return one.V, err
}
