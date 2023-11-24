package main

var vitess = `package {vitessDS}

import (
	"code.ts.co.ir/gaplib/golib/db/vitess"
    "sync"
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
`
