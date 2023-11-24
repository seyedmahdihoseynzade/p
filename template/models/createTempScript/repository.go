package main

var fllRepositoyTem = `package {lowerModelName}

import (
"sync"
"apiGolang/models/{lowerModelName}/{dataSources}"
)

type Repository struct {
	cacheDS         {dataSources}.{upperModelName}CacheDS
	dbDS            {dataSources}.{upperModelName}DBDS
}

var (
	repo   *Repository
	doOnce sync.Once
)

func initRepo() {
	repo = &Repository{}
}

func GetRepo() *Repository {
	doOnce.Do(initRepo)
	return repo
}

func (repo *Repository) db() {dataSources}.{upperModelName}DBDS{
	if repo.dbDS == nil {
		repo.dbDS = {dataSources}.Get{upperModelName}DBDS()
	}

	return repo.dbDS
}

func (repo *Repository) cache() {dataSources}.{upperModelName}CacheDS {
	if repo.cacheDS == nil {
		repo.cacheDS = dataSources.GetCacheDS()
	}

	return repo.cacheDS
}
`

var repositoyTemWithDB = `package {lowerModelName}

import (
"sync"
"apiGolang/models/{lowerModelName}/{dataSources}"
)

type Repository struct {
	dbDS            {dataSources}.{upperModelName}DBDS
}

var (
	repo   *Repository
	doOnce sync.Once
)

func initRepo() {
	repo = &Repository{}
}

func GetRepo() *Repository {
	doOnce.Do(initRepo)
	return repo
}

func (repo *Repository) db() {dataSources}.{upperModelName}DBDS{
	if repo.dbDS == nil {
		repo.dbDS = {dataSources}.Get{upperModelName}DBDS()
	}

	return repo.dbDS
}
`

var repositoyTemWithCache = `package {lowerModelName}

import (
"sync"
"apiGolang/models/{lowerModelName}/dataSources"
)

type Repository struct {
	cacheDS         {dataSources}.{upperModelName}CacheDS
}

var (
	repo   *Repository
	doOnce sync.Once
)

func initRepo() {
	repo = &Repository{}
}

func GetRepo() *Repository {
	doOnce.Do(initRepo)
	return repo
}

func (repo *Repository) cache() {dataSources}.{upperModelName}CacheDS {
	if repo.cacheDS == nil {
		repo.cacheDS = dataSources.GetCacheDS()
	}

	return repo.cacheDS
}
`

var normalRepositoyTem = `package {lowerModelName}

import (
"sync"
)

type Repository struct {
}

var (
	repo   *Repository
	doOnce sync.Once
)

func initRepo() {
	repo = &Repository{}
}

func GetRepo() *Repository {
	doOnce.Do(initRepo)
	return repo
}
`
