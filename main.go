package main

import (
	"github.com/spf13/viper"
	"fmt"
	"time"
)

func main(){

	getConfig()
	username := viper.GetString("username")
	getContribution(username,time.Now())
}

func getConfig(){
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
}