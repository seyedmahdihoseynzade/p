package vitessDS

import (
	"apiGolang/models/category/dataModel"
	"code.ts.co.ir/gaplib/golib/db/vitess"
	"context"
	"errors"
	"sync"
	"time"
)

type VitessDataSource struct {
	vitess.BaseVitessDB
}

var (
	vitessDS *VitessDataSource
	doOnce   sync.Once
)

func initVitessDS() {
	vitessDS = &VitessDataSource{
		BaseVitessDB: vitess.NewVitessDB(),
	}
}

func GetVitessDS() *VitessDataSource {
	doOnce.Do(initVitessDS)
	return vitessDS
}

func (v *VitessDataSource) Create(ctx context.Context, category dataModel.Category) error {
	if category.ParentID != 0 {
		c := v.GatByID(ctx, category.UserID, category.ParentID)
		if c.ID == 0 {
			return errors.New("دسته بندی اصلی وجود ندارد")
		}
	}

	category.ID = time.Now().UnixNano()
	return v.Connection.Create(&category).Error
}

func (v *VitessDataSource) GatByID(ctx context.Context, userID, ID int64) dataModel.Category {
	c := dataModel.Category{}
	v.Connection.Where("ID = ? and userID = ?", ID, userID).Take(&c)
	return c
}

func (v *VitessDataSource) GetMainCategory(ctx context.Context, userID int64) []dataModel.Category {
	c := []dataModel.Category{}
	v.Connection.Where("userID = ? and parentID = ?", userID, 0).Order("ID desc").Find(&c)
	return c
}

func (v *VitessDataSource) GetSubCategory(ctx context.Context, userID int64, prentID int64) []dataModel.Category {
	c := []dataModel.Category{}
	v.Connection.Where("userID = ? and parentID = ?", userID, prentID).Order("ID desc").Find(&c)
	return c
}
