package databases

import (
	config "baobaozhuan/config"
	"baobaozhuan/modules/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB connection
var DB *gorm.DB

// init mysql database
func init() {
	var err error
	DB, err = gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		log.ErrorLog.Println(err)
	}
	// set max idle connections
	DB.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	// set max open connections
	DB.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
}
