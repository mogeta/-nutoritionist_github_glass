
package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"fmt"
	"time"
)

func getContribution(username string,t time.Time){

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

	sprint := fmt.Sprintf("rect[data-date=\"%s\"]", dayString(t))
	s := doc.Find(sprint)
	rect,_ := s.Attr("data-count")
	fmt.Print(rect)

}

func dayString(t time.Time)string{
	const layout = "2006-01-02"
	return t.Format(layout)
}

