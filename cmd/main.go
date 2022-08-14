package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shashkomari/SnakeGame/internal/snakegame"
)

type Config struct {
	Board struct {
		Hight int `json:"hight"`
		Wight int `json:"width"`
	} `json:"board"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config

	configFile, err := os.Open(filename)

	if err != nil {
		configFile, err = os.Create(filename)
		jsonFile := json.NewEncoder(configFile)
		config.Board.Hight = 10
		config.Board.Wight = 20
		err = jsonFile.Encode(&config)
	} else {
		defer configFile.Close()
		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&config)
	}
	return config, err
}

func main() {
	config, _ := LoadConfiguration("config.json")

	var sg = snakegame.CreateSnakeGame(config.Board.Hight, config.Board.Wight)

	go sg.UserControl()
	score := sg.Run()

	fmt.Printf("Game Over!\nScore: %d\n", score)
}
