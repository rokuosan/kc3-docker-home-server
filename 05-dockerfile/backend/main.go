package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

type Article struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func main() {
	initDatabase()
	defer db.Close()

	startServer()
}

func initDatabase() {
	timezone, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	config := &mysql.Config{
		User:                 "test",
		Passwd:               "test_password",
		Net:                  "tcp",
		Addr:                 "database:3306",
		DBName:               "kc3_docker_home_server",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  timezone,
	}

	d, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	db = d
}

func startServer() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM article")
		if err != nil {
			return err
		}
		defer rows.Close()

		articles := []Article{}
		for rows.Next() {
			var article Article
			if err := rows.Scan(&article.Title, &article.Author, &article.Content); err != nil {
				panic(err)
			}
			articles = append(articles, article)
		}

		jsonBytes, err := json.Marshal(articles)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, string(jsonBytes))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
