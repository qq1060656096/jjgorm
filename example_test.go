package jjmgorm

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
	"time"
)

func ExampleConnection_GetDB() {
	manager := NewManager()
	dataSource := "root:root@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True&loc=Local"
	manager.Add("mysql_data_1", Config{
		DriverName: DriverNameMySql,
		DataSource: dataSource,
	})

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("osGetDirErr: ", err)
	}
	dataSource = dir + "/testdata/sqlite3.1.db"
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	manager.Add("sqlite3_data_2", conf)
	connection := manager.Get("mysql_data_1")
	if connection == nil {
		fmt.Println("connectionNotExist: ", connection)
	}

	db, err := connection.GetDB()
	if err != nil {
		fmt.Println("getDbErr: ", err)
	}
	sql := `insert into test(nickname) values(?)`
	db2 := db.Exec(sql, fmt.Sprintf("TestSqlite3Connection.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	if db2.Error != nil {
		fmt.Println("insertDataErr: ", db2.Error)
	}
	fmt.Printf("insertCount: %d", db2.RowsAffected)
	// Output:
	// insertCount: 1
}

// ExampleNewManagerGormReadWriteSplitting 读写分离
func ExampleNewManagerGormReadWriteSplitting() {
	manager := NewManager()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("osGetDirErr: ", err)
	}
	dataSource := dir + "/testdata/sqlite3-read-write-splitting-1.db"
	conf := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	manager.Add("sqlite3_data_1", conf)

	dataSource = dir + "/testdata/sqlite3-read-write-splitting-2.db"
	conf2 := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}
	dataSource = dir + "/testdata/sqlite3-read-write-splitting-3.db"
	conf3 := Config{
		DriverName: DriverNameSqlite3,
		DataSource: dataSource,
	}

	connection := manager.Get("sqlite3_data_1")
	if connection == nil {
		fmt.Println("connectionNotExist: ", connection)
	}

	db, err := connection.GetDB()
	if err != nil {
		fmt.Println("getDbErr: ", err)
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		// use `db4` as replicas
		Replicas: []gorm.Dialector{NewDialector(conf2)},
		Policy:   dbresolver.RandomPolicy{},
	}).Register(dbresolver.Config{
		// `db3` as replicas for `orders`
		Replicas: []gorm.Dialector{NewDialector(conf3)},
	}, "orders"))

	sql := `insert into test(nickname) values(?)`
	db2 := db.Exec(sql, fmt.Sprintf("TestSqlite3Connection.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	if db2.Error != nil {
		fmt.Println("insertDataErr: ", db2.Error)
	}
	fmt.Printf("insertCount: %d", db2.RowsAffected)
	// Output:
	// insertCount: 1
}
