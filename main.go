package main

import (
	"github.com/spf13/viper"
	"fmt"
	"time"
)

func main(){

	getConfig()
	username := viper.GetString("username")
	date := time.Now().AddDate(0,0,-1)
	count := getContribution(username,date)

	g := []Grass{{
		date,
		count,
	}}
	sendData("github_grass",g)
	fmt.Println(g)
}

func getConfig(){
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
}