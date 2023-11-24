package name

/*
type Repository struct {
	mongoDS dataSources.NameDBDS
	cacheDS dataSources.NameRedisDS
	mysqlDS dataSources.NameDBDS
	// define other data source
}

const (
// define cast
)

var (
	//define var
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

func (repo *Repository) mysql() dataSources.NameDBDS {
	if repo.mysqlDS == nil {
		repo.mysqlDS = dataSources.GetDBDS()
	}
	return repo.mysqlDS
}

func (repo *Repository) mongo() dataSources.NameDBDS {
	if repo.mongoDS == nil {
		repo.mongoDS = dataSources.GetDBDS()
	}
	return repo.mongoDS
}

func (repo *Repository) cache() dataSources.NameRedisDS {
	if repo.cacheDS == nil {
		repo.cacheDS = dataSources.GetCacheDS()
	}
	return repo.cacheDS
}


func (repo *Repository) CreateUser(spanCtx context.Context, data dataModel) (data dataModel, err error) {

}


func (repo *Repository) SomeValidation() { //like identifier validation ...
}
*/
