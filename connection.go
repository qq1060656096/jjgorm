package jjgorm

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"sync"
)

type Connection struct {
	db   *gorm.DB
	conf Config
	mutex sync.Mutex
}

func (conn *Connection) Connect() (*gorm.DB, error) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.noLockConnect()
}

func (conn *Connection) noLockConnect() (*gorm.DB, error) {
	db, err := gorm.Open(NewDialector(conn.conf), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	conn.db = db
	return conn.db, nil
}

func (conn *Connection) GetDB() (*gorm.DB, error) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	if conn.db == nil {
		return conn.noLockConnect()
	}
	return conn.db, nil
}

func (conn *Connection) Disconnect() bool {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	if conn.db != nil {
		conn.db = nil
	}
	return true
}

func (conn *Connection) Config() Config {
	return conn.conf
}

func NewConnection(conf Config) *Connection {
	return &Connection{conf: conf}
}

// NewDialector create gorm.Dialector
func NewDialector(config Config) gorm.Dialector {
	switch config.DriverName {
	case DriverNameMySql:
		return mysql.Open(config.DataSource)
	case DriverNamePostgreSql:
		return postgres.Open(config.DataSource)
	case DriverNameSqlite3:
		return sqlite.Open(config.DataSource)
	case DriverNameSqlServer:
		return sqlserver.Open(config.DataSource)
	}
	return nil
}
