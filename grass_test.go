package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
	"time"
)

func TestExampleSuccess(t *testing.T) {
	getContribution("mogeta",time.Now())

		//if err != nil {
		//	t.Fatalf("failed test %#v", err)
		//}
		//if result != 1 {
		//	t.Fatal("failed test")
		//}
}
