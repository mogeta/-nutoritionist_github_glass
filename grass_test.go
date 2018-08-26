package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
	"time"
	"fmt"
)

var username = "mogeta"

func TestContribution(t *testing.T) {
	count := getContribution(username,time.Now().AddDate(0,0,-1))
	fmt.Println(count)

		//if err != nil {
		//	t.Fatalf("failed test %#v", err)
		//}
		//if result != 1 {
		//	t.Fatal("failed test")
		//}
}

func TestGetYearContributions(t *testing.T) {
	getYearContributions(username)
}