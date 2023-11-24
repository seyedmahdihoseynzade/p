package vitessDS

import (
	"apiGolang/models/detailIncomeAndExpense/dataModel"
	"code.ts.co.ir/gaplib/golib/db/vitess"
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

func (v *VitessDataSource) GetList(prentID int64) []dataModel.DetailIncomeAndExpense {
	d := []dataModel.DetailIncomeAndExpense{}
	v.Connection.Where("parentID = ?", prentID).Order("ID desc").Find(&d)
	return d
}

func (v *VitessDataSource) Create(d dataModel.DetailIncomeAndExpense) error {
	d.ID = time.Now().UnixNano()
	return v.Connection.Create(&d).Error
}
func (v *VitessDataSource) SumOfAmount(parentID int64) float32 {
	var s float32
	v.Connection.Model(dataModel.DetailIncomeAndExpense{}).Where("parentID = ?", parentID).Select("SUM(amount)").Take(&s)
	return s
}
