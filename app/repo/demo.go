package repo

import (
	"github.com/jiebutech/datasource"
	"github.com/jiebutech/zin/app/entity"
)

type DemoRepo interface {
	GetDemoByID(id int64) (*entity.Demo, error)
}

type demoRepo struct {
	orm *datasource.Orm
}

func NewDemoRepo(db *datasource.Orm) DemoRepo {
	return &demoRepo{orm: db}
}

func (d demoRepo) GetDemoByID(id int64) (*entity.Demo, error) {
	info := new(entity.Demo)
	err := d.orm.Where("id = ?", id).First(info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}
