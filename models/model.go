package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"whacos/pkg/settings"
)

var (
	DB  *gorm.DB
	ERR error

	DelFlagYes = 1
	DelFlagNo  = 0
)

// 公共Model
type Model struct {
	Id          int        `json:"id" gorm:"primary_key"` // 主键ID
	CreatedBy   int        `json:"createdBy"`             // 创建人
	CreatedTime *time.Time `json:"createdTime"`           // 创建时间
	UpdatedBy   int        `json:"updatedBy"`             // 修改人
	UpdatedTime *time.Time `json:"updatedTime"`           // 修改时间
	DelFlag     int        `json:"delFlag" default:"1"`   // 删除标记，默认为1
}

// 分页Model
type Page struct {
	List      interface{} `json:"list"`
	PageTotal int         `json:"pageTotal"`
	Total     int         `json:"total"`
}

func PageUtils(dataList interface{}, countNum, pageSize int) (p Page) {
	p.List = dataList
	p.PageTotal = countNum / pageSize
	p.Total = countNum
	return
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

	//  默认数据表名称
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	DB.SingularTable(true)
	DB.LogMode(true)
	//DB.SetLogger()
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
