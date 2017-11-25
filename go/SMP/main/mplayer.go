package main

import (
	"SMP/library"
	"fmt"
	"strconv"
	"SMP/mp"
	"bufio"
	"os"
	"strings"
)

type Music struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

var lib *library.MusicManager
var id int = 0
//????
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
			id++
		} else {
			fmt.Println("Usage: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("Usage: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:",tokens[1])
	}
}


func handlePalayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "doesnt exist.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`Enter following commands to control the player:
lib list -- View the existing music lib 8
lib add <name><artist><source><type> -- Add a music to the music lib
lib remove <name> -- Remove the specified music from the lib
play <name> -- Play the specified music
`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command ->")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePalayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command", tokens[0])
		}
	}
}