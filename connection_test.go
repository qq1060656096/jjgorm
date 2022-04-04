package jjgorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestMysqlConnection(t *testing.T) {
	dataSource := "user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	dataSource = "root:root@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True&loc=Local"
	conf := Config{
		DriverName: DriverNameMySql,
		DataSource: dataSource,
	}
	connection := NewConnection(conf)
	conn, err := connection.Connect()
	if err != nil {
		assert.Error(t, err, "connection")
	}

	sql := `insert into test(nickname) values(?)`
	db := conn.Exec(sql, fmt.Sprintf("TestMysqlConnection.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "TestMysqlConnection.insertData.error", db.Error)
}

func TestSqlite3Connection(t *testing.T)  {
	dir, _ := os.Getwd()
	dataSource := dir + "/testdata/sqlite3.1.db"
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	connection := NewConnection(conf)
	conn, err := connection.Connect()
	if err != nil {
		assert.Error(t, err, "connection")
	}
	sql := `insert into test(nickname) values(?)`
	db := conn.Exec(sql, fmt.Sprintf("TestSqlite3Connection.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "TestSqlite3Connection.insertData.error", db.Error)
}

func TestConnection_GetDB(t *testing.T) {
	dir, _ := os.Getwd()
	dataSource := dir + "/testdata/sqlite3.1.db"
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	connection := NewConnection(conf)
	_, err := connection.Connect()
	if err != nil {
		assert.Error(t, err, "connection")
	}
	conn2,_ := connection.GetDB()
	connection.Disconnect()
	assert.NotEqual(t , nil, conn2)
}


func TestConnection_Disconnect(t *testing.T) {
	dir, _ := os.Getwd()
	dataSource := dir + "/testdata/sqlite3.1.db"
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	connection := NewConnection(conf)
	_, err := connection.Connect()
	if err != nil {
		assert.Error(t, err, "connection")
	}
	connection.Disconnect()
	conn, _ := connection.GetDB()
	assert.True(t, nil != conn)
}

func TestConnection_Config(t *testing.T) {
	dir, _ := os.Getwd()
	dataSource := dir + "/testdata/sqlite3.1.db"
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	connection := NewConnection(conf)
	_, err := connection.Connect()
	if err != nil {
		assert.Error(t, err, "connection")
	}
	assert.Equal(t, conf, connection.Config())
}
