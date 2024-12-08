package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	SeiyaDirectory string `json:"seiyaDirectory"`
	CurrentWalk    string `json:"currentWalk"`
}

func (config *Config) StartNewTaskDirectory(data []string) {
	entries, err := os.ReadDir(config.SeiyaDirectory)
	if err != nil {

	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasSuffix(entry.Name(), "(on going)") {
			fmt.Println("Can't start a new task directory, you have an ongoing task!")
			return
		}
	}

	newTaskDirectory := fmt.Sprintf("%s/WEEKLY-TASK-"+strconv.Itoa(len(entries))+"(on going)", config.SeiyaDirectory)
	if err := os.Mkdir(newTaskDirectory, fs.ModePerm); err != nil {
		return
	}
	fmt.Println("Started a new task!")

}

func (cfg *Config) NewTask(data []string) {
	if len(data) > 1 { // check if input data is valid
		path := cfg.GetCurrentWalkPath()

		entries, err := os.ReadDir(path + "/")
		CheckEror(err)

		var taskType string = data[1]

		// check if task name already exists
		for _, entry := range entries {
			if taskType == HEADER {
				if entry.IsDir() && entry.Name() == data[2] {
					fmt.Println("That name is already taken!")
					return
				}
			} else if taskType == TASK {
				if entry.Name() == data[2] {
					fmt.Println("That name is already taken!")
					return
				}
			}
		}

		if taskType == HEADER {
			// header is just a folder
			if err := os.Mkdir(path+"/"+data[2], fs.ModePerm); err != nil {
				return
			}
		} else if taskType == TASK {
			// normal task is just a normal txt file
			file, err := os.Create(path + "/" + data[2] + ".txt")
			CheckEror(err)

			defer file.Close()
		}
	}
}

func Edit(data []string) {

}

func Delete(data []string) {

}

func Undo() {

}

func Redo() {

}

func Done(data []string) {

}

func Reversal(data []string) {
}

func (cfg *Config) View() {
	path := cfg.GetCurrentWalkPath()
	entries, err := os.ReadDir(path + "/")
	if err != nil {

	}

	for index, entry := range entries {
		if entry.IsDir() {
			fmt.Println("[" + strconv.Itoa(index) + "]" + entry.Name() + " HEADER")
		} else {
			fmt.Println("[" + strconv.Itoa(index) + "]" + entry.Name())
		}
	}
}

func (cfg *Config) Use(data []string) *Config {
	if len(data) > 1 {
		newWalk := cfg.CurrentWalk
		depth := strings.Split(newWalk, "/")
		var color string = Blue
		if len(depth)%2 == 0 {
			color = Magenta
		}

		path := cfg.GetCurrentWalkPath()
		entries, err := os.ReadDir(path + "/")

		if err != nil {

		}

		var toUse string
		indexToUse, _ := strconv.Atoi(data[1])

		if len(entries)-1 >= indexToUse {
			for index, entry := range entries {
				if index == indexToUse {
					toUse = entry.Name()
				}
			}
		}
		cfg.CurrentWalk = newWalk + color + "/" + toUse
	}
	return cfg
}

func (cfg *Config) Back() *Config {
	newWalk := cfg.CurrentWalk

	newWalk = strings.TrimPrefix(newWalk, "/")
	var walked = strings.Split(newWalk, "/")
	var newWalkedString string = ""

	if len(walked) > 1 {
		for _, step := range walked[:len(walked)-1] {
			newWalkedString = newWalkedString + "/" + step
		}
	}
	cfg.CurrentWalk = strings.TrimLeft(newWalkedString, "/")
	return cfg
}
