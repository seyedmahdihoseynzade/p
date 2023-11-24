package main

var DBDS = `package {dataSources}

import (
	"apiGolang/models/{lowerModelName}/{dataSources}/{vitessDS}"
)

type {upperModelName}DBDS interface {

}

func Get{upperModelName}DBDS() {upperModelName}DBDS {
	return vitessDS.GetVitessDS()
}

`
