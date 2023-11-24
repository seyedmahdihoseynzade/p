package mongoDS

/*
type MongoDataSource struct {
	mongoDB.BaseMongoDB
}

var (
	mongoDS *MongoDataSource
	doOnce sync.Once
)

func initMongoDS() {
	mongoDS = &MongoDataSource{
		BaseMongoDB: mongoDB.NewMongoDB(),
	}
}

func GetMongoDS() *MongoDataSource {
	doOnce.Do(initMongoDS)
	return mongoDS
}


func (mangosDS *MongoDataSource) Insert(spanCtx context.Context, user dataModel.User) error {
	logger := log.StartSpan(spanCtx, mangosDS.Tracer, "insert user query")
	defer logger.FinishSpan()

	_, err := mangosDS.Connection.Collection("users").InsertOne(spanCtx, user)
	if err != nil {
		logger.Error("insert user err", zap.Error(err), zap.Any("insertField", user))
	}
	return err
}

//and other ...

*/
