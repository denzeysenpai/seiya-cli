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
	fmt.Println("Welcome to Seiya CLI!")
	fmt.Println()
	config, err := os.Open("./config.json")
	var cfg *Config

	if err != nil {

	}
	defer config.Close()

	bytes, err := io.ReadAll(config)
	if err != nil {

	}

	if err = json.Unmarshal(bytes, &cfg); err != nil {

	}

	for {
		data, hasInput := ConsoleLine(cfg)
		_ = data
		_ = hasInput

		if len(data) > 0 {
			switch data[0] {
			case "start":
				StartNewTaskDirectory(data, cfg)
			case "new":
				NewTask(data)
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
				if len(data) > 1 {
					cfg.Use(data)
				}
			case "back":
				cfg.Back()
			}
		}
	}
}

func ConsoleLine(cfg *Config) ([]string, bool) {
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
