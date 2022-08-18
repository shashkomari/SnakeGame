package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/shashkomari/SnakeGame/internal/snakegame"
)

const kDefaultConfigFilename = "config.json"

type Config struct {
	Board struct {
		Hight int `json:"hight"`
		Wight int `json:"width"`
	} `json:"board"`
}

func LoadConfiguration(filename string) (*Config, error) {
	var config Config

	configFile, err := os.Open(filename)

	if errors.Is(err, os.ErrNotExist) {
		configFile, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
		defer configFile.Close()

		jsonFile := json.NewEncoder(configFile)
		config.Board.Hight = 10
		config.Board.Wight = 20
		err = jsonFile.Encode(&config)
		if err != nil {
			return nil, err
		}
	} else {
		if err != nil {
			return nil, err
		}
		defer configFile.Close()

		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&config)
		if err != nil {
			return nil, err
		}
	}
	return &config, nil
}

func main() {
	config, err := LoadConfiguration(kDefaultConfigFilename)
	if err != nil {
		fmt.Println("failed to load configuration: ", err.Error())
		return
	}

	var sg = snakegame.CreateSnakeGame(config.Board.Hight, config.Board.Wight)

	go sg.UserControl()
	score := sg.Run()

	fmt.Printf("Game Over!\nScore: %d\n", score)
}
