package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"seiya-cli/src/functions"
	"seiya-cli/src/utils"
)

func Start() {
	var cfg *functions.Config
	//
	fmt.Println("Welcome to Seiya CLI!")
	fmt.Println()
	config, err := os.Open("./config.json")
	utils.CheckEror(err)
	defer config.Close()

	bytes, err := io.ReadAll(config)
	utils.CheckEror(err)

	err = json.Unmarshal(bytes, &cfg)
	utils.CheckEror(err)

}
