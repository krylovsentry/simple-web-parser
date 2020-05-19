package main

import (
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"simple-web-parser/rest"
)

func main() {
	db := initDataBase("storage.db")
	createDB(db)

	port := os.Getenv("PORT")

	if port == "" {
		panic("port not descibed")
	}

	e := echo.New()
	e.GET("/news", rests.GetNews(db))
	e.POST("/feeder", rests.PostFeeder(db))
	e.File("/", "app/index.html")
	e.File("/src/main.js", "app/src/main.js")

	e.Logger.Fatal(e.Start(":" + port))
}

func initDataBase(filePath string) *sql.DB {
	db, err := sql.Open("sqlite3", filePath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	return db
}

func createDB(db *sql.DB) {
	createQuery := `
    CREATE TABLE IF NOT EXISTS news(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title VARCHAR NOT NULL,
        content VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(createQuery)
	if err != nil {
		panic(err)
	}
}
