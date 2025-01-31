package functions

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"seiya-cli/src/utils"
	"strconv"
	"strings"
)

type Config struct {
	SeiyaDirectory string `json:"seiyaDirectory"`
	CurrentWalk    string `json:"currentWalk"`
}

func ConsoleLineStart(cfg *Config) {
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
	fmt.Print(utils.Green + "Seiya" + utils.Magenta + currentWalk + utils.Yellow + ">> " + utils.Reset)
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

func (cfg *Config) GetCurrentWalkPath() string {
	path := strings.ReplaceAll(cfg.CurrentWalk, utils.Magenta, "")
	path = strings.ReplaceAll(path, utils.Blue, "")
	path = cfg.SeiyaDirectory + path
	return path
}

func (config *Config) StartNewTaskDirectory(data []string) {
	entries, err := os.ReadDir(config.SeiyaDirectory)
	if err != nil {

	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasSuffix(entry.Name(), "(on going)") {
			fmt.Print(utils.Red)
			fmt.Println("Can't start a new task directory, you have an on going task!")
			fmt.Print(utils.Reset)
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

		if strings.HasSuffix(path, "seiya") {
			fmt.Print(utils.Red)
			fmt.Println("You can't create anything other than main task directories in this directory!")
			fmt.Println("Use: 'start' to create a new task directory instead!")
			fmt.Print(utils.Reset)
			return
		}

		entries, err := os.ReadDir(path + "/")
		utils.CheckEror(err)

		var taskType string = data[1]

		// check if task name already exists
		for _, entry := range entries {
			if taskType == utils.HEADER {
				if entry.IsDir() && entry.Name() == data[2] {
					fmt.Println("That name is already taken!")
					return
				}
			} else if taskType == utils.TASK {
				if entry.Name() == data[2] {
					fmt.Println("That name is already taken!")
					return
				}
			}
		}

		if taskType == utils.HEADER {
			// header is just a folder
			if err := os.Mkdir(path+"/"+data[2]+" (on going)", fs.ModePerm); err != nil {
				return
			}
		} else if taskType == utils.TASK {
			// normal task is just a normal txt file
			file, err := os.Create(path + "/" + data[2] + ".txt")
			utils.CheckEror(err)

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
	utils.CheckEror(err)
	if strings.HasSuffix(path, ".txt") {
		fmt.Println(utils.Red + "This type is invalid, can't view children!" + utils.Reset)
		return
	} else {
		if len(entries) == 0 {
			fmt.Println(utils.Red + "This directory currently has no children!" + utils.Reset)
			return
		}
	}

	// displays the children of the current walk path
	fmt.Println(utils.Yellow + "TYPE\tFILENAME")
	fmt.Println(utils.Yellow + "============================")
	for index, entry := range entries {
		if entry.IsDir() {
			fmt.Println(utils.Yellow + "HEADER\t" + utils.Reset + "[" + strconv.Itoa(index) + "] " + entry.Name() + "\t")
		} else {
			fmt.Println(utils.Yellow + "TASK\t" + utils.Reset + "[" + strconv.Itoa(index) + "] " + entry.Name())
		}
	}
	fmt.Println()
}

func (cfg *Config) Use(data []string) *Config {
	if len(data) > 1 {
		newWalk := cfg.CurrentWalk
		depth := strings.Split(newWalk, "/")
		var color string = utils.Blue
		if len(depth)%2 == 0 {
			color = utils.Magenta
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
