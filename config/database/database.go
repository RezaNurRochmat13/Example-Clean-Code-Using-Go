package database

import (
	"clean-arch/utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func MysqlDevelopmentConfiguration() (*sql.DB, error) {
	viper.SetConfigName("mysql")
	viper.AddConfigPath("config/database")

	errorReadingFileConfig := viper.ReadInConfig()

	if !utils.GlobalErrorWithBool(errorReadingFileConfig) {
		return nil, errorReadingFileConfig
	}

	var MySQLHost = viper.GetString("mysql.MySQLHost")
	var MySQLUsername = viper.GetString("mysql.MySQLUsername")
	var MySQLPassword = viper.GetString("mysql.MySQLPassword")
	var MySQLDatabaseName = viper.GetString("mysql.MySQLDatabaseName")

	var databaseConfig = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		MySQLUsername, MySQLPassword, MySQLHost, MySQLDatabaseName)

	databaseConnectionConfiguration, errorDatabaseConfiguration := sql.Open("mysql", databaseConfig)

	if !utils.GlobalErrorWithBool(errorDatabaseConfiguration) {
		return nil, errorDatabaseConfiguration
	}

	databaseConnectionConfiguration.SetMaxIdleConns(10)
	databaseConnectionConfiguration.SetMaxOpenConns(10)

	return databaseConnectionConfiguration, nil
}
