package main

import (
	"bufio"
	"fmt"
	"github.com/captncraig/gpiet/machine"
	"github.com/captncraig/gpiet/micropiet"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".mp") {
			continue
		}
		runTest(f.Name())
	}
}

func runTest(name string) {
	fmt.Println("--------------")
	fmt.Println(name)
	fmt.Println("--------------")

	vm := machine.NewMachine()

	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	text := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "@@") {
			prog, err := micropiet.Parse(text)
			if err != nil {
				log.Fatalln(err)
			}
			prog.Run(vm)
			text = ""
		} else {
			text = text + line + "\n"
		}
	}
}
