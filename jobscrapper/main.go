package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

// #4 JOB SCRAPPER
// https://nomadcoders.co/go-for-beginners/lectures/1528
func main() {
	var jobs []extractedJob
	var c = make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	// fmt.Println(jobs)
	writeJobs(jobs)
	fmt.Println("Extracting job was done. Extracted", len(jobs), "jobs.")

}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	var c = make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	// fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, selection *goquery.Selection) {
		go extractJob(selection, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
	// c <- jobs
	// return jobs

}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	// == finally
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, selection *goquery.Selection) { pages = selection.Find("a").Length() })

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	// fmt.Println(id)
	title := cleanString(card.Find(".title>a").Text())
	// fmt.Println(title)
	location := cleanString(card.Find(".sjcl").Text())
	// fmt.Println(location)
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	// fmt.Println(id, title, location, salary, summary)
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	wErr := writer.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlices := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.salary}
		jwErr := writer.Write(jobSlices)
		checkErr(jwErr)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
