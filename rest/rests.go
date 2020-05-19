package rests

import (
	"database/sql"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/labstack/echo"
	"net/http"
	"simple-web-parser/models"
)

func GetNews(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetNews(db, c.QueryParam("search")))
	}
}

func PostFeeder(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		seed := new(models.Seed)
		if err = c.Bind(seed); err != nil {
			return
		}

		geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: []string{seed.Url},
			ParseFunc: createParser(db, seed),
		}).Start()
		return c.JSON(http.StatusCreated, seed)
	}
}

func createParser(db *sql.DB, seed *models.Seed) func(g *geziyor.Geziyor, r *client.Response) {
	return func(g *geziyor.Geziyor, r *client.Response) {
		r.HTMLDoc.Find("." + seed.OneParent).Each(func(i int, s *goquery.Selection) {
			title := s.Find("." + seed.Title).Text()
			if !models.NewsExist(db, title) {
				models.AddNews(db, title, s.Find("."+seed.Content).Text())
			}
		})
	}
}
