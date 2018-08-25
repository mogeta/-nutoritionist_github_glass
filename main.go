package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	url := "https://github.com/users/mogeta/contributions"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer access-token")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}