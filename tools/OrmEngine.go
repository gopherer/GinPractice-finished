package tools

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/wonderivan/logger"
	"goWeb/model"
)

type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		logger.Error("xorm Engine 启动失败", err)
		return nil, err
	}
	engine.ShowSQL(database.ShowSql)
	err = engine.Sync2(new(model.UserInfo), new(model.UserRegister))
	if err != nil {
		logger.Error("结构体映射数据表失败", err)
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return DbEngine, nil
}
