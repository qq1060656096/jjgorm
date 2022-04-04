# jjgorm
jjgorm
https://github.com/qq1060656096/jjgorm


```sh
# 代码静态检查发现可能的bug或者可疑的构造
go vet .

# 竞态检测
go build -race -v .
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
