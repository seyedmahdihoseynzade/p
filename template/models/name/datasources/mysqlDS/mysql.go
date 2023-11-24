package mysqlDS

/*
type MysqlDataSource struct {
	mysqlDB.BaseMysqlDB
}

var (
	mysqlDS *MysqlDataSource
	doOnce  sync.Once
)

func initMysqlDS() {
	mysqlDS = &MysqlDataSource{
		BaseMysqlDB: mysqlDB.NewMysqlDB(),
	}
}

func GetMysqlDS() *MysqlDataSource {
	doOnce.Do(initMysqlDS)
	return mysqlDS
}

func (mysql *MysqlDataSource) Insert(spanCtx context.Context, otp dataModel.OTP) error {
	logger := log.StartSpan(spanCtx, mysql.Tracer, "insert otp err")
	defer logger.FinishSpan()
	_, err := mysql.Connection.Exec(...)
	if err != nil {
		logger.Error("insert otp err", zap.Error(err), zap.Any("otp", otp))
	}
	return err
}
*/
