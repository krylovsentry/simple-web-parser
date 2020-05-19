package models

import (
	"database/sql"
	"fmt"
	"net/url"
)

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsCollection struct {
	News []News `json:"items"`
}

func NewsExist(db *sql.DB, title string) bool {
	sqlQuery := `SELECT title FROM news WHERE title = ?`
	err := db.QueryRow(sqlQuery, title).Scan(&title)
	if err != nil {
		if err != sql.ErrNoRows {
			panic(err)
		}

		return false
	}

	return true
}

func GetNews(db *sql.DB, search string) NewsCollection {
	decodedValue, err := url.QueryUnescape(search)
	fmt.Println(search)
	if err != nil {
		panic(err)
	}
	sql := "SELECT * FROM news where title LIKE '%" + decodedValue + "%'"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := NewsCollection{}
	for rows.Next() {
		news := News{}
		err2 := rows.Scan(&news.ID, &news.Title, &news.Content)
		if err2 != nil {
			panic(err2)
		}
		result.News = append(result.News, news)
	}
	return result
}

func AddNews(db *sql.DB, title string, content string) (int64, error) {
	sql := "INSERT INTO news(title, content) VALUES(?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(title, content)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
