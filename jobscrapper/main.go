package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dorbae/golang-study/jobscrapper/scrapper"
	"github.com/labstack/echo"
)

const fileName = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":8999"))
	// scrapper.Scrape("java")
}

func handleHome(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	fmt.Println("[DEBUG] term=", term)
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}
