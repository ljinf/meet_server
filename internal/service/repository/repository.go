package repository

import (
	"context"
	"fmt"
	"github.com/ljinf/meet_server/internal/model"
	"github.com/ljinf/meet_server/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

const ctxTxKey = "MeetTx"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

// DB return tx
// If you need to create a Transaction, you must call DB(ctx) and Transaction(ctx,fn)
func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func (r *Repository) Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo < 1 {
			pageNo = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 30
		}

		return db.Offset((pageNo - 1) * pageSize).Limit(pageSize)
	}
}

func NewDB(conf *config.AppConfig) *gorm.DB {

	dbLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, //慢sql阀值
			LogLevel:                  logger.Info, //日志级别
			IgnoreRecordNotFoundError: true,        //忽略RecordNotFound(记录未找到)错误
			Colorful:                  true,        //禁用彩色打印
		},
	)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlDB.Username, conf.MysqlDB.Password,
		conf.MysqlDB.Host, conf.MysqlDB.Port, conf.MysqlDB.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLog,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //英文单数形式
		},
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Register{}, &model.UserInfo{})

	db = db.Debug()
	return db
}
