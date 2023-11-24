package vitessDS

import (
	"apiGolang/models/incomeAndExpense/dataModel"
	"code.ts.co.ir/gaplib/golib/db/vitess"
	"errors"
	"fmt"
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

func (vitessDS *VitessDataSource) GetList(userID int64) []dataModel.IncomeAndExpense {
	var a []dataModel.IncomeAndExpense
	err := vitessDS.Connection.Where("userID = ?", userID).Find(&a).Error
	fmt.Println(err)
	return a
}

func (a *VitessDataSource) Create(m dataModel.IncomeAndExpense) error {
	if m.Amount < 0 {
		x := 0.0
		a.Connection.Model(dataModel.IncomeAndExpense{}).Where("userID = ?", m.UserID).Select("SUM(amount)").Take(&x)
		if float32(x)+m.Amount < 0 {
			return errors.New("موجودی کافی نیست")
		}
	}
	m.ID = time.Now().UnixNano()
	return a.Connection.Create(&m).Error
}

func (a *VitessDataSource) GetByID(ID, userID int64) dataModel.IncomeAndExpense {
	i := dataModel.IncomeAndExpense{}
	a.Connection.Where("userID = ? and ID = ?", userID, ID).Take(&i)
	return i
}

func (v *VitessDataSource) TotalBalance(userID int64) float32 {
	x := 0.0
	v.Connection.Model(dataModel.IncomeAndExpense{}).Where("userID = ?", userID).Select("SUM(amount)").Take(&x)
	return float32(x)

}
