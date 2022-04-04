# jjmgorm
>  帮助你简单管理 gorm 连接，多数据库连接，Sass 多数据库切换


### 简单示例

```go
package main

import (
	"os"
	"time"
	"fmt"
	"github.com/qq1060656096/jjmgorm"
)

func main()  {
	manager := jjmgorm.NewManager()
	dataSource := "root:root@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True&loc=Local"
	manager.Add("mysql_data_1", jjmgorm.Config{
		DriverName: jjmgorm.DriverNameMySql,
		DataSource: dataSource,
	})

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("osGetDirErr: ", err)
	}
	dataSource = dir + "/testdata/sqlite3.1.db"
	conf := jjmgorm.Config{
		DriverName: jjmgorm.DriverNameSqlite3,
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

```

### 主从示例
> 请参考 example_test.go
> 或者参考 [gorm.io/plugin/dbresolver](https://github.com/go-gorm/dbresolver) 包使用
```go

```


```sh
# 代码静态检查发现可能的bug或者可疑的构造
go vet .

# 竞态检测
go build -race -v .

# 开启本地官网
go get -v -u golang.org/x/tools/cmd/godoc
godoc -http=:8080 
```

### 单元测试
```sh
# 运行所有单元测试
go test -count=1 -v . 

# 运行所有单元测试，并查看测试覆盖率
go test -count=1 -v -cover .

# 运行所有单元测试，并查看测试覆盖率，竞态检测
go test -count=1 -v -cover -race .

```

### 单元测试 sql
```sh
# mysql
CREATE TABLE `test` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
  
# sqlite3
CREATE TABLE "test" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "nickname" text NOT NULL DEFAULT ''
);
```
