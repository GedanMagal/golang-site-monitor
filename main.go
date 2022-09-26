package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
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
	fmt.Println("Olá,", nome)
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

	sites := []string{"http://patorjk.com", "https://random-status-code.herokuapp.com"}

	for _, site := range sites {
		fmt.Println("Site:", site)

		resp, _ := http.Get(site)
		fmt.Println(resp)

		if resp.StatusCode == 200 {
			fmt.Println(site, "successfully loaded!")
		} else {
			fmt.Println(site, "has error. Status code", resp.StatusCode)
		}
	}
}

func getNames() []string {

	return []string{"Testando", "teste2"}
}

func testSliceCapacity() {
	nomes := getNames()

	fmt.Println("Length:", len(nomes), "capacity:", cap(nomes))

	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println(nomes)

	nomes = append(nomes, "Teste 2", "teste 3")
	fmt.Println("Length:", len(nomes), "capacity:", cap(nomes))
	fmt.Println(nomes)
}
