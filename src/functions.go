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

func StartNewTaskDirectory(data []string, config *Config) {
	fmt.Println("Started new Weekly Task!")
	entries, err := os.ReadDir(config.SeiyaDirectory)
	if err != nil {

	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasSuffix(entry.Name(), "(on going)") {
			return
		}
	}

	newTaskDirectory := fmt.Sprintf("%s/TASK-"+strconv.Itoa(len(entries))+"(on going)", config.SeiyaDirectory)
	if err := os.Mkdir(newTaskDirectory, fs.ModePerm); err != nil {

	}
}

func NewTask(data []string) {

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

func (cfg *Config) Use(data []string) *Config {
	newWalk := cfg.CurrentWalk
	cfg.CurrentWalk = newWalk + "/" + data[1]
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
	cfg.CurrentWalk = newWalkedString
	return cfg
}
