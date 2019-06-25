package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"whacos/pkg/settings"
)

var DB *gorm.DB
var ERR error

type Model struct {
	Id          int       `json:"id" gorm:"primary_key"`
	CreatedBy   int       `json:"createdBy"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedBy   int       `json:"updatedBy"`
	UpdatedTime time.Time `json:"updatedTime"`
	DelFlag     int       `json:"delFlag"`
}

func init() {

	DB, ERR = gorm.Open(settings.DatabaseConfig.Dialect, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DatabaseConfig.Username,
		settings.DatabaseConfig.Password,
		settings.DatabaseConfig.Host,
		settings.DatabaseConfig.Name))

	if ERR != nil {
		log.Println(ERR)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(settings.DatabaseConfig.MaxIdleConns)
	DB.DB().SetMaxOpenConns(settings.DatabaseConfig.MaxOpenConns)
}

func CloseDB() {
	defer DB.Close()
}

type DBModel interface {
	// 统计数据
	Count(maps interface{}) (count int)

	ExistById(id int) (b bool)

	SelectById(id int)

	SelectList(maps interface{})

	SelectPage(pageNum int, pageSize int, maps interface{})

	UpdateById(maps interface{})

	Insert(maps interface{})

	BeforeInsert(scope *gorm.Scope) (err error)

	AfterInsert(scope *gorm.Scope) (err error)

	BeforeUpdate(scope *gorm.Scope) (err error)

	AfterUpdate(scope *gorm.Scope) (err error)
}
