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
	depth := strings.Split(newWalk, "/")
	var color string = Blue
	if len(depth)%2 == 0 {
		color = Magenta
	}
	cfg.CurrentWalk = newWalk + color + "/" + data[1]
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
