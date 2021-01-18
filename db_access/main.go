package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Task is ...
type Task struct {
	Content string `json:"content"`
}

func execDB(db *sql.DB, q string) {
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}

/*
	前回はPOSTして送られてきたデータをただただ返す処理の学習をした。
	今回は送られてきたデータをデータベースに格納する処理を記述してみよう。
	Relational Databaseと呼ばれる形式のデータベースを使用する。
	データベースには様々な種類がある
	MySQL
	PostgreSQL
	SQLite
	今回は手軽にファイルを使用して構築することができるSQLiteを使用する。
	余力があれば自分のPCにMySQLを用意してMySQLで行ってみて欲しい。
*/
func main() {
	router := gin.Default()

	db, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)
	defer db.Close()
	query := `
        CREATE TABLE memo (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          body VARCHAR(255) NOT NULL,
          created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))
        )
    `
	execDB(db, query)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})

	router.POST("/post", func(ctx *gin.Context) {
		var req Task
		ctx.BindJSON(&req)
		ctx.JSON(200, req)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	endless.ListenAndServe(":"+port, router)
}
