package logic

import (
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/pkg/pagination"
	"gorm.io/gorm"
)

// Paginator 分页
// rule 分页的规则
func Paginator(paginate *pagination.Paginator, rows any, scopes func(*gorm.DB) *gorm.DB, filter ...func(desc any) (any, error)) error {
	var err error

	err = app.DB.Scopes(scopes).Count(&paginate.Total).Error
	if err == nil {
		err = app.DB.Scopes(scopes).
			Offset(paginate.Offset()).
			Limit(paginate.Size).
			Find(&rows).
			Error
	}

	if err == nil {
		for _, f := range filter {
			rows, err = f(rows)
			if err != nil {
				return err
			}
		}
		paginate.Load(rows)
	}

	return err
}
