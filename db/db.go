package db

import (
	"easy-search/config"
	"fmt"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Conn is the database connection
var Conn *gorm.DB

func Init() {
	logLevel := logger.Error
	if gin.IsDebugging() {
		logLevel = logger.Info
	}

	// custom logger
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)

	// data source name
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.DbConfig.Username,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Database,
	)

	// Init conn
	var err error
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: customLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
}
