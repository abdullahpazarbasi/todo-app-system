package driven_adapter_db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
	"todo-app-service/configs"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

func ConfigureDB(eva corePort.EnvironmentVariableAccessor) (*gorm.DB, error) {
	var err error

	var dbHost string
	dbHost, err = eva.GetOrThrowError(configs.EnvironmentVariableNameTodoDbHost)
	if err != nil {
		return nil, err
	}
	var dbPort int64
	dbPort, err = strconv.ParseInt(eva.Get(configs.EnvironmentVariableNameTodoDbPort, "3306"), 10, 32)
	if err != nil {
		return nil, err
	}
	var dbUser string
	dbUser, err = eva.GetOrThrowError(configs.EnvironmentVariableNameTodoDbUser)
	if err != nil {
		return nil, err
	}
	var dbPass string
	dbPass, err = eva.GetOrThrowError(configs.EnvironmentVariableNameTodoDbPass)
	if err != nil {
		return nil, err
	}
	var dbName string
	dbName, err = eva.GetOrThrowError(configs.EnvironmentVariableNameTodoDbName)
	if err != nil {
		return nil, err
	}
	var maximumNumberOfOpenConnections int64
	maximumNumberOfOpenConnections, err = strconv.ParseInt(eva.Get(
		"TODO_DB_MAX_OPEN_CONNECTIONS",
		"100",
	), 10, 32)
	if err != nil {
		return nil, err
	}
	var maximumLifetimeOfConnectionInSeconds int64
	maximumLifetimeOfConnectionInSeconds, err = strconv.ParseInt(eva.Get(
		"TODO_DB_MAX_LIFETIME_OF_CONNECTION",
		"3600",
	), 10, 32)
	if err != nil {
		return nil, err
	}
	var maximumNumberOfIdleConnections int64
	maximumNumberOfIdleConnections, err = strconv.ParseInt(eva.Get(
		"TODO_DB_MAX_IDLE_CONNECTIONS",
		"10",
	), 10, 32)
	if err != nil {
		return nil, err
	}
	var maximumIdleTimeOfConnectionInSeconds int64
	maximumIdleTimeOfConnectionInSeconds, err = strconv.ParseInt(eva.Get(
		"TODO_DB_MAX_IDLE_TIME_OF_CONNECTION",
		"3600",
	), 10, 32)
	if err != nil {
		return nil, err
	}
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
		),
	), &gorm.Config{
		TranslateError:                           false,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(int(maximumNumberOfOpenConnections))
	sqlDB.SetConnMaxLifetime(time.Duration(maximumLifetimeOfConnectionInSeconds) * time.Second)
	sqlDB.SetMaxIdleConns(int(maximumNumberOfIdleConnections))
	sqlDB.SetConnMaxIdleTime(time.Duration(maximumIdleTimeOfConnectionInSeconds) * time.Second)

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Todo{},
		&TodoTag{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
