package main

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/xtraice/pokedexcli/internal/cli"
)

var exit bool = false

const termPrefix = "pokedex > "

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	for !exit {
		execCmd := false
		fmt.Print("pokedex > ")
		for !execCmd && !exit {
			command := readCommand()
			if command != nil {
				execCmd = true
				cli.AddPrevCmd(strings.Join(command, " "))
				executeCommand(command)
			}
		}
	}
}

// function that reads in any keypress and returns the byte value
func readInput() rune {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch != 0 {
				return ev.Ch
			} else {
				return rune(ev.Key)
			}
		}
	}
}

// print string on terminal input line
func printOnTerminalInputLine(str string) {
	fmt.Print("\r\033[K")
	fmt.Print(termPrefix + str)
}

// readCommand reads user input from the terminal and returns a slice of strings representing the command entered.
// It continuously reads input until the user presses the Enter key.
// The function handles various key events, such as arrow keys for command history traversal and backspace for deleting characters.
// If the user presses the Esc key, the function sets the `exit` flag to true and returns nil.
// The returned command is split into individual strings based on space characters.
func readCommand() []string {
	var command []string
	str := ""
	input := []rune{}
	for {
		c := readInput()
		switch c {
		default:
			input = append(input, c)
			fmt.Print(string(c))
		case rune(termbox.KeyArrowUp):
			str = cli.TraversePrevCmds()
			printOnTerminalInputLine(str)
			input = []rune(str)
		case rune(termbox.KeyArrowDown):
			str = cli.TraverseNextCmds()
			printOnTerminalInputLine(str)
			input = []rune(str)
		case rune(termbox.KeyEsc):
			exit = true
			return nil
		case rune(termbox.KeyBackspace2), rune(termbox.KeyBackspace):
			if len(input) > 0 {
				input = input[:len(input)-1]
				fmt.Print("\b \b")
			}
		case rune(termbox.KeyEnter):
			command = strings.Split(string(input), " ")
			fmt.Println()
			return command
		}
		termbox.SetCursor(len(input), 0)
	}
}

func executeCommand(command []string) {
	if _, ok := cli.CliCommands[command[0]]; !ok {
		fmt.Println("Command not found")
		return
	}
	if err := cli.CliCommands[command[0]].Callback(command); err != nil {
		exit = true
	}
}
