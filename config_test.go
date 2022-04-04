package jjgorm

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMySQLConfig(t *testing.T) {
	dataSource := "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	c := &Config{
		DriverName: DriverNameMySql,
		DataSource: dataSource,
	}
	assert.Equal(t, dataSource, c.DataSource)
}

func TestSqliteConfig(t *testing.T) {
	dir, _ := os.Getwd()
	dataSource := dir + "/tmp/sqlite3.1.db"
	c := &Config{
		DriverName: DriverNameMySql,
		DataSource: dataSource,
	}
	assert.Equal(t, dataSource, c.DataSource)
}
