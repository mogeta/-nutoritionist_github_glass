
package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"fmt"
	"time"
)

func getContribution(){
	// Request the HTML page.
	res, err := http.Get("https://github.com/users/mogeta/contributions")
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

	sprint := fmt.Sprintf("rect[data-date=\"%s\"]", previousDayString(time.Now()))
	s := doc.Find(sprint)
	rect,_ := s.Attr("data-count")
	fmt.Print(rect)

}

func previousDayString(t time.Time)string{
	t = t.AddDate(0,0,-1)
	const layout = "2006-01-02"
	return t.Format(layout)
}

