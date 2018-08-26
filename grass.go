
package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"fmt"
	"time"
	"strconv"
)

func getContribution(username string,t time.Time)int{

	doc := getDocument(username)

	sprint := fmt.Sprintf("rect[data-date=\"%s\"]", dayString(t))
	s := doc.Find(sprint)
	data,_ := s.Attr("data-count")
	count ,_ := strconv.Atoi(data)
	return count
}

func getYearContributions(username string)[]Grass{
	layout := "2006-01-02"

	doc := getDocument(username)
	fmt.Println(len(doc.Find("rect").Nodes))
	grasses := make([]Grass, len(doc.Find("rect").Nodes))
	doc.Find("rect").Each(func(i int, s *goquery.Selection) {

		dateStr,_ := s.Attr("data-date")
		t, _ := time.Parse(layout, dateStr)
		grasses[i].CreatedAt = t

		data,_ := s.Attr("data-count")
		count ,_ := strconv.Atoi(data)
		grasses[i].DataCount = count
	})
	return grasses
}


func getDocument(username string)*goquery.Document{
	url := fmt.Sprintf("https://github.com/users/%s/contributions",username)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func dayString(t time.Time)string{
	const layout = "2006-01-02"
	return t.Format(layout)
}

