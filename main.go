package main

import (
	"fmt"
	"net/http"
	"os"
)

const introductionApplicationName = `

							______   _______  _        _        _______  _        _______ 
							(  ___ \ (  ____ \( \      ( \      (  ___  )( (    /|(  ___  )
							| (   ) )| (    \/| (      | (      | (   ) ||  \  ( || (   ) |
							| (__/ / | (__    | |      | |      | |   | ||   \ | || (___) |
							|  __ (  |  __)   | |      | |      | |   | || (\ \) ||  ___  |
							| (  \ \ | (      | |      | |      | |   | || | \   || (   ) |
							| )___) )| (____/\| (____/\| (____/\| (___) || )  \  || )   ( |
							|/ \___/ (_______/(_______/(_______/(_______)|/    )_)|/     \|
																													

`

func main() {

	displayIntroduction()

	for {
		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
			break
		case 2:
			fmt.Println("Displaying logs...")
			break
		case 0:
			fmt.Println("Leaving the program...")
			os.Exit(0)
			break
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
			break
		}
	}
}

func displayMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Display logs")
	fmt.Println("0- Exit")
}

func displayIntroduction() {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, introductionApplicationName)
	fmt.Println(colored)

	nome := "Gedan"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Current program version:", versao)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The selected command was:", command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")

	// site := "http://patorjk.com"
	site := "https://random-status-code.herokuapp.com"

	resp, _ := http.Get(site)
	fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "successfully loaded!")
	} else {
		fmt.Println("Site", site, "has error. Status code", resp.StatusCode)
	}
}
