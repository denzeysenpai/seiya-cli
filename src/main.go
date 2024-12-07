package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var cfg *Config
	//
	fmt.Println("Welcome to Seiya CLI!")
	fmt.Println()
	config, err := os.Open("./config.json")
	CheckEror(err)

	defer config.Close()

	bytes, err := io.ReadAll(config)
	CheckEror(err)

	err = json.Unmarshal(bytes, &cfg)
	CheckEror(err)

	for {
		data, hasInput := cfg.ConsoleLine()
		_ = data
		_ = hasInput

		cfg.InputProcessing(data)
	}
}

func (cfg *Config) InputProcessing(data []string) {
	if len(data) > 0 {
		switch data[0] {
		case "start":
			cfg.StartNewTaskDirectory(data)
		case "new":
			cfg.NewTask(data)
		case "edit":
			Edit(data)
		case "delete":
			Delete(data)
		case "undo":
			Undo()
		case "redo":
			Redo()
		case "done":
			Done(data)
		case "reversal":
			Reversal(data)
		case "use":
			cfg.Use(data)
		case "back":
			cfg.Back()
		case "view":
			cfg.View()
		}
	}
}

func (cfg *Config) ConsoleLine() ([]string, bool) {
	var output []string = []string{}
	var currentWalk string = cfg.CurrentWalk

	input := bufio.NewReader(os.Stdin)
	fmt.Print(Green + "Seiya" + Magenta + currentWalk + Yellow + ">> " + Reset)
	line, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("...")
	}

	line = strings.TrimSpace(line)

	if len(line) > 0 {
		var split = strings.Split(line, " /")
		for _, data := range split {
			if len(data) > 0 {
				output = append(output, data)
			}
		}
	}

	return output, len(output) > 0
}
