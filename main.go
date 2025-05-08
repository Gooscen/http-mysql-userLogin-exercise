package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 修改为你本地的数据库连接配置
	dsn := "root:12345678@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	InitDB(dsn)

	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)

	log.Println("服务器启动于 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
