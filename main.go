package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/SC2Analyser/Analyser"
	"github.com/SC2Analyser/Player"
	"github.com/spf13/viper"
)

func main() {
	p0 := &Player.Player{}
	p1 := &Player.Player{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %w \n", err))
	}
	p0.Names = viper.GetStringSlice("player1Name")
	p1.Names = viper.GetStringSlice("player2Name")
	p0.Race = viper.GetStringSlice("player1Race")
	p1.Race = viper.GetStringSlice("player2Race")

	dir := viper.GetString("replayFolder")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".SC2Replay") {
				err = Analyser.Analyse(dir+"/"+file.Name(), p0, p1)
				if err != nil {
					fmt.Println(err.Error)
					panic(err)
				}
			}
		}
	}
	fmt.Printf("%v wins %d\n", p0.Names[0], p0.Winning)
	fmt.Printf("%v wins %d\n", p1.Names[0], p1.Winning)
}
